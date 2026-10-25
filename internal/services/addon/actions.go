// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addon

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// NewAddonUnarchiveAction returns the dodopayments_addon_unarchive action.
func NewAddonUnarchiveAction() action.Action {
	return engine.NewAction(actionAddonUnarchiveSpec)()
}

var actionAddonUnarchiveSpec = &engine.ActionSpec{
	Name:    "addon_unarchive",
	Method:  "POST",
	Path:    "addons/%s/unarchive",
	Params:  []string{"id"},
	Success: "POST request succeeded",
	Schema: schema.Schema{
		MarkdownDescription: "Sends `POST /addons/{id}/unarchive`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the addon.",
				Required:            true,
			},
		},
	},
}
