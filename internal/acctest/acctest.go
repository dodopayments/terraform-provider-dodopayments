// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acctest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/provider"
)

// ProtoV6ProviderFactories starts the provider in the test process.
var ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"dodopayments": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// PreCheck stops the test when the environment does not have an API key for a test environment.
func PreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("DODO_PAYMENTS_API_KEY") == "" {
		t.Fatal("Set DODO_PAYMENTS_API_KEY to run acceptance tests.")
	}
	if os.Getenv("DODO_PAYMENTS_ENVIRONMENT") != "test_mode" {
		t.Fatal("Acceptance tests create real objects. Set DODO_PAYMENTS_ENVIRONMENT=test_mode.")
	}
}
