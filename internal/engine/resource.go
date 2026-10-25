// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
)

// DecimalStringType is the type of a string attribute that holds a decimal number. Terraform treats "10.50" and "10.5"
// as equal, so a configuration and an API value in another form do not cause a change in the plan.
type DecimalStringType struct {
	basetypes.StringType
}

var _ basetypes.StringTypable = DecimalStringType{}

func (t DecimalStringType) Equal(o attr.Type) bool {
	_, ok := o.(DecimalStringType)
	return ok
}

func (t DecimalStringType) String() string {
	return "DecimalStringType"
}

func (t DecimalStringType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return DecimalString{StringValue: in}, nil
}

func (t DecimalStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	v, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	s, ok := v.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	return DecimalString{StringValue: s}, nil
}

func (t DecimalStringType) ValueType(context.Context) attr.Value {
	return DecimalString{}
}

// DecimalString is a value of DecimalStringType.
type DecimalString struct {
	basetypes.StringValue
}

var _ basetypes.StringValuableWithSemanticEquals = DecimalString{}

func (v DecimalString) Type(context.Context) attr.Type {
	return DecimalStringType{}
}

func (v DecimalString) Equal(o attr.Value) bool {
	other, ok := o.(DecimalString)
	return ok && v.StringValue.Equal(other.StringValue)
}

func (v DecimalString) StringSemanticEquals(_ context.Context, other basetypes.StringValuable) (bool, diag.Diagnostics) {
	o, ok := other.(DecimalString)
	if !ok {
		return false, nil
	}
	return v.ValueString() == o.ValueString() || decimalEqual(v.ValueString(), o.ValueString()), nil
}

// ResourceSpec describes one managed resource. Each path parameter in a path is a %s verb: the parent IDs, then the ID.
type ResourceSpec struct {
	Name           string
	Schema         schema.Schema
	IdentitySchema identityschema.Schema
	IDProperty     string
	// ParentParams are the attributes that fill the path verbs before the ID.
	ParentParams []string
	CreateMethod string
	CreatePath   string
	// CreatePathParam is the attribute that fills the last %s verb in CreatePath. It is empty when CreatePath has no such verb.
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
	Rules      Rules
	// WriteOnly attributes are sent to the API but the API does not return them as sent.
	WriteOnly []string
	// MergeMaps are map attributes that an update merges into the stored map. An empty string removes a key.
	MergeMaps map[string]bool
	// ReadAfterWrite makes the provider read the resource again after a create or an update.
	ReadAfterWrite bool
	// ArchivedProperty is a response property that is set when the object is archived. Such an object is gone.
	ArchivedProperty string
	// DestroyWarning is the warning for a plan that destroys or replaces the object.
	DestroyWarning string
	// ReplaceFields are the top-level attributes that destroy and create the object again when they change.
	ReplaceFields []string
}

// RequiresReplaceIfChanged forces a replacement when the plan changes a nested value of the attribute.
// A computed child that the configuration does not set does not force a replacement.
func RequiresReplaceIfChanged() replaceIfChanged {
	return replaceIfChanged{}
}

type replaceIfChanged struct{}

func (replaceIfChanged) Description(context.Context) string {
	return "A change to the configured value destroys the resource and creates a new resource."
}

func (m replaceIfChanged) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (replaceIfChanged) decide(ctx context.Context, stateRaw, planRaw tftypes.Value, cfg, plan, state attr.Value) bool {
	if stateRaw.IsNull() || planRaw.IsNull() {
		return false
	}
	c, errC := cfg.ToTerraformValue(ctx)
	p, errP := plan.ToTerraformValue(ctx)
	s, errS := state.ToTerraformValue(ctx)
	if errC != nil || errP != nil || errS != nil {
		return true
	}
	return valueChanged(c, p, s)
}

func (m replaceIfChanged) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	resp.RequiresReplace = m.decide(ctx, req.State.Raw, req.Plan.Raw, req.ConfigValue, req.PlanValue, req.StateValue)
}

