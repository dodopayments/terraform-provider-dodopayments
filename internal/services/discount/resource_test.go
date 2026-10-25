// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package discount_test

import (
	"strings"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/acctest"
)

func TestAccDiscount_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	config := strings.NewReplacer("RANDOM", strings.ToUpper(suffix)).Replace("resource \"dodopayments_discount\" \"test\" {\n  amount = 100\n  type = \"percentage\"\n}\n\ndata \"dodopayments_discount\" \"test\" {\n  id = dodopayments_discount.test.id\n}\n\ndata \"dodopayments_discounts\" \"test\" {\n  max_items = 1\n  depends_on = [dodopayments_discount.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_discount.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_discount.test", "id", "dodopayments_discount.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_discounts.test", "items.#", "1"),
				),
			},
			{
				ResourceName:            "dodopayments_discount.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
