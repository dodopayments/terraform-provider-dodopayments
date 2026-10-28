// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package provider

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/list"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/addon"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/brand"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/creditentitlement"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/discount"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/entitlement"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/meter"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/product"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/productcollection"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/productlocalizedprice"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/webhook"
)

const environmentEnv = "DODO_PAYMENTS_ENVIRONMENT"

// defaultMaxRetries is the number of retries for a request that fails with HTTP 408, 409, 429 or 5xx, or with a connection error.
const defaultMaxRetries = 5

type Provider struct {
	version string
}

var (
	_ provider.Provider                       = (*Provider)(nil)
	_ provider.ProviderWithEphemeralResources = (*Provider)(nil)
	_ provider.ProviderWithActions            = (*Provider)(nil)
	_ provider.ProviderWithListResources      = (*Provider)(nil)
)

func New(version string) func() provider.Provider {
	return func() provider.Provider { return &Provider{version: version} }
}

func (p *Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dodopayments"
	resp.Version = p.version
}

func (p *Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages Dodo Payments resources.",
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "The API key. You can also set the `DODO_PAYMENTS_API_KEY` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"environment": schema.StringAttribute{
				MarkdownDescription: "The API environment: `live_mode` or `test_mode`. The default is `live_mode`. You can also set the `DODO_PAYMENTS_ENVIRONMENT` environment variable.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("live_mode", "test_mode")},
			},
			"base_url": schema.StringAttribute{
				MarkdownDescription: "A custom base URL for the API. This value overrides `environment`.",
				Optional:            true,
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of retries for a request that fails with HTTP 408, 409, 429 or 5xx, or with a connection error. The default is 5. On HTTP 429, the provider waits until the rate limit window resets.",
				Optional:            true,
				Validators:          []validator.Int64{int64validator.Between(0, 20)},
			},
		},
		Blocks: map[string]schema.Block{
			"workload_identity": workloadIdentityBlock(),
		},
	}
}

func configString(cfg map[string]tftypes.Value, name string) string {
	var s string
	if v, ok := cfg[name]; ok && v.IsKnown() && !v.IsNull() {
		_ = v.As(&s)
	}
	return s
}

func configInt(cfg map[string]tftypes.Value, name string, fallback int) int {
	var n big.Float
	if v, ok := cfg[name]; ok && v.IsKnown() && !v.IsNull() && v.As(&n) == nil {
		if i, acc := n.Int64(); acc == big.Exact {
			return int(i)
		}
	}
	return fallback
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg map[string]tftypes.Value
	if err := req.Config.Raw.As(&cfg); err != nil {
		resp.Diagnostics.AddError("Invalid provider configuration", err.Error())
		return
	}
	auth, diags := authOption(cfg)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	environment := configString(cfg, "environment")
	if environment == "" {
		environment = os.Getenv(environmentEnv)
	}

	opts := []option.RequestOption{
		auth,
		option.WithMaxRetries(configInt(cfg, "max_retries", defaultMaxRetries)),
		option.WithMiddleware(engine.RateLimitMiddleware),
		option.WithHeader("User-Agent", fmt.Sprintf("terraform-provider-dodopayments/%s Terraform/%s", p.version, req.TerraformVersion)),
	}
	switch environment {
	case "live_mode":
		opts = append(opts, option.WithBaseURL("https://live.dodopayments.com"))
	case "test_mode":
		opts = append(opts, option.WithBaseURL("https://test.dodopayments.com"))
	case "":
	default:
		resp.Diagnostics.AddError("Invalid environment", "The environment "+environment+" is not known.")
		return
	}
	if baseURL := configString(cfg, "base_url"); baseURL != "" {
		opts = append(opts, option.WithBaseURL(baseURL))
	}

	client := sdk.NewClient(opts...)
	resp.ResourceData = client
	resp.DataSourceData = client
	resp.EphemeralResourceData = client
	resp.ActionData = client
	resp.ListResourceData = client
}

func (p *Provider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		webhook.NewEphemeralWebhookSecret,
	}
}

func (p *Provider) Actions(_ context.Context) []func() action.Action {
	return []func() action.Action{
		product.NewProductUnarchiveAction,
		addon.NewAddonUnarchiveAction,
		meter.NewMeterUnarchiveAction,
		creditentitlement.NewCreditEntitlementUndeleteAction,
		productcollection.NewProductCollectionUnarchiveAction,
	}
}

func (p *Provider) ListResources(_ context.Context) []func() list.ListResource {
	return []func() list.ListResource{
		product.NewListResource,
		productlocalizedprice.NewListResource,
		discount.NewListResource,
		addon.NewListResource,
		brand.NewListResource,
		webhook.NewListResource,
		meter.NewListResource,
		creditentitlement.NewListResource,
		entitlement.NewListResource,
		productcollection.NewListResource,
	}
}

func (p *Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		product.NewResource,
		productlocalizedprice.NewResource,
		discount.NewResource,
		addon.NewResource,
		brand.NewResource,
		webhook.NewResource,
		meter.NewResource,
		creditentitlement.NewResource,
		entitlement.NewResource,
		productcollection.NewResource,
	}
}

func (p *Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		product.NewDataSource,
		product.NewListDataSource,
		productlocalizedprice.NewDataSource,
		productlocalizedprice.NewListDataSource,
		discount.NewDataSource,
		discount.NewListDataSource,
		addon.NewDataSource,
		addon.NewListDataSource,
		brand.NewDataSource,
		brand.NewListDataSource,
		webhook.NewDataSource,
		webhook.NewListDataSource,
		meter.NewDataSource,
		meter.NewListDataSource,
		creditentitlement.NewDataSource,
		creditentitlement.NewListDataSource,
		entitlement.NewDataSource,
		entitlement.NewListDataSource,
		productcollection.NewDataSource,
		productcollection.NewListDataSource,
	}
}
