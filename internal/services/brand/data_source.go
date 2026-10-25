// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NewDataSource returns the dodopayments_brand data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "brand",
	ReadPath:     "brands/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{"statement_descriptor": "DODOPAY_"},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one brand by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the brand.",
				Required:            true,
			},
			"brand_id": schema.StringAttribute{
				Computed: true,
			},
			"business_id": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"statement_descriptor": schema.StringAttribute{
				Computed: true,
			},
			"verification_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"verification_status": schema.StringAttribute{
				Computed: true,
			},
			"archived_at": schema.StringAttribute{
				MarkdownDescription: "Time the brand was archived. Null for an active brand.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"image": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"reason_for_hold": schema.StringAttribute{
				MarkdownDescription: "Incase the brand verification fails or is put on hold.",
				Computed:            true,
			},
			"support_email": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_brands data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "brands",
	ListPath:        "brands",
	ParentParams:    []string{},
	QueryParams:     []string{"include_archived"},
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
		StripPrefix: map[string]string{"statement_descriptor": "DODOPAY_"},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists brands.",
		Attributes: map[string]schema.Attribute{
			"include_archived": schema.BoolAttribute{
				MarkdownDescription: "Set to true to also list archived brands. Default false.",
				Optional:            true,
			},
			"max_items": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of items to return. The default is 1000.",
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The brands that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"brand_id": schema.StringAttribute{
							Computed: true,
						},
						"business_id": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"statement_descriptor": schema.StringAttribute{
							Computed: true,
						},
						"verification_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"verification_status": schema.StringAttribute{
							Computed: true,
						},
						"archived_at": schema.StringAttribute{
							MarkdownDescription: "Time the brand was archived. Null for an active brand.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"image": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"reason_for_hold": schema.StringAttribute{
							MarkdownDescription: "Incase the brand verification fails or is put on hold.",
							Computed:            true,
						},
						"support_email": schema.StringAttribute{
							Computed: true,
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	},
}
