// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	sdk "github.com/dodopayments/dodopayments-go"
)

const (
	testSubjectToken = "subject-token-for-tests"
	testAccessToken  = "access-token-for-tests"
	testProviderID   = "wip_test"
)

// clearAuthEnv removes the environment variables that select an authentication method.
func clearAuthEnv(t *testing.T) {
	t.Helper()
	workloadIdentityOptions.Range(func(key, _ any) bool {
		workloadIdentityOptions.Delete(key)
		return true
	})
	for _, name := range []string{
		"DODO_PAYMENTS_API_KEY", identityProviderIDEnv, identityTokenFileEnv,
		"ACTIONS_ID_TOKEN_REQUEST_URL", "ACTIONS_ID_TOKEN_REQUEST_TOKEN",
		"TFC_WORKLOAD_IDENTITY_TOKEN", "TFC_WORKLOAD_IDENTITY_TOKEN_DODO",
	} {
		t.Setenv(name, "")
		if err := os.Unsetenv(name); err != nil {
			t.Fatal(err)
		}
	}
}

func providerSchema(t *testing.T) schema.Schema {
	t.Helper()
	var resp provider.SchemaResponse
	New("test")().Schema(context.Background(), provider.SchemaRequest{}, &resp)
	return resp.Schema
}

// providerConfig builds a provider configuration. A nil block means that the configuration has no workload_identity block.
func providerConfig(t *testing.T, attrs map[string]string, block map[string]string) tfsdk.Config {
	t.Helper()
	s := providerSchema(t)
	typ := s.Type().TerraformType(context.Background()).(tftypes.Object)
	values := map[string]tftypes.Value{}
	for name, attrType := range typ.AttributeTypes {
		values[name] = tftypes.NewValue(attrType, nil)
	}
	for name, v := range attrs {
		values[name] = tftypes.NewValue(tftypes.String, v)
	}
	if block != nil {
		blockType := typ.AttributeTypes["workload_identity"].(tftypes.Object)
		blockValues := map[string]tftypes.Value{}
		for name := range blockType.AttributeTypes {
			blockValues[name] = tftypes.NewValue(tftypes.String, nil)
		}
		for name, v := range block {
			blockValues[name] = tftypes.NewValue(tftypes.String, v)
		}
		values["workload_identity"] = tftypes.NewValue(blockType, blockValues)
	}
	values["max_retries"] = tftypes.NewValue(tftypes.Number, 0)
	return tfsdk.Config{Schema: s, Raw: tftypes.NewValue(typ, values)}
}

func configure(t *testing.T, cfg tfsdk.Config) provider.ConfigureResponse {
	t.Helper()
	var resp provider.ConfigureResponse
	New("test")().Configure(context.Background(), provider.ConfigureRequest{Config: cfg, TerraformVersion: "1.12.0"}, &resp)
	return resp
}

// fakeAPI serves the token exchange and one API path. It counts the exchanges and the API calls that have the access token.
type fakeAPI struct {
	server     *httptest.Server
	exchanges  atomic.Int32
	authorized atomic.Int32
	audience   atomic.Value
}

