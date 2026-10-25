// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewDataSource returns the dodopayments_product data source.
func NewDataSource() datasource.DataSource {
	return engine.NewDataSource(dataSourceSpec)()
}

var dataSourceSpec = &engine.DataSourceSpec{
	Name:         "product",
	ReadPath:     "products/%s",
	ParentParams: []string{},
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{"entitlements.integration_config.digital_files.external_url": true, "entitlements.integration_config.digital_files.files": true, "entitlements.integration_config.digital_files.instructions": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{"credit_entitlements": "credit_entitlement_id", "entitlements": "entitlement_id", "price.meters": "meter_id"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{"entitlements.entitlement_id": "id"},
		Decimals:    map[string]bool{"credit_entitlements.credits_amount": true, "credit_entitlements.overage_limit": true, "credit_entitlements.price_per_unit": true, "credit_entitlements.trial_credits": true, "price.meters.meter_units_per_credit": true, "price.meters.price_per_unit": true},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Reads one product by its ID.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the product.",
				Required:            true,
			},
			"brand_id": schema.StringAttribute{
				Computed: true,
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the business to which the product belongs.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the product was created.",
				Computed:            true,
			},
			"credit_entitlements": schema.ListNestedAttribute{
				MarkdownDescription: "Attached credit entitlements with settings.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique ID of this mapping.",
							Computed:            true,
						},
						"credit_entitlement_id": schema.StringAttribute{
							MarkdownDescription: "ID of the credit entitlement.",
							Computed:            true,
						},
						"credit_entitlement_name": schema.StringAttribute{
							MarkdownDescription: "Name of the credit entitlement.",
							Computed:            true,
						},
						"credit_entitlement_unit": schema.StringAttribute{
							MarkdownDescription: "Unit label for the credit entitlement.",
							Computed:            true,
						},
						"credits_amount": schema.StringAttribute{
							MarkdownDescription: "Number of credits granted.",
							Computed:            true,
						},
						"overage_behavior": schema.StringAttribute{
							MarkdownDescription: "Controls how overage is handled at billing cycle end.",
							Computed:            true,
						},
						"overage_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether overage is enabled.",
							Computed:            true,
						},
						"proration_behavior": schema.StringAttribute{
							MarkdownDescription: "Proration behavior for credit grants during plan changes.",
							Computed:            true,
						},
						"rollover_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether rollover is enabled.",
							Computed:            true,
						},
						"trial_credits_expire_after_trial": schema.BoolAttribute{
							MarkdownDescription: "Whether trial credits expire after trial.",
							Computed:            true,
						},
						"currency": schema.StringAttribute{
							MarkdownDescription: "Currency.",
							Computed:            true,
						},
						"expires_after_days": schema.Int64Attribute{
							MarkdownDescription: "Days until credits expire.",
							Computed:            true,
						},
						"low_balance_threshold_percent": schema.Int64Attribute{
							MarkdownDescription: "Low balance threshold percentage.",
							Computed:            true,
						},
						"max_rollover_count": schema.Int64Attribute{
							MarkdownDescription: "Maximum rollover cycles.",
							Computed:            true,
						},
						"overage_limit": schema.StringAttribute{
							MarkdownDescription: "Overage limit.",
							Computed:            true,
						},
						"price_per_unit": schema.StringAttribute{
							MarkdownDescription: "Price per unit.",
							Computed:            true,
						},
						"rollover_percentage": schema.Int64Attribute{
							MarkdownDescription: "Rollover percentage.",
							Computed:            true,
						},
						"rollover_timeframe_count": schema.Int64Attribute{
							MarkdownDescription: "Rollover timeframe count.",
							Computed:            true,
						},
						"rollover_timeframe_interval": schema.StringAttribute{
							MarkdownDescription: "Rollover timeframe interval.",
							Computed:            true,
						},
						"trial_credits": schema.StringAttribute{
							MarkdownDescription: "Trial credits.",
							Computed:            true,
						},
					},
				},
			},
			"entitlements": schema.ListNestedAttribute{
				MarkdownDescription: "Attached entitlements (integration-based access grants).",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"integration_config": schema.SingleNestedAttribute{
							MarkdownDescription: "Integration-specific configuration on an entitlement read response.\n\nFor `digital_files` entitlements the response includes presigned\ndownload URLs for each attached file; other integrations match the\nshape supplied at creation.",
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
										"files": schema.StringAttribute{
											MarkdownDescription: "One entry per attached file. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"external_url": schema.StringAttribute{
											MarkdownDescription: "Optional external URL, passed through from the entitlement\nconfiguration. A JSON document. Use `jsonencode()` to set it.",
											Computed:            true,
										},
										"instructions": schema.StringAttribute{
											MarkdownDescription: "Optional human-readable delivery instructions, passed through from\nthe entitlement configuration. A JSON document. Use `jsonencode()` to set it.",
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
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"is_recurring": schema.BoolAttribute{
				MarkdownDescription: "Indicates if the product is recurring (e.g., subscriptions).",
				Computed:            true,
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Additional custom data associated with the product.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"price": schema.SingleNestedAttribute{
				MarkdownDescription: "Pricing information for the product.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						MarkdownDescription: "The currency in which the payment is made.",
						Computed:            true,
					},
					"price": schema.Int64Attribute{
						MarkdownDescription: "The payment amount, in the smallest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.\n\nIf [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field represents\nthe **minimum** amount the customer must pay. Only for `type` `one_time_price` or `recurring_price`. Required when `type` is `one_time_price` or `recurring_price`.",
						Computed:            true,
					},
					"type": schema.StringAttribute{
						Computed: true,
					},
					"discount": schema.Int64Attribute{
						MarkdownDescription: "Deprecated: use `discount_bps` instead.\n\nDiscount applied to the price, represented as a percentage (0 to 100).\nA response rounds this value to the nearest whole percent.\nDefaults to `0`.",
						Computed:            true,
					},
					"discount_bps": schema.Int64Attribute{
						MarkdownDescription: "Discount applied to the price, in basis points. 100 basis points make\none percent, so `1250` is a discount of 12.5%.\n\nUse this field for a discount with a fraction of a percent. A request\nthat sends this field ignores `discount`. A value of `0` gives no\ndiscount.",
						Computed:            true,
					},
					"pay_what_you_want": schema.BoolAttribute{
						MarkdownDescription: "Indicates whether the customer can pay any amount they choose.\nIf set to `true`, the [`price`](Self::price) field is the minimum amount. Only for `type` `one_time_price`.",
						Computed:            true,
					},
					"purchasing_power_parity": schema.BoolAttribute{
						MarkdownDescription: "Opts this price in to purchasing power parity. The business must also\nenable purchasing power parity. The discount percentage per country is\nalways business-wide. Defaults to `false`.",
						Computed:            true,
					},
					"suggested_price": schema.Int64Attribute{
						MarkdownDescription: "A suggested price for the user to pay. This value is only considered if\n[`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is ignored. Only for `type` `one_time_price`.",
						Computed:            true,
					},
					"tax_inclusive": schema.BoolAttribute{
						MarkdownDescription: "Indicates if the price is tax inclusive.",
						Computed:            true,
					},
					"payment_frequency_count": schema.Int64Attribute{
						MarkdownDescription: "Number of `payment_frequency_interval` units between two charges.\nThis field sets the billing cycle.\nFor example, `1` with a `payment_frequency_interval` of `month` charges the customer every month. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
						Computed:            true,
					},
					"payment_frequency_interval": schema.StringAttribute{
						MarkdownDescription: "The time unit of the billing cycle (e.g., day, month, year).\nUse it with `payment_frequency_count`. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
						Computed:            true,
					},
					"subscription_period_count": schema.Int64Attribute{
						MarkdownDescription: "Number of `subscription_period_interval` units for the total duration of the subscription.\nThis field does not set the billing cycle. Use `payment_frequency_count` for the billing cycle.\nWhen this period ends, the subscription expires. Dodo Payments does not start a new\nbilling cycle after the subscription expires. Dodo Payments can still retry a failed\npayment for an earlier billing cycle.\n\nFor a subscription that renews until the customer cancels it, set this to `20` with a\n`subscription_period_interval` of `year`. This is the recommended value and also the maximum.\nThe subscription expires after 20 years if the customer does not cancel it before then.\n\nFor example, a monthly plan uses a `payment_frequency_count` of `1`, a `payment_frequency_interval`\nof `month`, and a subscription period of `20` `year`. Do not set the subscription period to\n`1` `month` for a monthly plan: the subscription then expires after one month. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
						Computed:            true,
					},
					"subscription_period_interval": schema.StringAttribute{
						MarkdownDescription: "The time unit of the total subscription duration (e.g., day, month, year).\nUse it with `subscription_period_count`. The recommended value is `year`, with a\n`subscription_period_count` of `20`. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
						Computed:            true,
					},
					"trial_amount": schema.Int64Attribute{
						MarkdownDescription: "Amount charged today for a paid trial, in the price currency's minor units.\nRequires `trial_period_days > 0`. Omit or null for a free trial (the default). Only for `type` `recurring_price`.",
						Computed:            true,
					},
					"trial_apply_discounts": schema.BoolAttribute{
						MarkdownDescription: "Whether discount codes reduce the trial charge. Defaults to false. Only meaningful\nwhen a paid trial is configured. Only for `type` `recurring_price`.",
						Computed:            true,
					},
					"trial_payment_method_optional": schema.BoolAttribute{
						MarkdownDescription: "Let a customer start a free trial with no card. Defaults to false. Only for `type` `recurring_price`.",
						Computed:            true,
					},
					"trial_period_days": schema.Int64Attribute{
						MarkdownDescription: "Number of days for the trial period. A value of `0` indicates no trial period. Only for `type` `recurring_price`.",
						Computed:            true,
					},
					"zero_amount_payment_method_optional": schema.BoolAttribute{
						MarkdownDescription: "Let a customer start a subscription with no card, when the amount due\ntoday is `0` (a native `0` price, or a 100% discount). Defaults to false. Only for `type` `recurring_price`.",
						Computed:            true,
					},
					"fixed_price": schema.Int64Attribute{
						MarkdownDescription: "The fixed payment amount. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`. Only for `type` `usage_based_price`. Required when `type` is `usage_based_price`.",
						Computed:            true,
					},
					"meters": schema.ListNestedAttribute{
						MarkdownDescription: "Only for `type` `usage_based_price`.",
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"meter_id": schema.StringAttribute{
									Computed: true,
								},
								"credit_entitlement_id": schema.StringAttribute{
									MarkdownDescription: "Optional credit entitlement ID to link this meter to for credit-based billing.",
									Computed:            true,
								},
								"description": schema.StringAttribute{
									MarkdownDescription: "Meter description. Will ignored on Request, but will be shown in response.",
									Computed:            true,
								},
								"free_threshold": schema.Int64Attribute{
									Computed: true,
								},
								"measurement_unit": schema.StringAttribute{
									MarkdownDescription: "Meter measurement unit. Will ignored on Request, but will be shown in response.",
									Computed:            true,
								},
								"meter_units_per_credit": schema.StringAttribute{
									MarkdownDescription: "Number of meter units that equal one credit. Required when credit_entitlement_id is set.",
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Meter name. Will ignored on Request, but will be shown in response.",
									Computed:            true,
								},
								"price_per_unit": schema.StringAttribute{
									MarkdownDescription: "The price per unit in lowest denomination. Must be greater than zero. Supports up to 5 digits before decimal point and 12 decimal places.",
									Computed:            true,
								},
							},
						},
					},
				},
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the product.",
				Computed:            true,
			},
			"tax_category": schema.StringAttribute{
				MarkdownDescription: "Tax category associated with the product.",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the product was last updated.",
				Computed:            true,
			},
			"addons": schema.SetAttribute{
				MarkdownDescription: "Available Addons for subscription products.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the product, optional.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the product, optional.",
				Computed:            true,
			},
			"pricing_mode": schema.StringAttribute{
				MarkdownDescription: "Pricing mode for localized pricing. NULL means base-only (no localized rules apply).",
				Computed:            true,
			},
			"product_collection_id": schema.StringAttribute{
				MarkdownDescription: "The product collection ID this product belongs to, if any.",
				Computed:            true,
			},
		},
	},
}

