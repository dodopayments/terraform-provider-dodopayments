// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"testing"

	sdk "github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
)

func TestSweepDeletesOnlyMarkedObjects(t *testing.T) {
	pages := map[string]map[string]any{
		"": {"data": []any{
			map[string]any{"id": "a", "name": "tf-acc-ONE"},
			map[string]any{"id": "b", "name": "production"},
		}, "next": "p2"},
		"p2": {"data": []any{
			map[string]any{"id": "c", "code": "TFACCTWO"},
			map[string]any{"id": "d", "name": "tf-acc-FAIL"},
			map[string]any{"id": "e", "count": 3},
		}, "next": ""},
	}
	var mu sync.Mutex
	var deleted []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/items":
			_ = json.NewEncoder(w).Encode(pages[r.URL.Query().Get("cursor")])
		case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/items/"):
			id := strings.TrimPrefix(r.URL.Path, "/items/")
			mu.Lock()
			deleted = append(deleted, id)
			mu.Unlock()
			if id == "d" {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"message":"cannot delete"}`))
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	c := sdk.NewClient(option.WithBaseURL(srv.URL), option.WithBearerToken("test"), option.WithMaxRetries(0))
	spec := SweepSpec{
		Name:         "item",
		IDProperty:   "id",
		DeleteMethod: http.MethodDelete,
		DeletePath:   "items/%s",
		List:         &ListDataSourceSpec{ListPath: "items", ItemsProperty: "data", CursorParam: "cursor", CursorField: "next"},
	}
	err := Sweep(context.Background(), c, spec, "tf-acc-", "tfacc")
	if err == nil || !strings.Contains(err.Error(), "item d") {
		t.Errorf("Sweep error = %v, want an error for item d", err)
	}
	sort.Strings(deleted)
	if got := strings.Join(deleted, ","); got != "a,c,d" {
		t.Errorf("deleted %q, want \"a,c,d\"", got)
	}
}
