resource "dodopayments_brand" "example" {
  description   = "Managed by Terraform"
  name          = "example-brand"
  support_email = "support@example.com"
  url           = "https://example.com/example-brand"
}

data "dodopayments_brand" "example" {
  id = dodopayments_brand.example.id
}
