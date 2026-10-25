// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productlocalizedprice

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
)

// NewListResource returns the dodopayments_product_localized_price list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "product_localized_price",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "id",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds product localized prices to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"product_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the product.",
				Required:            true,
			},
		},
	},
}
