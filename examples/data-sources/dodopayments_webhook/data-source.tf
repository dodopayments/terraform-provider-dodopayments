resource "dodopayments_webhook" "example" {
  url = "https://example.com/example-webhook"
}

data "dodopayments_webhook" "example" {
  id = dodopayments_webhook.example.id
}
