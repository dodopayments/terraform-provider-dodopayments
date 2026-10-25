// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package discount

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewDataSource returns the dodopayments_discount data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "discount",
	ReadPath:     "discounts/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one discount by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the discount.",
				Required:            true,
			},
			"amount": schema.Int64Attribute{
				MarkdownDescription: "The discount amount in **basis points** (e.g., 540 => 5.4%).",
				Computed:            true,
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "The business this discount belongs to.",
				Computed:            true,
			},
			"code": schema.StringAttribute{
				MarkdownDescription: "The discount code (up to 16 chars).",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the discount is created.",
				Computed:            true,
			},
			"customer_eligibility": schema.StringAttribute{
				MarkdownDescription: "Who may redeem this discount code.",
				Computed:            true,
			},
			"discount_id": schema.StringAttribute{
				MarkdownDescription: "The unique discount ID.",
				Computed:            true,
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Arbitrary key-value metadata. Values can be string, integer, number, or boolean.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"preserve_on_plan_change": schema.BoolAttribute{
				MarkdownDescription: "Whether this discount should be preserved when a subscription changes plans.\nDefault: false (discount is removed on plan change).",
				Computed:            true,
			},
			"restricted_to": schema.ListAttribute{
				MarkdownDescription: "List of product IDs to which this discount is restricted.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"times_used": schema.Int64Attribute{
				MarkdownDescription: "How many times this discount has been used.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of discount (`percentage` or `flat`).",
				Computed:            true,
			},
			"currency_options": schema.ListNestedAttribute{
				MarkdownDescription: "Per-currency options (flat deduction / percentage cap + minimum subtotal).\nEmpty for discounts without any configured currency options.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"currency": schema.StringAttribute{
							MarkdownDescription: "The currency this option applies to.",
							Computed:            true,
						},
						"is_default": schema.BoolAttribute{
							MarkdownDescription: "Whether this is the default row FX conversions pivot from.",
							Computed:            true,
						},
						"minimum_subtotal": schema.Int64Attribute{
							MarkdownDescription: "Eligible-cart threshold in this currency's subunits (0 = no minimum).",
							Computed:            true,
						},
						"max_amount_possible": schema.Int64Attribute{
							MarkdownDescription: "The most this code discounts in this currency's subunits (flat deduction\nor percentage cap).",
							Computed:            true,
						},
					},
				},
			},
			"expires_at": schema.StringAttribute{
				MarkdownDescription: "Optional date/time after which discount is expired.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name for the Discount.",
				Computed:            true,
			},
			"per_customer_usage_limit": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of times a single customer may redeem this discount, if any.",
				Computed:            true,
			},
			"starts_at": schema.StringAttribute{
				MarkdownDescription: "Optional date/time before which the discount is not yet active. NULL = active immediately.",
				Computed:            true,
			},
			"subscription_cycles": schema.Int64Attribute{
				MarkdownDescription: "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
				Computed:            true,
			},
			"usage_limit": schema.Int64Attribute{
				MarkdownDescription: "Usage limit for this discount, if any.",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_discounts data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "discounts",
	ListPath:        "discounts",
	ParentParams:    []string{},
	QueryParams:     []string{"active", "code", "discount_type", "product_id"},
	ItemsProperty:   "items",
	PageNumberParam: "page_number",
	PageSizeParam:   "page_size",
	CursorParam:     "",
	CursorField:     "",
	LastPageField:   "",
	ItemsAttr:       "items",
	MaxItemsAttr:    "max_items",
	DefaultMaxItems: 1000,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists discounts.",
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
			"max_items": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of items to return. The default is 1000.",
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The discounts that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"amount": schema.Int64Attribute{
							MarkdownDescription: "The discount amount in **basis points** (e.g., 540 => 5.4%).",
							Computed:            true,
						},
						"business_id": schema.StringAttribute{
							MarkdownDescription: "The business this discount belongs to.",
							Computed:            true,
						},
						"code": schema.StringAttribute{
							MarkdownDescription: "The discount code (up to 16 chars).",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the discount is created.",
							Computed:            true,
						},
						"customer_eligibility": schema.StringAttribute{
							MarkdownDescription: "Who may redeem this discount code.",
							Computed:            true,
						},
						"discount_id": schema.StringAttribute{
							MarkdownDescription: "The unique discount ID.",
							Computed:            true,
						},
						"metadata": schema.MapAttribute{
							MarkdownDescription: "Arbitrary key-value metadata. Values can be string, integer, number, or boolean.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"preserve_on_plan_change": schema.BoolAttribute{
							MarkdownDescription: "Whether this discount should be preserved when a subscription changes plans.\nDefault: false (discount is removed on plan change).",
							Computed:            true,
						},
						"restricted_to": schema.ListAttribute{
							MarkdownDescription: "List of product IDs to which this discount is restricted.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"times_used": schema.Int64Attribute{
							MarkdownDescription: "How many times this discount has been used.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The type of discount (`percentage` or `flat`).",
							Computed:            true,
						},
						"currency_options": schema.ListNestedAttribute{
							MarkdownDescription: "Per-currency options (flat deduction / percentage cap + minimum subtotal).\nEmpty for discounts without any configured currency options.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"currency": schema.StringAttribute{
										MarkdownDescription: "The currency this option applies to.",
										Computed:            true,
									},
									"is_default": schema.BoolAttribute{
										MarkdownDescription: "Whether this is the default row FX conversions pivot from.",
										Computed:            true,
									},
									"minimum_subtotal": schema.Int64Attribute{
										MarkdownDescription: "Eligible-cart threshold in this currency's subunits (0 = no minimum).",
										Computed:            true,
									},
									"max_amount_possible": schema.Int64Attribute{
										MarkdownDescription: "The most this code discounts in this currency's subunits (flat deduction\nor percentage cap).",
										Computed:            true,
									},
								},
							},
						},
						"expires_at": schema.StringAttribute{
							MarkdownDescription: "Optional date/time after which discount is expired.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name for the Discount.",
							Computed:            true,
						},
						"per_customer_usage_limit": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of times a single customer may redeem this discount, if any.",
							Computed:            true,
						},
						"starts_at": schema.StringAttribute{
							MarkdownDescription: "Optional date/time before which the discount is not yet active. NULL = active immediately.",
							Computed:            true,
						},
						"subscription_cycles": schema.Int64Attribute{
							MarkdownDescription: "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
							Computed:            true,
						},
						"usage_limit": schema.Int64Attribute{
							MarkdownDescription: "Usage limit for this discount, if any.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
