// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewResource returns the dodopayments_webhook resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "webhook",
	IDProperty:       "id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "webhooks",
	CreatePathParam:  "",
	ReadPath:         "webhooks/%s",
	UpdatePath:       "webhooks/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "webhooks/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"description", "disabled", "filter_types", "headers", "idempotency_key", "metadata", "rate_limit", "url"},
	UpdateFields:     []string{"description", "disabled", "filter_types", "metadata", "rate_limit", "url"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{"headers": true, "idempotency_key": true},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	WriteOnly:        []string{"headers", "idempotency_key"},
	MergeMaps:        map[string]bool{},
	ReadAfterWrite:   false,
	ArchivedProperty: "",
	DestroyWarning:   "",
	ReplaceFields:    []string{"headers", "idempotency_key"},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the webhook.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a webhook.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The webhook's ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Created at timestamp.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "An example webhook name. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"disabled": schema.BoolAttribute{
				MarkdownDescription: "Create the webhook in a disabled state.\n\nDefault is false. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"filter_types": schema.ListAttribute{
				MarkdownDescription: "Filter events to the webhook.\n\nWebhook event will only be sent for events in the list.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"headers": schema.MapAttribute{
				MarkdownDescription: "Custom headers to be passed. A change to this value destroys the resource and creates a new resource.",
				Optional:            true,
				Sensitive:           true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"idempotency_key": schema.StringAttribute{
				MarkdownDescription: "The request's idempotency key. A change to this value destroys the resource and creates a new resource.",
				Optional:            true,
				Sensitive:           true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Metadata to be passed to the webhook\nDefaut is {}. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
			},
			"rate_limit": schema.Int64Attribute{
				MarkdownDescription: "Configured rate limit.",
				Optional:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Updated at timestamp.",
				Computed:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "Url of the webhook.",
				Required:            true,
			},
		},
	},
}
