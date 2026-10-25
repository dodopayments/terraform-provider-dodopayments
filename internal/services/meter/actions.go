// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
)

// NewMeterUnarchiveAction returns the dodopayments_meter_unarchive action.
func NewMeterUnarchiveAction() action.Action {
	return engine.NewAction(actionMeterUnarchiveSpec)()
}

var actionMeterUnarchiveSpec = &engine.ActionSpec{
	Name:    "meter_unarchive",
	Method:  "POST",
	Path:    "meters/%s/unarchive",
	Params:  []string{"id"},
	Success: "POST request succeeded",
	Schema: schema.Schema{
		MarkdownDescription: "Sends `POST /meters/{id}/unarchive`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the meter.",
				Required:            true,
			},
		},
	},
}
