// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productcollection

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// NewProductCollectionUnarchiveAction returns the dodopayments_product_collection_unarchive action.
func NewProductCollectionUnarchiveAction() action.Action {
	return engine.NewAction(actionProductCollectionUnarchiveSpec)()
}

var actionProductCollectionUnarchiveSpec = &engine.ActionSpec{
	Name:    "product_collection_unarchive",
	Method:  "POST",
	Path:    "product-collections/%s/unarchive",
	Params:  []string{"id"},
	Success: "POST request succeeded",
	Schema: schema.Schema{
		MarkdownDescription: "Sends `POST /product-collections/{id}/unarchive`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the product collection.",
				Required:            true,
			},
		},
	},
}
