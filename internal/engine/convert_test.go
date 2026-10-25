// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

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
