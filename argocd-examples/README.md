# ArgoCD Examples for Crossplane OCI Provider

This directory contains ArgoCD applications and Kubernetes manifests to deploy Crossplane with the OCI provider sub-packages using GitOps workflows. This setup uses Crossplane Composite Resource Definitions (XRDs) and Compositions to create and manage OKE (Oracle Kubernetes Engine) clusters through high-level, declarative APIs.

> **Note**: All commands in this guide should be run from the `crossplane-provider-oci` project root directory, not from within the `argocd-examples` subdirectory.

## Prerequisites

Before starting, ensure you have:

### Software Requirements

- [Git](https://git-scm.com/downloads)
- [Terraform](https://developer.hashicorp.com/terraform/downloads) 1.4.6+ (required by provider)
- [Go](https://go.dev/doc/install) 1.21+
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Docker](https://docs.docker.com/engine/install/)

### Environment Setup

- Kubernetes cluster (OKE, Rancher Desktop, or any K8s 1.30+)
- Correct Kubernetes context selected
- `$GOPATH` configured properly
- OCI tenancy with proper credentials configured

### Container Registry Access

- Container registry access (Docker Hub, OCIR, GitHub Container Registry, etc.)
- Registry credentials configured locally
- Docker logged in to your chosen registry:

  ```bash
  # Examples for different registries:
  docker login                           # Docker Hub
  docker login <region>.ocir.io          # OCI Registry (OCIR)
  docker login ghcr.io                   # GitHub Container Registry
  docker login <your-registry>           # Custom registry
  ```

## Quick Start Guide

### Step 1: Build and Publish OCI Provider Sub-packages

1. **Generate CRDs and APIs** (required after any API changes):

   ```bash
   make generate
   ```

2. **Build and publish sub-packages to your registry**:

   **For OKE cluster deployment**, use these specific packages:

   ```bash
   VERSION=v0.0.1-test \
   XPKG_REG_ORGS=your-registry/your-org \
   make publish-subpackages \
   SUBPACKAGES_FOR_BATCH="config,objectstorage,identity,networkloadbalancer,loadbalancer,containerengine,networking" \
   BATCH_PLATFORMS=linux_amd64
   ```

   **Required Variables:**
   - `VERSION`: Your package version (e.g., `v0.0.1-test`, `v1.0.0`)
   - `XPKG_REG_ORGS`: Your registry URL and organization (examples below)
   - `SUBPACKAGES_FOR_BATCH`: Comma-separated list of sub-packages to build
   - `BATCH_PLATFORMS`: Target platform(s) - usually `linux_amd64` or `linux_amd64,linux_arm64`

   **Registry Examples:**

   ```bash
   # Docker Hub
   XPKG_REG_ORGS=docker.io/myusername
   
   # OCI Registry (OCIR) 
   XPKG_REG_ORGS=iad.ocir.io/my-tenancy-namespace
   
   # GitHub Container Registry
   XPKG_REG_ORGS=ghcr.io/myusername
   
   # Private registry
   XPKG_REG_ORGS=registry.company.com/team
   ```

   **Available Sub-packages:**
   - `config`: Configuration and authentication (family provider) - **REQUIRED**
   - `networking`: VCN, subnets, gateways, security lists
   - `containerengine`: OKE clusters and node pools
   - `compute`: Instances, images, dedicated hosts
   - `blockstorage`: Volumes, backups, volume groups
   - `identity`: Compartments, policies, groups
   - `objectstorage`: Buckets and objects
   - `loadbalancer`: Layer 7 load balancers
   - `networkloadbalancer`: Layer 4 load balancers
   - `networkconnectivity`: DRG, IPSec, FastConnect
   - And 13 more services (dns, kms, functions, logging, monitoring, etc.)

### Step 2: Install ArgoCD

Install ArgoCD in your Kubernetes cluster:

```bash
# Create ArgoCD namespace
kubectl create namespace argocd

# Install ArgoCD
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Wait for ArgoCD to be ready
kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd
```

### Step 3: Configure OCI Credentials

**IMPORTANT**: Configure OCI credentials before deploying providers, as they're required for OCI API authentication.

1. **Create OCI credentials secret**:

   ```bash
   # Copy and fill the secret template from main repo
   cp examples/providerconfig/secret.yaml.tmpl oci-secret.yaml
   # Edit oci-secret.yaml with your OCI credentials (tenancy, user, key, etc.)
   kubectl apply -f oci-secret.yaml
   ```

   > **Note**: The secret must be named `oci-creds` in the `crossplane-system` namespace as expected by the provider configuration.

### Step 4: Configure ArgoCD for Crossplane

Apply the custom ArgoCD configuration that optimizes ArgoCD for Crossplane resources:

```bash
kubectl apply -f argocd-examples/argocd/config/argocd-crossplane-config.yaml
```

This configuration provides:

- Resource tracking optimization for Crossplane
- Custom health checks for OCI resources
- Ignore patterns for resolved references
- Exclusions for ProviderConfigUsage resources

### Step 5: Update ArgoCD Applications

Before applying the ArgoCD applications, you need to update them to point to your repository and registry:

1. **Update `argocd-examples/argocd/applications/oci-providers-suite.yaml`**:

   ```yaml
   spec:
     source:
       repoURL: https://github.com/YOUR-USERNAME/YOUR-REPO  # Update this
       targetRevision: main  # Update branch if needed
       path: argocd-examples/crossplane/providers/oci-sub-packages  # Update path
   ```

   **Similarly, update `argocd-examples/argocd/applications/oke-platform.yaml` and `argocd-examples/argocd/applications/oke-claims.yaml`**:

   ```yaml
   spec:
     source:
       repoURL: https://github.com/YOUR-USERNAME/YOUR-REPO  # Update this
       targetRevision: main  # Update branch if needed
       path: argocd-examples/oke-platform  # Update path for platform
       # OR
       path: argocd-examples/oke-claims    # Update path for claims
   ```

2. **Update provider configurations** in `argocd-examples/crossplane/providers/oci-sub-packages/*/provider.yaml` files to use your registry:

   ```yaml
   spec:
     package: YOUR-REGISTRY/provider-family-oci:YOUR-VERSION
     # or
     package: YOUR-REGISTRY/provider-oci-networking:YOUR-VERSION
   ```

### Step 6: Deploy ArgoCD Applications

Deploy the ArgoCD applications in the correct order using sync waves:

1. **Install Crossplane** (sync-wave: 1):

   ```bash
   kubectl apply -f argocd-examples/argocd/applications/crossplane.yaml
   ```

2. **Install OCI Provider Suite** (sync-wave: 2 - includes provider configuration automatically):

   ```bash
   kubectl apply -f argocd-examples/argocd/applications/oci-providers-suite.yaml
   ```

   > **Note**: This automatically deploys the provider configuration that references the `oci-creds` secret created in Step 3.

3. **Deploy OKE Platform** (sync-wave: 25 - XRDs and Compositions):

   ```bash
   kubectl apply -f argocd-examples/argocd/applications/oke-platform.yaml
   ```

4. **Deploy OKE Claims** (sync-wave: 30 - Cluster instances):

   ```bash
   kubectl apply -f argocd-examples/argocd/applications/oke-claims.yaml
   ```

   > **Note**: Claims will only succeed if OCI credentials are properly configured (Step 3).

### Step 7: Verify OKE Cluster Creation

After all applications are deployed and synced, verify your OKE cluster is being created:

1. **Check the cluster claim**:

   ```bash
   kubectl get clusterclaims
   kubectl describe clusterclaim oke-demo-cluster
   ```

2. **Monitor cluster creation progress**:

   ```bash
   # Check composite resource (XR)
   kubectl get xokeclusters
   
   # Check individual managed resources
   kubectl get managed
   ```

3. **Access OKE cluster when ready**:
   - The composition will create all necessary OCI resources (VCN, subnets, OKE cluster)
   - Once ready, you can get kubeconfig from OCI console or CLI

### Step 8: Access ArgoCD UI

1. **Get ArgoCD admin password**:

   ```bash
   kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
   ```

2. **Port forward to ArgoCD server**:

   ```bash
   kubectl port-forward svc/argocd-server -n argocd 8080:443
   ```

3. **Access ArgoCD UI**:
   - URL: <https://localhost:8080>
   - Username: `admin`
   - Password: (from step 1)

## Architecture Overview

### Application Structure

```
crossplane-provider-oci/       # Project root directory
├── argocd-examples/
│   ├── argocd/
│   │   ├── applications/           # ArgoCD application definitions
│   │   │   ├── crossplane.yaml           # Installs Crossplane (sync-wave: 1)
│   │   │   ├── oci-providers-suite.yaml  # Installs OCI providers (sync-wave: 2)
│   │   │   ├── oke-platform.yaml         # Deploys XRDs & Compositions (sync-wave: 25)
│   │   │   └── oke-claims.yaml           # Deploys cluster instances (sync-wave: 30)
│   │   └── config/
│   │       └── argocd-crossplane-config.yaml  # ArgoCD configuration for Crossplane
│   ├── crossplane/
│   │   └── providers/
│   │       └── oci-sub-packages/   # Provider configurations
│   │           ├── family/         # Family provider (config) - auth & ProviderConfig
│   │           ├── networking/     # Networking provider
│   │           ├── containerengine/# OKE provider
│   │           ├── identity/       # Identity provider
│   │           └── ...            # Other service providers
│   ├── oke-platform/              # Platform definitions (XRDs & Compositions)
│   │   ├── xrd-okecluster.yaml           # OKE Cluster XRD definition
│   │   ├── composition-okecluster.yaml   # OKE Cluster Composition
│   │   └── kustomization.yaml            # Kustomize configuration
│   └── oke-claims/                # Example cluster claims
│       ├── claim-okecluster.yaml         # Example OKE cluster claim
│       ├── kustomization.yaml            # Kustomize configuration
│       └── README.md                     # Claims documentation
├── Makefile                   # Build system with sub-package support
├── README.md                  # Main project documentation
└── ...                       # Other project files
```

### Deployment Flow

1. **Crossplane Installation**: ArgoCD installs Crossplane core using Helm chart
2. **Provider Suite**: ArgoCD deploys OCI providers as sub-packages (family, networking, containerengine, etc.)
3. **Platform Definition**: ArgoCD deploys XRDs and Compositions that define the OKE platform API
4. **Infrastructure Claims**: ArgoCD deploys example claims that create actual OKE clusters using the platform API

### XRD-Based Architecture Benefits

This setup uses **Composite Resource Definitions (XRDs)** and **Compositions** instead of individual OCI resources:

- **Higher-Level APIs**: Create entire OKE clusters with a single `ClusterClaim` resource
- **Abstraction**: Hide complexity of individual VCN, subnets, security groups, etc.
- **Reusability**: Define OKE platform once, create multiple clusters easily
- **Best Practices**: Enforce organizational standards through compositions
- **Simplified Management**: Manage clusters as single units rather than dozens of individual resources

### Sub-package Architecture Benefits

- **Modular Deployments**: Only install needed OCI services
- **Reduced Resource Usage**: 60-80% smaller deployments for targeted use cases
- **Independent Scaling**: Each service provider scales independently
- **Zero Circular Dependencies**: Services are completely independent

## Troubleshooting

### Common Issues

1. **Provider Not Ready**:

   ```bash
   kubectl get providers -n crossplane-system
   kubectl describe provider <provider-name> -n crossplane-system
   
   # Check if OCI credentials secret exists
   kubectl get secret oci-creds -n crossplane-system
   kubectl describe secret oci-creds -n crossplane-system
   ```

2. **ArgoCD Sync Issues**:
   - Check ArgoCD application status in UI
   - Verify repository access and credentials
   - Check sync waves are applied in order

3. **OKE Cluster Creation Issues**:

   ```bash
   # Check cluster claim status
   kubectl get clusterclaims
   kubectl describe clusterclaim <claim-name>
   
   # Check composite resource (XR)
   kubectl get xokeclusters
   kubectl describe xokecluster <xr-name>
   
   # Check individual managed resources
   kubectl get managed
   kubectl describe <managed-resource-name>
   ```

4. **OCI Credential Issues**:
   - Verify `oci-creds` secret exists in `crossplane-system` namespace
   - Check secret contains valid OCI credentials (tenancy, user, key, region, fingerprint)
   - Ensure OCI user has proper permissions for creating resources
   - Verify private key format and passphrase (if applicable)

5. **Registry Access Issues**:
   - Verify Docker login to your chosen registry
   - Check registry credentials and permissions
   - Ensure proper RBAC for pulling images from your registry

### Useful Commands

```bash
# Check all Crossplane resources
kubectl get crossplane

# Check provider status
kubectl get providers -n crossplane-system

# Check OCI credentials secret
kubectl get secret oci-creds -n crossplane-system

# Check XRDs and Compositions
kubectl get xrd
kubectl get compositions

# Check cluster claims and composite resources
kubectl get clusterclaims
kubectl get xokeclusters

# Check managed resources (actual OCI resources)
kubectl get managed

# Check ArgoCD applications
kubectl get applications -n argocd

# View provider logs
kubectl logs -n crossplane-system deployment/provider-oci-networking
kubectl logs -n crossplane-system deployment/provider-family-oci
```

## Next Steps

After successful deployment:

1. **Monitor ArgoCD Applications**: Check all applications are synced and healthy
2. **Verify OKE Cluster Creation**: Monitor the cluster claim until resources are ready
3. **Access OKE Cluster**: Download kubeconfig and test cluster connectivity
4. **Customize Compositions**: Modify compositions to match your requirements
5. **Create Additional Claims**: Deploy more clusters using the platform API
6. **Extend Platform**: Add more XRDs for other OCI services (databases, storage, etc.)
7. **Implement CI/CD**: Integrate with your existing CI/CD pipelines

## Additional Resources

- [Crossplane Documentation](https://docs.crossplane.io/)
- [ArgoCD Documentation](https://argoproj.github.io/argo-cd/)
- [OCI Provider Documentation](../README.md)
- [OCI Documentation](https://docs.oracle.com/en-us/iaas/)
