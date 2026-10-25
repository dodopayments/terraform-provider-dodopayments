// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
)

// DataSourceSpec describes a data source that reads one object by its ID.
type DataSourceSpec struct {
	Name      string
	Schema    schema.Schema
	ReadPath  string
	JSONAttrs map[string]bool
}

// ListDataSourceSpec describes a data source that lists objects.
type ListDataSourceSpec struct {
	Name        string
	Schema      schema.Schema
	ListPath    string
	QueryParams []string
	// ItemsProperty is the response property that holds the page items.
	ItemsProperty string
	// ItemsAttr and MaxItemsAttr are the names of the result list attribute and the limit attribute.
	ItemsAttr    string
	MaxItemsAttr string
	// PageNumberParam is set for page-number pagination. CursorParam and CursorField are set for cursor pagination.
	// LastPageField names an optional response flag that is true on the last page.
	PageNumberParam string
	PageSizeParam   string
	CursorParam     string
	CursorField     string
	LastPageField   string
	DefaultMaxItems int64
	JSONAttrs       map[string]bool
}

type dataSourceBase struct {
	client *sdk.Client
}

func (d *dataSourceBase) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*sdk.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", fmt.Sprintf("Expected *sdk.Client, got %T.", req.ProviderData))
		return
	}
	d.client = client
}

type DataSource struct {
	dataSourceBase
	spec *DataSourceSpec
}

var _ datasource.DataSourceWithConfigure = (*DataSource)(nil)

func NewDataSource(spec *DataSourceSpec) func() datasource.DataSource {
	return func() datasource.DataSource { return &DataSource{spec: spec} }
}

func (d *DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.spec.Name
}

func (d *DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = d.spec.Schema
}

func (d *DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	id := stringAttr(req.Config.Raw, "id")
	out, _, err := call(ctx, d.client, http.MethodGet, fmt.Sprintf(d.spec.ReadPath, url.PathEscape(id)), nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to read %s %s", d.spec.Name, id), err.Error())
		return
	}
	typ := d.spec.Schema.Type().TerraformType(ctx)
	state, err := fromJSON(typ, out, "", d.spec.JSONAttrs)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	var attrs map[string]tftypes.Value
	if err := state.As(&attrs); err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	attrs["id"] = tftypes.NewValue(tftypes.String, id)
	resp.State.Raw = tftypes.NewValue(typ, attrs)
}

type ListDataSource struct {
	dataSourceBase
	spec *ListDataSourceSpec
}

var _ datasource.DataSourceWithConfigure = (*ListDataSource)(nil)

func NewListDataSource(spec *ListDataSourceSpec) func() datasource.DataSource {
	return func() datasource.DataSource { return &ListDataSource{spec: spec} }
}

func (d *ListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.spec.Name
}

func (d *ListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = d.spec.Schema
}

func queryValue(v tftypes.Value) (string, bool) {
	if !v.IsKnown() || v.IsNull() {
		return "", false
	}
	j, err := toJSON(v, "", nil, false)
	if err != nil {
		return "", false
	}
	switch x := j.(type) {
	case string:
		return x, true
	case bool:
		return strconv.FormatBool(x), true
	case fmt.Stringer:
		return x.String(), true
	}
	return "", false
}

func (d *ListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var cfg map[string]tftypes.Value
	if err := req.Config.Raw.As(&cfg); err != nil {
		resp.Diagnostics.AddError("Invalid configuration", err.Error())
		return
	}
	query := url.Values{}
	for _, name := range d.spec.QueryParams {
		if s, ok := queryValue(cfg[name]); ok {
			query.Set(name, s)
		}
	}
	maxItems := d.spec.DefaultMaxItems
	if v, ok := cfg[d.spec.MaxItemsAttr]; ok && v.IsKnown() && !v.IsNull() {
		if s, ok := queryValue(v); ok {
			if n, err := strconv.ParseInt(s, 10, 64); err == nil && n > 0 {
				maxItems = n
			}
		}
	}
	const pageSize = 100
	items := []any{}
	cursor := ""
	for page := 0; int64(len(items)) < maxItems; page++ {
		if d.spec.PageNumberParam != "" {
			query.Set(d.spec.PageNumberParam, strconv.Itoa(page))
		}
		if d.spec.PageSizeParam != "" {
			query.Set(d.spec.PageSizeParam, strconv.Itoa(pageSize))
		}
		if cursor != "" {
			query.Set(d.spec.CursorParam, cursor)
		}
		out, _, err := call(ctx, d.client, http.MethodGet, d.spec.ListPath+"?"+query.Encode(), nil)
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to list %s", d.spec.Name), err.Error())
			return
		}
		m, _ := out.(map[string]any)
		pageItems, _ := m[d.spec.ItemsProperty].([]any)
		items = append(items, pageItems...)
		if len(pageItems) == 0 {
			break
		}
		if d.spec.PageNumberParam != "" {
			// Without a page size parameter, the API selects the page size. Only an empty page ends the list.
			if d.spec.PageSizeParam != "" && len(pageItems) < pageSize {
				break
			}
			continue
		}
		if d.spec.CursorParam == "" {
			break
		}
		if last, _ := m[d.spec.LastPageField].(bool); last {
			break
		}
		next, _ := m[d.spec.CursorField].(string)
		if next == "" || next == cursor {
			break
		}
		cursor = next
	}
	if int64(len(items)) > maxItems {
		items = items[:maxItems]
	}

	typ, ok := d.spec.Schema.Type().TerraformType(ctx).(tftypes.Object)
	if !ok {
		resp.Diagnostics.AddError("Internal error", "The data source schema is not an object.")
		return
	}
	vals := make(map[string]tftypes.Value, len(typ.AttributeTypes))
	for name, at := range typ.AttributeTypes {
		if v, ok := cfg[name]; ok && name != d.spec.ItemsAttr {
			vals[name] = v
			continue
		}
		vals[name] = tftypes.NewValue(at, nil)
	}
	list, err := fromJSON(typ.AttributeTypes[d.spec.ItemsAttr], items, d.spec.ItemsAttr, d.spec.JSONAttrs)
	if err != nil {
		resp.Diagnostics.AddError("Failed to decode the API response", err.Error())
		return
	}
	vals[d.spec.ItemsAttr] = list
	resp.State.Raw = tftypes.NewValue(typ, vals)
}
