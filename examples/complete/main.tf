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

resource "dodopayments_addon" "example" {
  currency     = "USD"
  name         = "example-addon"
  price        = 100
  tax_category = "digital_products"
}

resource "dodopayments_credit_entitlement" "example" {
  currency         = "USD"
  name             = "example-credit-entitlement"
  overage_enabled  = false
  precision        = 2
  price_per_unit   = "10.5"
  rollover_enabled = false
  unit             = "credits"
}

resource "dodopayments_entitlement" "example" {
  integration_config = {
    feature_id   = "example-entitlement"
    feature_type = "boolean"
  }
  integration_type = "feature_flag"
  metadata = {
    tier = "pro"
  }
  name = "example-entitlement"
}

resource "dodopayments_product" "example" {
  addons = [dodopayments_addon.example.id]
  credit_entitlements = [{
    credit_entitlement_id = dodopayments_credit_entitlement.example.id
    credits_amount        = "100.0"
  }]
  entitlements = [{
    entitlement_id = dodopayments_entitlement.example.id
  }]
  metadata = {
    plan = "pro"
  }
  name = "example-product"
  price = {
    type                         = "recurring_price"
    currency                     = "USD"
    price                        = 1000
    payment_frequency_count      = 1
    payment_frequency_interval   = "Month"
    subscription_period_count    = 1
    subscription_period_interval = "Year"
  }
  pricing_mode = "by_currency"
  tax_category = "saas"
}

resource "dodopayments_product_localized_price" "example" {
  amount     = 900
  currency   = "EUR"
  product_id = dodopayments_product.example.id
}

resource "dodopayments_discount" "example" {
  amount = 100
  name   = "example-discount"
  type   = "percentage"
}

resource "dodopayments_brand" "example" {
  description   = "Managed by Terraform"
  name          = "example-brand"
  support_email = "support@example.com"
  url           = "https://example.com/example-brand"
}

resource "dodopayments_webhook" "example" {
  url = "https://example.com/example-webhook"
}

resource "dodopayments_meter" "example" {
  aggregation = {
    type = "count"
  }
  event_name       = "example-meter"
  measurement_unit = "example-meter"
  name             = "example-meter"
}

resource "dodopayments_product_collection" "example" {
  groups = [{
    products = [{
      product_id = dodopayments_product.example.id
    }]
  }]
  name = "example-product-collection"
}

output "product_id" {
  value = dodopayments_product.example.id
}

data "dodopayments_product" "example" {
  id = dodopayments_product.example.id
}

data "dodopayments_products" "all" {
  max_items  = 10
  depends_on = [dodopayments_product.example]
}

output "products_count" {
  value = length(data.dodopayments_products.all.items)
}

output "product_localized_price_id" {
  value = dodopayments_product_localized_price.example.id
}

data "dodopayments_product_localized_price" "example" {
  product_id = dodopayments_product_localized_price.example.product_id
  id         = dodopayments_product_localized_price.example.id
}

data "dodopayments_product_localized_prices" "all" {
  product_id = dodopayments_product_localized_price.example.product_id
  max_items  = 10
  depends_on = [dodopayments_product_localized_price.example]
}

output "product_localized_prices_count" {
  value = length(data.dodopayments_product_localized_prices.all.items)
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

output "brand_id" {
  value = dodopayments_brand.example.id
}

data "dodopayments_brand" "example" {
  id = dodopayments_brand.example.id
}

data "dodopayments_brands" "all" {
  max_items  = 10
  depends_on = [dodopayments_brand.example]
}

output "brands_count" {
  value = length(data.dodopayments_brands.all.items)
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

# Terraform reads the value for each run and does not store it. Use it in a write-only argument or a provider block.
ephemeral "dodopayments_webhook_secret" "example" {
  webhook_id = dodopayments_webhook.example.id
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

output "credit_entitlement_id" {
  value = dodopayments_credit_entitlement.example.id
}

data "dodopayments_credit_entitlement" "example" {
  id = dodopayments_credit_entitlement.example.id
}

data "dodopayments_credit_entitlements" "all" {
  max_items  = 10
  depends_on = [dodopayments_credit_entitlement.example]
}

output "credit_entitlements_count" {
  value = length(data.dodopayments_credit_entitlements.all.items)
}

output "entitlement_id" {
  value = dodopayments_entitlement.example.id
}

data "dodopayments_entitlement" "example" {
  id = dodopayments_entitlement.example.id
}

data "dodopayments_entitlements" "all" {
  max_items  = 10
  depends_on = [dodopayments_entitlement.example]
}

output "entitlements_count" {
  value = length(data.dodopayments_entitlements.all.items)
}

output "product_collection_id" {
  value = dodopayments_product_collection.example.id
}

data "dodopayments_product_collection" "example" {
  id = dodopayments_product_collection.example.id
}

data "dodopayments_product_collections" "all" {
  max_items  = 10
  depends_on = [dodopayments_product_collection.example]
}

output "product_collections_count" {
  value = length(data.dodopayments_product_collections.all.items)
}
