// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
)

// NewEphemeralWebhookSecret returns the dodopayments_webhook_secret ephemeral resource.
func NewEphemeralWebhookSecret() ephemeral.EphemeralResource {
	return engine.NewEphemeralResource(ephemeralWebhookSecretSpec)()
}

var ephemeralWebhookSecretSpec = &engine.EphemeralSpec{
	Name:      "webhook_secret",
	Path:      "webhooks/%s/secret",
	Params:    []string{"webhook_id"},
	JSONAttrs: map[string]bool{},
	Schema: schema.Schema{
		MarkdownDescription: "Get webhook secret by id Terraform reads this value for each run and does not write it to the state or the plan. Use it in write-only arguments and in provider configuration.",
		Attributes: map[string]schema.Attribute{
			"webhook_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the webhook.",
				Required:            true,
			},
			"secret": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
			},
		},
	},
}
