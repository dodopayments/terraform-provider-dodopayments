// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package engine

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	sdk "github.com/dodopayments/dodopayments-go"
)

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
			return nil, apiErr.StatusCode, err
		}
		return nil, 0, err
	}
	out, err := decodeJSON(raw)
	return out, 0, err
}

func isGone(status int) bool {
	return status == 404 || status == 410
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
