# Sample OKE Environment

## Introduction
This is a sample OKE environment that is deployed using Crossplane and an alpha version of the OCI Crossplane provider.

### Topology

![OCI topology diagram](./docs/topology.png "OCI Topology Diagram")

All resources are deployed into an existing compartment (or at the root compartment level), which OCID is provided in the claim.

Three subnets are created (with the CIDR for each being specified by the user in the claim):

| Name | Type | Description |
|------|------|-------------|
| `oke` | Private | The OKE cluster and node pools are placed in this subnet. |
| `lb` | Private | This is where LBs are intended to go. |
| `bastion` | Public | A "jump box" (aka bastion) instance is placed in this Subnet. |

### Security Policies
An empty Security List is used with the different Subnets (LB, OKE and bastion).  Network Security Groups (NSGs) are created, permitting different traffic flows within the environmnet.

Look at the `composition-okenetwork.yml` to see the different security rules.  Below is a summary:

| Src NSG | Dst | Direction | Protocol | Description |
|---------|-----|-----------|----------|-------------|
| `oke` | `oke` NSG | EGRESS | all | Permit intra-Subnet traffic (outbound). |
| `oke` | `oke` NSG | INGRESS | all | Permit intra-Subnet traffic (inbound). |
| `oke` | `bastion` NSG | INGRESS | all | Permit traffic from bastion to OKE compute instances. |
| `oke` | `lb` NSG | INGRESS | TCP | Permit all TCP traffic from the LBs. |
| `oke` | `0.0.0.0/0` | EGRESS | TCP/443 | Permit outbound HTTPS traffic from the OKE compute instances. |
| `bastion` | `oke` NSG | EGRESS | all | Permit all traffic to the OKE compute instances from the bastion. |
| `bastion` | `0.0.0.0/0` | EGRESS | TCP | Permit all outbound TCP traffic from the bastion. |
| `bastion` | `lb` NSG | EGRESS | TCP | Permit all TCP traffic to the LBs. |
| `bastion` | `permittedAccessCidr` value | INGRESS | TCP/22 | Permit SSH traffic from a user-provided CIDR. |
| `lb` | `bastion` NSG | INGRESS | TCP | Permit inbound TCP traffic from the bastion. |
| `lb` | `oke` NSG | EGRESS | TCP | Permit all outbound TCP traffic to the OKE compute instances. |

### Routing Policies
There are two Route Tables (RTs): `public` and `private`.  As the names imply, one is used for Subnets that permit public IPs, the other for those using only private IP addresses.  An Internet Gateway is used for the `public` RT and a NAT Gateway is used for the `private` RT.

## Getting Started
It's assumed that the user is using Rancher Desktop, Docker Desktop or some other locally running K8s implementation.  Since this is using an alpha (or better yet, pre-alpha) OCI Crossplane provider, it's a little rough, but works.

Start by installing Crossplane:

```
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```

Clone the Crossplane repo:
```
git clone git@github.com:crossplane/crossplane.git
```

Next install the different CRDs used by Crossplane:

```
kubectl apply -f crossplane/cluster/crds
```

Install role binding for `crossplane-system` account (this isn't the right way to do this, but for a demo, it works):

```
kubectl apply -f crossplane-cluster-role-binding.yml
```

> NOTE: The above file also maps for the K8s dashboard user.

Copy `examples/providerconfig/secret.yaml.tmpl` to `secret.yaml` and modify its contents for your environment.

Install the OCI Crossplane provider:

```
kubectl apply -f examples/providerconfig/crossplane-namespace.yaml
kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f package/crds
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Apply the XRD and Composition for the OkeNetwork:

```
kubectl apply -f xrd-okenetwork.yml
kubectl apply -f composition-okenetwork.yml
```

Now for the fun part.  Copy `claim-okenetwork.yml` to `claim.yml` and modify it for your environment (put proper values in it) and apply it, which will get the gears turning:

```
kubectl apply -f claim.yml
```

Lastly, get the OCI Crossplane provider running:

```
make run
```

### Interacting With It

`kubectl describe okenetwork oketest`

This will show several useful bits of info in the status, including the bastion public IP, the OKE cluster OCID (along with other potentially needed OCIDs), etc.

This topology is designed to tunnel (port forward) traffic via the bastion to/from the OKE cluster (which only has a private IP), as well as the LB (as services are deployed).

Here's an example of how to setup an SSH port forwarding session for accessing the OKE K8s API:

```
ssh -L 6443:<OKE_cluster_private_IP>:6443 opc@<bastion_pub_ip>
```

Here's an example of getting to a LB-based app (service):

```
ssh -L 8081:<LB_private_IP>:8081 opc@<bastion_public_ip>
```

To access the OKE cluster, you'll need to follow the directions (run an OCI CLI command) given in the OKE cluster's Quick Start screen.  This will setup a new config context in `~/.kube/config` that you can switch between.  Use `oci ce cluster create-kubeconfig --cluster-id <oke_cluster_ocid> --file $HOME/.kube/config --region <region> --token-version 2.0.0  --kube-endpoint PRIVATE_ENDPOINT`

When interacting with the local, K8s, `kubectl config use-context rancher-desktop` (or whatever is the proper context name for your environment), then to interact with the OKE cluster, change to its context.

## Notes/Issues
This is a really basic/rough example.  It's not meant for production use, but is a sample.

## Contributing
This project is open source.  Please submit your contributions by forking this repository and submitting a pull request!  Oracle appreciates any contributions that are made by the open source community.

## License
Copyright (c) 2022 Oracle and/or its affiliates.

Licensed under the Universal Permissive License (UPL), Version 1.0.

See [LICENSE](LICENSE) for more details.

ORACLE AND ITS AFFILIATES DO NOT PROVIDE ANY WARRANTY WHATSOEVER, EXPRESS OR IMPLIED, FOR ANY SOFTWARE, MATERIAL OR CONTENT OF ANY KIND CONTAINED OR PRODUCED WITHIN THIS REPOSITORY, AND IN PARTICULAR SPECIFICALLY DISCLAIM ANY AND ALL IMPLIED WARRANTIES OF TITLE, NON-INFRINGEMENT, MERCHANTABILITY, AND FITNESS FOR A PARTICULAR PURPOSE.  FURTHERMORE, ORACLE AND ITS AFFILIATES DO NOT REPRESENT THAT ANY CUSTOMARY SECURITY REVIEW HAS BEEN PERFORMED WITH RESPECT TO ANY SOFTWARE, MATERIAL OR CONTENT CONTAINED OR PRODUCED WITHIN THIS REPOSITORY. IN ADDITION, AND WITHOUT LIMITING THE FOREGOING, THIRD PARTIES MAY HAVE POSTED SOFTWARE, MATERIAL OR CONTENT TO THIS REPOSITORY WITHOUT ANY REVIEW. USE AT YOUR OWN RISK. 