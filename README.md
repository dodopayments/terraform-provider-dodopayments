# Dodo Payments Terraform provider

The Dodo Payments provider manages Dodo Payments objects with Terraform or OpenTofu 1.12 or later.

## Usage

```hcl
terraform {
  required_providers {
    dodopayments = {
      source = "dodopayments/dodopayments"
    }
  }
}

provider "dodopayments" {
  # Set the DODO_PAYMENTS_API_KEY environment variable, or set api_key here.
  environment = "test_mode"
}
```

## Workload identity

The provider can use workload identity federation instead of an API key.
It exchanges an OIDC token from the CI system or the cloud platform for a short-lived access token.
Create a workload identity provider in the dashboard first.

```hcl
provider "dodopayments" {
  environment = "test_mode"
  workload_identity {
    identity_provider_id = "wip_..."
    token_source         = "github_actions"
  }
}
```

`token_source` is `auto` (the default), `github_actions`, `gitlab`, `terraform_cloud`, `gcp`, `vercel`, `kubernetes` or `file`.
Do not set `api_key` together with `workload_identity`.
If you set neither, `DODO_PAYMENTS_API_KEY` is not set and `DODO_PAYMENTS_IDENTITY_PROVIDER_ID` is set, the provider uses workload identity with the `auto` token source.
Then a GitHub Actions job needs no change to the provider configuration:

```yaml
permissions:
  contents: read
  id-token: write
jobs:
  apply:
    runs-on: ubuntu-latest
    env:
      DODO_PAYMENTS_IDENTITY_PROVIDER_ID: wip_...
    steps:
      - uses: actions/checkout@11d5960a326750d5838078e36cf38b85af677262 # v4
      - uses: hashicorp/setup-terraform@b9cd54a3c349d3f38e8881555d616ced269862dd # v3
      - run: terraform init && terraform apply -auto-approve
```

[`examples/complete`](examples/complete) has one configuration that uses each resource. The registry documentation has guides for test mode and live mode, for import, and for destroy behavior.

## Resources, data sources, ephemeral resources and actions

### Products

- resource `dodopayments_product`, data source `dodopayments_product`, data source `dodopayments_products`, list resource `dodopayments_product`, action `dodopayments_product_unarchive`
- resource `dodopayments_product_collection`, data source `dodopayments_product_collection`, data source `dodopayments_product_collections`, list resource `dodopayments_product_collection`, action `dodopayments_product_collection_unarchive`

### Pricing

- resource `dodopayments_product_localized_price`, data source `dodopayments_product_localized_price`, data source `dodopayments_product_localized_prices`, list resource `dodopayments_product_localized_price`

### Discounts

- resource `dodopayments_discount`, data source `dodopayments_discount`, data source `dodopayments_discounts`, list resource `dodopayments_discount`

### Subscriptions

- resource `dodopayments_addon`, data source `dodopayments_addon`, data source `dodopayments_addons`, list resource `dodopayments_addon`, action `dodopayments_addon_unarchive`

### Account Management

- resource `dodopayments_brand`, data source `dodopayments_brand`, data source `dodopayments_brands`, list resource `dodopayments_brand`

### Webhooks

- resource `dodopayments_webhook`, data source `dodopayments_webhook`, data source `dodopayments_webhooks`, list resource `dodopayments_webhook`, ephemeral resource `dodopayments_webhook_secret`

### Usage Based Billing

- resource `dodopayments_meter`, data source `dodopayments_meter`, data source `dodopayments_meters`, list resource `dodopayments_meter`, action `dodopayments_meter_unarchive`

### Credit-Based Billing

- resource `dodopayments_credit_entitlement`, data source `dodopayments_credit_entitlement`, data source `dodopayments_credit_entitlements`, list resource `dodopayments_credit_entitlement`, action `dodopayments_credit_entitlement_undelete`

### Entitlements

- resource `dodopayments_entitlement`, data source `dodopayments_entitlement`, data source `dodopayments_entitlements`, list resource `dodopayments_entitlement`

## Import

You can import each resource with an `import` block, by its ID or by its resource identity:

```hcl
import {
  to = dodopayments_product.example
  identity = {
    id = "<id>"
  }
}
```

With Terraform 1.14 or later, `terraform query` finds the objects that you can import and writes the import blocks for you.

## Destroy behavior

| Resource                               | When Terraform destroys the resource                                                                                                              | Action that restores the object             |
| -------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| `dodopayments_product`                 | The provider archives the object.                                                                                                                 | `dodopayments_product_unarchive`            |
| `dodopayments_product_localized_price` | The provider archives the object.                                                                                                                 | None                                        |
| `dodopayments_discount`                | The provider deletes the object.                                                                                                                  | None                                        |
| `dodopayments_addon`                   | The provider deletes the object.                                                                                                                  | `dodopayments_addon_unarchive`              |
| `dodopayments_brand`                   | The API does not delete a brand. When Terraform destroys the resource, the provider archives the brand. The API cannot restore an archived brand. | None                                        |
| `dodopayments_webhook`                 | The provider deletes the object.                                                                                                                  | None                                        |
| `dodopayments_meter`                   | The provider archives the object.                                                                                                                 | `dodopayments_meter_unarchive`              |
| `dodopayments_credit_entitlement`      | The provider deletes the object.                                                                                                                  | `dodopayments_credit_entitlement_undelete`  |
| `dodopayments_entitlement`             | The provider deletes the object.                                                                                                                  | None                                        |
| `dodopayments_product_collection`      | The provider deletes the object.                                                                                                                  | `dodopayments_product_collection_unarchive` |

## Development

```sh
go build ./...
./scripts/format   # go mod tidy, gofmt, the registry docs and terraform fmt
TF_ACC=1 DODO_PAYMENTS_API_KEY=<test mode key> DODO_PAYMENTS_ENVIRONMENT=test_mode go test ./... -v
```
