// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
)

// NewListResource returns the dodopayments_webhook list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "webhook",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "id",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds webhooks to import with `terraform query`.",
		Attributes:          map[string]schema.Attribute{},
	},
}
