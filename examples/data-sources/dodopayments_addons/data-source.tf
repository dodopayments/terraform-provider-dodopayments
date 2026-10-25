data "dodopayments_addons" "example" {
  max_items = 10
}

output "addons" {
  value = data.dodopayments_addons.example.items
}