func (m replaceIfChanged) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	resp.RequiresReplace = m.decide(ctx, req.State.Raw, req.Plan.Raw, req.ConfigValue, req.PlanValue, req.StateValue)
}

func (m replaceIfChanged) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	resp.RequiresReplace = m.decide(ctx, req.State.Raw, req.Plan.Raw, req.ConfigValue, req.PlanValue, req.StateValue)
}

// valueChanged reports whether plan differs from state. An unknown plan value is a change only when the configuration
// sets it: otherwise the API computes it and it keeps the state value.
func valueChanged(cfg, plan, state tftypes.Value) bool {
	if !plan.IsKnown() {
		return !cfg.IsFullyKnown()
	}
	if plan.Equal(state) {
		return false
	}
	if plan.IsNull() || state.IsNull() || !state.IsKnown() {
		return true
	}
	switch plan.Type().(type) {
	case tftypes.Object:
		var c, p, s map[string]tftypes.Value
		if plan.As(&p) != nil || state.As(&s) != nil {
			return true
		}
		if cfg.IsKnown() && !cfg.IsNull() {
			_ = cfg.As(&c)
		}
		for k, pv := range p {
			cv, ok := c[k]
			if !ok {
				cv = tftypes.NewValue(pv.Type(), nil)
			}
			sv, ok := s[k]
			if !ok {
				sv = tftypes.NewValue(pv.Type(), nil)
			}
			if valueChanged(cv, pv, sv) {
				return true
			}
		}
		return false
	case tftypes.Set:
		// A set has no order. It is unchanged when each plan element matches a different state element.
		// An unknown value that the configuration sets, for example a changed reference, is a change.
		var p, s []tftypes.Value
		if plan.As(&p) != nil || state.As(&s) != nil || len(p) != len(s) || !cfg.IsFullyKnown() {
			return true
		}
		used := make([]bool, len(s))
		for _, pv := range p {
			found := false
			for j, sv := range s {
				if !used[j] && !valueChanged(tftypes.NewValue(pv.Type(), nil), pv, sv) {
					used[j], found = true, true
					break
				}
			}
			if !found {
				return true
			}
		}
		return false
	case tftypes.List:
		var c, p, s []tftypes.Value
		if plan.As(&p) != nil || state.As(&s) != nil || len(p) != len(s) {
			return true
		}
		if cfg.IsKnown() && !cfg.IsNull() {
			_ = cfg.As(&c)
		}
		for i := range p {
			cv := tftypes.NewValue(p[i].Type(), nil)
			if i < len(c) {
				cv = c[i]
			}
			if valueChanged(cv, p[i], s[i]) {
				return true
			}
		}
		return false
	}
	return true
}

type Resource struct {
	spec   *ResourceSpec
	client *sdk.Client
}

var (
	_ resource.ResourceWithConfigure   = (*Resource)(nil)
	_ resource.ResourceWithImportState = (*Resource)(nil)
	_ resource.ResourceWithIdentity    = (*Resource)(nil)
	_ resource.ResourceWithModifyPlan  = (*Resource)(nil)
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

func (r *Resource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = r.spec.IdentitySchema
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

// importNames are the identity attributes, in the order of the import ID parts.
func (r *Resource) importNames() []string {
	return append(append([]string{}, r.spec.ParentParams...), "id")
}

func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	names := r.importNames()
	if req.ID == "" {
		for _, name := range names {
			var v types.String
			resp.Diagnostics.Append(req.Identity.GetAttribute(ctx, path.Root(name), &v)...)
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(name), v)...)
		}
		return
	}
	parts := []string{req.ID}
	if len(names) > 1 {
		parts = strings.Split(req.ID, "/")
	}
	if len(parts) != len(names) {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("The import ID %q is not valid. Use the format %q.", req.ID, strings.Join(names, "/")))
		return
	}
	for i, name := range names {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(name), parts[i])...)
		if resp.Identity != nil {
			resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root(name), parts[i])...)
		}
	}
}

