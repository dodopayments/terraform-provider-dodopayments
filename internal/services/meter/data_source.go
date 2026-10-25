// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NewDataSource returns the dodopayments_meter data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "meter",
	ReadPath:     "meters/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{"filter.clauses": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one meter by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"aggregation": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						MarkdownDescription: "Aggregation type for the meter.",
						Computed:            true,
					},
					"key": schema.StringAttribute{
						MarkdownDescription: "Required when type is not COUNT.",
						Computed:            true,
					},
				},
			},
			"business_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"event_name": schema.StringAttribute{
				Computed: true,
			},
			"measurement_unit": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				MarkdownDescription: "A filter structure that combines multiple conditions with logical conjunctions (AND/OR).\n\nSupports up to 3 levels of nesting to create complex filter expressions.\nEach filter has a conjunction (and/or) and clauses that can be either direct conditions or nested filters.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"clauses": schema.StringAttribute{
						MarkdownDescription: "Filter clauses - can be direct conditions or nested filters (up to 3 levels deep). A JSON document. Use `jsonencode()` to set it.",
						Computed:            true,
					},
					"conjunction": schema.StringAttribute{
						MarkdownDescription: "Logical conjunction to apply between clauses (and/or).",
						Computed:            true,
					},
				},
			},
		},
	},
}

// NewListDataSource returns the dodopayments_meters data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "meters",
	ListPath:        "meters",
	ParentParams:    []string{},
	QueryParams:     []string{"archived", "locale"},
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
		JSONAttrs:   map[string]bool{"items.filter.clauses": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists meters.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived meters.",
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
				MarkdownDescription: "The meters that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"aggregation": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									MarkdownDescription: "Aggregation type for the meter.",
									Computed:            true,
								},
								"key": schema.StringAttribute{
									MarkdownDescription: "Required when type is not COUNT.",
									Computed:            true,
								},
							},
						},
						"business_id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"event_name": schema.StringAttribute{
							Computed: true,
						},
						"measurement_unit": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"filter": schema.SingleNestedAttribute{
							MarkdownDescription: "A filter structure that combines multiple conditions with logical conjunctions (AND/OR).\n\nSupports up to 3 levels of nesting to create complex filter expressions.\nEach filter has a conjunction (and/or) and clauses that can be either direct conditions or nested filters.",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"clauses": schema.StringAttribute{
									MarkdownDescription: "Filter clauses - can be direct conditions or nested filters (up to 3 levels deep). A JSON document. Use `jsonencode()` to set it.",
									Computed:            true,
								},
								"conjunction": schema.StringAttribute{
									MarkdownDescription: "Logical conjunction to apply between clauses (and/or).",
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	},
}
