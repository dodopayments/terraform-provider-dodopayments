// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"

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
// When keepNull is true, a null value becomes a JSON null. Otherwise it is omitted. A null list, set or map
// entry stays a JSON null. An unset child of an entry is omitted, because a merge patch replaces the entry.
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
			if item.IsKnown() && item.IsNull() {
				out = append(out, nil)
				continue
			}
			c, err := toJSON(item, path, jsonAttrs, false)
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
			if child.IsKnown() && child.IsNull() {
				out[k] = nil
				continue
			}
			c, err := toJSON(child, path, jsonAttrs, false)
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

var decimalPattern = regexp.MustCompile(`^[+-]?(\d+\.?\d*|\.\d+)([eE][+-]?\d+)?$`)

// decimalEqual reports whether two strings hold the same decimal number, for example "10.50" and "10.5".
func decimalEqual(a, b string) bool {
	if !decimalPattern.MatchString(a) || !decimalPattern.MatchString(b) {
		return false
	}
	x, okA := new(big.Rat).SetString(a)
	y, okB := new(big.Rat).SetString(b)
	return okA && okB && x.Cmp(y) == 0
}

// Rules tell how an API value maps to an attribute value. Paths are dotted attribute names. List items have no index.
type Rules struct {
	// JSONAttrs are the attribute paths that hold a JSON document as a string.
	JSONAttrs map[string]bool
	// KeepPrior are input paths that the API does not return. The provider keeps the configured value.
	KeepPrior map[string]bool
	// ListKeys maps a list path to the attribute that identifies each item. The provider keeps the configured order.
	ListKeys map[string]string
	// StripPrefix maps a string path to a prefix that the API adds to the configured value.
	StripPrefix map[string]string
	// Aliases maps an attribute path to the response property in the same object that holds its value.
	Aliases map[string]string
	// Decimals are the string paths that hold a decimal number. The provider keeps "10.50" when the API returns "10.5".
	Decimals map[string]bool
}

// applyAliases copies each aliased response property to the attribute name, in place.
func (rules *Rules) applyAliases(out any) {
	for target, source := range rules.Aliases {
		setAlias(out, strings.Split(target, "."), source)
	}
}

func setAlias(v any, path []string, source string) {
	switch x := v.(type) {
	case []any:
		for _, item := range x {
			setAlias(item, path, source)
		}
	case map[string]any:
		if len(path) == 1 {
			if cur, ok := x[path[0]]; !ok || cur == nil {
				if s, ok := x[source]; ok {
					x[path[0]] = s
				}
			}
			return
		}
		setAlias(x[path[0]], path[1:], source)
	}
}

// reconcile returns the value to store at path. prior is the planned or the stored value, or null on import.
// It removes the prefix that the API adds to a value. It keeps prior when next holds the same value in another
// form: equivalent JSON, an equal decimal, or the same list items in another order. It also keeps an input that
// the API does not return.
func (rules *Rules) reconcile(prior, next tftypes.Value, path string) tftypes.Value {
	hasPrior := prior.IsKnown() && !prior.IsNull() && prior.Type().Equal(next.Type())
	if !next.IsKnown() {
		return next
	}
	if next.IsNull() {
		if hasPrior && rules.KeepPrior[path] {
			return prior
		}
		return next
	}
	switch t := next.Type().(type) {
	case tftypes.Object:
		var p, n map[string]tftypes.Value
		if next.As(&n) != nil || hasPrior && prior.As(&p) != nil {
			return next
		}
		// As returns the internal map of next, so the result goes in a new map.
		out := make(map[string]tftypes.Value, len(n))
		for k, v := range n {
			pv, ok := p[k]
			if !ok {
				pv = tftypes.NewValue(v.Type(), nil)
			}
			out[k] = rules.reconcile(pv, v, childPath(path, k))
		}
		return tftypes.NewValue(t, out)
	case tftypes.List, tftypes.Set:
		var p, n []tftypes.Value
		if next.As(&n) != nil || hasPrior && prior.As(&p) != nil {
			return next
		}
		_, isList := t.(tftypes.List)
		key := rules.ListKeys[path]
		if isList && key != "" {
			n = alignByKey(p, n, key)
		} else {
			n = append([]tftypes.Value(nil), n...)
		}
		byKey := map[string]tftypes.Value{}
		if key != "" {
			for _, item := range p {
				if k := stringAttr(item, key); k != "" {
					byKey[k] = item
				}
			}
		}
		used := make([]bool, len(p))
		for i := range n {
			pv := tftypes.NewValue(n[i].Type(), nil)
			switch {
			case key != "":
				if item, ok := byKey[stringAttr(n[i], key)]; ok {
					pv = item
				}
			case isList:
				if len(p) == len(n) {
					pv = p[i]
				}
			default:
				// A set has no order. The prior element is the one that the API value matches after reconciliation.
				for j := range p {
					if !used[j] && rules.reconcile(p[j], n[i], path).Equal(p[j]) {
						pv, used[j] = p[j], true
						break
					}
				}
			}
			n[i] = rules.reconcile(pv, n[i], path)
		}
		return tftypes.NewValue(t, n)
	case tftypes.Map:
		var p, n map[string]tftypes.Value
		if next.As(&n) != nil || hasPrior && prior.As(&p) != nil {
			return next
		}
		out := make(map[string]tftypes.Value, len(n))
		for k, v := range n {
			pv, ok := p[k]
			if !ok {
				pv = tftypes.NewValue(v.Type(), nil)
			}
			out[k] = rules.reconcile(pv, v, path)
		}
		return tftypes.NewValue(t, out)
	}
	if !next.Type().Is(tftypes.String) {
		return next
	}
	var ps, ns string
	if next.As(&ns) != nil {
		return next
	}
	if prefix := rules.StripPrefix[path]; prefix != "" && strings.HasPrefix(ns, prefix) {
		ns = strings.TrimPrefix(ns, prefix)
		next = tftypes.NewValue(tftypes.String, ns)
	}
	if !hasPrior || prior.As(&ps) != nil || ps == ns {
		return next
	}
	if rules.JSONAttrs[path] && jsonEqual(ps, ns) || rules.Decimals[path] && decimalEqual(ps, ns) {
		return prior
	}
	return next
}

// alignByKey orders the items of next like the items of prior with the same key. New items go last, in API order.
func alignByKey(prior, next []tftypes.Value, key string) []tftypes.Value {
	used := make([]bool, len(next))
	out := make([]tftypes.Value, 0, len(next))
	for _, p := range prior {
		k := stringAttr(p, key)
		if k == "" {
			continue
		}
		for i, n := range next {
			if !used[i] && stringAttr(n, key) == k {
				out = append(out, n)
				used[i] = true
				break
			}
		}
	}
	for i, n := range next {
		if !used[i] {
			out = append(out, n)
		}
	}
	return out
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

func stringAttr(v tftypes.Value, name string) string {
	if !v.IsKnown() || v.IsNull() {
		return ""
	}
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

// pathArgs returns the escaped values for the path verbs: the parent attributes of v and then id.
func pathArgs(template string, v tftypes.Value, parents []string, id string) []any {
	args := make([]any, 0, len(parents)+1)
	for _, p := range parents {
		args = append(args, escape(stringAttr(v, p)))
	}
	if strings.Count(template, "%s") > len(args) {
		args = append(args, escape(id))
	}
	return args
}
