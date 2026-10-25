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

resource "dodopayments_product_collection" "example" {
  groups = [{
    products = [{
      product_id = dodopayments_product.example.id
    }]
  }]
  name = "example-product-collection"
}
