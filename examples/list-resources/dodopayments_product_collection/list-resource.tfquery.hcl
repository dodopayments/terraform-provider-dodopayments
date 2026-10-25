# Run "terraform query -generate-config-out=generated.tf" to write a resource block and an import block for each object.
list "dodopayments_product_collection" "all" {
  provider         = dodopayments
  include_resource = true
}