// NewListDataSource returns the dodopayments_products data source.
func NewListDataSource() datasource.DataSource {
	return engine.NewListDataSource(listDataSourceSpec)()
}

var listDataSourceSpec = &engine.ListDataSourceSpec{
	Name:            "products",
	ListPath:        "products",
	ParentParams:    []string{},
	QueryParams:     []string{"archived", "brand_id", "locale", "recurring"},
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
		JSONAttrs:   map[string]bool{"items.entitlements.integration_config.digital_files.external_url": true, "items.entitlements.integration_config.digital_files.files": true, "items.entitlements.integration_config.digital_files.instructions": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{"credit_entitlements": "credit_entitlement_id", "entitlements": "entitlement_id", "price.meters": "meter_id"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{"entitlements.entitlement_id": "id"},
		Decimals:    map[string]bool{},
	},
	Schema: schema.Schema{
		MarkdownDescription: "Lists products.",
		Attributes: map[string]schema.Attribute{
			"archived": schema.BoolAttribute{
				MarkdownDescription: "List archived products.",
				Optional:            true,
			},
			"brand_id": schema.StringAttribute{
				MarkdownDescription: "filter by Brand id.",
				Optional:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "Shows the translated name, description and unit of the catalog entities in this\nlocale. The value must be a code from `GET /locales`, for example `fr`.\nA field with no translation, an unknown code, or no code shows the base text.\n\nDo not send a text from a localized response back in a create or update request.\nThe base text then changes to the translated text.",
				Optional:            true,
			},
			"recurring": schema.BoolAttribute{
				MarkdownDescription: "Filter products by pricing type:\n- `true`: Show only recurring pricing products (e.g. subscriptions)\n- `false`: Show only one-time price products\n- `null` or absent: Show both types of products.",
				Optional:            true,
			},
			"max_items": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of items to return. The default is 1000.",
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The products that the API returned.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"business_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for the business to which the product belongs.",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the product was created.",
							Computed:            true,
						},
						"entitlements": schema.ListNestedAttribute{
							MarkdownDescription: "Entitlements linked to this product.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"integration_config": schema.SingleNestedAttribute{
										MarkdownDescription: "Integration-specific configuration on an entitlement read response.\n\nFor `digital_files` entitlements the response includes presigned\ndownload URLs for each attached file; other integrations match the\nshape supplied at creation.",
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
													"files": schema.StringAttribute{
														MarkdownDescription: "One entry per attached file. A JSON document. Use `jsonencode()` to set it.",
														Computed:            true,
													},
													"external_url": schema.StringAttribute{
														MarkdownDescription: "Optional external URL, passed through from the entitlement\nconfiguration. A JSON document. Use `jsonencode()` to set it.",
														Computed:            true,
													},
													"instructions": schema.StringAttribute{
														MarkdownDescription: "Optional human-readable delivery instructions, passed through from\nthe entitlement configuration. A JSON document. Use `jsonencode()` to set it.",
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
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"is_recurring": schema.BoolAttribute{
							MarkdownDescription: "Indicates if the product is recurring (e.g., subscriptions).",
							Computed:            true,
						},
						"metadata": schema.MapAttribute{
							MarkdownDescription: "Additional custom data associated with the product.",
							Computed:            true,
							ElementType:         types.StringType,
						},
						"product_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for the product.",
							Computed:            true,
						},
						"tax_category": schema.StringAttribute{
							MarkdownDescription: "Tax category associated with the product.",
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp when the product was last updated.",
							Computed:            true,
						},
						"currency": schema.StringAttribute{
							MarkdownDescription: "Currency of the price.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the product, optional.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the product, optional.",
							Computed:            true,
						},
						"price": schema.Int64Attribute{
							MarkdownDescription: "Price of the product, optional.\n\nThe price is represented in the lowest denomination of the currency.\nFor example:\n- In USD, a price of `$12.34` would be represented as `1234` (cents).\n- In JPY, a price of `¥1500` would be represented as `1500` (yen).\n- In INR, a price of `₹1234.56` would be represented as `123456` (paise).\n\nThis ensures precision and avoids floating-point rounding errors.",
							Computed:            true,
						},
						"price_detail": schema.SingleNestedAttribute{
							MarkdownDescription: "Details of the price.",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"currency": schema.StringAttribute{
									MarkdownDescription: "The currency in which the payment is made.",
									Computed:            true,
								},
								"price": schema.Int64Attribute{
									MarkdownDescription: "The payment amount, in the smallest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.\n\nIf [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field represents\nthe **minimum** amount the customer must pay. Only for `type` `one_time_price` or `recurring_price`. Required when `type` is `one_time_price` or `recurring_price`.",
									Computed:            true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
								"discount": schema.Int64Attribute{
									MarkdownDescription: "Deprecated: use `discount_bps` instead.\n\nDiscount applied to the price, represented as a percentage (0 to 100).\nA response rounds this value to the nearest whole percent.\nDefaults to `0`.",
									Computed:            true,
								},
								"discount_bps": schema.Int64Attribute{
									MarkdownDescription: "Discount applied to the price, in basis points. 100 basis points make\none percent, so `1250` is a discount of 12.5%.\n\nUse this field for a discount with a fraction of a percent. A request\nthat sends this field ignores `discount`. A value of `0` gives no\ndiscount.",
									Computed:            true,
								},
								"pay_what_you_want": schema.BoolAttribute{
									MarkdownDescription: "Indicates whether the customer can pay any amount they choose.\nIf set to `true`, the [`price`](Self::price) field is the minimum amount. Only for `type` `one_time_price`.",
									Computed:            true,
								},
								"purchasing_power_parity": schema.BoolAttribute{
									MarkdownDescription: "Opts this price in to purchasing power parity. The business must also\nenable purchasing power parity. The discount percentage per country is\nalways business-wide. Defaults to `false`.",
									Computed:            true,
								},
								"suggested_price": schema.Int64Attribute{
									MarkdownDescription: "A suggested price for the user to pay. This value is only considered if\n[`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is ignored. Only for `type` `one_time_price`.",
									Computed:            true,
								},
								"tax_inclusive": schema.BoolAttribute{
									MarkdownDescription: "Indicates if the price is tax inclusive.",
									Computed:            true,
								},
								"payment_frequency_count": schema.Int64Attribute{
									MarkdownDescription: "Number of `payment_frequency_interval` units between two charges.\nThis field sets the billing cycle.\nFor example, `1` with a `payment_frequency_interval` of `month` charges the customer every month. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
									Computed:            true,
								},
								"payment_frequency_interval": schema.StringAttribute{
									MarkdownDescription: "The time unit of the billing cycle (e.g., day, month, year).\nUse it with `payment_frequency_count`. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
									Computed:            true,
								},
								"subscription_period_count": schema.Int64Attribute{
									MarkdownDescription: "Number of `subscription_period_interval` units for the total duration of the subscription.\nThis field does not set the billing cycle. Use `payment_frequency_count` for the billing cycle.\nWhen this period ends, the subscription expires. Dodo Payments does not start a new\nbilling cycle after the subscription expires. Dodo Payments can still retry a failed\npayment for an earlier billing cycle.\n\nFor a subscription that renews until the customer cancels it, set this to `20` with a\n`subscription_period_interval` of `year`. This is the recommended value and also the maximum.\nThe subscription expires after 20 years if the customer does not cancel it before then.\n\nFor example, a monthly plan uses a `payment_frequency_count` of `1`, a `payment_frequency_interval`\nof `month`, and a subscription period of `20` `year`. Do not set the subscription period to\n`1` `month` for a monthly plan: the subscription then expires after one month. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
									Computed:            true,
								},
								"subscription_period_interval": schema.StringAttribute{
									MarkdownDescription: "The time unit of the total subscription duration (e.g., day, month, year).\nUse it with `subscription_period_count`. The recommended value is `year`, with a\n`subscription_period_count` of `20`. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`.",
									Computed:            true,
								},
								"trial_amount": schema.Int64Attribute{
									MarkdownDescription: "Amount charged today for a paid trial, in the price currency's minor units.\nRequires `trial_period_days > 0`. Omit or null for a free trial (the default). Only for `type` `recurring_price`.",
									Computed:            true,
								},
								"trial_apply_discounts": schema.BoolAttribute{
									MarkdownDescription: "Whether discount codes reduce the trial charge. Defaults to false. Only meaningful\nwhen a paid trial is configured. Only for `type` `recurring_price`.",
									Computed:            true,
								},
								"trial_payment_method_optional": schema.BoolAttribute{
									MarkdownDescription: "Let a customer start a free trial with no card. Defaults to false. Only for `type` `recurring_price`.",
									Computed:            true,
								},
								"trial_period_days": schema.Int64Attribute{
									MarkdownDescription: "Number of days for the trial period. A value of `0` indicates no trial period. Only for `type` `recurring_price`.",
									Computed:            true,
								},
								"zero_amount_payment_method_optional": schema.BoolAttribute{
									MarkdownDescription: "Let a customer start a subscription with no card, when the amount due\ntoday is `0` (a native `0` price, or a 100% discount). Defaults to false. Only for `type` `recurring_price`.",
									Computed:            true,
								},
								"fixed_price": schema.Int64Attribute{
									MarkdownDescription: "The fixed payment amount. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`. Only for `type` `usage_based_price`. Required when `type` is `usage_based_price`.",
									Computed:            true,
								},
								"meters": schema.ListNestedAttribute{
									MarkdownDescription: "Only for `type` `usage_based_price`.",
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"meter_id": schema.StringAttribute{
												Computed: true,
											},
											"credit_entitlement_id": schema.StringAttribute{
												MarkdownDescription: "Optional credit entitlement ID to link this meter to for credit-based billing.",
												Computed:            true,
											},
											"description": schema.StringAttribute{
												MarkdownDescription: "Meter description. Will ignored on Request, but will be shown in response.",
												Computed:            true,
											},
											"free_threshold": schema.Int64Attribute{
												Computed: true,
											},
											"measurement_unit": schema.StringAttribute{
												MarkdownDescription: "Meter measurement unit. Will ignored on Request, but will be shown in response.",
												Computed:            true,
											},
											"meter_units_per_credit": schema.StringAttribute{
												MarkdownDescription: "Number of meter units that equal one credit. Required when credit_entitlement_id is set.",
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: "Meter name. Will ignored on Request, but will be shown in response.",
												Computed:            true,
											},
											"price_per_unit": schema.StringAttribute{
												MarkdownDescription: "The price per unit in lowest denomination. Must be greater than zero. Supports up to 5 digits before decimal point and 12 decimal places.",
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"pricing_mode": schema.StringAttribute{
							MarkdownDescription: "Pricing mode for localized pricing. NULL means base-only (no localized rules apply).",
							Computed:            true,
						},
						"tax_inclusive": schema.BoolAttribute{
							MarkdownDescription: "Indicates if the price is tax inclusive.",
							Computed:            true,
						},
					},
				},
			},
		},
	},
}
