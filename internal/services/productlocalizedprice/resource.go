// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productlocalizedprice

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

// NewResource returns the dodopayments_product_localized_price resource.
func NewResource() resource.Resource {
	return engine.NewResource(resourceSpec)()
}

var resourceSpec = &engine.ResourceSpec{
	Name:             "product_localized_price",
	IDProperty:       "id",
	ParentParams:     []string{"product_id"},
	CreateMethod:     "POST",
	CreatePath:       "products/%s/localized-prices",
	CreatePathParam:  "",
	ReadPath:         "products/%s/localized-prices/%s",
	UpdatePath:       "products/%s/localized-prices/%s",
	UpdateMethod:     "PATCH",
	DeletePath:       "products/%s/localized-prices/%s",
	DeleteMethod:     "DELETE",
	CreateFields:     []string{"amount", "country_code", "currency"},
	UpdateFields:     []string{"amount"},
	UpdateOnlyFields: []string{},
	UpdateFull:       false,
	Rules: engine.Rules{
		JSONAttrs:   map[string]bool{},
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
	ReplaceFields:    []string{"country_code", "currency"},
	IdentitySchema: identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"product_id": identityschema.StringAttribute{
				Description:       "The ID of the product.",
				RequiredForImport: true,
			},
			"id": identityschema.StringAttribute{
				Description:       "The ID of the product localized price.",
				RequiredForImport: true,
			},
		},
	},
	Schema: schema.Schema{
		Version:             0,
		MarkdownDescription: "Manages a product localized price. The API does not delete this object. When Terraform destroys the resource, the provider archives the object.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for the localized price.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"amount": schema.Int64Attribute{
				MarkdownDescription: "Amount in the smallest currency unit (e.g., cents). Must be greater than zero.",
				Required:            true,
			},
			"country_code": schema.StringAttribute{
				MarkdownDescription: "Required when the product's pricing_mode is by_country; forbidden when by_currency. A change to this value destroys the resource and creates a new resource.",
				Optional:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:          []validator.String{stringvalidator.OneOf("AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "KH", "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "UM", "US", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW")},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the localized price was created.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				MarkdownDescription: "Currency to charge in. Must be a supported currency. A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:          []validator.String{stringvalidator.OneOf("AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP", "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN", "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW")},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Pricing mode of the rule: by_currency or by_country.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "Product this localized price belongs to. A change to this value destroys the resource and creates a new resource.",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Timestamp when the localized price was last updated.",
				Computed:            true,
			},
		},
	},
}
