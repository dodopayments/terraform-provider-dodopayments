// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestToStateKeepsUpsertKey(t *testing.T) {
	ctx := context.Background()
	r := &Resource{spec: &ResourceSpec{
		IDProperty:      "key",
		CreatePathParam: "key",
		Schema: schema.Schema{Attributes: map[string]schema.Attribute{
			"id":    schema.StringAttribute{Computed: true},
			"key":   schema.StringAttribute{Required: true},
			"value": schema.StringAttribute{Required: true},
		}},
	}}
	typ := r.spec.Schema.Type().TerraformType(ctx)
	plan := tftypes.NewValue(typ, map[string]tftypes.Value{
		"id":    tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"key":   tftypes.NewValue(tftypes.String, "k1"),
		"value": tftypes.NewValue(tftypes.String, "v"),
	})
	// The response does not contain the key or the ID.
	got, err := r.toState(ctx, map[string]any{"value": "v"}, plan, "k1")
	if err != nil {
		t.Fatal(err)
	}
	want := tftypes.NewValue(typ, map[string]tftypes.Value{
		"id":    tftypes.NewValue(tftypes.String, "k1"),
		"key":   tftypes.NewValue(tftypes.String, "k1"),
		"value": tftypes.NewValue(tftypes.String, "v"),
	})
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestJSONEqual(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{`{"a": 1, "b": [true, null]}`, `{"b":[true,null],"a":1}`, true},
		{`{"a": 1.0}`, `{"a":1}`, true},
		{`{"a": 1}`, `{"a": 2}`, false},
		{`{"a": 1}`, `not json`, false},
	}
	for _, c := range cases {
		if got := jsonEqual(c.a, c.b); got != c.want {
			t.Errorf("jsonEqual(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestKeepEquivalentJSON(t *testing.T) {
	typ := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"clauses": tftypes.String, "name": tftypes.String}}
	value := func(clauses, name string) tftypes.Value {
		return tftypes.NewValue(typ, map[string]tftypes.Value{
			"clauses": tftypes.NewValue(tftypes.String, clauses),
			"name":    tftypes.NewValue(tftypes.String, name),
		})
	}
	jsonAttrs := map[string]bool{"filter.clauses": true}
	configured := "{\n  \"key\": \"a\"\n}"

	got := keepEquivalentJSON(value(configured, `{"x": 1}`), value(`{"key":"a"}`, `{"x":1}`), "filter", jsonAttrs)
	if want := value(configured, `{"x":1}`); !got.Equal(want) {
		t.Errorf("equivalent JSON: got %v, want %v", got, want)
	}

	next := value(`{"key":"b"}`, `{"x":1}`)
	if got := keepEquivalentJSON(value(configured, `{"x":1}`), next, "filter", jsonAttrs); !got.Equal(next) {
		t.Errorf("changed JSON: got %v, want %v", got, next)
	}
}
