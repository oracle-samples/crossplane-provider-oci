# Terrajet Oci Provider

`provider-jet-oci` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the
Oci API.

## Installing the provider

Install the provider by using the following steps and commands

Build the provider and push to docker registry
```
make all
```

After pushing to the docker registry run the below command to install the provider
`repoName` is the docker registry repo to which the image is pushed.
```
kubectl crossplane install provider <repoName> /provider-jet-oci:v0.1.0
```

Alternatively, you can use declarative installation:
```
kubectl apply -f examples/install.yaml
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.


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

## Contributing

This project is not accepting external contributions at this time. For bugs or enhancement requests, please file a GitHub issue unless it’s security related. When filing a bug remember that the better written the bug is, the more likely it is to be fixed. If you think you’ve found a security vulnerability, do not raise a GitHub issue and follow the instructions in our [security policy](./SECURITY.md).

## Code of Conduct

provider-jet-oci adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Security

Please consult the [security guide](./SECURITY.md) for our responsible security vulnerability disclosure process

## License

provider-jet-oci is under the Apache 2.0 license.
