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
