// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/echoprovider"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	sdk "github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/provider"
)

// ProtoV6ProviderFactories starts the provider in the test process.
var ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"dodopayments": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// ProtoV6ProviderFactoriesWithEcho also starts the echo provider. A test uses it to examine an ephemeral value.
func ProtoV6ProviderFactoriesWithEcho() map[string]func() (tfprotov6.ProviderServer, error) {
	out := maps.Clone(ProtoV6ProviderFactories)
	out["echo"] = echoprovider.NewProviderServer()
	return out
}

// PreCheck stops the test when the environment does not have an API key or a workload identity provider for a test environment.
func PreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("DODO_PAYMENTS_API_KEY") == "" && os.Getenv("DODO_PAYMENTS_IDENTITY_PROVIDER_ID") == "" {
		t.Fatal("Set DODO_PAYMENTS_API_KEY or DODO_PAYMENTS_IDENTITY_PROVIDER_ID to run acceptance tests.")
	}
	if os.Getenv("DODO_PAYMENTS_ENVIRONMENT") != "test_mode" {
		t.Fatal("Acceptance tests create real objects. Set DODO_PAYMENTS_ENVIRONMENT=test_mode.")
	}
}

// AuthOption uses the API key. Without an API key, it uses workload identity with the token source that the environment has.
func AuthOption() option.RequestOption {
	if key := os.Getenv("DODO_PAYMENTS_API_KEY"); key != "" {
		return option.WithBearerToken(key)
	}
	return option.WithWorkloadIdentity(option.WorkloadIdentity{})
}

// client is shared, so that the tests use one access token.
var client = sync.OnceValue(func() *sdk.Client {
	return sdk.NewClient(
		AuthOption(),
		option.WithBaseURL("https://test.dodopayments.com"),
		option.WithMaxRetries(5),
		option.WithMiddleware(engine.RateLimitMiddleware),
	)
})

// CheckDestroy returns a check that each object of the resource type is gone or archived in the API.
// readPath has a %s verb for each parent attribute and then for the ID.
func CheckDestroy(resourceType, readPath string, parents []string, archivedProperty string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		c := client()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			args := make([]any, 0, len(parents)+1)
			for _, p := range parents {
				args = append(args, url.PathEscape(rs.Primary.Attributes[p]))
			}
			args = append(args, url.PathEscape(rs.Primary.ID))
			var raw []byte
			err := c.Execute(context.Background(), http.MethodGet, fmt.Sprintf(readPath, args...), nil, &raw)
			var apiErr *sdk.Error
			if errors.As(err, &apiErr) && (apiErr.StatusCode == http.StatusNotFound || apiErr.StatusCode == http.StatusGone) {
				continue
			}
			if err != nil {
				return err
			}
			if archivedProperty != "" {
				var body map[string]any
				if json.Unmarshal(raw, &body) == nil && body[archivedProperty] != nil {
					continue
				}
			}
			return fmt.Errorf("%s %s still exists after destroy", resourceType, rs.Primary.ID)
		}
		return nil
	}
}
