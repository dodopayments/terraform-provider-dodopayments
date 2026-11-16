// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
)

// NewListResource returns the dodopayments_meter list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "meter",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "name",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds meters to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived meters.",
				Optional:            true,
			},
		},
	},
}
