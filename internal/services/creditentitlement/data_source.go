// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package creditentitlement

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NewDataSource returns the dodopayments_credit_entitlement data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "credit_entitlement",
	ReadPath:     "credit-entitlements/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{"price_per_unit": true},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one credit entitlement by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"business_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"overage_behavior": schema.StringAttribute{
				MarkdownDescription: "Controls how overage is handled at billing cycle end.",
				Computed:            true,
			},
			"overage_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"precision": schema.Int64Attribute{
				Computed: true,
			},
			"rollover_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"unit": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"currency": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"expires_after_days": schema.Int64Attribute{
				Computed: true,
			},
			"max_rollover_count": schema.Int64Attribute{
				Computed: true,
			},
			"overage_limit": schema.Int64Attribute{
				Computed: true,
			},
			"price_per_unit": schema.StringAttribute{
				MarkdownDescription: "Price per credit unit.",
				Computed:            true,
			},
			"rollover_percentage": schema.Int64Attribute{
				Computed: true,
			},
			"rollover_timeframe_count": schema.Int64Attribute{
				Computed: true,
			},
			"rollover_timeframe_interval": schema.StringAttribute{
				MarkdownDescription: "Unit of a duration count (e.g. license-key validity period).",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_credit_entitlements data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "credit_entitlements",
	ListPath:        "credit-entitlements",
	ParentParams:    []string{},
	QueryParams:     []string{"deleted", "locale"},
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
		Decimals:    map[string]bool{"items.price_per_unit": true},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists credit entitlements.",
		Attributes: map[string]schema.Attribute{
			"deleted": schema.BoolAttribute{
				MarkdownDescription: "List deleted credit entitlements.",
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
				MarkdownDescription: "The credit entitlements that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"business_id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"overage_behavior": schema.StringAttribute{
							MarkdownDescription: "Controls how overage is handled at billing cycle end.",
							Computed:            true,
						},
						"overage_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"precision": schema.Int64Attribute{
							Computed: true,
						},
						"rollover_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"unit": schema.StringAttribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
						"currency": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"expires_after_days": schema.Int64Attribute{
							Computed: true,
						},
						"max_rollover_count": schema.Int64Attribute{
							Computed: true,
						},
						"overage_limit": schema.Int64Attribute{
							Computed: true,
						},
						"price_per_unit": schema.StringAttribute{
							MarkdownDescription: "Price per credit unit.",
							Computed:            true,
						},
						"rollover_percentage": schema.Int64Attribute{
							Computed: true,
						},
						"rollover_timeframe_count": schema.Int64Attribute{
							Computed: true,
						},
						"rollover_timeframe_interval": schema.StringAttribute{
							MarkdownDescription: "Unit of a duration count (e.g. license-key validity period).",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
