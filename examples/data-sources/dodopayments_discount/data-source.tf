resource "dodopayments_discount" "example" {
  amount = 100
  name   = "example-discount"
  type   = "percentage"
}

data "dodopayments_discount" "example" {
  id = dodopayments_discount.example.id
}
