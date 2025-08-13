# OKE Cluster XRD (Composite Resource Definition)

This directory contains a complete XRD-based solution for deploying Oracle Kubernetes Engine (OKE) clusters using Crossplane. This approach provides a high-level, declarative API for creating production-ready OKE clusters with all necessary networking infrastructure.

## ArgoCD Sync Waves

The resources in this directory are deployed in a specific order using ArgoCD sync-wave annotations:

1. **Wave 3**: `xrd-okecluster.yaml` - CompositeResourceDefinition (defines the API)
2. **Wave 4**: `composition-okecluster.yaml` - Composition (implements the API)  
3. **Wave 5**: `claim-okecluster.yaml` - Example claim (namespace-scoped resource)

This ensures proper ordering when deployed via ArgoCD, building on:
- Wave 1: Crossplane core
- Wave 2: OCI Provider

## Overview

This XRD creates a complete OKE environment including:

### Networking Infrastructure
- **VCN (Virtual Cloud Network)** - Isolated network for the cluster
- **Internet Gateway** - For public internet access
- **NAT Gateway** - For private subnet outbound access
- **Route Tables** - Public and private routing
- **Security Lists** - Network security rules for OKE
- **Subnets** - Separate subnets for cluster endpoint, worker nodes, and services

### Kubernetes Infrastructure
- **OKE Cluster** - Enhanced cluster with private endpoint
- **Node Pool** - Configurable worker nodes with flexible compute shapes

## Files

- `claim-okecluster.yaml` - Example claim showing how to use the XRD
- `README.md` - This documentation

The XRD and Composition are located in the `oke-platform` directory.

## Prerequisites

1. **Crossplane** installed with OCI Provider configured
2. **OCI Provider Config** set up with proper credentials
3. **Compartment OCID** where resources will be created
4. **OKE Node Image OCID** for the target region
5. **SSH Key Secret** containing public key for node access

## Quick Start

### 1. Deploy the XRD and Composition

```bash
# Apply the XRD and Composition
kubectl apply -f xrd-okecluster.yaml
kubectl apply -f composition-okecluster.yaml

# Verify they're installed
kubectl get xrd
kubectl get compositions
```

### 2. Prepare SSH Key

```bash
# Generate SSH key pair (if you don't have one)
ssh-keygen -t rsa -b 4096 -f ~/.ssh/oke-nodes -N ""

# Copy your public key - you'll need to paste this in the claim
cat ~/.ssh/oke-nodes.pub
```

**Note**: Due to OCI NodePool API limitations, the SSH public key must be provided directly in the claim as plain text, not via a Kubernetes secret.

### 3. Get Required Information

```bash
# Get OKE-compatible node image OCID
oci ce node-pool-options get --node-pool-option-id all

# Get availability domains for your region
oci iam availability-domain list
```

### 4. Create Your Cluster

1. Copy and modify `claim-okecluster.yaml`:
   - Update `compartmentId` with your compartment OCID
   - Update `nodeImageId` with a valid OKE node image OCID
   - Update `availabilityDomain` for your region
   - Update `sshPublicKey` with your actual SSH public key (from step 2)
   - Adjust network CIDRs if needed

2. Apply the claim:
```bash
kubectl apply -f claim-okecluster.yaml
```

### 5. Monitor Deployment

```bash
# Check the claim status
kubectl get okeclusterclaim production-oke-cluster

# Check individual resources being created
kubectl get vcn,internetgateway,natgateway,routetable,securitylist,subnet,cluster,nodepool

# Get detailed status
kubectl describe okeclusterclaim production-oke-cluster
```

### 6. Access Your Cluster

Once deployment is complete, get the kubeconfig command from the claim status:

```bash
# Get the kubeconfig command
kubectl get okeclusterclaim production-oke-cluster -o jsonpath='{.status.kubeconfigCommand}'

# Run the command (example output):
oci ce cluster create-kubeconfig --cluster-id ocid1.cluster.oc1... --file $HOME/.kube/config --region us-ashburn-1 --token-version 2.0.0 --kube-endpoint PRIVATE_ENDPOINT
```

