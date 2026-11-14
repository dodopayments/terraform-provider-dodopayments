// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package creditentitlement

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
)

// NewListResource returns the dodopayments_credit_entitlement list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "credit_entitlement",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "name",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds credit entitlements to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"deleted": schema.BoolAttribute{
				MarkdownDescription: "List deleted credit entitlements.",
				Optional:            true,
			},
		},
	},
}
