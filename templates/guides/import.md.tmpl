---
page_title: "Import existing objects"
subcategory: "Guides"
description: |-
  Bring objects that you created in the dashboard under Terraform management.
---

# Import existing objects

## Find objects with terraform query

Terraform 1.14 and later can list the objects in your account and write the configuration for you.
Write a file with the `.tfquery.hcl` extension:

```hcl
list "dodopayments_product" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_discount" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_addon" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_brand" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_webhook" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_meter" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_credit_entitlement" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_entitlement" "all" {
  provider         = dodopayments
  include_resource = true
}

list "dodopayments_product_collection" "all" {
  provider         = dodopayments
  include_resource = true
}
```

Then run:

```sh
terraform query -generate-config-out=generated.tf
terraform plan
```

The plan must show only imports. Move the blocks from `generated.tf` to your configuration.

## Import one object

Use an `import` block with the identity of the object:

```hcl
import {
  to = dodopayments_product.example
  identity = {
    id = "<id>"
  }
}
```

A resource that belongs to a parent object has the parent ID in its identity. For the `terraform import` command, join the IDs with a slash:

```sh
terraform import dodopayments_product_localized_price.example <product_id>/<id>
```

Run `terraform plan -generate-config-out=generated.tf` to write the resource block.

## Objects with the same unique value

When a create fails with HTTP 409, an object with the same unique value exists. Import that object. Do not create a second object.
