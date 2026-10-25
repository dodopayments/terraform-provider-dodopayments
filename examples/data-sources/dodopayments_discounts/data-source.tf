data "dodopayments_discounts" "example" {
  max_items = 10
}

output "discounts" {
  value = data.dodopayments_discounts.example.items
}
