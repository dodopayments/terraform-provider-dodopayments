data "dodopayments_product_collections" "example" {
  max_items = 10
}

output "product_collections" {
  value = data.dodopayments_product_collections.example.items
}