func newFakeAPI(t *testing.T) *fakeAPI {
	t.Helper()
	api := &fakeAPI{}
	api.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/oidc":
			if r.Header.Get("Authorization") != "bearer request-token" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			api.audience.Store(r.URL.Query().Get("audience"))
			_ = json.NewEncoder(w).Encode(map[string]string{"value": testSubjectToken})
		case r.URL.Path == "/workload-identity/token":
			var body map[string]string
			if r.Method != http.MethodPost || r.Header.Get("Authorization") != "" || json.NewDecoder(r.Body).Decode(&body) != nil ||
				body["subject_token"] != testSubjectToken || body["identity_provider_id"] != testProviderID ||
				body["grant_type"] != "urn:ietf:params:oauth:grant-type:token-exchange" {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"error":"invalid_request","error_description":"unexpected exchange request"}`))
				return
			}
			api.exchanges.Add(1)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": testAccessToken, "issued_token_type": "urn:ietf:params:oauth:token-type:access_token",
				"token_type": "Bearer", "expires_in": 3600,
			})
		default:
			if r.Header.Get("Authorization") != "Bearer "+testAccessToken {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"code":"UNAUTHORIZED","message":"unexpected Authorization header"}`))
				return
			}
			api.authorized.Add(1)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"items":[]}`))
		}
	}))
	t.Cleanup(api.server.Close)
	return api
}

// call configures the provider, sends two API requests with its client and checks that the provider exchanged the token once.
func (api *fakeAPI) call(t *testing.T, cfg tfsdk.Config) {
	t.Helper()
	resp := configure(t, cfg)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure: %v", resp.Diagnostics)
	}
	client, ok := resp.ResourceData.(*sdk.Client)
	if !ok {
		t.Fatalf("ResourceData is %T, want *sdk.Client", resp.ResourceData)
	}
	for range 2 {
		var raw []byte
		if err := client.Execute(context.Background(), http.MethodGet, "items", nil, &raw); err != nil {
			t.Fatalf("API call: %v", err)
		}
	}
	if n := api.exchanges.Load(); n != 1 {
		t.Fatalf("got %d token exchanges, want 1", n)
	}
	if n := api.authorized.Load(); n != 2 {
		t.Fatalf("got %d API calls with the access token, want 2", n)
	}
}

func TestWorkloadIdentitySchema(t *testing.T) {
	block, ok := providerSchema(t).Blocks["workload_identity"].(schema.SingleNestedBlock)
	if !ok {
		t.Fatal("the provider schema has no workload_identity single nested block")
	}
	for _, name := range []string{"identity_provider_id", "token_source", "token_file", "audience"} {
		a, ok := block.Attributes[name].(schema.StringAttribute)
		if !ok || !a.Optional || a.Required {
			t.Errorf("workload_identity.%s must be an optional string attribute", name)
		}
	}
}

func TestWorkloadIdentityConflictsWithAPIKey(t *testing.T) {
	clearAuthEnv(t)
	resp := configure(t, providerConfig(t, map[string]string{"api_key": "key"}, map[string]string{"identity_provider_id": testProviderID}))
	if !resp.Diagnostics.HasError() || resp.Diagnostics.Errors()[0].Summary() != "Conflicting authentication" {
		t.Fatalf("got %v, want a Conflicting authentication error", resp.Diagnostics)
	}
}

func TestWorkloadIdentityNeedsProviderID(t *testing.T) {
	clearAuthEnv(t)
	resp := configure(t, providerConfig(t, nil, map[string]string{"token_source": "vercel"}))
	if !resp.Diagnostics.HasError() || resp.Diagnostics.Errors()[0].Summary() != "Missing identity provider ID" {
		t.Fatalf("got %v, want a Missing identity provider ID error", resp.Diagnostics)
	}
}

func TestWorkloadIdentityFileNeedsPath(t *testing.T) {
	clearAuthEnv(t)
	resp := configure(t, providerConfig(t, nil, map[string]string{"identity_provider_id": testProviderID, "token_source": "file"}))
	if !resp.Diagnostics.HasError() || resp.Diagnostics.Errors()[0].Summary() != "Missing token file" {
		t.Fatalf("got %v, want a Missing token file error", resp.Diagnostics)
	}
}

func TestMissingAPIKey(t *testing.T) {
	clearAuthEnv(t)
	resp := configure(t, providerConfig(t, nil, nil))
	if !resp.Diagnostics.HasError() || resp.Diagnostics.Errors()[0].Summary() != "Missing API key" {
		t.Fatalf("got %v, want a Missing API key error", resp.Diagnostics)
	}
}

// TestWorkloadIdentityEnvGitHubActions has no provider configuration for authentication, like a GitHub Actions job that
// sets only the identity provider ID. An empty API key variable is the value of a repository secret that is not set.
func TestWorkloadIdentityEnvGitHubActions(t *testing.T) {
	clearAuthEnv(t)
	api := newFakeAPI(t)
	t.Setenv("DODO_PAYMENTS_API_KEY", "")
	t.Setenv(identityProviderIDEnv, testProviderID)
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", api.server.URL+"/oidc?api-version=2.0")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "request-token")
	api.call(t, providerConfig(t, map[string]string{"base_url": api.server.URL}, nil))
	if got := api.audience.Load(); got != api.server.URL {
		t.Fatalf("got audience %v, want the origin of the base URL", got)
	}
}

func TestWorkloadIdentityBlockGitHubActionsAudience(t *testing.T) {
	clearAuthEnv(t)
	api := newFakeAPI(t)
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", api.server.URL+"/oidc?api-version=2.0")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "request-token")
	api.call(t, providerConfig(t, map[string]string{"base_url": api.server.URL}, map[string]string{
		"identity_provider_id": testProviderID, "token_source": "github_actions", "audience": "custom-audience",
	}))
	if got := api.audience.Load(); got != "custom-audience" {
		t.Fatalf("got audience %v, want custom-audience", got)
	}
}

// TestWorkloadIdentityBlockFile also has an API key in the environment. The workload_identity block takes precedence.
func TestWorkloadIdentityBlockFile(t *testing.T) {
	clearAuthEnv(t)
	api := newFakeAPI(t)
	t.Setenv("DODO_PAYMENTS_API_KEY", "key-from-environment")
	t.Setenv(identityProviderIDEnv, testProviderID)
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte(testSubjectToken+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	api.call(t, providerConfig(t, map[string]string{"base_url": api.server.URL}, map[string]string{
		"token_source": "file", "token_file": tokenFile,
	}))
}

func TestAPIKeyFromEnvironmentWinsOverProviderIDFromEnvironment(t *testing.T) {
	clearAuthEnv(t)
	var authorization atomic.Value
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization.Store(r.Header.Get("Authorization"))
		if strings.HasPrefix(r.URL.Path, "/workload-identity/token") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte(`{}`))
	}))
	t.Cleanup(srv.Close)
	t.Setenv("DODO_PAYMENTS_API_KEY", "key-from-environment")
	t.Setenv(identityProviderIDEnv, testProviderID)
	resp := configure(t, providerConfig(t, map[string]string{"base_url": srv.URL}, nil))
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure: %v", resp.Diagnostics)
	}
	var raw []byte
	if err := resp.ResourceData.(*sdk.Client).Execute(context.Background(), http.MethodGet, "items", nil, &raw); err != nil {
		t.Fatalf("API call: %v", err)
	}
	if authorization.Load() != "Bearer key-from-environment" {
		t.Fatal("the request does not have the API key from the environment")
	}
}
