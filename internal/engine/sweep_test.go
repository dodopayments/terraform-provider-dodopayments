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
				_, _ = w.Write([]byte(`{"code":"CANNOT_DELETE","message":"cannot delete"}`))
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
	if err == nil || !strings.Contains(err.Error(), "item d") || !strings.Contains(err.Error(), "cannot delete (HTTP 400, error code CANNOT_DELETE, request DELETE /items/d)") {
		t.Errorf("Sweep error = %v, want an API error for item d", err)
	}
	sort.Strings(deleted)
	if got := strings.Join(deleted, ","); got != "a,c,d" {
		t.Errorf("deleted %q, want \"a,c,d\"", got)
	}
}

func TestRateLimitMiddlewareRetriesAfterReset(t *testing.T) {
	var mu sync.Mutex
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		calls++
		n := calls
		mu.Unlock()
		if n == 1 {
			w.Header().Set("X-RateLimit-Reset", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		_, _ = w.Write([]byte(`{"id":"a"}`))
	}))
	defer srv.Close()
	c := sdk.NewClient(option.WithBaseURL(srv.URL), option.WithBearerToken("test"), option.WithMaxRetries(2), option.WithMiddleware(RateLimitMiddleware))
	out, _, err := call(context.Background(), c, http.MethodGet, "items/a", nil)
	if err != nil || idOf(out, "id") != "a" || calls != 2 {
		t.Errorf("call = %v, %v after %d calls, want the item after 2 calls", out, err, calls)
	}
}
