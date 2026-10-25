// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

// listItems reads pages from the list endpoint until it has maxItems items or the list ends.
func listItems(ctx context.Context, c *sdk.Client, spec *ListDataSourceSpec, query url.Values, maxItems int64) ([]any, error) {
	const pageSize = 100
	items := []any{}
	cursor := ""
	for page := 0; int64(len(items)) < maxItems; page++ {
		if spec.PageNumberParam != "" {
			query.Set(spec.PageNumberParam, strconv.Itoa(page))
		}
		if spec.PageSizeParam != "" {
			query.Set(spec.PageSizeParam, strconv.Itoa(pageSize))
		}
		if cursor != "" {
			query.Set(spec.CursorParam, cursor)
		}
		out, _, err := call(ctx, c, http.MethodGet, spec.ListPath+"?"+query.Encode(), nil)
		if err != nil {
			return nil, err
		}
		m, _ := out.(map[string]any)
		pageItems, _ := m[spec.ItemsProperty].([]any)
		items = append(items, pageItems...)
		if len(pageItems) == 0 {
			break
		}
		if spec.PageNumberParam != "" {
			// Without a page size parameter, the API selects the page size. Only an empty page ends the list.
			if spec.PageSizeParam != "" && len(pageItems) < pageSize {
				break
			}
			continue
		}
		if spec.CursorParam == "" {
			break
		}
		if last, _ := m[spec.LastPageField].(bool); last {
			break
		}
		next, _ := m[spec.CursorField].(string)
		if next == "" || next == cursor {
			break
		}
		cursor = next
	}
	return items, nil
}

// SweepSpec describes how to find and delete the objects that acceptance tests left behind.
type SweepSpec struct {
	Name         string
	List         *ListDataSourceSpec
	IDProperty   string
	DeleteMethod string
	DeletePath   string
}

// Sweep reads the full list and deletes each object that has a top-level string value that contains one of the markers.
func Sweep(ctx context.Context, c *sdk.Client, spec SweepSpec, markers ...string) error {
	items, err := listItems(ctx, c, spec.List, url.Values{}, math.MaxInt64)
	if err != nil {
		return fmt.Errorf("list %s: %w", spec.Name, err)
	}
	var failed []string
	for _, item := range items {
		m, _ := item.(map[string]any)
		if !hasMarker(m, markers) {
			continue
		}
		id := idOf(m, spec.IDProperty)
		if id == "" {
			continue
		}
		_, status, err := call(ctx, c, spec.DeleteMethod, fmt.Sprintf(spec.DeletePath, url.PathEscape(id)), nil)
		if err != nil && !isGone(status) {
			failed = append(failed, fmt.Sprintf("%s %s: %v", spec.Name, id, err))
		}
	}
	if len(failed) > 0 {
		return fmt.Errorf("failed to delete %d objects: %s", len(failed), strings.Join(failed, "; "))
	}
	return nil
}

func hasMarker(m map[string]any, markers []string) bool {
	for _, v := range m {
		s, ok := v.(string)
		if !ok {
			continue
		}
		for _, marker := range markers {
			if strings.Contains(strings.ToLower(s), strings.ToLower(marker)) {
				return true
			}
		}
	}
	return false
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
	items, err := listItems(ctx, d.client, d.spec, query, maxItems)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to list %s", d.spec.Name), err.Error())
		return
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
