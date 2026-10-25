// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package entitlement

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewDataSource returns the dodopayments_entitlement data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "entitlement",
	ReadPath:     "entitlements/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{"integration_config.digital_files.files.content_type": true, "integration_config.digital_files.files.download_url": true, "integration_config.digital_files.files.expires_in": true, "integration_config.digital_files.files.file_id": true, "integration_config.digital_files.files.file_size": true, "integration_config.digital_files.files.filename": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one entitlement by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier of the entitlement.",
				Required:            true,
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "Identifier of the business that owns this entitlement.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the entitlement was created.",
				Computed:            true,
			},
			"integration_config": schema.SingleNestedAttribute{
				MarkdownDescription: "Integration-specific configuration. For `digital_files` entitlements\nthis includes presigned download URLs for each attached file.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"feature_id": schema.StringAttribute{
						MarkdownDescription: "Merchant-chosen identifier for the capability this entitlement unlocks.",
						Computed:            true,
					},
					"feature_type": schema.StringAttribute{
						MarkdownDescription: "Type of capability conferred. Only `boolean` is supported today.",
						Computed:            true,
					},
					"permission": schema.StringAttribute{
						MarkdownDescription: "Permission to grant on the repository.",
						Computed:            true,
					},
					"target_id": schema.StringAttribute{
						MarkdownDescription: "Repository or organisation slug to grant access to.",
						Computed:            true,
					},
					"guild_id": schema.StringAttribute{
						MarkdownDescription: "Discord guild (server) ID.",
						Computed:            true,
					},
					"role_id": schema.StringAttribute{
						MarkdownDescription: "Optional Discord role to assign within the guild.",
						Computed:            true,
					},
					"chat_id": schema.StringAttribute{
						MarkdownDescription: "Telegram chat ID. For groups this is typically a negative integer.",
						Computed:            true,
					},
					"figma_file_id": schema.StringAttribute{
						MarkdownDescription: "Figma file identifier to grant access to.",
						Computed:            true,
					},
					"framer_template_id": schema.StringAttribute{
						MarkdownDescription: "Framer template identifier to grant access to.",
						Computed:            true,
					},
					"notion_template_id": schema.StringAttribute{
						MarkdownDescription: "Notion template identifier to grant access to.",
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
					"activation_message": schema.StringAttribute{
						MarkdownDescription: "Optional message displayed when a customer activates the license\nkey (≤ 2500 characters).",
						Computed:            true,
					},
					"activations_limit": schema.Int64Attribute{
						MarkdownDescription: "Maximum activations allowed per issued license key. Omit for unlimited.",
						Computed:            true,
					},
					"duration_count": schema.Int64Attribute{
						MarkdownDescription: "Validity duration of issued license keys. Provide both\n`duration_count` and `duration_interval` together for a fixed\nduration; omit both for non-expiring keys.",
						Computed:            true,
					},
					"duration_interval": schema.StringAttribute{
						MarkdownDescription: "Unit of `duration_count`.",
						Computed:            true,
					},
					"fulfillment_mode": schema.StringAttribute{
						MarkdownDescription: "How license keys are fulfilled. `auto` (default) generates and delivers\nkeys to customers automatically; `manual` creates pending grants that you\nfulfill with the supplied key via `POST /grants/{grant_id}/license-key`.",
						Computed:            true,
					},
				},
			},
			"integration_type": schema.StringAttribute{
				MarkdownDescription: "Platform integration this entitlement uses.",
				Computed:            true,
			},
			"is_active": schema.BoolAttribute{
				MarkdownDescription: "Always `true` for entitlements returned by the public API;\nsoft-deleted entitlements are not returned.",
				Computed:            true,
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Arbitrary key-value metadata supplied at creation or via PATCH.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name supplied at creation.",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the entitlement was last modified.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description supplied at creation.",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_entitlements data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "entitlements",
	ListPath:        "entitlements",
	ParentParams:    []string{},
	QueryParams:     []string{"integration_type", "locale"},
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
		JSONAttrs:   map[string]bool{"items.integration_config.digital_files.files.content_type": true, "items.integration_config.digital_files.files.download_url": true, "items.integration_config.digital_files.files.expires_in": true, "items.integration_config.digital_files.files.file_id": true, "items.integration_config.digital_files.files.file_size": true, "items.integration_config.digital_files.files.filename": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists entitlements.",
		Attributes: map[string]schema.Attribute{
			"integration_type": schema.StringAttribute{
				MarkdownDescription: "Filter by integration type.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("discord", "telegram", "github", "figma", "framer", "notion", "digital_files", "license_key", "feature_flag")},
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
				MarkdownDescription: "The entitlements that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier of the entitlement.",
							Computed:            true,
						},
						"business_id": schema.StringAttribute{
							MarkdownDescription: "Identifier of the business that owns this entitlement.",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the entitlement was created.",
							Computed:            true,
						},
						"integration_config": schema.SingleNestedAttribute{
							MarkdownDescription: "Integration-specific configuration. For `digital_files` entitlements\nthis includes presigned download URLs for each attached file.",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"feature_id": schema.StringAttribute{
									MarkdownDescription: "Merchant-chosen identifier for the capability this entitlement unlocks.",
									Computed:            true,
								},
								"feature_type": schema.StringAttribute{
									MarkdownDescription: "Type of capability conferred. Only `boolean` is supported today.",
									Computed:            true,
								},
								"permission": schema.StringAttribute{
									MarkdownDescription: "Permission to grant on the repository.",
									Computed:            true,
								},
								"target_id": schema.StringAttribute{
									MarkdownDescription: "Repository or organisation slug to grant access to.",
									Computed:            true,
								},
								"guild_id": schema.StringAttribute{
									MarkdownDescription: "Discord guild (server) ID.",
									Computed:            true,
								},
								"role_id": schema.StringAttribute{
									MarkdownDescription: "Optional Discord role to assign within the guild.",
									Computed:            true,
								},
								"chat_id": schema.StringAttribute{
									MarkdownDescription: "Telegram chat ID. For groups this is typically a negative integer.",
									Computed:            true,
								},
								"figma_file_id": schema.StringAttribute{
									MarkdownDescription: "Figma file identifier to grant access to.",
									Computed:            true,
								},
								"framer_template_id": schema.StringAttribute{
									MarkdownDescription: "Framer template identifier to grant access to.",
									Computed:            true,
								},
								"notion_template_id": schema.StringAttribute{
									MarkdownDescription: "Notion template identifier to grant access to.",
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
								"activation_message": schema.StringAttribute{
									MarkdownDescription: "Optional message displayed when a customer activates the license\nkey (≤ 2500 characters).",
									Computed:            true,
								},
								"activations_limit": schema.Int64Attribute{
									MarkdownDescription: "Maximum activations allowed per issued license key. Omit for unlimited.",
									Computed:            true,
								},
								"duration_count": schema.Int64Attribute{
									MarkdownDescription: "Validity duration of issued license keys. Provide both\n`duration_count` and `duration_interval` together for a fixed\nduration; omit both for non-expiring keys.",
									Computed:            true,
								},
								"duration_interval": schema.StringAttribute{
									MarkdownDescription: "Unit of `duration_count`.",
									Computed:            true,
								},
								"fulfillment_mode": schema.StringAttribute{
									MarkdownDescription: "How license keys are fulfilled. `auto` (default) generates and delivers\nkeys to customers automatically; `manual` creates pending grants that you\nfulfill with the supplied key via `POST /grants/{grant_id}/license-key`.",
									Computed:            true,
								},
							},
						},
						"integration_type": schema.StringAttribute{
							MarkdownDescription: "Platform integration this entitlement uses.",
							Computed:            true,
						},
						"is_active": schema.BoolAttribute{
							MarkdownDescription: "Always `true` for entitlements returned by the public API;\nsoft-deleted entitlements are not returned.",
							Computed:            true,
						},
						"metadata": schema.MapAttribute{
							MarkdownDescription: "Arbitrary key-value metadata supplied at creation or via PATCH.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Display name supplied at creation.",
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the entitlement was last modified.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Optional description supplied at creation.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
