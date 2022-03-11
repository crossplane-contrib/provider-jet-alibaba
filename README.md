# Terrajet Alicloud Provider

`provider-jet-alibaba` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the 
Alibaba Cloud API.

## Getting Started

- Install Alibaba Cloud provider

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/crossplane-contrib/provider-jet-alibaba/releases):

```shell
$ kubectl crossplane install provider crossplane/provider-jet-alibaba:v0.1.0
```

- Authenticate Alibaba Cloud

Rename [secret.yaml.tmpl](./examples/providerconfig/secret.yaml.tmpl) to `secret.yaml`. Update values for `accessKeyID`
and `accessKeySecret` in [secret.yaml.tmpl](./examples/providerconfig/secret.yaml.tmpl) with your Alibaba Cloud credentials and
apply it.

```shell
$ kubectl apply -f examples/providerconfig/secret.yaml
```

- Provision Cloud resources

Let's take Alibaba Cloud VPC as an example.

```shell
$ kubectl apply -f examples/sample/vpc.yaml
......
vpc-example   True    True     vpc-2zexzf5ocsgm0vc4jriyr   48s
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-jet-alibaba) to set cloud resources.

- Destroy Cloud resources

```shell
$ kubectl delete vpc.vpc.alicloud.jet.crossplane.io vpc-example
vpc.vpc.alicloud.jet.crossplane.io "vpc-example" deleted
```

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-jet-alibaba/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-jet-alibaba is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-alibaba adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-alibaba is under the Apache 2.0 license.
