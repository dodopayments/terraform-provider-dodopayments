// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type omitted struct{}

// omit marks a value that must not be sent: an unknown value, or a null value inside a create body.
var omit any = omitted{}

func childPath(parent, name string) string {
	if parent == "" {
		return name
	}
	return parent + "." + name
}

func decodeJSON(raw []byte) (any, error) {
	if len(bytes.TrimSpace(raw)) == 0 {
		return nil, nil
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var out any
	if err := dec.Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// toJSON converts a Terraform value to a value for encoding/json.
// When keepNull is true, a null value becomes a JSON null. Otherwise it is omitted.
func toJSON(v tftypes.Value, path string, jsonAttrs map[string]bool, keepNull bool) (any, error) {
	if !v.IsKnown() {
		return omit, nil
	}
	if v.IsNull() {
		if keepNull {
			return nil, nil
		}
		return omit, nil
	}
	switch t := v.Type().(type) {
	case tftypes.Object:
		var m map[string]tftypes.Value
		if err := v.As(&m); err != nil {
			return nil, err
		}
		out := map[string]any{}
		for k, child := range m {
			c, err := toJSON(child, childPath(path, k), jsonAttrs, keepNull)
			if err != nil {
				return nil, err
			}
			if c == omit {
				continue
			}
			out[k] = c
		}
		return out, nil
	case tftypes.List, tftypes.Set:
		var items []tftypes.Value
		if err := v.As(&items); err != nil {
			return nil, err
		}
		out := make([]any, 0, len(items))
		for _, item := range items {
			c, err := toJSON(item, path, jsonAttrs, true)
			if err != nil {
				return nil, err
			}
			if c == omit {
				continue
			}
			out = append(out, c)
		}
		return out, nil
	case tftypes.Map:
		var m map[string]tftypes.Value
		if err := v.As(&m); err != nil {
			return nil, err
		}
		out := map[string]any{}
		for k, child := range m {
			c, err := toJSON(child, path, jsonAttrs, true)
			if err != nil {
				return nil, err
			}
			if c == omit {
				continue
			}
			out[k] = c
		}
		return out, nil
	default:
		switch {
		case t.Is(tftypes.String):
			var s string
			if err := v.As(&s); err != nil {
				return nil, err
			}
			if jsonAttrs[path] {
				out, err := decodeJSON([]byte(s))
				if err != nil {
					return nil, fmt.Errorf("attribute %q must contain valid JSON: %w", path, err)
				}
				return out, nil
			}
			return s, nil
		case t.Is(tftypes.Bool):
			var b bool
			err := v.As(&b)
			return b, err
		case t.Is(tftypes.Number):
			var f big.Float
			if err := v.As(&f); err != nil {
				return nil, err
			}
			if f.IsInt() {
				i, _ := f.Int(nil)
				return json.Number(i.String()), nil
			}
			return json.Number(f.Text('g', -1)), nil
		}
	}
	return nil, fmt.Errorf("attribute %q has an unsupported type %s", path, v.Type())
}

// fromJSON converts a decoded JSON value to a Terraform value of type t.
func fromJSON(t tftypes.Type, raw any, path string, jsonAttrs map[string]bool) (tftypes.Value, error) {
	if raw == nil {
		return tftypes.NewValue(t, nil), nil
	}
	switch tt := t.(type) {
	case tftypes.Object:
		m, ok := raw.(map[string]any)
		if !ok {
			return tftypes.Value{}, fmt.Errorf("attribute %q: expected a JSON object, got %T", path, raw)
		}
		vals := make(map[string]tftypes.Value, len(tt.AttributeTypes))
		for k, at := range tt.AttributeTypes {
			cv, err := fromJSON(at, m[k], childPath(path, k), jsonAttrs)
			if err != nil {
				return tftypes.Value{}, err
			}
			vals[k] = cv
		}
		return tftypes.NewValue(tt, vals), nil
	case tftypes.List:
		elems, err := fromJSONArray(tt.ElementType, raw, path, jsonAttrs)
		if err != nil {
			return tftypes.Value{}, err
		}
		return tftypes.NewValue(tt, elems), nil
	case tftypes.Set:
		elems, err := fromJSONArray(tt.ElementType, raw, path, jsonAttrs)
		if err != nil {
			return tftypes.Value{}, err
		}
		return tftypes.NewValue(tt, elems), nil
	case tftypes.Map:
		m, ok := raw.(map[string]any)
		if !ok {
			return tftypes.Value{}, fmt.Errorf("attribute %q: expected a JSON object, got %T", path, raw)
		}
		vals := make(map[string]tftypes.Value, len(m))
		for k, item := range m {
			cv, err := fromJSON(tt.ElementType, item, path, jsonAttrs)
			if err != nil {
				return tftypes.Value{}, err
			}
			vals[k] = cv
		}
		return tftypes.NewValue(tt, vals), nil
	}
	switch {
	case t.Is(tftypes.String):
		if jsonAttrs[path] {
			b, err := json.Marshal(raw)
			if err != nil {
				return tftypes.Value{}, err
			}
			return tftypes.NewValue(t, string(b)), nil
		}
		switch r := raw.(type) {
		case string:
			return tftypes.NewValue(t, r), nil
		case json.Number:
			return tftypes.NewValue(t, r.String()), nil
		case bool:
			return tftypes.NewValue(t, strconv.FormatBool(r)), nil
		default:
			b, err := json.Marshal(r)
			if err != nil {
				return tftypes.Value{}, err
			}
			return tftypes.NewValue(t, string(b)), nil
		}
	case t.Is(tftypes.Number):
		var s string
		switch r := raw.(type) {
		case json.Number:
			s = r.String()
		case string:
			s = r
		default:
			return tftypes.Value{}, fmt.Errorf("attribute %q: expected a number, got %T", path, raw)
		}
		f, _, err := big.ParseFloat(s, 10, 512, big.ToNearestEven)
		if err != nil {
			return tftypes.Value{}, fmt.Errorf("attribute %q: %w", path, err)
		}
		return tftypes.NewValue(t, f), nil
	case t.Is(tftypes.Bool):
		b, ok := raw.(bool)
		if !ok {
			return tftypes.Value{}, fmt.Errorf("attribute %q: expected a boolean, got %T", path, raw)
		}
		return tftypes.NewValue(t, b), nil
	}
	return tftypes.Value{}, fmt.Errorf("attribute %q has an unsupported type %s", path, t)
}

func fromJSONArray(elem tftypes.Type, raw any, path string, jsonAttrs map[string]bool) ([]tftypes.Value, error) {
	arr, ok := raw.([]any)
	if !ok {
		return nil, fmt.Errorf("attribute %q: expected a JSON array, got %T", path, raw)
	}
	out := make([]tftypes.Value, 0, len(arr))
	for _, item := range arr {
		v, err := fromJSON(elem, item, path, jsonAttrs)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

// jsonEqual reports whether two JSON documents have the same value.
func jsonEqual(a, b string) bool {
	var va, vb any
	if json.Unmarshal([]byte(a), &va) != nil || json.Unmarshal([]byte(b), &vb) != nil {
		return false
	}
	return reflect.DeepEqual(va, vb)
}

// keepEquivalentJSON returns prior for a JSON attribute when next holds the same JSON value.
// This keeps the text that the configuration set, for example its whitespace and its key order.
func keepEquivalentJSON(prior, next tftypes.Value, path string, jsonAttrs map[string]bool) tftypes.Value {
	if !prior.IsKnown() || prior.IsNull() || !next.IsKnown() || next.IsNull() || !prior.Type().Equal(next.Type()) {
		return next
	}
	switch t := next.Type().(type) {
	case tftypes.Object:
		var p, n map[string]tftypes.Value
		if prior.As(&p) != nil || next.As(&n) != nil {
			return next
		}
		for k, v := range n {
			if pv, ok := p[k]; ok {
				n[k] = keepEquivalentJSON(pv, v, childPath(path, k), jsonAttrs)
			}
		}
		return tftypes.NewValue(t, n)
	case tftypes.List:
		var p, n []tftypes.Value
		if prior.As(&p) != nil || next.As(&n) != nil || len(p) != len(n) {
			return next
		}
		for i := range n {
			n[i] = keepEquivalentJSON(p[i], n[i], path, jsonAttrs)
		}
		return tftypes.NewValue(t, n)
	}
	if jsonAttrs[path] && next.Type().Is(tftypes.String) {
		var ps, ns string
		if prior.As(&ps) == nil && next.As(&ns) == nil && jsonEqual(ps, ns) {
			return prior
		}
	}
	return next
}

func isEmptyValue(v tftypes.Value) bool {
	if v.IsNull() || !v.IsKnown() {
		return false
	}
	switch v.Type().(type) {
	case tftypes.List, tftypes.Set:
		var items []tftypes.Value
		return v.As(&items) == nil && len(items) == 0
	case tftypes.Map:
		var m map[string]tftypes.Value
		return v.As(&m) == nil && len(m) == 0
	}
	if v.Type().Is(tftypes.String) {
		var s string
		return v.As(&s) == nil && s == ""
	}
	if v.Type().Is(tftypes.Bool) {
		var b bool
		return v.As(&b) == nil && !b
	}
	return false
}
