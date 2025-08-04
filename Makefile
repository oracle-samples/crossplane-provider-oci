# ====================================================================================
# Setup Project

PROJECT_NAME := provider-oci
PROJECT_REPO := github.com/oracle/$(PROJECT_NAME)

export TERRAFORM_VERSION := 1.4.6

export TERRAFORM_PROVIDER_SOURCE := oracle/oci
export TERRAFORM_PROVIDER_REPO := https://github.com/oracle/terraform-provider-oci
export TERRAFORM_PROVIDER_VERSION := 7.12.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-oci
export TERRAFORM_NATIVE_PROVIDER_BINARY := terraform-provider-oci_v$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX := https://releases.hashicorp.com/terraform-provider-oci/$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_DOCS_PATH := website/docs/r

export CROSSPLANE_PROVIDER_VERSION := 1.0
# Insert Oracle-CrossplaneProvider/<version> to terraform oci User-Agent using
# USER_AGENT_PROVIDER_NAME (default=Oracle-TerraformProvider).
# The value will be inserted as <USER_AGENT_PROVIDER_NAME>/<terraform oci version>
# Sample from terraform oci verbose log:
# User-Agent: Oracle-GoSDK/65.37.1 ... Oracle-CrossplaneProvider/1.0 Oracle-TerraformProvider/4.120.0
export USER_AGENT_PROVIDER_NAME := Oracle-CrossplaneProvider/$(CROSSPLANE_PROVIDER_VERSION) Oracle-TerraformProvider

PLATFORMS ?= darwin_amd64 linux_arm64 linux_amd64

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

GO_REQUIRED_VERSION ?= 1.24
GOLANGCILINT_VERSION ?= 1.50.0
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal apis
-include build/makelib/golang.mk

# ====================================================================================
# Setup Sub-packages
# SUBPACKAGES can be:
# - "monolith" (default): builds the traditional single provider binary
# - comma-separated service names: builds only specified service sub-packages
#   e.g., "config,compute,networking"
SUBPACKAGES ?= monolith

# Helper variable for comma
comma := ,

# Configure GO_STATIC_PACKAGES based on SUBPACKAGES value
ifeq ($(SUBPACKAGES),monolith)
# Monolith build - traditional behavior
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
else
# Sub-package build - parse comma-separated list and build service-specific binaries
SUBPACKAGE_LIST := $(subst $(comma), ,$(SUBPACKAGES))
GO_STATIC_PACKAGES = $(foreach pkg,$(SUBPACKAGE_LIST),$(GO_PROJECT)/cmd/provider/$(pkg))
endif

# ====================================================================================
# Setup Kubernetes tools

KIND_VERSION = v0.29.0
UP_VERSION = v0.39.0
UP_CHANNEL = stable
UPTEST_VERSION = v0.2.1
-include build/makelib/k8s_tools.mk

# ====================================================================================
# Setup Images

REGISTRY_ORGS ?= xpkg.upbound.io/upbound
IMAGES = $(PROJECT_NAME)
-include build/makelib/imagelight.mk

# ====================================================================================
# Setup XPKG

XPKG_REG_ORGS ?= xpkg.upbound.io/upbound
# NOTE(hasheddan): skip promoting on xpkg.upbound.io as channel tags are
# inferred.
XPKG_REG_ORGS_NO_PROMOTE ?= xpkg.upbound.io/upbound
XPKGS = $(PROJECT_NAME)
-include build/makelib/xpkg.mk

# NOTE(hasheddan): we force image building to happen prior to xpkg build so that
# we ensure image is present in daemon.
xpkg.build.provider-oci: do.build.images

# NOTE(hasheddan): we ensure up is installed prior to running platform-specific
# build steps in parallel to avoid encountering an installation race condition.
build.init: $(UP)

# Override build behavior when SUBPACKAGES is set to non-monolith
ifneq ($(SUBPACKAGES),monolith)
# When building sub-packages, override the build target
build:
	@$(INFO) Building sub-packages: $(SUBPACKAGE_LIST)
	@for pkg in $(SUBPACKAGE_LIST); do \
		if [ -d "cmd/provider/$$pkg" ]; then \
			echo "Building $$pkg..."; \
			if [ "$$pkg" = "config" ]; then \
				output_name="provider-family-oci"; \
			else \
				output_name="provider-oci-$$pkg"; \
			fi; \
			CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -ldflags '$(GO_LDFLAGS)' \
				-o $(GO_OUT_DIR)/$$output_name $(GO_PROJECT)/cmd/provider/$$pkg/ || exit 1; \
		else \
			echo "Warning: Sub-package directory cmd/provider/$$pkg does not exist."; \
			echo "Note: Sub-package directories will be created by the code generator in Phase 2."; \
		fi; \
	done
	@$(OK) Built sub-packages: $(SUBPACKAGE_LIST)
