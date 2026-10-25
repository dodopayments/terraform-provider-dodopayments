// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewResource returns the dodopayments_product resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "product",
	IDProperty:       "product_id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "products",
	CreatePathParam:  "",
	ReadPath:         "products/%s",
	UpdatePath:       "products/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "products/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"addons", "brand_id", "credit_entitlements", "description", "entitlements", "metadata", "name", "price", "pricing_mode", "tax_category"},
	UpdateFields:     []string{"addons", "brand_id", "credit_entitlements", "description", "entitlements", "metadata", "name", "price", "pricing_mode", "tax_category"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{"entitlements.integration_config.digital_files.external_url": true, "entitlements.integration_config.digital_files.files": true, "entitlements.integration_config.digital_files.instructions": true},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{"credit_entitlements": "credit_entitlement_id", "entitlements": "entitlement_id", "price.meters": "meter_id"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{"entitlements.entitlement_id": "id"},
		Decimals:    map[string]bool{"credit_entitlements.credits_amount": true, "credit_entitlements.overage_limit": true, "credit_entitlements.price_per_unit": true, "credit_entitlements.trial_credits": true, "price.meters.meter_units_per_credit": true, "price.meters.price_per_unit": true},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{"metadata": true},
	ReadAfterWrite:   true,
	ArchivedProperty: "",
	DestroyWarning:   "The API does not delete this object. When Terraform destroys the resource, the provider archives the object.",
	ReplaceFields:    []string{},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the product.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a product. The API does not delete this object. When Terraform destroys the resource, the provider archives the object.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the object.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"addons": schema.SetAttribute{
				MarkdownDescription: "Addons available for subscription product.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"brand_id": schema.StringAttribute{
				MarkdownDescription: "Brand id for the product, if not provided will default to primary brand. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the business to which the product belongs.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the product was created.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"credit_entitlements": schema.ListNestedAttribute{
				MarkdownDescription: "Optional credit entitlements to attach (max 5). If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"credit_entitlement_id": schema.StringAttribute{
							MarkdownDescription: "ID of the credit entitlement to attach.",
							Required:            true,
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
							MarkdownDescription: "Number of credits to grant when this product is purchased.",
							Required:            true,
							CustomType:          engine.DecimalStringType{},
						},
						"currency": schema.StringAttribute{
							MarkdownDescription: "Currency for credit-related pricing. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							Validators:          []validator.String{stringvalidator.OneOf("AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP", "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW")},
						},
						"expires_after_days": schema.Int64Attribute{
							MarkdownDescription: "Number of days after which credits expire. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique ID of this mapping.",
							Computed:            true,
						},
						"low_balance_threshold_percent": schema.Int64Attribute{
							MarkdownDescription: "Balance threshold percentage for low balance notifications (0-100). If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"max_rollover_count": schema.Int64Attribute{
							MarkdownDescription: "Maximum number of rollover cycles allowed. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"overage_behavior": schema.StringAttribute{
							MarkdownDescription: "Controls how overage is handled at billing cycle end. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							Validators:          []validator.String{stringvalidator.OneOf("forgive_at_reset", "invoice_at_billing", "carry_deficit", "carry_deficit_auto_repay")},
						},
						"overage_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether overage usage is allowed beyond credit balance. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"overage_limit": schema.StringAttribute{
							MarkdownDescription: "Maximum amount of overage allowed. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							CustomType:          engine.DecimalStringType{},
						},
						"price_per_unit": schema.StringAttribute{
							MarkdownDescription: "Price per credit unit for purchasing additional credits. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							CustomType:          engine.DecimalStringType{},
						},
						"proration_behavior": schema.StringAttribute{
							MarkdownDescription: "Proration behavior for credit grants during plan changes. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							Validators:          []validator.String{stringvalidator.OneOf("prorate", "no_prorate")},
						},
						"rollover_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether unused credits can roll over to the next billing period. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"rollover_percentage": schema.Int64Attribute{
							MarkdownDescription: "Percentage of unused credits that can roll over (0-100). If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"rollover_timeframe_count": schema.Int64Attribute{
							MarkdownDescription: "Number of timeframe units for rollover window. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"rollover_timeframe_interval": schema.StringAttribute{
							MarkdownDescription: "Time interval for rollover window (day, week, month, year). If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							Validators:          []validator.String{stringvalidator.OneOf("Day", "Week", "Month", "Year")},
						},
						"trial_credits": schema.StringAttribute{
							MarkdownDescription: "Credits granted during trial period. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
							CustomType:          engine.DecimalStringType{},
						},
						"trial_credits_expire_after_trial": schema.BoolAttribute{
							MarkdownDescription: "Whether trial credits expire when trial ends. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the product. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"entitlements": schema.ListNestedAttribute{
				MarkdownDescription: "Optional entitlements to attach to this product (max 50). If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Computed: true,
						},
						"entitlement_id": schema.StringAttribute{
							MarkdownDescription: "ID of the entitlement to attach to the product.",
							Required:            true,
						},
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
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"is_recurring": schema.BoolAttribute{
				MarkdownDescription: "Indicates if the product is recurring (e.g., subscriptions).",
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Additional metadata for the product.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the product.",
				Required:            true,
			},
			"price": schema.SingleNestedAttribute{
				MarkdownDescription: "Price configuration for the product.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						MarkdownDescription: "The currency in which the payment is made.",
						Required:            true,
						Validators:          []validator.String{stringvalidator.OneOf("AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP", "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW")},
					},
					"discount": schema.Int64Attribute{
						MarkdownDescription: "Deprecated: use `discount_bps` instead.\n\nDiscount applied to the price, represented as a percentage (0 to 100).\nA response rounds this value to the nearest whole percent.\nDefaults to `0`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"discount_bps": schema.Int64Attribute{
						MarkdownDescription: "Discount applied to the price, in basis points. 100 basis points make\none percent, so `1250` is a discount of 12.5%.\n\nUse this field for a discount with a fraction of a percent. A request\nthat sends this field ignores `discount`. A value of `0` gives no\ndiscount. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"fixed_price": schema.Int64Attribute{
						MarkdownDescription: "The fixed payment amount. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`. Only for `type` `usage_based_price`. Required when `type` is `usage_based_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"meters": schema.ListNestedAttribute{
						MarkdownDescription: "Only for `type` `usage_based_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"credit_entitlement_id": schema.StringAttribute{
									MarkdownDescription: "Optional credit entitlement ID to link this meter to for credit-based billing. If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
								},
								"description": schema.StringAttribute{
									MarkdownDescription: "Meter description. Will ignored on Request, but will be shown in response. If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
								},
								"free_threshold": schema.Int64Attribute{
									MarkdownDescription: "If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
								},
								"measurement_unit": schema.StringAttribute{
									MarkdownDescription: "Meter measurement unit. Will ignored on Request, but will be shown in response. If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
								},
								"meter_id": schema.StringAttribute{
									Required: true,
								},
								"meter_units_per_credit": schema.StringAttribute{
									MarkdownDescription: "Number of meter units that equal one credit. Required when credit_entitlement_id is set. If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
									CustomType:          engine.DecimalStringType{},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: "Meter name. Will ignored on Request, but will be shown in response. If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
								},
								"price_per_unit": schema.StringAttribute{
									MarkdownDescription: "The price per unit in lowest denomination. Must be greater than zero. Supports up to 5 digits before decimal point and 12 decimal places. If you remove this attribute from the configuration, the API keeps the last value.",
									Optional:            true,
									Computed:            true,
									CustomType:          engine.DecimalStringType{},
								},
							},
						},
					},
					"pay_what_you_want": schema.BoolAttribute{
						MarkdownDescription: "Indicates whether the customer can pay any amount they choose.\nIf set to `true`, the [`price`](Self::price) field is the minimum amount. Only for `type` `one_time_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"payment_frequency_count": schema.Int64Attribute{
						MarkdownDescription: "Number of `payment_frequency_interval` units between two charges.\nThis field sets the billing cycle.\nFor example, `1` with a `payment_frequency_interval` of `month` charges the customer every month. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"payment_frequency_interval": schema.StringAttribute{
						MarkdownDescription: "The time unit of the billing cycle (e.g., day, month, year).\nUse it with `payment_frequency_count`. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{stringvalidator.OneOf("Day", "Week", "Month", "Year")},
					},
					"price": schema.Int64Attribute{
						MarkdownDescription: "The payment amount, in the smallest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.\n\nIf [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field represents\nthe **minimum** amount the customer must pay. Only for `type` `one_time_price` or `recurring_price`. Required when `type` is `one_time_price` or `recurring_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"purchasing_power_parity": schema.BoolAttribute{
						MarkdownDescription: "Opts this price in to purchasing power parity. The business must also\nenable purchasing power parity. The discount percentage per country is\nalways business-wide. Defaults to `false`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"subscription_period_count": schema.Int64Attribute{
						MarkdownDescription: "Number of `subscription_period_interval` units for the total duration of the subscription.\nThis field does not set the billing cycle. Use `payment_frequency_count` for the billing cycle.\nWhen this period ends, the subscription expires. Dodo Payments does not start a new\nbilling cycle after the subscription expires. Dodo Payments can still retry a failed\npayment for an earlier billing cycle.\n\nFor a subscription that renews until the customer cancels it, set this to `20` with a\n`subscription_period_interval` of `year`. This is the recommended value and also the maximum.\nThe subscription expires after 20 years if the customer does not cancel it before then.\n\nFor example, a monthly plan uses a `payment_frequency_count` of `1`, a `payment_frequency_interval`\nof `month`, and a subscription period of `20` `year`. Do not set the subscription period to\n`1` `month` for a monthly plan: the subscription then expires after one month. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"subscription_period_interval": schema.StringAttribute{
						MarkdownDescription: "The time unit of the total subscription duration (e.g., day, month, year).\nUse it with `subscription_period_count`. The recommended value is `year`, with a\n`subscription_period_count` of `20`. Only for `type` `recurring_price` or `usage_based_price`. Required when `type` is `recurring_price` or `usage_based_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{stringvalidator.OneOf("Day", "Week", "Month", "Year")},
					},
					"suggested_price": schema.Int64Attribute{
						MarkdownDescription: "A suggested price for the user to pay. This value is only considered if\n[`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is ignored. Only for `type` `one_time_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"tax_inclusive": schema.BoolAttribute{
						MarkdownDescription: "Indicates if the price is tax inclusive. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"trial_amount": schema.Int64Attribute{
						MarkdownDescription: "Amount charged today for a paid trial, in the price currency's minor units.\nRequires `trial_period_days > 0`. Omit or null for a free trial (the default). Only for `type` `recurring_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"trial_apply_discounts": schema.BoolAttribute{
						MarkdownDescription: "Whether discount codes reduce the trial charge. Defaults to false. Only meaningful\nwhen a paid trial is configured. Only for `type` `recurring_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"trial_payment_method_optional": schema.BoolAttribute{
						MarkdownDescription: "Let a customer start a free trial with no card. Defaults to false. Only for `type` `recurring_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"trial_period_days": schema.Int64Attribute{
						MarkdownDescription: "Number of days for the trial period. A value of `0` indicates no trial period. Only for `type` `recurring_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
					"type": schema.StringAttribute{
						Required:   true,
						Validators: []validator.String{stringvalidator.OneOf("one_time_price", "recurring_price", "usage_based_price")},
					},
					"zero_amount_payment_method_optional": schema.BoolAttribute{
						MarkdownDescription: "Let a customer start a subscription with no card, when the amount due\ntoday is `0` (a native `0` price, or a 100% discount). Defaults to false. Only for `type` `recurring_price`. If you remove this attribute from the configuration, the API keeps the last value.",
						Optional:            true,
						Computed:            true,
					},
				},
			},
			"pricing_mode": schema.StringAttribute{
				MarkdownDescription: "Pricing mode for localized pricing. When set, rules from\n/products/{id}/localized-prices apply at checkout. NULL means base-only\n(existing behavior).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("by_currency", "by_country")},
			},
			"product_collection_id": schema.StringAttribute{
				MarkdownDescription: "The product collection ID this product belongs to, if any.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the product.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tax_category": schema.StringAttribute{
				MarkdownDescription: "Tax category applied to this product.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("digital_products", "saas", "e_book", "edtech", "live_tutoring")},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the product was last updated.",
				Computed:            true,
			},
		},
	},
}
