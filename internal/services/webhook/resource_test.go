// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook_test

import (
	"strings"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/acctest"
)

func TestAccWebhook_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	config := strings.NewReplacer("RANDOM", strings.ToUpper(suffix)).Replace("resource \"dodopayments_webhook\" \"test\" {\n  url = \"https://example.com/tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_webhook\" \"test\" {\n  id = dodopayments_webhook.test.id\n}\n\ndata \"dodopayments_webhooks\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_webhook.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_webhook.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_webhook.test", "id", "dodopayments_webhook.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_webhooks.test", "items.#", "1"),
				),
			},
			{
				ResourceName:            "dodopayments_webhook.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"headers", "idempotency_key"},
			},
		},
	})
}
