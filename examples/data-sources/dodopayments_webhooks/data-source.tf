data "dodopayments_webhooks" "example" {
  max_items = 10
}

output "webhooks" {
  value = data.dodopayments_webhooks.example.items
}
