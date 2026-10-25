// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewDataSource returns the dodopayments_webhook data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "webhook",
	ReadPath:     "webhooks/%s",
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
		MarkdownDescription: "Reads one webhook by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The webhook's ID.",
				Required:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Created at timestamp.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "An example webhook name.",
				Computed:            true,
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Metadata of the webhook.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Updated at timestamp.",
				Computed:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "Url endpoint of the webhook.",
				Computed:            true,
			},
			"disabled": schema.BoolAttribute{
				MarkdownDescription: "Status of the webhook.\n\nIf true, events are not sent.",
				Computed:            true,
			},
			"filter_types": schema.ListAttribute{
				MarkdownDescription: "Filter events to the webhook.\n\nWebhook event will only be sent for events in the list.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"rate_limit": schema.Int64Attribute{
				MarkdownDescription: "Configured rate limit.",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_webhooks data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "webhooks",
	ListPath:        "webhooks",
	ParentParams:    []string{},
	QueryParams:     []string{},
	ItemsProperty:   "data",
	PageNumberParam: "",
	PageSizeParam:   "limit",
	CursorParam:     "iterator",
	CursorField:     "iterator",
	LastPageField:   "done",
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
		MarkdownDescription: "Lists webhooks.",
		Attributes: map[string]schema.Attribute{

			"max_items": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of items to return. The default is 1000.",
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The webhooks that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The webhook's ID.",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Created at timestamp.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "An example webhook name.",
							Computed:            true,
						},
						"metadata": schema.MapAttribute{
							MarkdownDescription: "Metadata of the webhook.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "Updated at timestamp.",
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "Url endpoint of the webhook.",
							Computed:            true,
						},
						"disabled": schema.BoolAttribute{
							MarkdownDescription: "Status of the webhook.\n\nIf true, events are not sent.",
							Computed:            true,
						},
						"filter_types": schema.ListAttribute{
							MarkdownDescription: "Filter events to the webhook.\n\nWebhook event will only be sent for events in the list.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"rate_limit": schema.Int64Attribute{
							MarkdownDescription: "Configured rate limit.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
