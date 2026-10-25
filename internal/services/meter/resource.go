// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NewResource returns the dodopayments_meter resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "meter",
	IDProperty:       "id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "meters",
	CreatePathParam:  "",
	ReadPath:         "meters/%s",
	UpdatePath:       "meters/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "meters/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"aggregation", "description", "event_name", "filter", "measurement_unit", "name"},
	UpdateFields:     []string{"description", "measurement_unit", "name"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{"filter.clauses": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{},
	ReadAfterWrite:   false,
	ArchivedProperty: "",
	DestroyWarning:   "The API does not delete this object. When Terraform destroys the resource, the provider archives the object.",
	ReplaceFields:    []string{"aggregation", "event_name", "filter"},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the meter.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a meter. The API does not delete this object. When Terraform destroys the resource, the provider archives the object.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"aggregation": schema.SingleNestedAttribute{
				MarkdownDescription: "Aggregation configuration for the meter. A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						MarkdownDescription: "Required when type is not COUNT. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						MarkdownDescription: "Aggregation type for the meter.",
						Required:            true,
						Validators:          []validator.String{stringvalidator.OneOf("count", "sum", "max", "last")},
					},
				},
				PlanModifiers: []planmodifier.Object{engine.RequiresReplaceIfChanged()},
			},
			"business_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the meter.",
				Optional:            true,
			},
			"event_name": schema.StringAttribute{
				MarkdownDescription: "Event name to track. A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"filter": schema.SingleNestedAttribute{
				MarkdownDescription: "Optional filter to apply to the meter. A change to this value destroys the resource and creates a new resource.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"clauses": schema.StringAttribute{
						MarkdownDescription: "Filter clauses - can be direct conditions or nested filters (up to 3 levels deep). A JSON document. Use `jsonencode()` to set it.",
						Required:            true,
					},
					"conjunction": schema.StringAttribute{
						MarkdownDescription: "Logical conjunction to apply between clauses (and/or).",
						Required:            true,
						Validators:          []validator.String{stringvalidator.OneOf("and", "or")},
					},
				},
				PlanModifiers: []planmodifier.Object{engine.RequiresReplaceIfChanged()},
			},
			"measurement_unit": schema.StringAttribute{
				MarkdownDescription: "measurement unit.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the meter.",
				Required:            true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	},
}
