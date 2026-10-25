// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// NewProductUnarchiveAction returns the dodopayments_product_unarchive action.
func NewProductUnarchiveAction() action.Action {
	return engine.NewAction(actionProductUnarchiveSpec)()
}

var actionProductUnarchiveSpec = &engine.ActionSpec{
	Name:    "product_unarchive",
	Method:  "POST",
	Path:    "products/%s/unarchive",
	Params:  []string{"id"},
	Success: "POST request succeeded",
	Schema: schema.Schema{
		MarkdownDescription: "Sends `POST /products/{id}/unarchive`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the product.",
				Required:            true,
			},
		},
	},
}
