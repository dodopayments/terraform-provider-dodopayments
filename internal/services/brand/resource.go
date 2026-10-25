// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// NewResource returns the dodopayments_brand resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "brand",
	IDProperty:       "brand_id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "brands",
	CreatePathParam:  "",
	ReadPath:         "brands/%s",
	UpdatePath:       "brands/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "brands/%s/archive",
	DeleteMethod:     "POST",
	CreateFields:     []string{"description", "name", "statement_descriptor", "support_email", "url"},
	UpdateFields:     []string{"description", "name", "statement_descriptor", "support_email", "url"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{"statement_descriptor": "DODOPAY_"},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{},
	ReadAfterWrite:   false,
	ArchivedProperty: "archived_at",
	DestroyWarning:   "The API does not delete a brand. When Terraform destroys the resource, the provider archives the brand. The API cannot restore an archived brand.",
	ReplaceFields:    []string{},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the brand.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a brand. The API does not delete a brand. When Terraform destroys the resource, the provider archives the brand. The API cannot restore an archived brand.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the object.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"archived_at": schema.StringAttribute{
				MarkdownDescription: "Time the brand was archived. Null for an active brand.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"brand_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"business_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"image": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"reason_for_hold": schema.StringAttribute{
				MarkdownDescription: "Incase the brand verification fails or is put on hold.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				MarkdownDescription: "If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"support_email": schema.StringAttribute{
				MarkdownDescription: "If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"verification_enabled": schema.BoolAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"verification_status": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	},
}
