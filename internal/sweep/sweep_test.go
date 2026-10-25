// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sweep

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	sdk "github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/engine"
)

// markers are in a string attribute of each object that the acceptance tests create.
var markers = []string{"tf-acc-", "TFACC"}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func client() (*sdk.Client, error) {
	if os.Getenv("DODO_PAYMENTS_ENVIRONMENT") != "test_mode" {
		return nil, fmt.Errorf("sweepers delete objects. Set DODO_PAYMENTS_ENVIRONMENT=test_mode.")
	}
	key := os.Getenv("DODO_PAYMENTS_API_KEY")
	if key == "" {
		return nil, fmt.Errorf("set DODO_PAYMENTS_API_KEY to run the sweepers.")
	}
	return sdk.NewClient(option.WithBearerToken(key), option.WithBaseURL("https://test.dodopayments.com")), nil
}

func addSweeper(spec engine.SweepSpec) {
	resource.AddTestSweepers(spec.Name, &resource.Sweeper{
		Name: spec.Name,
		F: func(string) error {
			c, err := client()
			if err != nil {
				return err
			}
			return engine.Sweep(context.Background(), c, spec, markers...)
		},
	})
}

func init() {
	addSweeper(engine.SweepSpec{
		Name:         "dodopayments_discount",
		IDProperty:   "discount_id",
		DeleteMethod: "DELETE",
		DeletePath:   "discounts/%s",
		List: &engine.ListDataSourceSpec{
			ListPath:        "discounts",
			ItemsProperty:   "items",
			PageNumberParam: "page_number",
			PageSizeParam:   "page_size",
			CursorParam:     "",
			CursorField:     "",
			LastPageField:   "",
		},
	})
	addSweeper(engine.SweepSpec{
		Name:         "dodopayments_addon",
		IDProperty:   "id",
		DeleteMethod: "DELETE",
		DeletePath:   "addons/%s",
		List: &engine.ListDataSourceSpec{
			ListPath:        "addons",
			ItemsProperty:   "items",
			PageNumberParam: "page_number",
			PageSizeParam:   "page_size",
			CursorParam:     "",
			CursorField:     "",
			LastPageField:   "",
		},
	})
	addSweeper(engine.SweepSpec{
		Name:         "dodopayments_webhook",
		IDProperty:   "id",
		DeleteMethod: "DELETE",
		DeletePath:   "webhooks/%s",
		List: &engine.ListDataSourceSpec{
			ListPath:        "webhooks",
			ItemsProperty:   "data",
			PageNumberParam: "",
			PageSizeParam:   "limit",
			CursorParam:     "iterator",
			CursorField:     "iterator",
			LastPageField:   "done",
		},
	})
	addSweeper(engine.SweepSpec{
		Name:         "dodopayments_meter",
		IDProperty:   "id",
		DeleteMethod: "DELETE",
		DeletePath:   "meters/%s",
		List: &engine.ListDataSourceSpec{
			ListPath:        "meters",
			ItemsProperty:   "items",
			PageNumberParam: "page_number",
			PageSizeParam:   "page_size",
			CursorParam:     "",
			CursorField:     "",
			LastPageField:   "",
		},
	})
}
