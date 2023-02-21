# Terrajet Oci Provider

`provider-jet-oci` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the
Oci API.

## Getting Started
### Install the provider

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/crossplane-contrib/provider-jet-oci/releases):
```
kubectl crossplane install provider crossplane/provider-jet-oci:v0.1.0
```

Alternatively, you can use declarative installation:
```
kubectl apply -f examples/install.yaml
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-jet-oci).

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
open an [issue](https://github.com/crossplane-contrib/provider-jet-oci/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-jet-oci is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-oci adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-oci is under the Apache 2.0 license.
