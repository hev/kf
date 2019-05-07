# kf

`kf` is like `cf` for Knative. It combines the best of
the best of the `cf` CLI, the OSBAPI spec, [service-catalog](svcat)
and Knative into a unified Platform as a Service Experience.

## Getting started.
Build a `kf` binary and follow our [install instructions](./docs/install.md) 
for Knative. 

## How to build

**Dependencies:**

[go mod](https://github.com/golang/go/wiki/Modules#quick-start) 
is used and required for dependencies

**Requirements:**

  - Golang `1.12`

**Building:**

```sh
$ go build ./cmd/...
```

**Notes:**

- `kf` CLI must be built outside of the `$GOPATH` folder unless 
you explicitly use `export GO111MODULE=on`.

## Development and releasing

We use [ko](https://github.com/google/ko) for rapid development 
and during the release process to build a full set of `kf` images 
and installation YAML.