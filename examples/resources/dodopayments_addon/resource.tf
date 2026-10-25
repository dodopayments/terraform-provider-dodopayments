resource "dodopayments_addon" "example" {
  currency     = "USD"
  name         = "example-addon"
  price        = 100
  tax_category = "digital_products"
}
