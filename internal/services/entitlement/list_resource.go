// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package entitlement

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NewListResource returns the dodopayments_entitlement list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "entitlement",
	List:            listDataSourceSpec,
	IDProperty:      "id",
	DisplayProperty: "name",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds entitlements to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"integration_type": schema.StringAttribute{
				MarkdownDescription: "Filter by integration type.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("discord", "telegram", "github", "figma", "framer", "notion", "digital_files", "license_key", "feature_flag")},
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "Shows the translated name, description and unit of the catalog entities in this\nlocale. The value must be a code from `GET /locales`, for example `fr`.\nA field with no translation, an unknown code, or no code shows the base text.\n\nDo not send a text from a localized response back in a create or update request.\nThe base text then changes to the translated text.",
				Optional:            true,
			},
		},
	},
}
