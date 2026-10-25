// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/addon"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/discount"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/meter"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/services/webhook"
)

const environmentEnv = "DODO_PAYMENTS_ENVIRONMENT"

type Provider struct {
	version string
}

var _ provider.Provider = (*Provider)(nil)

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

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var cfg map[string]tftypes.Value
	if err := req.Config.Raw.As(&cfg); err != nil {
		resp.Diagnostics.AddError("Invalid provider configuration", err.Error())
		return
	}
	apiKey := configString(cfg, "api_key")
	if apiKey == "" {
		apiKey = os.Getenv("DODO_PAYMENTS_API_KEY")
	}
	if apiKey == "" {
		resp.Diagnostics.AddError("Missing API key", "Set the api_key provider attribute or the DODO_PAYMENTS_API_KEY environment variable.")
		return
	}
	environment := configString(cfg, "environment")
	if environment == "" {
		environment = os.Getenv(environmentEnv)
	}

	opts := []option.RequestOption{option.WithBearerToken(apiKey)}
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
}

func (p *Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		discount.NewResource,
		addon.NewResource,
		webhook.NewResource,
		meter.NewResource,
	}
}

func (p *Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		discount.NewDataSource,
		discount.NewListDataSource,
		addon.NewDataSource,
		addon.NewListDataSource,
		webhook.NewDataSource,
		webhook.NewListDataSource,
		meter.NewDataSource,
		meter.NewListDataSource,
	}
}
