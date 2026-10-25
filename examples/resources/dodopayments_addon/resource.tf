resource "dodopayments_addon" "example" {
  currency     = "AED"
  name         = "example-addon"
  price        = 100
  tax_category = "digital_products"
}
