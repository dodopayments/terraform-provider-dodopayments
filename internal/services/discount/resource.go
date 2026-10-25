// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package discount

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NewResource returns the dodopayments_discount resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "discount",
	IDProperty:       "discount_id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "discounts",
	CreatePathParam:  "",
	ReadPath:         "discounts/%s",
	UpdatePath:       "discounts/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "discounts/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"amount", "code", "currency_options", "customer_eligibility", "expires_at", "metadata", "name", "per_customer_usage_limit", "preserve_on_plan_change", "restricted_to", "starts_at", "subscription_cycles", "type", "usage_limit"},
	UpdateFields:     []string{"amount", "code", "currency_options", "customer_eligibility", "expires_at", "metadata", "name", "per_customer_usage_limit", "preserve_on_plan_change", "restricted_to", "starts_at", "subscription_cycles", "type", "usage_limit"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{"currency_options": "currency"},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{},
	ReadAfterWrite:   false,
	ArchivedProperty: "",
	DestroyWarning:   "",
	ReplaceFields:    []string{},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the discount.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a discount.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the object.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"amount": schema.Int64Attribute{
				MarkdownDescription: "The discount amount in **basis points** (e.g. `540` means `5.4%`, `10000` means `100%`).\n\nMust be at least 1.",
				Required:            true,
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "The business this discount belongs to.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"code": schema.StringAttribute{
				MarkdownDescription: "Optionally supply a code (will be uppercased).\n- Must be at least 3 characters if provided.\n- If omitted, a random 16-character code is generated. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the discount is created.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"currency_options": schema.ListNestedAttribute{
				MarkdownDescription: "Per-currency options (flat deduction / percentage cap + minimum subtotal).\nCheckout uses the row for the currency the buyer pays in. For any other\ncurrency it converts the default row. Required for `flat` codes (must\ninclude a resolvable default); optional per-currency caps for\n`percentage` codes. Per-row invariants are checked in\n`normalize_currency_options`, not via `#[validate(nested)]`. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"currency": schema.StringAttribute{
							MarkdownDescription: "The currency this option applies to. The row applies when the buyer pays\nin this currency.",
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP", "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW")},
						},
						"is_default": schema.BoolAttribute{
							MarkdownDescription: "Whether this row is the default to convert from when the buyer pays in a\ncurrency that has no row. At most one row per discount may be default. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"max_amount_possible": schema.Int64Attribute{
							MarkdownDescription: "The most this code discounts in this currency's subunits. For `flat` codes\nthis is the deduction; for `percentage` codes it is the max-discount cap.\nMust be > 0 if provided. If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
						"minimum_subtotal": schema.Int64Attribute{
							MarkdownDescription: "Eligible-cart threshold in this currency's subunits (0 = no minimum). If you remove this attribute from the configuration, the API keeps the last value.",
							Optional:            true,
							Computed:            true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"customer_eligibility": schema.StringAttribute{
				MarkdownDescription: "Who may redeem this discount code. Defaults to `any` (unrestricted).\n`specific` starts with zero attached customers (fails closed) until\ncustomers are attached via `POST /discounts/{id}/customers`. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:          []validator.String{stringvalidator.OneOf("any", "first_time", "existing", "specific")},
			},
			"discount_id": schema.StringAttribute{
				MarkdownDescription: "The unique discount ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"expires_at": schema.StringAttribute{
				MarkdownDescription: "When the discount expires, if ever.",
				Optional:            true,
			},
			"metadata": schema.MapAttribute{
				MarkdownDescription: "Additional metadata for the discount. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name for the Discount.",
				Optional:            true,
			},
			"per_customer_usage_limit": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of times a single customer may redeem this discount.\nMust be `<= usage_limit` when both are set.",
				Optional:            true,
			},
			"preserve_on_plan_change": schema.BoolAttribute{
				MarkdownDescription: "Whether this discount should be preserved when a subscription changes plans.\nDefault: false (discount is removed on plan change). If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"restricted_to": schema.SetAttribute{
				MarkdownDescription: "List of product IDs to restrict usage (if any). If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"starts_at": schema.StringAttribute{
				MarkdownDescription: "When the discount becomes active, if scheduled for the future.\nNULL = active immediately. Must be strictly before `expires_at` when both are set.",
				Optional:            true,
			},
			"subscription_cycles": schema.Int64Attribute{
				MarkdownDescription: "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
				Optional:            true,
			},
			"times_used": schema.Int64Attribute{
				MarkdownDescription: "How many times this discount has been used.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The discount type: `percentage` or `flat` (`flat_per_unit` stays blocked).",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("flat", "percentage")},
			},
			"usage_limit": schema.Int64Attribute{
				MarkdownDescription: "How many times this discount can be used (if any).\nMust be >= 1 if provided.",
				Optional:            true,
			},
		},
	},
}
