resource "dodopayments_addon" "example" {
  currency     = "USD"
  name         = "example-addon"
  price        = 100
  tax_category = "digital_products"
}

data "dodopayments_addon" "example" {
  id = dodopayments_addon.example.id
}
