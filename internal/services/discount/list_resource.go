// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package discount

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NewListResource returns the dodopayments_discount list resource for terraform query.
func NewListResource() list.ListResource {
	return engine.NewListResource(listResourceSpec)()
}

var listResourceSpec = &engine.ListResourceSpec{
	Name:            "discount",
	List:            listDataSourceSpec,
	IDProperty:      "discount_id",
	DisplayProperty: "name",
	Resource:        resourceSpec,
	Schema: schema.Schema{
		MarkdownDescription: "Finds discounts to import with `terraform query`.",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				MarkdownDescription: "Filter by active status. `true` = currently redeemable (started, not\nexpired, not usage-exhausted). `false` = not currently redeemable\n(expired, usage-exhausted, or pending a future `starts_at`).",
				Optional:            true,
			},
			"code": schema.StringAttribute{
				MarkdownDescription: "Filter by discount code (partial match, case-insensitive).",
				Optional:            true,
			},
			"discount_type": schema.StringAttribute{
				Optional:   true,
				Validators: []validator.String{stringvalidator.OneOf("flat", "percentage")},
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "Filter by product restriction (only discounts that apply to this product).",
				Optional:            true,
			},
		},
	},
}
