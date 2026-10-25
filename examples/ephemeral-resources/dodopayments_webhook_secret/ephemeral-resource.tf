resource "dodopayments_webhook" "example" {
  url = "https://example.com/example-webhook"
}

ephemeral "dodopayments_webhook_secret" "example" {
  webhook_id = dodopayments_webhook.example.id
}

# Use the value in a write-only argument, a provider block or another ephemeral context.
# For example: ephemeral.dodopayments_webhook_secret.example.secret