func (r *Resource) setIdentity(ctx context.Context, identity *tfsdk.ResourceIdentity, state tftypes.Value) diag.Diagnostics {
	var diags diag.Diagnostics
	if identity == nil {
		return diags
	}
	for _, name := range r.importNames() {
		diags.Append(identity.SetAttribute(ctx, path.Root(name), stringAttr(state, name))...)
	}
	return diags
}

// ModifyPlan warns when a plan destroys or replaces an object whose delete does not work like a delete, for example an archive.
func (r *Resource) ModifyPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if r.spec.DestroyWarning == "" || req.State.Raw.IsNull() {
		return
	}
	id := stringAttr(req.State.Raw, "id")
	if req.Plan.Raw.IsNull() {
		resp.Diagnostics.AddWarning(fmt.Sprintf("Terraform will destroy %s %s", r.spec.Name, id), r.spec.DestroyWarning)
		return
	}
	var plan, state map[string]tftypes.Value
	if req.Plan.Raw.As(&plan) != nil || req.State.Raw.As(&state) != nil {
		return
	}
	var changed []string
	for _, name := range r.spec.ReplaceFields {
		if p, ok := plan[name]; ok && !p.Equal(state[name]) {
			changed = append(changed, name)
		}
	}
	if len(changed) > 0 {
		resp.Diagnostics.AddWarning(
			fmt.Sprintf("Terraform will replace %s %s", r.spec.Name, id),
			fmt.Sprintf("A change to %s needs a new object. Terraform creates a new object and destroys the old object. %s", strings.Join(changed, ", "), r.spec.DestroyWarning),
		)
	}
}

func (r *Resource) path(template string, v tftypes.Value, id string) string {
	return fmt.Sprintf(template, pathArgs(template, v, r.spec.ParentParams, id)...)
}

func (r *Resource) get(ctx context.Context, v tftypes.Value, id string) (any, int, error) {
	return call(ctx, r.client, http.MethodGet, r.path(r.spec.ReadPath, v, id), nil)
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
		j, err := toJSON(v, name, r.spec.Rules.JSONAttrs, prior != nil)
		if err != nil {
			return nil, err
		}
		if prior != nil && r.spec.MergeMaps[name] {
			j = mergeMapBody(priorAttrs[name], j)
		}
		if j == omit {
			continue
		}
		body[name] = j
	}
	return body, nil
}

// mergeMapBody adds an empty string for each key that the new map removed, because the API merges the map.
func mergeMapBody(old tftypes.Value, next any) any {
	if next == omit {
		return next
	}
	out, _ := next.(map[string]any)
	if out == nil {
		out = map[string]any{}
	}
	var oldMap map[string]tftypes.Value
	if old.IsKnown() && !old.IsNull() && old.As(&oldMap) == nil {
		for k := range oldMap {
			if _, ok := out[k]; !ok {
				out[k] = ""
			}
		}
	}
	return out
}

// toState converts an API response to state. It reconciles each value with the prior value, keeps the prior value
// of write-only attributes, and keeps a null for an attribute that the configuration does not set when the API
// returns an empty value. It uses id when the response does not contain the ID.
func (r *Resource) toState(ctx context.Context, out any, prior tftypes.Value, id string) (tftypes.Value, error) {
	typ := r.spec.Schema.Type().TerraformType(ctx)
	r.spec.Rules.applyAliases(out)
	state, err := fromJSON(typ, out, "", r.spec.Rules.JSONAttrs)
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
	if prior.IsKnown() && !prior.IsNull() {
		if err := prior.As(&priorAttrs); err != nil {
			return state, err
		}
	}
	for _, name := range r.spec.ParentParams {
		if v, ok := stateAttrs[name]; !ok || v.IsNull() {
			if p, ok := priorAttrs[name]; ok {
				stateAttrs[name] = p
			}
		}
	}
	for name, v := range stateAttrs {
		p, ok := priorAttrs[name]
		if !ok {
			p = tftypes.NewValue(v.Type(), nil)
		}
		stateAttrs[name] = r.spec.Rules.reconcile(p, v, name)
	}
	for _, name := range r.spec.WriteOnly {
		if v, ok := priorAttrs[name]; ok && v.IsKnown() {
			stateAttrs[name] = v
		}
	}
	// After an import, the prior state has only the IDs. Then each API value goes to the state.
	if !r.onlyIDs(priorAttrs) {
		for name, v := range priorAttrs {
			attr, ok := r.spec.Schema.Attributes[name]
			if ok && !attr.IsComputed() && v.IsNull() && isEmptyValue(stateAttrs[name]) {
				stateAttrs[name] = v
			}
		}
	}
	return tftypes.NewValue(typ, stateAttrs), nil
}

