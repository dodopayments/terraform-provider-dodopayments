terraform {
  required_version = ">= 1.12"
  required_providers {
    dodopayments = {
      source = "dodopayments/dodopayments"
    }
  }
}

# The provider reads the API key from the DODO_PAYMENTS_API_KEY environment variable.
provider "dodopayments" {
  environment = "test_mode"
}

resource "dodopayments_discount" "example" {
  amount = 100
  name   = "example-discount"
  type   = "percentage"
}

output "discount_id" {
  value = dodopayments_discount.example.id
}

data "dodopayments_discount" "example" {
  id = dodopayments_discount.example.id
}

data "dodopayments_discounts" "all" {
  max_items  = 10
  depends_on = [dodopayments_discount.example]
}

output "discounts_count" {
  value = length(data.dodopayments_discounts.all.items)
}

resource "dodopayments_addon" "example" {
  currency     = "AED"
  name         = "example-addon"
  price        = 100
  tax_category = "digital_products"
}

output "addon_id" {
  value = dodopayments_addon.example.id
}

data "dodopayments_addon" "example" {
  id = dodopayments_addon.example.id
}

data "dodopayments_addons" "all" {
  max_items  = 10
  depends_on = [dodopayments_addon.example]
}

output "addons_count" {
  value = length(data.dodopayments_addons.all.items)
}

resource "dodopayments_webhook" "example" {
  url = "https://example.com/example-webhook"
}

output "webhook_id" {
  value = dodopayments_webhook.example.id
}

data "dodopayments_webhook" "example" {
  id = dodopayments_webhook.example.id
}

data "dodopayments_webhooks" "all" {
  max_items  = 10
  depends_on = [dodopayments_webhook.example]
}

output "webhooks_count" {
  value = length(data.dodopayments_webhooks.all.items)
}

resource "dodopayments_meter" "example" {
  aggregation = {
    type = "count"
  }
  event_name       = "example-meter"
  measurement_unit = "example-meter"
  name             = "example-meter"
}

output "meter_id" {
  value = dodopayments_meter.example.id
}

data "dodopayments_meter" "example" {
  id = dodopayments_meter.example.id
}

data "dodopayments_meters" "all" {
  max_items  = 10
  depends_on = [dodopayments_meter.example]
}

output "meters_count" {
  value = length(data.dodopayments_meters.all.items)
}
