// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productcollection

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NewDataSource returns the dodopayments_product_collection data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "product_collection",
	ReadPath:     "product-collections/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{"groups.products": "product_id"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one product collection by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the product collection.",
				Required:            true,
			},
			"brand_id": schema.StringAttribute{
				MarkdownDescription: "Brand ID for the collection.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the collection was created.",
				Computed:            true,
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: "Groups in this collection.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Computed: true,
						},
						"products": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"addons_count": schema.Int64Attribute{
										Computed: true,
									},
									"files_count": schema.Int64Attribute{
										Computed: true,
									},
									"has_credit_entitlements": schema.BoolAttribute{
										MarkdownDescription: "Whether this product has any credit entitlements attached.",
										Computed:            true,
									},
									"is_recurring": schema.BoolAttribute{
										Computed: true,
									},
									"license_key_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"meters_count": schema.Int64Attribute{
										Computed: true,
									},
									"product_id": schema.StringAttribute{
										Computed: true,
									},
									"status": schema.BoolAttribute{
										Computed: true,
									},
									"currency": schema.StringAttribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"price": schema.Int64Attribute{
										Computed: true,
									},
									"price_detail": schema.StringAttribute{
										MarkdownDescription: "One-time price details.",
										Computed:            true,
									},
									"tax_category": schema.StringAttribute{
										MarkdownDescription: "Represents the different categories of taxation applicable to various products and services.",
										Computed:            true,
									},
									"tax_inclusive": schema.BoolAttribute{
										Computed: true,
									},
								},
							},
						},
						"status": schema.BoolAttribute{
							Computed: true,
						},
						"group_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the collection.",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the collection was last updated.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the collection.",
				Computed:            true,
			},
			"effective_at_on_downgrade": schema.StringAttribute{
				MarkdownDescription: "Default effective_at setting for subscription plan downgrades (null = inherit from business).",
				Computed:            true,
			},
			"effective_at_on_upgrade": schema.StringAttribute{
				MarkdownDescription: "Default effective_at setting for subscription plan upgrades (null = inherit from business).",
				Computed:            true,
			},
			"on_payment_failure": schema.StringAttribute{
				MarkdownDescription: "Default behavior for subscription plan changes on payment failure (null = inherit from business).",
				Computed:            true,
			},
			"proration_billing_mode_on_downgrade": schema.StringAttribute{
				MarkdownDescription: "Default proration billing mode for subscription plan downgrades (null = inherit from business).",
				Computed:            true,
			},
			"proration_billing_mode_on_upgrade": schema.StringAttribute{
				MarkdownDescription: "Default proration billing mode for subscription plan upgrades (null = inherit from business).",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_product_collections data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "product_collections",
	ListPath:        "product-collections",
	ParentParams:    []string{},
	QueryParams:     []string{"archived", "brand_id", "locale"},
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
		ListKeys:    map[string]string{"groups.products": "product_id"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists product collections.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived collections.",
				Optional:            true,
			},
			"brand_id": schema.StringAttribute{
				MarkdownDescription: "Filter by Brand id.",
				Optional:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "Shows the translated name, description and unit of the catalog entities in this\nlocale. The value must be a code from `GET /locales`, for example `fr`.\nA field with no translation, an unknown code, or no code shows the base text.\n\nDo not send a text from a localized response back in a create or update request.\nThe base text then changes to the translated text.",
				Optional:            true,
			},
			"max_items": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of items to return. The default is 1000.",
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The product collections that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Collection ID.",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when created.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Collection name.",
							Computed:            true,
						},
						"products_count": schema.Int64Attribute{
							MarkdownDescription: "Number of products in the collection.",
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when last updated.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Collection description.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
