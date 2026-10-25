// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package entitlement

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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewResource returns the dodopayments_entitlement resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "entitlement",
	IDProperty:       "id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "entitlements",
	CreatePathParam:  "",
	ReadPath:         "entitlements/%s",
	UpdatePath:       "entitlements/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "entitlements/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"description", "integration_config", "integration_type", "metadata", "name"},
	UpdateFields:     []string{"description", "integration_config", "metadata", "name"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{"integration_config.digital_files.files.content_type": true, "integration_config.digital_files.files.download_url": true, "integration_config.digital_files.files.expires_in": true, "integration_config.digital_files.files.file_id": true, "integration_config.digital_files.files.file_size": true, "integration_config.digital_files.files.filename": true},
		KeepPrior:   map[string]bool{"integration_config.external_url": true, "integration_config.instructions": true},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{"metadata": true},
	ReadAfterWrite:   false,
	ArchivedProperty: "",
	DestroyWarning:   "",
	ReplaceFields:    []string{"integration_type"},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the entitlement.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a entitlement.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier of the entitlement.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "Identifier of the business that owns this entitlement.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the entitlement was created.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description.",
				Optional:            true,
			},
			"integration_config": schema.SingleNestedAttribute{
				MarkdownDescription: "Platform-specific configuration (validated per integration_type).",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"activation_message": schema.StringAttribute{
						MarkdownDescription: "Optional message displayed when a customer activates the license\nkey (≤ 2500 characters). If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"activations_limit": schema.Int64Attribute{
						MarkdownDescription: "Maximum activations allowed per issued license key. Omit for unlimited. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"chat_id": schema.StringAttribute{
						MarkdownDescription: "Telegram chat ID. For groups this is typically a negative integer. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"digital_files": schema.SingleNestedAttribute{
						MarkdownDescription: "Populated digital-files payload with each file's metadata and a\nshort-lived presigned download URL.",
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"files": schema.ListNestedAttribute{
								MarkdownDescription: "One entry per attached file.",
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"download_url": schema.StringAttribute{
											MarkdownDescription: "Short-lived presigned URL for downloading the file. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"expires_in": schema.StringAttribute{
											MarkdownDescription: "Seconds until `download_url` expires. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"file_id": schema.StringAttribute{
											MarkdownDescription: "Identifier of the attached file. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"filename": schema.StringAttribute{
											MarkdownDescription: "Original filename of the attached file. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"content_type": schema.StringAttribute{
											MarkdownDescription: "Optional content-type declared at upload. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"file_size": schema.StringAttribute{
											MarkdownDescription: "Optional size of the file in bytes. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
									},
								},
							},
							"external_url": schema.StringAttribute{
								MarkdownDescription: "Optional external URL, passed through from the entitlement\nconfiguration.",
								Computed:            true,
							},
							"instructions": schema.StringAttribute{
								MarkdownDescription: "Optional human-readable delivery instructions, passed through from\nthe entitlement configuration.",
								Computed:            true,
							},
						},
					},
					"duration_count": schema.Int64Attribute{
						MarkdownDescription: "Validity duration of issued license keys. Provide both\n`duration_count` and `duration_interval` together for a fixed\nduration; omit both for non-expiring keys. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"duration_interval": schema.StringAttribute{
						MarkdownDescription: "Unit of `duration_count`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{stringvalidator.OneOf("Day", "Week", "Month", "Year")},
					},
					"external_url": schema.StringAttribute{
						MarkdownDescription: "Optional external URL shown to the customer alongside the files.",
						Optional:            true,
					},
					"feature_id": schema.StringAttribute{
						MarkdownDescription: "Merchant-chosen identifier for the capability this entitlement\nunlocks. Not unique across entitlements. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"feature_type": schema.StringAttribute{
						MarkdownDescription: "Type of capability conferred. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{stringvalidator.OneOf("boolean")},
					},
					"figma_file_id": schema.StringAttribute{
						MarkdownDescription: "Figma file identifier to grant access to. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"framer_template_id": schema.StringAttribute{
						MarkdownDescription: "Framer template identifier to grant access to. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"fulfillment_mode": schema.StringAttribute{
						MarkdownDescription: "How license keys are fulfilled. `auto` (default) generates and delivers\nkeys to customers automatically; `manual` creates pending grants that you\nfulfill with the supplied key via `POST /grants/{grant_id}/license-key`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{stringvalidator.OneOf("auto", "manual")},
					},
					"guild_id": schema.StringAttribute{
						MarkdownDescription: "Discord guild (server) ID. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"instructions": schema.StringAttribute{
						MarkdownDescription: "Optional human-readable delivery instructions shown to the customer\nalongside the files.",
						Optional:            true,
					},
					"notion_template_id": schema.StringAttribute{
						MarkdownDescription: "Notion template identifier to grant access to. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"permission": schema.StringAttribute{
						MarkdownDescription: "Permission to grant on the repository. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{stringvalidator.OneOf("pull", "push", "admin", "maintain", "triage")},
					},
					"role_id": schema.StringAttribute{
						MarkdownDescription: "Optional Discord role to assign within the guild. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"target_id": schema.StringAttribute{
						MarkdownDescription: "Repository or organisation slug to grant access to. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
				},
			},
			"integration_type": schema.StringAttribute{
				MarkdownDescription: "Which platform integration this entitlement uses. A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:          []validator.String{stringvalidator.OneOf("discord", "telegram", "github", "figma", "framer", "notion", "digital_files", "license_key", "feature_flag")},
			},
			"is_active": schema.BoolAttribute{
				MarkdownDescription: "Always `true` for entitlements returned by the public API;\nsoft-deleted entitlements are not returned.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Additional metadata for the entitlement.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name for this entitlement.",
				Required:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the entitlement was last modified.",
				Computed:            true,
			},
		},
	},
}
