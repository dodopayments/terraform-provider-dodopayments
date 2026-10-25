// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package creditentitlement

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NewResource returns the dodopayments_credit_entitlement resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "credit_entitlement",
	IDProperty:       "id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "credit-entitlements",
	CreatePathParam:  "",
	ReadPath:         "credit-entitlements/%s",
	UpdatePath:       "credit-entitlements/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "credit-entitlements/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"currency", "description", "expires_after_days", "max_rollover_count", "name", "overage_behavior", "overage_enabled", "overage_limit", "precision", "price_per_unit", "rollover_enabled", "rollover_percentage", "rollover_timeframe_count", "rollover_timeframe_interval", "unit"},
	UpdateFields:     []string{"currency", "description", "expires_after_days", "max_rollover_count", "name", "overage_behavior", "overage_enabled", "overage_limit", "price_per_unit", "rollover_enabled", "rollover_percentage", "rollover_timeframe_count", "rollover_timeframe_interval", "unit"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{},
		ListKeys:    map[string]string{},
		StripPrefix: map[string]string{},
		Aliases:     map[string]string{},
		Decimals:    map[string]bool{"price_per_unit": true},
	},
	WriteOnly:        []string{},
	MergeMaps:        map[string]bool{},
	ReadAfterWrite:   false,
	ArchivedProperty: "",
	DestroyWarning:   "",
	ReplaceFields:    []string{"precision"},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				Description:       "The ID of the credit entitlement.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a credit entitlement.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"business_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				MarkdownDescription: "Currency for pricing (required if price_per_unit is set).",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP", "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW")},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the credit entitlement.",
				Optional:            true,
			},
			"expires_after_days": schema.Int64Attribute{
				MarkdownDescription: "Number of days after which credits expire (optional).",
				Optional:            true,
			},
			"max_rollover_count": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of times credits can be rolled over.",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the credit entitlement.",
				Required:            true,
			},
			"overage_behavior": schema.StringAttribute{
				MarkdownDescription: "Controls how overage is handled at billing cycle end.\nDefaults to forgive_at_reset if not specified. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:          []validator.String{stringvalidator.OneOf("forgive_at_reset", "invoice_at_billing", "carry_deficit", "carry_deficit_auto_repay")},
			},
			"overage_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether overage charges are enabled when credits run out.",
				Required:            true,
			},
			"overage_limit": schema.Int64Attribute{
				MarkdownDescription: "Maximum overage units allowed (optional).",
				Optional:            true,
			},
			"precision": schema.Int64Attribute{
				MarkdownDescription: "Precision for credit amounts (0-10 decimal places). A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"price_per_unit": schema.StringAttribute{
				MarkdownDescription: "Price per credit unit.",
				Optional:            true,
				CustomType:          engine.DecimalStringType{},
			},
			"rollover_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether rollover is enabled for unused credits.",
				Required:            true,
			},
			"rollover_percentage": schema.Int64Attribute{
				MarkdownDescription: "Percentage of unused credits that can rollover (0-100).",
				Optional:            true,
			},
			"rollover_timeframe_count": schema.Int64Attribute{
				MarkdownDescription: "Count of timeframe periods for rollover limit.",
				Optional:            true,
			},
			"rollover_timeframe_interval": schema.StringAttribute{
				MarkdownDescription: "Interval type for rollover timeframe.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("Day", "Week", "Month", "Year")},
			},
			"unit": schema.StringAttribute{
				MarkdownDescription: "Unit of measurement for the credit (e.g., \"API Calls\", \"Tokens\", \"Credits\").",
				Required:            true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	},
}
