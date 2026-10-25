---
page_title: "Destroy behavior"
subcategory: "Guides"
description: |-
  What the API does with each object when Terraform destroys the resource.
---

# Destroy behavior

Some Dodo Payments objects cannot be deleted, because payments and subscriptions refer to them. The API archives these objects.
A plan that destroys or replaces such a resource shows a warning.

| Resource | When Terraform destroys the resource | Action that restores the object |
|---|---|---|
| `dodopayments_product` | The provider archives the object. | `dodopayments_product_unarchive` |
| `dodopayments_product_localized_price` | The provider archives the object. | None |
| `dodopayments_discount` | The provider deletes the object. | None |
| `dodopayments_addon` | The provider deletes the object. | `dodopayments_addon_unarchive` |
| `dodopayments_brand` | The API does not delete a brand. When Terraform destroys the resource, the provider archives the brand. The API cannot restore an archived brand. | None |
| `dodopayments_webhook` | The provider deletes the object. | None |
| `dodopayments_meter` | The provider archives the object. | `dodopayments_meter_unarchive` |
| `dodopayments_credit_entitlement` | The provider deletes the object. | `dodopayments_credit_entitlement_undelete` |
| `dodopayments_entitlement` | The provider deletes the object. | None |
| `dodopayments_product_collection` | The provider deletes the object. | `dodopayments_product_collection_unarchive` |

When the API archives an object outside Terraform, the next plan shows that Terraform must create the resource again.

To keep an object when you remove its resource block, use a `removed` block:

```hcl
removed {
  from = dodopayments_product.example
  lifecycle {
    destroy = false
  }
}
```

To make sure that Terraform does not destroy an object, set `prevent_destroy`:

```hcl
lifecycle {
  prevent_destroy = true
}
```
