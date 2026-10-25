// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productlocalizedprice

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NewDataSource returns the dodopayments_product_localized_price data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "product_localized_price",
	ReadPath:     "products/%s/localized-prices/%s",
	ParentParams: []string{"product_id"},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one product localized price by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the localized price.",
				Required:            true,
			},
			"amount": schema.Int64Attribute{
				MarkdownDescription: "Amount in the smallest currency unit (e.g., cents).",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the localized price was created.",
				Computed:            true,
			},
			"currency": schema.StringAttribute{
				MarkdownDescription: "Currency to charge in.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Pricing mode of the rule: by_currency or by_country.",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the localized price was last updated.",
				Computed:            true,
			},
			"country_code": schema.StringAttribute{
				MarkdownDescription: "Country the rule applies to. Only set when mode is by_country.",
				Computed:            true,
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the product.",
				Required:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_product_localized_prices data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "product_localized_prices",
	ListPath:        "products/%s/localized-prices",
	ParentParams:    []string{"product_id"},
	QueryParams:     []string{},
	ItemsProperty:   "items",
	PageNumberParam: "",
	PageSizeParam:   "",
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
		MarkdownDescription: "Lists product localized prices.",
		Attributes: map[string]schema.Attribute{
			"product_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the product.",
				Required:            true,
			},
			"max_items": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of items to return. The default is 1000.",
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The product localized prices that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for the localized price.",
							Computed:            true,
						},
						"amount": schema.Int64Attribute{
							MarkdownDescription: "Amount in the smallest currency unit (e.g., cents).",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the localized price was created.",
							Computed:            true,
						},
						"currency": schema.StringAttribute{
							MarkdownDescription: "Currency to charge in.",
							Computed:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: "Pricing mode of the rule: by_currency or by_country.",
							Computed:            true,
						},
						"product_id": schema.StringAttribute{
							MarkdownDescription: "Product this localized price belongs to.",
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the localized price was last updated.",
							Computed:            true,
						},
						"country_code": schema.StringAttribute{
							MarkdownDescription: "Country the rule applies to. Only set when mode is by_country.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
