data "dodopayments_entitlements" "example" {
  max_items = 10
}

output "entitlements" {
  value = data.dodopayments_entitlements.example.items
}
