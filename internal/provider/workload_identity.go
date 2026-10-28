// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package provider

import (
	"os"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/workloadidentity"
)

const (
	identityProviderIDEnv = "DODO_PAYMENTS_IDENTITY_PROVIDER_ID"
	identityTokenFileEnv  = "DODO_PAYMENTS_IDENTITY_TOKEN_FILE"
)

var tokenSources = []string{"auto", "github_actions", "gitlab", "terraform_cloud", "gcp", "vercel", "kubernetes", "file"}

func workloadIdentityBlock() schema.Block {
	return schema.SingleNestedBlock{
		MarkdownDescription: "Authenticates with workload identity federation instead of an API key. The provider exchanges a token from the CI system or the cloud platform for a short-lived access token. Do not set `api_key` together with this block. If you set neither, `DODO_PAYMENTS_API_KEY` is not set and `DODO_PAYMENTS_IDENTITY_PROVIDER_ID` is set, the provider uses workload identity with the `auto` token source.",
		Attributes: map[string]schema.Attribute{
			"identity_provider_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the workload identity provider (`wip_...`). You can also set the `DODO_PAYMENTS_IDENTITY_PROVIDER_ID` environment variable.",
				Optional:            true,
			},
			"token_source": schema.StringAttribute{
				MarkdownDescription: "The source of the subject token: `auto`, `github_actions`, `gitlab`, `terraform_cloud`, `gcp`, `vercel`, `kubernetes`, `file`. The default is `auto`. `auto` examines GitHub Actions, Terraform Cloud, `DODO_PAYMENTS_IDENTITY_TOKEN_FILE` and the Kubernetes token path, in that order.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf(tokenSources...)},
			},
			"token_file": schema.StringAttribute{
				MarkdownDescription: "The path of the subject token file for the `file` and `kubernetes` token sources. `file` needs this attribute or the `DODO_PAYMENTS_IDENTITY_TOKEN_FILE` environment variable. The default for `kubernetes` is `/var/run/secrets/tokens/dodo-token`.",
				Optional:            true,
			},
			"audience": schema.StringAttribute{
				MarkdownDescription: "The audience of the subject token for the `github_actions` and `gcp` token sources. The default is the origin of the API base URL.",
				Optional:            true,
			},
		},
	}
}

// authOption returns the client option that authenticates the provider: an API key or workload identity.
func authOption(cfg map[string]tftypes.Value) (option.RequestOption, diag.Diagnostics) {
	var diags diag.Diagnostics
	apiKey := configString(cfg, "api_key")
	if block, ok := workloadIdentityConfig(cfg); ok {
		if apiKey != "" {
			diags.AddAttributeError(path.Root("workload_identity"), "Conflicting authentication", "Set api_key or the workload_identity block, not both.")
			return nil, diags
		}
		return workloadIdentityOption(block)
	}
	if apiKey == "" {
		apiKey = os.Getenv("DODO_PAYMENTS_API_KEY")
	}
	if apiKey != "" {
		return option.WithBearerToken(apiKey), diags
	}
	if os.Getenv(identityProviderIDEnv) != "" {
		return workloadIdentityOption(nil)
	}
	diags.AddError("Missing API key", "Set the api_key provider attribute or the DODO_PAYMENTS_API_KEY environment variable. To use workload identity, set the workload_identity block or the DODO_PAYMENTS_IDENTITY_PROVIDER_ID environment variable.")
	return nil, diags
}

// workloadIdentityConfig returns the attributes of the workload_identity block. It returns false when the configuration has no block.
func workloadIdentityConfig(cfg map[string]tftypes.Value) (map[string]tftypes.Value, bool) {
	v, ok := cfg["workload_identity"]
	if !ok || !v.IsKnown() || v.IsNull() {
		return nil, false
	}
	var block map[string]tftypes.Value
	if err := v.As(&block); err != nil {
		return nil, false
	}
	return block, true
}

type workloadIdentityKey struct{ id, source, tokenFile, audience string }

// workloadIdentityOptions keeps one option for each configuration. Terraform configures the provider many times in
// one process, and each new option exchanges a new access token. The API sets a limit on the active tokens of a mapping.
var workloadIdentityOptions sync.Map

func workloadIdentityOption(block map[string]tftypes.Value) (option.RequestOption, diag.Diagnostics) {
	var diags diag.Diagnostics
	attr := func(name string) path.Path { return path.Root("workload_identity").AtName(name) }
	id := configString(block, "identity_provider_id")
	if id == "" {
		id = os.Getenv(identityProviderIDEnv)
	}
	if id == "" {
		diags.AddAttributeError(attr("identity_provider_id"), "Missing identity provider ID", "Set identity_provider_id or the "+identityProviderIDEnv+" environment variable.")
		return nil, diags
	}
	source := configString(block, "token_source")
	if source == "" {
		source = "auto"
	}
	tokenFile := configString(block, "token_file")
	audience := configString(block, "audience")
	if tokenFile != "" && source != "file" && source != "kubernetes" {
		diags.AddAttributeWarning(attr("token_file"), "Unused token_file", "The "+source+" token source does not use token_file.")
	}
	if audience != "" && source != "github_actions" && source != "gcp" {
		diags.AddAttributeWarning(attr("audience"), "Unused audience", "The "+source+" token source does not use audience.")
	}
	if source == "file" && tokenFile == "" {
		tokenFile = os.Getenv(identityTokenFileEnv)
		if tokenFile == "" {
			diags.AddAttributeError(attr("token_file"), "Missing token file", "The file token source needs token_file or the "+identityTokenFileEnv+" environment variable.")
			return nil, diags
		}
	}
	var provider option.SubjectTokenProvider
	switch source {
	case "auto":
		provider = workloadidentity.Auto()
	case "github_actions":
		provider = workloadidentity.GitHubActions(audience)
	case "gitlab":
		provider = workloadidentity.GitLab("")
	case "terraform_cloud":
		provider = workloadidentity.TerraformCloud()
	case "gcp":
		provider = workloadidentity.GCP(audience)
	case "vercel":
		provider = workloadidentity.Vercel()
	case "kubernetes":
		provider = workloadidentity.Kubernetes(tokenFile)
	case "file":
		provider = workloadidentity.File(tokenFile)
	default:
		diags.AddAttributeError(attr("token_source"), "Invalid token source", "The token source "+source+" is not known.")
		return nil, diags
	}
	key := workloadIdentityKey{id: id, source: source, tokenFile: tokenFile, audience: audience}
	if opt, ok := workloadIdentityOptions.Load(key); ok {
		return opt.(option.RequestOption), diags
	}
	opt, _ := workloadIdentityOptions.LoadOrStore(key, option.WithWorkloadIdentity(option.WorkloadIdentity{IdentityProviderID: id, Provider: provider}))
	return opt.(option.RequestOption), diags
}
