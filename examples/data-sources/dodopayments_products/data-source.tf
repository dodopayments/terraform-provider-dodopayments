data "dodopayments_products" "example" {
  max_items = 10
}

output "products" {
  value = data.dodopayments_products.example.items
}
