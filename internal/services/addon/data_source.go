// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addon

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// NewDataSource returns the dodopayments_addon data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "addon",
	ReadPath:     "addons/%s",
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
		MarkdownDescription: "Reads one addon by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "id of the Addon.",
				Required:            true,
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the business to which the addon belongs.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Created time.",
				Computed:            true,
			},
			"currency": schema.StringAttribute{
				MarkdownDescription: "Currency of the Addon.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Addon.",
				Computed:            true,
			},
			"price": schema.Int64Attribute{
				MarkdownDescription: "Amount of the addon.",
				Computed:            true,
			},
			"tax_category": schema.StringAttribute{
				MarkdownDescription: "Tax category applied to this Addon.",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Updated time.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the Addon.",
				Computed:            true,
			},
			"image": schema.StringAttribute{
				MarkdownDescription: "Image of the Addon.",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_addons data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "addons",
	ListPath:        "addons",
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
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists addons.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived addons. The default is false.",
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
				MarkdownDescription: "The addons that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "id of the Addon.",
							Computed:            true,
						},
						"business_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for the business to which the addon belongs.",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Created time.",
							Computed:            true,
						},
						"currency": schema.StringAttribute{
							MarkdownDescription: "Currency of the Addon.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the Addon.",
							Computed:            true,
						},
						"price": schema.Int64Attribute{
							MarkdownDescription: "Amount of the addon.",
							Computed:            true,
						},
						"tax_category": schema.StringAttribute{
							MarkdownDescription: "Tax category applied to this Addon.",
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "Updated time.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Optional description of the Addon.",
							Computed:            true,
						},
						"image": schema.StringAttribute{
							MarkdownDescription: "Image of the Addon.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
