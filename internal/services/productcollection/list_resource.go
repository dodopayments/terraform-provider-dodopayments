// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productcollection

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
)

// NewListResource returns the dodopayments_product_collection list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "product_collection",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "name",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds product collections to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived collections.",
				Optional:            true,
			},
			"brand_id": schema.StringAttribute{
				MarkdownDescription: "Filter by Brand id.",
				Optional:            true,
			},
		},
	},
}
