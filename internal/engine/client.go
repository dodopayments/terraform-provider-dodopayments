// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	sdk "github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
)

// APIError is an error response from the API. Its text has the API message, the HTTP status, the error code and the request.
type APIError struct {
	Method  string
	Path    string
	Status  int
	Code    string
	Message string
}

func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = http.StatusText(e.Status)
	}
	if e.Code != "" {
		return fmt.Sprintf("%s (HTTP %d, error code %s, request %s /%s)", msg, e.Status, e.Code, e.Method, e.Path)
	}
	return fmt.Sprintf("%s (HTTP %d, request %s /%s)", msg, e.Status, e.Method, e.Path)
}

func newAPIError(method, path string, e *sdk.Error) *APIError {
	out := &APIError{Method: method, Path: strings.TrimPrefix(path, "/"), Status: e.StatusCode}
	var body struct {
		Code    string  `json:"code"`
		Message *string `json:"message"`
	}
	if json.Unmarshal([]byte(e.JSON.RawJSON()), &body) == nil {
		out.Code = body.Code
		if body.Message != nil {
			out.Message = *body.Message
		}
	}
	return out
}

// call sends one API request. It returns the decoded JSON body and, for an API error, the HTTP status.
func call(ctx context.Context, c *sdk.Client, method, path string, body any) (any, int, error) {
	var params any
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, 0, err
		}
		params = b
	}
	var raw []byte
	err := c.Execute(ctx, method, strings.TrimPrefix(path, "/"), params, &raw)
	if err != nil {
		var apiErr *sdk.Error
		if errors.As(err, &apiErr) {
			return nil, apiErr.StatusCode, newAPIError(method, path, apiErr)
		}
		return nil, 0, err
	}
	out, err := decodeJSON(raw)
	return out, 0, err
}

// bodyFor returns an empty JSON object for a method that needs a body, and nil for GET and DELETE.
func bodyFor(method string) any {
	if method == http.MethodGet || method == http.MethodDelete {
		return nil
	}
	return map[string]any{}
}

func isGone(status int) bool {
	return status == http.StatusNotFound || status == http.StatusGone
}

func escape(s string) string {
	return url.PathEscape(s)
}

func idOf(out any, property string) string {
	m, ok := out.(map[string]any)
	if !ok {
		return ""
	}
	if s, ok := m[property].(string); ok {
		return s
	}
	return ""
}

// RateLimitMiddleware lets the SDK wait until the rate limit window resets. On HTTP 429 the API sends
// X-RateLimit-Reset (seconds) without Retry-After, so the middleware copies it, with jitter, to Retry-After.
func RateLimitMiddleware(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
	res, err := next(req)
	if err != nil || res == nil || res.StatusCode != http.StatusTooManyRequests || res.Header.Get("Retry-After") != "" {
		return res, err
	}
	if reset, perr := strconv.ParseFloat(res.Header.Get("X-RateLimit-Reset"), 64); perr == nil && reset >= 0 {
		res.Header.Set("Retry-After", strconv.FormatFloat(reset+rand.Float64(), 'f', 3, 64))
	} else {
		// An edge rate limit sends neither header and blocks for up to 60 seconds. The retries must last longer than that.
		res.Header.Set("Retry-After", strconv.FormatFloat(15+rand.Float64()*5, 'f', 3, 64))
	}
	return res, err
}
