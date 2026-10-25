---
page_title: "Test mode and live mode"
subcategory: "Guides"
description: |-
  Manage test mode and live mode objects with separate configurations or provider aliases.
---

# Test mode and live mode

Each Dodo Payments API key works in one mode. A test mode key cannot change live mode objects.
Make sure that the `environment` of each provider configuration agrees with its API key.

## Use a separate workspace or directory for each mode

Use this method when you keep the two modes apart. Each state then has the objects of one mode only.

```hcl
variable "environment" {
  type = string
  validation {
    condition     = contains(["test_mode", "live_mode"], var.environment)
    error_message = "Set environment to test_mode or live_mode."
  }
}

provider "dodopayments" {
  environment = var.environment
}
```

Set `DODO_PAYMENTS_API_KEY` to the key of the same mode before you run Terraform.

## Use provider aliases in one configuration

Use this method when one configuration must manage objects in the two modes.
Give each alias its own key. Use variables that you mark as sensitive.

```hcl
variable "test_api_key" {
  type      = string
  sensitive = true
}

variable "live_api_key" {
  type      = string
  sensitive = true
}

provider "dodopayments" {
  alias       = "test"
  environment = "test_mode"
  api_key     = var.test_api_key
}

provider "dodopayments" {
  alias       = "live"
  environment = "live_mode"
  api_key     = var.live_api_key
}

resource "dodopayments_product" "test" {
  provider = dodopayments.test
  # ...
}
```

## Copy the test mode configuration to live mode

The IDs of objects are different in each mode. Do not write IDs in the configuration. Refer to resources, for example `dodopayments_product.example.id`.
The `DODO_PAYMENTS_ENVIRONMENT` environment variable sets the environment when the configuration does not set it.
