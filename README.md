# Terraform provider for Istio

- This is a [Terraform](https://www.terraform.io/) provider for [Istio](https://istio.io/).

- This provider is a work in progress, and is not yet ready for use.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)
- [Istio](https://istio.io/docs/setup/getting-started/) 1.6.x
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) 1.18.x

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/moabukar/terraform-provider-istio`

```sh
make build
```

## Development Environment Setup

## Requirements

- Terraform `>= 0.14`
   * We recommend you to use Terraform version manager [tfutils/tfenv](https://github.com/tfutils/tfenv)
      * `tfenv install 0.15.3`, `tfenv use 0.15.3`
- Go `1.18` (to build the provider plugin)
   * We recommend you to use Go version manager [syndbg/goenv](https://github.com/syndbg/goenv/blob/master/INSTALL.md)
      * `goenv install 1.18.0`

## Quick Start

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `bin` directory.

```shell
$ make deps
$ make build
```

## Testing the Provider

The easiest way to run acceptance tests is to use the built in `make` step `testacc`:

*Note:* acceptance tests won't create the actual cloud resources since they will be run against a mock server.

```shell
$ make testacc
```

## Using the Provider

With Terraform v0.14 and later, [development overrides for provider developers](https://www.terraform.io/docs/cli/config/config-file.html#development-overrides-for-provider-developers) can be leveraged in order to use the provider built from source.

To do this, populate a Terraform CLI configuration file (`~/.terraformrc` for all platforms other than Windows; `terraform.rc` in the `%APPDATA%` directory when using Windows):

```hcl
provider_installation {

  dev_overrides {
      "registry.terraform.io/moabukar/istio" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.

  direct {}
}
```

