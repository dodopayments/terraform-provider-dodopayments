// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"encoding/json"
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

func TestDecimalEqual(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"10.50", "10.5", true},
		{"100", "100.0", true},
		{"1e2", "100", true},
		{"10.5", "10.51", false},
		{"0x10", "16", false},
		{"abc", "abc", false},
	}
	for _, c := range cases {
		if got := decimalEqual(c.a, c.b); got != c.want {
			t.Errorf("decimalEqual(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func str(s string) tftypes.Value { return tftypes.NewValue(tftypes.String, s) }

func TestReconcile(t *testing.T) {
	item := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"key": tftypes.String, "amount": tftypes.String, "secret": tftypes.String}}
	list := tftypes.List{ElementType: item}
	typ := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"clauses": tftypes.String, "descriptor": tftypes.String, "items": list}}
	obj := func(key, amount, secret string) tftypes.Value {
		v := map[string]tftypes.Value{"key": str(key), "amount": str(amount), "secret": tftypes.NewValue(tftypes.String, nil)}
		if secret != "" {
			v["secret"] = str(secret)
		}
		return tftypes.NewValue(item, v)
	}
	value := func(clauses, descriptor string, items ...tftypes.Value) tftypes.Value {
		return tftypes.NewValue(typ, map[string]tftypes.Value{
			"clauses":    str(clauses),
			"descriptor": str(descriptor),
			"items":      tftypes.NewValue(list, items),
		})
	}
	rules := &Rules{
		JSONAttrs:   map[string]bool{"x.clauses": true},
		KeepPrior:   map[string]bool{"x.items.secret": true},
		ListKeys:    map[string]string{"x.items": "key"},
		StripPrefix: map[string]string{"x.descriptor": "PFX_"},
		Decimals:    map[string]bool{"x.items.amount": true},
	}
	configured := "{\n  \"key\": \"a\"\n}"
	prior := value(configured, "SHOP", obj("a", "10.50", "s1"), obj("b", "1", "s2"))
	next := value(`{"key":"a"}`, "PFX_SHOP", obj("b", "1.0", ""), obj("a", "10.5", ""))
	if got := rules.reconcile(prior, next, "x"); !got.Equal(prior) {
		t.Errorf("equivalent value: got %v, want %v", got, prior)
	}

	changed := value(`{"key":"b"}`, "PFX_OTHER", obj("a", "11", ""))
	want := value(`{"key":"b"}`, "OTHER", obj("a", "11", "s1"))
	if got := rules.reconcile(prior, changed, "x"); !got.Equal(want) {
		t.Errorf("changed value: got %v, want %v", got, want)
	}

	imported := rules.reconcile(tftypes.NewValue(typ, nil), value("{}", "PFX_SHOP"), "x")
	if got := stringAttr(imported, "descriptor"); got != "SHOP" {
		t.Errorf("imported descriptor = %q, want SHOP", got)
	}
}

func TestReconcileKeysSetsAndDecimals(t *testing.T) {
	item := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"key": tftypes.String, "secret": tftypes.String}}
	obj := func(key, secret string) tftypes.Value {
		v := map[string]tftypes.Value{"key": str(key), "secret": tftypes.NewValue(tftypes.String, nil)}
		if secret != "" {
			v["secret"] = str(secret)
		}
		return tftypes.NewValue(item, v)
	}
	list := tftypes.List{ElementType: item}
	set := tftypes.Set{ElementType: item}
	rules := &Rules{KeepPrior: map[string]bool{"items.secret": true}, ListKeys: map[string]string{"items": "key"}}

	// The API dropped item "a". Item "b" keeps its own prior value.
	prior := tftypes.NewValue(list, []tftypes.Value{obj("a", "s1"), obj("b", "s2")})
	got := rules.reconcile(prior, tftypes.NewValue(list, []tftypes.Value{obj("b", "")}), "items")
	if want := tftypes.NewValue(list, []tftypes.Value{obj("b", "s2")}); !got.Equal(want) {
		t.Errorf("dropped item: got %v, want %v", got, want)
	}

	// A set element gets the prior value of the element that it matches.
	noKey := &Rules{KeepPrior: map[string]bool{"items.secret": true}}
	priorSet := tftypes.NewValue(set, []tftypes.Value{obj("a", "s1"), obj("b", "s2")})
	got = noKey.reconcile(priorSet, tftypes.NewValue(set, []tftypes.Value{obj("b", ""), obj("a", "")}), "items")
	if !got.Equal(priorSet) {
		t.Errorf("set: got %v, want %v", got, priorSet)
	}

	// Decimal equality applies only to the configured paths.
	if got := (&Rules{}).reconcile(str("001"), str("1"), "code"); !got.Equal(str("1")) {
		t.Errorf("code: got %v, want 1", got)
	}
	if got := (&Rules{Decimals: map[string]bool{"amount": true}}).reconcile(str("10.50"), str("10.5"), "amount"); !got.Equal(str("10.50")) {
		t.Errorf("amount: got %v, want 10.50", got)
	}
}

func TestToJSONNulls(t *testing.T) {
	item := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"a": tftypes.String}}
	list := tftypes.List{ElementType: item}
	v := tftypes.NewValue(list, []tftypes.Value{
		tftypes.NewValue(item, map[string]tftypes.Value{"a": tftypes.NewValue(tftypes.String, nil)}),
		tftypes.NewValue(item, nil),
	})
	got, err := toJSON(v, "x", nil, true)
	if err != nil {
		t.Fatal(err)
	}
	if b, _ := json.Marshal(got); string(b) != `[{},null]` {
		t.Errorf("list: got %s, want [{},null]", b)
	}
	m := tftypes.NewValue(tftypes.Map{ElementType: tftypes.String}, map[string]tftypes.Value{"k": tftypes.NewValue(tftypes.String, nil)})
	got, _ = toJSON(m, "x", nil, false)
	if b, _ := json.Marshal(got); string(b) != `{"k":null}` {
		t.Errorf("map: got %s, want a null value for k", b)
	}
}

func TestValueChangedSet(t *testing.T) {
	set := tftypes.Set{ElementType: tftypes.String}
	a, b := str("a"), str("b")
	state := tftypes.NewValue(set, []tftypes.Value{a, b})
	plan := tftypes.NewValue(set, []tftypes.Value{b, a})
	if valueChanged(plan, plan, state) {
		t.Error("the same set in another order is a change")
	}
	// A changed reference is unknown in the configuration and in the plan.
	unknown := tftypes.NewValue(set, []tftypes.Value{a, tftypes.NewValue(tftypes.String, tftypes.UnknownValue)})
	if !valueChanged(unknown, unknown, state) {
		t.Error("an unknown configured element is not a change")
	}
}

func TestApplyAliases(t *testing.T) {
	out := map[string]any{"entitlements": []any{map[string]any{"id": "ent_1"}, map[string]any{"id": "ent_2", "entitlement_id": "ent_x"}}}
	rules := &Rules{Aliases: map[string]string{"entitlements.entitlement_id": "id"}}
	rules.applyAliases(out)
	items := out["entitlements"].([]any)
	if got := items[0].(map[string]any)["entitlement_id"]; got != "ent_1" {
		t.Errorf("alias = %v, want ent_1", got)
	}
	if got := items[1].(map[string]any)["entitlement_id"]; got != "ent_x" {
		t.Errorf("alias kept = %v, want ent_x", got)
	}
}
