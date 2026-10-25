resource "dodopayments_meter" "example" {
  aggregation = {
    type = "count"
  }
  event_name       = "example-meter"
  measurement_unit = "example-meter"
  name             = "example-meter"
}