## Configuration Options

### Network Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `network.name` | Name prefix for network resources | Required |
| `network.vcnCidr` | VCN CIDR block | `10.0.0.0/16` |
| `network.clusterSubnetCidr` | Cluster endpoint subnet | `10.0.1.0/24` |
| `network.nodeSubnetCidr` | Worker node subnet | `10.0.2.0/24` |
| `network.serviceSubnetCidr` | Load balancer subnet | `10.0.3.0/24` |

### Cluster Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `cluster.name` | OKE cluster name | Required |
| `cluster.kubernetesVersion` | Kubernetes version | `v1.33.1` |
| `cluster.type` | Cluster type | `ENHANCED_CLUSTER` |
| `cluster.isPublicEndpoint` | Public endpoint | `false` |
| `cluster.podsCidr` | Pod networking CIDR | `10.244.0.0/16` |
| `cluster.servicesCidr` | Service networking CIDR | `10.96.0.0/16` |

### Node Pool Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `nodePool.name` | Node pool name | Required |
| `nodePool.nodeShape` | Compute shape | `VM.Standard.E4.Flex` |
| `nodePool.nodeCount` | Number of nodes | `3` |
| `nodePool.memoryInGbs` | Memory per node | `16` |
| `nodePool.ocpus` | OCPUs per node | `2` |
| `nodePool.availabilityDomain` | Availability domain | Required |
| `nodePool.nodeImageId` | OKE node image OCID | Required |
| `nodePool.sshPublicKey` | SSH public key | Required |

**IMPORTANT SSH Key Note**: The OCI NodePool API requires the SSH public key to be provided as a plain text string in the claim, not as a secret reference. Replace the placeholder in the claim with your actual SSH public key.

## ArgoCD Integration

This XRD can be deployed via ArgoCD using the provided application definition:

```bash
kubectl apply -f ../argocd/applications/oke-xrd.yaml
```

The ArgoCD application will:
- Deploy the XRD and Composition to `crossplane-system` namespace
- Handle sync waves to ensure proper ordering
- Ignore status changes on Crossplane resources

## Security Considerations

### Network Security
- **Private Cluster Endpoint**: Cluster API is only accessible from within the VCN
- **Private Worker Nodes**: Worker nodes have no direct internet access
- **NAT Gateway**: Provides secure outbound internet access for nodes
- **Security Lists**: Restrict traffic to necessary ports only

### Recommended Improvements
1. **Restrict SSH Access**: Update security list to allow SSH only from specific CIDRs
2. **Use Bastion Host**: Deploy a bastion host for secure cluster access
3. **Network Security Groups**: Consider using NSGs for more granular security rules
4. **Secrets Management**: Use OCI Vault or Kubernetes secrets for sensitive data

## Troubleshooting

### Common Issues

1. **Image OCID Invalid**: Ensure you're using a valid OKE node image for your region
2. **Availability Domain**: Verify the AD name format for your region
3. **Network CIDR Conflicts**: Ensure subnet CIDRs don't overlap and fit within VCN CIDR
4. **Resource Limits**: Check OCI service limits for your tenancy

### Useful Commands

```bash
# Check Crossplane provider logs
kubectl logs -n crossplane-system -l pkg.crossplane.io/provider=oci-provider

# Get detailed resource status
kubectl describe <resource-type> <resource-name>

# Check events
kubectl get events --sort-by=.metadata.creationTimestamp
```

## Cleanup

To delete the cluster and all associated resources:

```bash
kubectl delete okeclusterclaim production-oke-cluster
```

This will trigger Crossplane to delete all managed resources in the proper order.

## Contributing

This XRD can be extended to support additional features:
- Multiple node pools with different configurations
- Auto-scaling configurations
- Additional networking options (service mesh, etc.)
- Integration with OCI monitoring and logging services
- Support for GPU node pools
- Cluster add-ons management