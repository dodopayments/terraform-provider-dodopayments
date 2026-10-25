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

## Resources and data sources

- resource `dodopayments_discount`, data source `dodopayments_discount`, data source `dodopayments_discounts`
- resource `dodopayments_addon`, data source `dodopayments_addon`, data source `dodopayments_addons`
- resource `dodopayments_webhook`, data source `dodopayments_webhook`, data source `dodopayments_webhooks`
- resource `dodopayments_meter`, data source `dodopayments_meter`, data source `dodopayments_meters`

## Import

You can import each resource by its ID. Use an `import` block or the `terraform import` command:

```hcl
import {
  to = dodopayments_discount.example
  id = "<id>"
}
```

## Development

```sh
go build ./...
TF_ACC=1 DODO_PAYMENTS_API_KEY=<test mode key> DODO_PAYMENTS_ENVIRONMENT=test_mode go test ./... -v
```
