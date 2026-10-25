# Contributing

The generator makes most files in this repository from the OpenAPI spec. Do not edit a generated file by hand. The next generation run replaces your change.

## Build

```sh
./scripts/bootstrap
./scripts/build
./scripts/lint
```

## Test against the API

The acceptance tests create real objects. Use an API key for a test environment only. The `Development` section of `README.md` gives the command and the environment variables.

The tests run only when `TF_ACC=1` is set. Without `TF_ACC`, `./scripts/test` skips them.

## Use a local build

Add a `dev_overrides` block to `~/.terraformrc`. Then run `go build -o <dir>/<provider binary name> .`.

```hcl
provider_installation {
  dev_overrides {
    "<namespace>/<provider name>" = "<dir>"
  }
  direct {}
}
```
