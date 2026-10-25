resource "dodopayments_entitlement" "example" {
  integration_config = {
    feature_id   = "example-entitlement"
    feature_type = "boolean"
  }
  integration_type = "feature_flag"
  metadata = {
    tier = "pro"
  }
  name = "example-entitlement"
}
