// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook_test

import (
	"strings"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/acctest"
)

func TestAccWebhook_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))
	config := replacer.Replace("resource \"dodopayments_webhook\" \"test\" {\n  url = \"https://example.com/tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_webhook\" \"test\" {\n  id = dodopayments_webhook.test.id\n}\n\ndata \"dodopayments_webhooks\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_webhook.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithEcho(),
		CheckDestroy:             acctest.CheckDestroy("dodopayments_webhook", "webhooks/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_webhook.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_webhook.test", "id", "dodopayments_webhook.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_webhooks.test", "items.#", "1"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentityValueMatchesState("dodopayments_webhook.test", tfjsonpath.New("id")),
				},
			},
			{
				// The echo provider copies the ephemeral value to its state. Terraform does not store the value of dodopayments_webhook_secret.
				Config: replacer.Replace("resource \"dodopayments_webhook\" \"test\" {\n  url = \"https://example.com/tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_webhook\" \"test\" {\n  id = dodopayments_webhook.test.id\n}\n\ndata \"dodopayments_webhooks\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_webhook.test]\n}\n\nephemeral \"dodopayments_webhook_secret\" \"test\" {\n  webhook_id = dodopayments_webhook.test.id\n}\n\nprovider \"echo\" {\n  data = ephemeral.dodopayments_webhook_secret.test\n}\n\nresource \"echo\" \"test\" {}\n"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("echo.test", tfjsonpath.New("data").AtMapKey("secret"), knownvalue.NotNull()),
				},
			},
			{
				// Import by ID with the terraform import command.
				ResourceName:            "dodopayments_webhook.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"headers", "idempotency_key"},
			},
			{
				// Import by resource identity with an import block. The plan after the import must be empty.
				ResourceName:    "dodopayments_webhook.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
			},
		},
	})
}

// TestAccWebhook_query finds the object with terraform query. List resources need Terraform 1.14 or later.
func TestAccWebhook_query(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_14_0)},
		CheckDestroy:             acctest.CheckDestroy("dodopayments_webhook", "webhooks/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{Config: replacer.Replace("resource \"dodopayments_webhook\" \"test\" {\n  url = \"https://example.com/tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_webhook\" \"test\" {\n  id = dodopayments_webhook.test.id\n}\n\ndata \"dodopayments_webhooks\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_webhook.test]\n}\n")},
			{
				Query:             true,
				Config:            "provider \"dodopayments\" {}\n\nlist \"dodopayments_webhook\" \"test\" {\n  provider = dodopayments\n}\n",
				QueryResultChecks: []querycheck.QueryResultCheck{querycheck.ExpectLengthAtLeast("dodopayments_webhook.test", 1)},
			},
		},
	})
}