// onlyIDs reports whether each attribute other than the ID, the parent IDs and the create path key is null.
func (r *Resource) onlyIDs(attrs map[string]tftypes.Value) bool {
	ids := map[string]bool{"id": true, r.spec.CreatePathParam: true}
	for _, p := range r.spec.ParentParams {
		ids[p] = true
	}
	for name, v := range attrs {
		if !ids[name] && !v.IsNull() {
			return false
		}
	}
	return true
}

func (r *Resource) isArchived(out any) bool {
	if r.spec.ArchivedProperty == "" {
		return false
	}
	m, ok := out.(map[string]any)
	return ok && m[r.spec.ArchivedProperty] != nil
}

func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	body, err := r.buildBody(req.Plan.Raw, r.spec.CreateFields, nil)
	if err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
		return
	}
	key := ""
	if r.spec.CreatePathParam != "" {
		key = stringAttr(req.Plan.Raw, r.spec.CreatePathParam)
	}
	out, status, err := call(ctx, r.client, r.spec.CreateMethod, r.path(r.spec.CreatePath, req.Plan.Raw, key), body)
	if err != nil {
		detail := err.Error()
		if status == http.StatusConflict {
			detail += "\n\nAn object with the same unique value exists. To manage it with Terraform, import it with an import block or the terraform import command."
		}
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to create %s", r.spec.Name), detail)
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
		if out, _, err = call(ctx, r.client, r.spec.UpdateMethod, r.path(r.spec.UpdatePath, req.Plan.Raw, id), updateOnly); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to update %s %s after create", r.spec.Name, id), err.Error())
			return
		}
	}
	if out == nil || len(updateOnly) > 0 || r.spec.ReadAfterWrite {
		if out, _, err = r.get(ctx, req.Plan.Raw, id); err != nil {
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
	resp.Diagnostics.Append(r.setIdentity(ctx, resp.Identity, state)...)
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	id := stringAttr(req.State.Raw, "id")
	out, status, err := r.get(ctx, req.State.Raw, id)
	if isGone(status) || (err == nil && r.isArchived(out)) {
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
	resp.Diagnostics.Append(r.setIdentity(ctx, resp.Identity, state)...)
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
		if out, _, err = call(ctx, r.client, r.spec.UpdateMethod, r.path(r.spec.UpdatePath, req.Plan.Raw, id), body); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to update %s %s", r.spec.Name, id), err.Error())
			return
		}
	}
	if out == nil || r.spec.ReadAfterWrite {
		if out, _, err = r.get(ctx, req.Plan.Raw, id); err != nil {
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
	resp.Diagnostics.Append(r.setIdentity(ctx, resp.Identity, state)...)
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	id := stringAttr(req.State.Raw, "id")
	if r.spec.DeletePath == "" {
		resp.Diagnostics.AddWarning(fmt.Sprintf("%s %s was not deleted", r.spec.Name, id), "The API has no delete endpoint for this resource. Terraform removed it from the state only.")
		return
	}
	_, status, err := call(ctx, r.client, r.spec.DeleteMethod, r.path(r.spec.DeletePath, req.State.Raw, id), bodyFor(r.spec.DeleteMethod))
	if err != nil && !isGone(status) {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to delete %s %s", r.spec.Name, id), err.Error())
	}
}
