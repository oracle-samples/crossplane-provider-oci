# Provider Oci

`provider-oci` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Oci API.

## Getting Started

### Testing the provider 

Follow the below steps for generating the code and test.
Make sure go is installed and path is set for the go binaries
```
Git clone the repo 
cd provider-oci/
make generate # This will generate the binaries and CRDs.
kubectl apply -f package/crds  #  A k8s cluster must be available running locally
make run # This will start crossplane provider against the local kubernetes cluster .
```
In another terminal run the below commands to test the provider. 
Make sure to setup the OCI credentials in examples/providerconfig/secret.yaml using the template examples/providerconfig/secret.yaml.tmpl
```

# Create "crossplane-system" namespace if not exists
kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -

kubectl apply -f examples/providerconfig/
kubectl apply -f examples/identity/compartment.yaml  # open the compartment.yaml and update the compartment id in wich the new compartment to be created
kubectl get managed # To check the status
```
## Building the provider and package as docker image
```
make build # This will build docker image in local machie. This can be pushed to any container registry

```

## Installing the provider from a container registry 
using declarative api
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-oci
spec:
  package: <registry>/provider-oci:<tag>
EOF
```

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/oracle-samples/provider-oci/issues).


## Contributing

This project is not accepting external contributions at this time. For bugs or enhancement requests, please file a GitHub issue unless it’s security related. When filing a bug remember that the better written the bug is, the more likely it is to be fixed. If you think you’ve found a security vulnerability, do not raise a GitHub issue and follow the instructions in our [security policy](./SECURITY.md).

## Code of Conduct

provider-jet-oci adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Security

Please consult the [security guide](./SECURITY.md) for our responsible security vulnerability disclosure process

## License

provider-oci is under the Apache 2.0 license.
