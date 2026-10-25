// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addon

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
)

// NewListResource returns the dodopayments_addon list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "addon",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "name",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds addons to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived addons. The default is false.",
				Optional:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "Shows the translated name, description and unit of the catalog entities in this\nlocale. The value must be a code from `GET /locales`, for example `fr`.\nA field with no translation, an unknown code, or no code shows the base text.\n\nDo not send a text from a localized response back in a create or update request.\nThe base text then changes to the translated text.",
				Optional:            true,
			},
		},
	},
}