endif

# ====================================================================================
# Fallthrough

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# ====================================================================================
# Setup Terraform for fetching provider schema
TERRAFORM := $(TOOLS_HOST_DIR)/terraform-$(TERRAFORM_VERSION)
TERRAFORM_WORKDIR := $(WORK_DIR)/terraform
TERRAFORM_PROVIDER_SCHEMA := config/schema.json

$(TERRAFORM):
	@$(INFO) installing terraform $(HOSTOS)-$(HOSTARCH)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-terraform
	@curl -fsSL https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/terraform_$(TERRAFORM_VERSION)_$(SAFEHOST_PLATFORM).zip -o $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip
	@unzip $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip -d $(TOOLS_HOST_DIR)/tmp-terraform
	@mv $(TOOLS_HOST_DIR)/tmp-terraform/terraform $(TERRAFORM)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-terraform
	@$(OK) installing terraform $(HOSTOS)-$(HOSTARCH)

$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)
	@$(INFO) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)
	@mkdir -p $(TERRAFORM_WORKDIR)
	@echo '{"terraform":[{"required_providers":[{"provider":{"source":"'"$(TERRAFORM_PROVIDER_SOURCE)"'","version":"'"$(TERRAFORM_PROVIDER_VERSION)"'"}}],"required_version":"'"$(TERRAFORM_VERSION)"'"}]}' > $(TERRAFORM_WORKDIR)/main.tf.json
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) init > $(TERRAFORM_WORKDIR)/terraform-logs.txt 2>&1
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) providers schema -json=true > $(TERRAFORM_PROVIDER_SCHEMA) 2>> $(TERRAFORM_WORKDIR)/terraform-logs.txt
	@$(OK) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)

pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" ]; then \
  		mkdir -p "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" && \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)"; \
	fi
	@git -C "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"

generate.init: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs

.PHONY: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs
# ====================================================================================
# Targets

# NOTE: the build submodule currently overrides XDG_CACHE_HOME in order to
# force the Helm 3 to use the .work/helm directory. This causes Go on Linux
# machines to use that directory as the build cache as well. We should adjust
# this behavior in the build submodule because it is also causing Linux users
# to duplicate their build cache, but for now we just make it easier to identify
# its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_ | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	UPBOUND_CONTEXT="local" $(GO_OUT_DIR)/provider --debug

run-only:
	@$(INFO) Running Crossplane locally out-of-cluster w/o build . . .
	@# To see other arguments that can be provided, run the command with --help instead
	UPBOUND_CONTEXT="local" $(GO_OUT_DIR)/provider --debug

# ====================================================================================
# End to End Testing
CROSSPLANE_NAMESPACE = upbound-system
-include build/makelib/local.xpkg.mk
-include build/makelib/controlplane.mk

uptest: $(UPTEST) $(KUBECTL) $(KUTTL)
	@$(INFO) running automated tests
	@KUBECTL=$(KUBECTL) KUTTL=$(KUTTL) $(UPTEST) e2e "${UPTEST_EXAMPLE_LIST}" --setup-script=cluster/test/setup.sh || $(FAIL)
	@$(OK) running automated tests

local-deploy: build controlplane.up local.xpkg.deploy.provider.$(PROJECT_NAME)
	@$(INFO) running locally built provider
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Healthy --timeout 5m
	@$(KUBECTL) -n upbound-system wait --for=condition=Available deployment --all --timeout=5m
	@$(OK) running locally built provider

e2e: local-deploy uptest

.PHONY: cobertura submodules fallthrough run crds.clean

# ====================================================================================
# Sub-package Targets
# These targets support building and packaging individual service sub-packages

# Build a specific sub-package
# Usage: make build.subpackage.compute
build.subpackage.%:
	@if [ "$*" = "config" ]; then \
		$(INFO) Building sub-package provider-family-oci binary; \
		output_name="provider-family-oci"; \
	else \
		$(INFO) Building sub-package provider-oci-$* binary; \
		output_name="provider-oci-$*"; \
	fi
	@if [ ! -d "cmd/provider/$*" ]; then \
		echo "Error: Sub-package directory cmd/provider/$* does not exist."; \
		echo "Note: Sub-package directories will be created by the code generator in Phase 2."; \
		echo "For now, this is expected behavior as we're implementing the build system first."; \
		exit 1; \
	fi
	@if [ "$*" = "config" ]; then \
		output_name="provider-family-oci"; \
	else \
		output_name="provider-oci-$*"; \
	fi; \
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -ldflags '$(GO_LDFLAGS)' -o $(GO_OUT_DIR)/$$output_name $(GO_PROJECT)/cmd/provider/$*/
	@if [ "$*" = "config" ]; then \
		$(OK) Built sub-package provider-family-oci binary; \
	else \
		$(OK) Built sub-package provider-oci-$* binary; \
	fi

