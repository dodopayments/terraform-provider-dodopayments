// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productcollection

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NewResource returns the dodopayments_product_collection resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "product_collection",
	IDProperty:       "id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "product-collections",
	CreatePathParam:  "",
	ReadPath:         "product-collections/%s",
	UpdatePath:       "product-collections/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "product-collections/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"brand_id", "description", "effective_at_on_downgrade", "effective_at_on_upgrade", "groups", "name", "on_payment_failure", "proration_billing_mode_on_downgrade", "proration_billing_mode_on_upgrade"},
	UpdateFields:     []string{"brand_id", "description", "effective_at_on_downgrade", "effective_at_on_upgrade", "name", "on_payment_failure", "proration_billing_mode_on_downgrade", "proration_billing_mode_on_upgrade"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{"groups.products": "product_id"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{},
	ReadAfterWrite:   false,
	ArchivedProperty: "",
	DestroyWarning:   "",
	ReplaceFields:    []string{"groups"},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the product collection.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a product collection.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the product collection.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"brand_id": schema.StringAttribute{
				MarkdownDescription: "Brand id for the collection, if not provided will default to primary brand. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the collection was created.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the product collection.",
				Optional:            true,
			},
			"effective_at_on_downgrade": schema.StringAttribute{
				MarkdownDescription: "Default effective_at setting for subscription plan downgrades (NULL = inherit from business).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("immediately", "next_billing_date")},
			},
			"effective_at_on_upgrade": schema.StringAttribute{
				MarkdownDescription: "Default effective_at setting for subscription plan upgrades (NULL = inherit from business).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("immediately", "next_billing_date")},
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: "Groups of products in this collection. A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Computed:      true,
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Optional group name. Multiple groups can have null names, but named groups must be unique per collection. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"products": schema.ListNestedAttribute{
							MarkdownDescription: "Products in this group.",
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"addons_count": schema.Int64Attribute{
										Computed: true,
									},
									"currency": schema.StringAttribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"files_count": schema.Int64Attribute{
										Computed: true,
									},
									"has_credit_entitlements": schema.BoolAttribute{
										MarkdownDescription: "Whether this product has any credit entitlements attached.",
										Computed:            true,
									},
									"id": schema.StringAttribute{
										Computed: true,
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
									"product_id": schema.StringAttribute{
										MarkdownDescription: "Product ID to include in the group.",
										Required:            true,
									},
									"status": schema.BoolAttribute{
										MarkdownDescription: "Status of the product in this group (defaults to true if not provided). If you remove this attribute from the configuration, the API keeps the last value.",
										Optional:            true,
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
							MarkdownDescription: "Status of the group (defaults to true if not provided). If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
					},
				},
				PlanModifiers: []planmodifier.List{engine.RequiresReplaceIfChanged()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the product collection.",
				Required:            true,
			},
			"on_payment_failure": schema.StringAttribute{
				MarkdownDescription: "Default behavior for subscription plan changes on payment failure (NULL = inherit from business).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("prevent_change", "apply_change")},
			},
			"proration_billing_mode_on_downgrade": schema.StringAttribute{
				MarkdownDescription: "Default proration billing mode for subscription plan downgrades (NULL = inherit from business).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("prorated_immediately", "full_immediately", "difference_immediately", "do_not_bill", "prorated_immediately_keep_cycle")},
			},
			"proration_billing_mode_on_upgrade": schema.StringAttribute{
				MarkdownDescription: "Default proration billing mode for subscription plan upgrades (NULL = inherit from business).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("prorated_immediately", "full_immediately", "difference_immediately", "do_not_bill", "prorated_immediately_keep_cycle")},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the collection was last updated.",
				Computed:            true,
			},
		},
	},
}
