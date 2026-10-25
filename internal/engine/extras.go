// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/action"
	aschema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	eschema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/list"
	lschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
)

func clientOf(data any, diags *diag.Diagnostics) *sdk.Client {
	if data == nil {
		return nil
	}
	client, ok := data.(*sdk.Client)
	if !ok {
		diags.AddError("Unexpected provider data", fmt.Sprintf("Expected *sdk.Client, got %T.", data))
	}
	return client
}

// EphemeralSpec describes an ephemeral resource. Terraform reads its value for each run and never stores it.
type EphemeralSpec struct {
	Name      string
	Schema    eschema.Schema
	Path      string
	Params    []string
	JSONAttrs map[string]bool
}

type EphemeralResource struct {
	spec   *EphemeralSpec
	client *sdk.Client
}

var _ ephemeral.EphemeralResourceWithConfigure = (*EphemeralResource)(nil)

func NewEphemeralResource(spec *EphemeralSpec) func() ephemeral.EphemeralResource {
	return func() ephemeral.EphemeralResource { return &EphemeralResource{spec: spec} }
}

func (e *EphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + e.spec.Name
}

func (e *EphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = e.spec.Schema
}

func (e *EphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	e.client = clientOf(req.ProviderData, &resp.Diagnostics)
}

func (e *EphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	out, _, err := call(ctx, e.client, http.MethodGet, fmt.Sprintf(e.spec.Path, pathArgs(e.spec.Path, req.Config.Raw, e.spec.Params, "")...), nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to read %s", e.spec.Name), err.Error())
		return
	}
	typ := e.spec.Schema.Type().TerraformType(ctx)
	result, err := fromJSON(typ, out, "", e.spec.JSONAttrs)
	var attrs, cfg map[string]tftypes.Value
	if err == nil {
		err = result.As(&attrs)
	}
	if err == nil {
		err = req.Config.Raw.As(&cfg)
	}
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	for _, name := range e.spec.Params {
		attrs[name] = cfg[name]
	}
	resp.Result.Raw = tftypes.NewValue(typ, attrs)
}

// ActionSpec describes an action that sends one request with the path parameters from its configuration.
type ActionSpec struct {
	Name    string
	Schema  aschema.Schema
	Method  string
	Path    string
	Params  []string
	Success string
}

type Action struct {
	spec   *ActionSpec
	client *sdk.Client
}

var _ action.ActionWithConfigure = (*Action)(nil)

func NewAction(spec *ActionSpec) func() action.Action {
	return func() action.Action { return &Action{spec: spec} }
}

func (a *Action) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + a.spec.Name
}

func (a *Action) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = a.spec.Schema
}

func (a *Action) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	a.client = clientOf(req.ProviderData, &resp.Diagnostics)
}

func (a *Action) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	target := fmt.Sprintf(a.spec.Path, pathArgs(a.spec.Path, req.Config.Raw, a.spec.Params, "")...)
	if _, _, err := call(ctx, a.client, a.spec.Method, target, bodyFor(a.spec.Method)); err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to run %s", a.spec.Name), err.Error())
		return
	}
	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("%s: /%s", a.spec.Success, target)})
	}
}

// ListResourceSpec describes the list resource that "terraform query" uses to find objects to import.
type ListResourceSpec struct {
	Name            string
	Schema          lschema.Schema
	List            *ListDataSourceSpec
	IDProperty      string
	DisplayProperty string
	Resource        *ResourceSpec
}

type ListResource struct {
	spec   *ListResourceSpec
	client *sdk.Client
}

var _ list.ListResourceWithConfigure = (*ListResource)(nil)

func NewListResource(spec *ListResourceSpec) func() list.ListResource {
	return func() list.ListResource { return &ListResource{spec: spec} }
}

func (l *ListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + l.spec.Name
}

func (l *ListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = l.spec.Schema
}

func (l *ListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	l.client = clientOf(req.ProviderData, &resp.Diagnostics)
}

func (l *ListResource) List(ctx context.Context, req list.ListRequest, stream *list.ListResultsStream) {
	fail := func(summary string, err error) {
		var diags diag.Diagnostics
		diags.AddError(summary, err.Error())
		stream.Results = list.ListResultsStreamDiagnostics(diags)
	}
	listPath, query, err := listQuery(l.spec.List, req.Config.Raw)
	if err != nil {
		fail("Invalid configuration", err)
		return
	}
	limit := req.Limit
	if limit <= 0 {
		limit = math.MaxInt64
	}
	items, err := listItems(ctx, l.client, l.spec.List, listPath, query, limit)
	if err != nil {
		fail(fmt.Sprintf("Failed to list %s", l.spec.Name), err)
		return
	}
	r := &Resource{spec: l.spec.Resource, client: l.client}
	stream.Results = func(push func(list.ListResult) bool) {
		for _, item := range items {
			result := req.NewListResult(ctx)
			id := idOf(item, l.spec.IDProperty)
			result.DisplayName = id
			if m, ok := item.(map[string]any); ok {
				if name, ok := m[l.spec.DisplayProperty].(string); ok && name != "" {
					result.DisplayName = name
				}
			}
			typ := l.spec.Resource.Schema.Type().TerraformType(ctx)
			known, err := l.knownAttrs(ctx, typ, req.Config.Raw, id)
			if err == nil {
				result.Diagnostics.Append(r.setIdentity(ctx, result.Identity, known)...)
			}
			if err == nil && req.IncludeResource {
				var out any
				if out, _, err = r.get(ctx, known, id); err == nil {
					var state tftypes.Value
					if state, err = r.toState(ctx, out, known, id); err == nil {
						result.Resource.Raw = state
					}
				}
			}
			if err != nil {
				result.Diagnostics.AddAttributeError(path.Root("id"), fmt.Sprintf("Failed to read %s %s", l.spec.Name, id), err.Error())
			}
			if !push(result) {
				return
			}
		}
	}
}

// knownAttrs returns a resource value with only the ID and the parent IDs set.
func (l *ListResource) knownAttrs(_ context.Context, typ tftypes.Type, cfg tftypes.Value, id string) (tftypes.Value, error) {
	obj, ok := typ.(tftypes.Object)
	if !ok {
		return tftypes.Value{}, fmt.Errorf("the resource schema is not an object")
	}
	vals := make(map[string]tftypes.Value, len(obj.AttributeTypes))
	for name, at := range obj.AttributeTypes {
		vals[name] = tftypes.NewValue(at, nil)
	}
	vals["id"] = tftypes.NewValue(tftypes.String, id)
	for _, name := range l.spec.List.ParentParams {
		vals[name] = tftypes.NewValue(tftypes.String, stringAttr(cfg, name))
	}
	return tftypes.NewValue(obj, vals), nil
}
