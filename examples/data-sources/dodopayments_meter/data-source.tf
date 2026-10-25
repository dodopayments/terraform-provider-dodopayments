resource "dodopayments_meter" "example" {
  aggregation = {
    type = "count"
  }
  event_name       = "example-meter"
  measurement_unit = "example-meter"
  name             = "example-meter"
}

data "dodopayments_meter" "example" {
  id = dodopayments_meter.example.id
}
