// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
)

// ResourceSpec describes one managed resource. Paths with an ID contain one %s verb.
type ResourceSpec struct {
	Name         string
	Schema       schema.Schema
	IDProperty   string
	CreateMethod string
	CreatePath   string
	// CreatePathParam is the attribute that fills the %s verb in CreatePath. It is empty when CreatePath has no verb.
	CreatePathParam string
	ReadPath        string
	UpdatePath      string
	UpdateMethod    string
	DeletePath      string
	DeleteMethod    string
	// CreateFields and UpdateFields are the top-level attributes that the API accepts.
	CreateFields []string
	UpdateFields []string
	// UpdateOnlyFields are in UpdateFields but not in CreateFields. Create sends them with a second request.
	UpdateOnlyFields []string
	// UpdateFull sends every update field (PUT). Otherwise update sends only the changed fields (JSON merge patch).
	UpdateFull bool
	// JSONAttrs are the attribute paths that hold a JSON document as a string.
	JSONAttrs map[string]bool
	// WriteOnly attributes are sent to the API but the API does not return them as sent.
	WriteOnly []string
	// ReadAfterWrite makes the provider read the resource again after a create or an update.
	ReadAfterWrite bool
}

type Resource struct {
	spec   *ResourceSpec
	client *sdk.Client
}

var (
	_ resource.ResourceWithConfigure   = (*Resource)(nil)
	_ resource.ResourceWithImportState = (*Resource)(nil)
)

func NewResource(spec *ResourceSpec) func() resource.Resource {
	return func() resource.Resource { return &Resource{spec: spec} }
}

func (r *Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.spec.Name
}

func (r *Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.spec.Schema
}

func (r *Resource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*sdk.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", fmt.Sprintf("Expected *sdk.Client, got %T.", req.ProviderData))
		return
	}
	r.client = client
}

func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Resource) itemPath(template, id string) string {
	return fmt.Sprintf(template, url.PathEscape(id))
}

func (r *Resource) get(ctx context.Context, id string) (any, int, error) {
	return call(ctx, r.client, http.MethodGet, r.itemPath(r.spec.ReadPath, id), nil)
}

// buildBody returns the request body. With a nil prior value it sends every known, non-null field.
// With a prior value it sends a JSON null for a removed field. It sends only the changed fields, unless the spec has UpdateFull.
func (r *Resource) buildBody(plan tftypes.Value, fields []string, prior *tftypes.Value) (map[string]any, error) {
	var planAttrs, priorAttrs map[string]tftypes.Value
	if err := plan.As(&planAttrs); err != nil {
		return nil, err
	}
	if prior != nil {
		if err := prior.As(&priorAttrs); err != nil {
			return nil, err
		}
	}
	body := map[string]any{}
	for _, name := range fields {
		v, ok := planAttrs[name]
		if !ok {
			continue
		}
		if prior != nil && !r.spec.UpdateFull {
			if old, ok := priorAttrs[name]; ok && old.Equal(v) {
				continue
			}
		}
		j, err := toJSON(v, name, r.spec.JSONAttrs, prior != nil)
		if err != nil {
			return nil, err
		}
		if j == omit {
			continue
		}
		body[name] = j
	}
	return body, nil
}

// toState converts an API response to state. It keeps the prior value of write-only attributes, and
// keeps a null for an attribute that the configuration does not set when the API returns an empty value.
// It uses id when the response does not contain the ID.
func (r *Resource) toState(ctx context.Context, out any, prior tftypes.Value, id string) (tftypes.Value, error) {
	typ := r.spec.Schema.Type().TerraformType(ctx)
	state, err := fromJSON(typ, out, "", r.spec.JSONAttrs)
	if err != nil {
		return state, err
	}
	var stateAttrs, priorAttrs map[string]tftypes.Value
	if err := state.As(&stateAttrs); err != nil {
		return state, err
	}
	if found := idOf(out, r.spec.IDProperty); found != "" {
		id = found
	}
	if id != "" {
		stateAttrs["id"] = tftypes.NewValue(tftypes.String, id)
		// The create path key fills the same path segment as the ID, so it has the same value.
		if key := r.spec.CreatePathParam; key != "" {
			if v, ok := stateAttrs[key]; !ok || v.IsNull() {
				stateAttrs[key] = tftypes.NewValue(tftypes.String, id)
			}
		}
	}
	if prior.IsNull() || !prior.IsKnown() {
		return tftypes.NewValue(typ, stateAttrs), nil
	}
	if err := prior.As(&priorAttrs); err != nil {
		return state, err
	}
	for name, v := range stateAttrs {
		if p, ok := priorAttrs[name]; ok {
			stateAttrs[name] = keepEquivalentJSON(p, v, name, r.spec.JSONAttrs)
		}
	}
	for _, name := range r.spec.WriteOnly {
		if v, ok := priorAttrs[name]; ok && v.IsKnown() {
			stateAttrs[name] = v
		}
	}
	for name, v := range priorAttrs {
		attr, ok := r.spec.Schema.Attributes[name]
		if ok && !attr.IsComputed() && v.IsNull() && isEmptyValue(stateAttrs[name]) {
			stateAttrs[name] = v
		}
	}
	return tftypes.NewValue(typ, stateAttrs), nil
}

