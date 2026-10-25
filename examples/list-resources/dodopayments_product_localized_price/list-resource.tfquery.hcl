# Run "terraform query -generate-config-out=generated.tf" to write a resource block and an import block for each object.
list "dodopayments_product_localized_price" "all" {
  provider         = dodopayments
  include_resource = true

  config {
    product_id = "<product_id>"
  }
}
