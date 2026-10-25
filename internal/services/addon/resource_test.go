// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addon_test

import (
	"strings"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/acctest"
)

func TestAccAddon_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	config := strings.NewReplacer("RANDOM", strings.ToUpper(suffix)).Replace("resource \"dodopayments_addon\" \"test\" {\n  currency     = \"AED\"\n  name         = \"tf-acc-RANDOM\"\n  price        = 100\n  tax_category = \"digital_products\"\n}\n\ndata \"dodopayments_addon\" \"test\" {\n  id = dodopayments_addon.test.id\n}\n\ndata \"dodopayments_addons\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_addon.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_addon.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_addon.test", "id", "dodopayments_addon.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_addons.test", "items.#", "1"),
				),
			},
			{
				ResourceName:            "dodopayments_addon.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
