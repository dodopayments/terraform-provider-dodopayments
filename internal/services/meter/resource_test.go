// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter_test

import (
	"strings"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/acctest"
)

func TestAccMeter_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	config := strings.NewReplacer("RANDOM", strings.ToUpper(suffix)).Replace("resource \"dodopayments_meter\" \"test\" {\n  aggregation = {\n    type = \"count\"\n  }\n  event_name = \"tf-acc-RANDOM\"\n  measurement_unit = \"tf-acc-RANDOM\"\n  name = \"tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_meter\" \"test\" {\n  id = dodopayments_meter.test.id\n}\n\ndata \"dodopayments_meters\" \"test\" {\n  max_items = 1\n  depends_on = [dodopayments_meter.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_meter.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_meter.test", "id", "dodopayments_meter.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_meters.test", "items.#", "1"),
				),
			},
			{
				ResourceName:            "dodopayments_meter.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