func stringAttr(v tftypes.Value, name string) string {
	var attrs map[string]tftypes.Value
	if err := v.As(&attrs); err != nil {
		return ""
	}
	var s string
	if a, ok := attrs[name]; ok && a.IsKnown() && !a.IsNull() {
		_ = a.As(&s)
	}
	return s
}

func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	body, err := r.buildBody(req.Plan.Raw, r.spec.CreateFields, nil)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
		return
	}
	createPath, key := r.spec.CreatePath, ""
	if r.spec.CreatePathParam != "" {
		key = stringAttr(req.Plan.Raw, r.spec.CreatePathParam)
		createPath = r.itemPath(r.spec.CreatePath, key)
	}
	out, _, err := call(ctx, r.client, r.spec.CreateMethod, createPath, body)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to create %s", r.spec.Name), err.Error())
		return
	}
	id := idOf(out, r.spec.IDProperty)
	if id == "" {
		id = key
	}
	if id == "" {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to create %s", r.spec.Name), fmt.Sprintf("The API response has no %q property.", r.spec.IDProperty))
		return
	}
	// Save the ID first, so that Terraform can destroy the object if a later step fails.
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
	followUp := r.spec.UpdateOnlyFields
	if r.spec.UpdateFull {
		followUp = r.spec.UpdateFields
	}
	updateOnly := map[string]any{}
	if len(r.spec.UpdateOnlyFields) > 0 {
		updateOnly, err = r.buildBody(req.Plan.Raw, followUp, nil)
	}
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
		return
	}
	if len(updateOnly) > 0 {
		if out, _, err = call(ctx, r.client, r.spec.UpdateMethod, r.itemPath(r.spec.UpdatePath, id), updateOnly); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to update %s %s after create", r.spec.Name, id), err.Error())
			return
		}
	}
	if out == nil || len(updateOnly) > 0 || r.spec.ReadAfterWrite {
		if out, _, err = r.get(ctx, id); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to read %s %s after create", r.spec.Name, id), err.Error())
			return
		}
	}
	state, err := r.toState(ctx, out, req.Plan.Raw, id)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	resp.State.Raw = state
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	id := stringAttr(req.State.Raw, "id")
	out, status, err := r.get(ctx, id)
	if isGone(status) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to read %s %s", r.spec.Name, id), err.Error())
		return
	}
	state, err := r.toState(ctx, out, req.State.Raw, id)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	resp.State.Raw = state
}

func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	id := stringAttr(req.State.Raw, "id")
	if r.spec.UpdatePath == "" {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to update %s %s", r.spec.Name, id), "The API has no update endpoint for this resource.")
		return
	}
	prior := req.State.Raw
	body, err := r.buildBody(req.Plan.Raw, r.spec.UpdateFields, &prior)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
		return
	}
	var out any
	if len(body) > 0 {
		if out, _, err = call(ctx, r.client, r.spec.UpdateMethod, r.itemPath(r.spec.UpdatePath, id), body); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to update %s %s", r.spec.Name, id), err.Error())
			return
		}
	}
	if out == nil || r.spec.ReadAfterWrite {
		if out, _, err = r.get(ctx, id); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to read %s %s after update", r.spec.Name, id), err.Error())
			return
		}
	}
	state, err := r.toState(ctx, out, req.Plan.Raw, id)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	resp.State.Raw = state
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	id := stringAttr(req.State.Raw, "id")
	if r.spec.DeletePath == "" {
		resp.Diagnostics.AddWarning(fmt.Sprintf("%s %s was not deleted", r.spec.Name, id), "The API has no delete endpoint for this resource. Terraform removed it from the state only.")
		return
	}
	_, status, err := call(ctx, r.client, r.spec.DeleteMethod, r.itemPath(r.spec.DeletePath, id), nil)
	if err != nil && !isGone(status) {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to delete %s %s", r.spec.Name, id), err.Error())
	}
}
