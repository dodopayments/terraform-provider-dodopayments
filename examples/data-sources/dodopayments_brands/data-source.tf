data "dodopayments_brands" "example" {
  max_items = 10
}

output "brands" {
  value = data.dodopayments_brands.example.items
}
