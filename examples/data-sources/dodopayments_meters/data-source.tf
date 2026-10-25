data "dodopayments_meters" "example" {
  max_items = 10
}

output "meters" {
  value = data.dodopayments_meters.example.items
}