# Package a specific sub-package as a container
# Usage: make package.subpackage.compute
package.subpackage.%: build.subpackage.%
	@if [ "$*" = "config" ]; then \
		$(INFO) Packaging sub-package provider-family-oci container; \
		binary_name="provider-family-oci"; \
		package_name="provider-family-oci"; \
	else \
		$(INFO) Packaging sub-package provider-oci-$* container; \
		binary_name="provider-oci-$*"; \
		package_name="provider-oci-$*"; \
	fi; \
	docker build -t $(REGISTRY_ORGS)/$$package_name:$(VERSION) \
		--build-arg BINARY=$$binary_name \
		--build-arg PACKAGE_NAME=$$package_name \
		-f package/Dockerfile $(GO_OUT_DIR)
	@if [ "$*" = "config" ]; then \
		$(OK) Packaged sub-package provider-family-oci container; \
	else \
		$(OK) Packaged sub-package provider-oci-$* container; \
	fi

# Test a specific sub-package
# Usage: make test.subpackage.compute
test.subpackage.%:
	@$(INFO) Testing sub-package $*
	@if [ ! -d "cmd/provider/$*" ] && [ ! -d "internal/controller/$*" ]; then \
		echo "Error: Sub-package directories for '$*' do not exist."; \
		echo "Note: Sub-package directories will be created by the code generator in Phase 2."; \
		echo "For now, this is expected behavior as we're implementing the build system first."; \
		exit 1; \
	fi
	@if [ -d "cmd/provider/$*" ]; then \
		$(GO) test -v $(GO_PROJECT)/cmd/provider/$*/...; \
	fi
	@if [ -d "internal/controller/$*" ]; then \
		$(GO) test -v $(GO_PROJECT)/internal/controller/$*/...; \
	fi
	@$(OK) Testing sub-package $*

# Generate CRDs for specific sub-packages
# Usage: make generate.subpackage.crds SUBPACKAGES="compute,networking"
generate.subpackage.crds:
	@$(INFO) Generating CRDs for sub-packages: $(SUBPACKAGES)
	@if [ "$(SUBPACKAGES)" = "monolith" ]; then \
		echo "Error: SUBPACKAGES must specify service names, not 'monolith'"; \
		exit 1; \
	fi
	@# This target will be implemented when the generator supports sub-package CRD generation
	@$(OK) Generating CRDs for sub-packages: $(SUBPACKAGES)

# Debug target to verify SUBPACKAGES configuration
debug-subpackages:
	@echo "SUBPACKAGES: $(SUBPACKAGES)"
	@echo "SUBPACKAGE_LIST: $(SUBPACKAGE_LIST)"
	@echo "GO_STATIC_PACKAGES: $(GO_STATIC_PACKAGES)"

.PHONY: build.subpackage.% package.subpackage.% test.subpackage.% generate.subpackage.crds debug-subpackages

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    cobertura             Generate a coverage report for cobertura applying exclusions on generated files.
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.

Sub-package Build Support:
    SUBPACKAGES           Variable to control build behavior (default: "monolith")
                          Examples:
                            make build                                # builds monolithic provider (default)
                            make build SUBPACKAGES=config             # builds only config sub-package
                            make build SUBPACKAGES=config,compute     # builds multiple sub-packages

Sub-package Targets:
    build.subpackage.%    Build a specific sub-package binary
                          Example: make build.subpackage.compute
    package.subpackage.%  Package a specific sub-package as container
                          Example: make package.subpackage.networking
    test.subpackage.%     Run tests for a specific sub-package
                          Example: make test.subpackage.blockstorage
    generate.subpackage.crds  Generate CRDs for specified sub-packages
                          Example: make generate.subpackage.crds SUBPACKAGES="compute,networking"

Available sub-packages:
    config, compute, networking, blockstorage, networkconnectivity, containerengine,
    identity, objectstorage, loadbalancer, networkloadbalancer, dns, kms, functions,
    logging, monitoring, events, streaming, filestorage, artifacts, vault, ons,
    certificatesmanagement, networkfirewall, servicemesh, healthchecks

endef
# The reason CROSSPLANE_MAKE_HELP is used instead of CROSSPLANE_HELP is because the crossplane
# binary will try to use CROSSPLANE_HELP if it is set, and this is for something different.
export CROSSPLANE_MAKE_HELP

crossplane.help:
	@echo "$$CROSSPLANE_MAKE_HELP"

help-special: crossplane.help

.PHONY: crossplane.help help-special
