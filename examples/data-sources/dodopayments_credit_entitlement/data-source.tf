resource "dodopayments_credit_entitlement" "example" {
  currency         = "USD"
  name             = "example-credit-entitlement"
  overage_enabled  = false
  precision        = 2
  price_per_unit   = "10.5"
  rollover_enabled = false
  unit             = "credits"
}

data "dodopayments_credit_entitlement" "example" {
  id = dodopayments_credit_entitlement.example.id
}
