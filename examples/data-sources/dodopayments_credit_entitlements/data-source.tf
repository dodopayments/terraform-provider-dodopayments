data "dodopayments_credit_entitlements" "example" {
  max_items = 10
}

output "credit_entitlements" {
  value = data.dodopayments_credit_entitlements.example.items
}
