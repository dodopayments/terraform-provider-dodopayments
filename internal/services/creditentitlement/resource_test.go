// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package creditentitlement_test

import (
	"strings"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/dodopayments/terraform-provider-dodopayments/internal/acctest"
)

func TestAccCreditEntitlement_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))
	config := replacer.Replace("resource \"dodopayments_credit_entitlement\" \"test\" {\n  currency         = \"USD\"\n  name             = \"tf-acc-RANDOM\"\n  overage_enabled  = false\n  precision        = 2\n  price_per_unit   = \"10.5\"\n  rollover_enabled = false\n  unit             = \"credits\"\n}\n\ndata \"dodopayments_credit_entitlement\" \"test\" {\n  id = dodopayments_credit_entitlement.test.id\n}\n\ndata \"dodopayments_credit_entitlements\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_credit_entitlement.test]\n}\n")
	updateConfig := replacer.Replace("resource \"dodopayments_credit_entitlement\" \"test\" {\n  currency         = \"USD\"\n  name             = \"tf-acc-RANDOM-updated\"\n  overage_enabled  = false\n  precision        = 2\n  price_per_unit   = \"10.5\"\n  rollover_enabled = false\n  unit             = \"credits\"\n}\n\ndata \"dodopayments_credit_entitlement\" \"test\" {\n  id = dodopayments_credit_entitlement.test.id\n}\n\ndata \"dodopayments_credit_entitlements\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_credit_entitlement.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDestroy("dodopayments_credit_entitlement", "credit-entitlements/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_credit_entitlement.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_credit_entitlement.test", "id", "dodopayments_credit_entitlement.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_credit_entitlements.test", "items.#", "1"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentityValueMatchesState("dodopayments_credit_entitlement.test", tfjsonpath.New("id")),
				},
			},
			{
				Config: updateConfig,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{plancheck.ExpectResourceAction("dodopayments_credit_entitlement.test", plancheck.ResourceActionUpdate)},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				// Import by ID with the terraform import command.
				ResourceName:            "dodopayments_credit_entitlement.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			{
				// Import by resource identity with an import block. The plan after the import must be empty.
				ResourceName:    "dodopayments_credit_entitlement.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
			},
			{
				// Import with an import block and -generate-config-out. The generated configuration must plan with no changes.
				ResourceName:    "dodopayments_credit_entitlement.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				GenerateConfig:  true,
			},
		},
	})
}

// TestAccCreditEntitlement_query finds the object with terraform query. List resources need Terraform 1.14 or later.
func TestAccCreditEntitlement_query(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_14_0)},
		CheckDestroy:             acctest.CheckDestroy("dodopayments_credit_entitlement", "credit-entitlements/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{Config: replacer.Replace("resource \"dodopayments_credit_entitlement\" \"test\" {\n  currency         = \"USD\"\n  name             = \"tf-acc-RANDOM\"\n  overage_enabled  = false\n  precision        = 2\n  price_per_unit   = \"10.5\"\n  rollover_enabled = false\n  unit             = \"credits\"\n}\n\ndata \"dodopayments_credit_entitlement\" \"test\" {\n  id = dodopayments_credit_entitlement.test.id\n}\n\ndata \"dodopayments_credit_entitlements\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_credit_entitlement.test]\n}\n")},
			{
				Query:             true,
				Config:            "provider \"dodopayments\" {}\n\nlist \"dodopayments_credit_entitlement\" \"test\" {\n  provider = dodopayments\n}\n",
				QueryResultChecks: []querycheck.QueryResultCheck{querycheck.ExpectLengthAtLeast("dodopayments_credit_entitlement.test", 1)},
			},
		},
	})
}
