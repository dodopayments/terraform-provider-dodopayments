// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addon

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

// NewResource returns the dodopayments_addon resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "addon",
	IDProperty:       "id",
	ParentParams:     []string{},
	CreateMethod:     "POST",
	CreatePath:       "addons",
	CreatePathParam:  "",
	ReadPath:         "addons/%s",
	UpdatePath:       "addons/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "addons/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"currency", "description", "name", "price", "tax_category"},
	UpdateFields:     []string{"currency", "description", "image_id", "name", "price", "tax_category"},
	UpdateOnlyFields: []string{"image_id"},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
		KeepPrior:   map[string]bool{"image_id": true},
		ListKeys:    map[string]string{},
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
				Description:       "The ID of the addon.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a addon.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "id of the Addon.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"business_id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the business to which the addon belongs.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Created time.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				MarkdownDescription: "The currency of the Addon.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP", "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW")},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description of the Addon. If you remove this attribute from the configuration, the API keeps the last value.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"image": schema.StringAttribute{
				MarkdownDescription: "Image of the Addon.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: "Addon image id after its uploaded to S3.\nPass `null` to remove the existing image, omit to keep it unchanged.",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Addon.",
				Required:            true,
			},
			"price": schema.Int64Attribute{
				MarkdownDescription: "Amount of the addon.",
				Required:            true,
			},
			"tax_category": schema.StringAttribute{
				MarkdownDescription: "Tax category applied to this Addon.",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("digital_products", "saas", "e_book", "edtech", "live_tutoring")},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Updated time.",
				Computed:            true,
			},
		},
	},
}
