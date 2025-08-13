# ArgoCD Examples for Crossplane OCI Provider

Create OKE clusters using ArgoCD and Crossplane.

## Prerequisites

- Kubernetes cluster (1.30+), kubectl, Docker, Go 1.21+, Terraform 1.4.6+
- Container registry (Docker Hub, etc.) with `docker login` completed
- OCI tenancy with API credentials

## Quick Start

### 1. Build and Publish Providers

Replace `myusername` with your registry username:

```bash
make generate

VERSION=v0.0.1-test \
XPKG_REG_ORGS=docker.io/myusername \
make publish-subpackages \
SUBPACKAGES_FOR_BATCH="config,containerengine,networking,identity,objectstorage,loadbalancer,networkloadbalancer" \
BATCH_PLATFORMS=linux_amd64
```

### 2. Install ArgoCD

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd
```

### 3. Configure OCI Credentials

Create secret with your OCI API credentials:

```bash
cp examples/providerconfig/secret.yaml.tmpl oci-secret.yaml
# Edit oci-secret.yaml: add tenancy OCID, user OCID, private key, region, fingerprint
kubectl apply -f oci-secret.yaml
```

### 4. Configure ArgoCD

```bash
kubectl patch configmap argocd-cm -n argocd --patch-file argocd-examples/argocd/config/argocd-crossplane-config.yaml
kubectl rollout restart deployment/argocd-server -n argocd
kubectl wait --for=condition=available --timeout=300s deployment/argocd-server -n argocd
```

### 5. Update Repository and Registry URLs

Replace with your forked repository and registry:

```bash
# Update git repository URLs
sed -i 's|repoURL: https://github.com/.*|repoURL: https://github.com/YOUR-USERNAME/YOUR-REPO|g' argocd-examples/argocd/applications/*.yaml

# Update package registry URLs  
find argocd-examples/crossplane/providers -name "*.yaml" -exec sed -i 's|package: .*|package: docker.io/myusername/provider-family-oci:v0.0.1-test|g' {} \;
```

### 6. Customize Cluster Configuration

Edit cluster configuration with your OCI details:

```bash
# Update: compartmentId, region, availabilityDomain, nodeImageId
vim argocd-examples/oke-claims/claim-okecluster.yaml
```

### 7. Deploy Applications

Deploy in dependency order:

```bash
kubectl apply -f argocd-examples/argocd/applications/crossplane.yaml
kubectl apply -f argocd-examples/argocd/applications/oci-providers-suite.yaml
kubectl apply -f argocd-examples/argocd/applications/oke-platform.yaml
kubectl apply -f argocd-examples/argocd/applications/oke-claims.yaml
```

### 8. Monitor

```bash
kubectl get clusterclaims
kubectl get managed
```

## Cleanup

```bash
kubectl delete clusterclaim test-oke-cluster
```
