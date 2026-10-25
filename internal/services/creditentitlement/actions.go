// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package creditentitlement

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// NewCreditEntitlementUndeleteAction returns the dodopayments_credit_entitlement_undelete action.
func NewCreditEntitlementUndeleteAction() action.Action {
	return engine.NewAction(actionCreditEntitlementUndeleteSpec)()
}

var actionCreditEntitlementUndeleteSpec = &engine.ActionSpec{
	Name:    "credit_entitlement_undelete",
	Method:  "POST",
	Path:    "credit-entitlements/%s/undelete",
	Params:  []string{"id"},
	Success: "POST request succeeded",
	Schema: schema.Schema{
		MarkdownDescription: "Undeletes a soft-deleted credit entitlement by clearing `deleted_at`,\nmaking it available again through standard list and get endpoints.\n\n# Authentication\nRequires an API key with `Editor` role.\n\n# Path Parameters\n- `id` - The unique identifier of the credit entitlement to restore (format: `cde_...`)\n\n# Responses\n- `200 OK` - Credit entitlement restored successfully\n- `500 Internal Server Error` - Database error, entitlement not found, or entitlement is not deleted\n\n# Business Logic\n- Only deleted credit entitlements can be restored\n- The query filters for `deleted_at IS NOT NULL`, so non-deleted entitlements will result in 0 rows affected\n- If no rows are affected (entitlement doesn't exist, doesn't belong to business, or is not deleted), returns 500\n- The `updated_at` timestamp is automatically updated on successful restoration\n- Once restored, the entitlement becomes immediately available in the standard list and get endpoints\n- All configuration settings are preserved during delete/restore operations\n\n# Error Handling\nThis endpoint returns 500 Internal Server Error in several cases:\n- The credit entitlement does not exist\n- The credit entitlement belongs to a different business\n- The credit entitlement is not currently deleted (already active)\n\nCallers should verify the entitlement exists and is deleted before calling this endpoint.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the credit entitlement.",
				Required:            true,
			},
		},
	},
}
