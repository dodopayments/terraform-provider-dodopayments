// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter_test

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

func TestAccMeter_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))
	config := replacer.Replace("resource \"dodopayments_meter\" \"test\" {\n  aggregation = {\n    type = \"count\"\n  }\n  event_name       = \"tf-acc-RANDOM\"\n  measurement_unit = \"tf-acc-RANDOM\"\n  name             = \"tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_meter\" \"test\" {\n  id = dodopayments_meter.test.id\n}\n\ndata \"dodopayments_meters\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_meter.test]\n}\n")
	updateConfig := replacer.Replace("resource \"dodopayments_meter\" \"test\" {\n  aggregation = {\n    type = \"count\"\n  }\n  event_name       = \"tf-acc-RANDOM\"\n  measurement_unit = \"tf-acc-RANDOM\"\n  name             = \"tf-acc-RANDOM-updated\"\n}\n\ndata \"dodopayments_meter\" \"test\" {\n  id = dodopayments_meter.test.id\n}\n\ndata \"dodopayments_meters\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_meter.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDestroy("dodopayments_meter", "meters/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_meter.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_meter.test", "id", "dodopayments_meter.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_meters.test", "items.#", "1"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentityValueMatchesState("dodopayments_meter.test", tfjsonpath.New("id")),
				},
			},
			{
				Config: updateConfig,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{plancheck.ExpectResourceAction("dodopayments_meter.test", plancheck.ResourceActionUpdate)},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				// Import by ID with the terraform import command.
				ResourceName:            "dodopayments_meter.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			{
				// Import by resource identity with an import block. The plan after the import must be empty.
				ResourceName:    "dodopayments_meter.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
			},
			{
				// Import with an import block and -generate-config-out. The generated configuration must plan with no changes.
				ResourceName:    "dodopayments_meter.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				GenerateConfig:  true,
			},
		},
	})
}

// TestAccMeter_query finds the object with terraform query. List resources need Terraform 1.14 or later.
func TestAccMeter_query(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_14_0)},
		CheckDestroy:             acctest.CheckDestroy("dodopayments_meter", "meters/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{Config: replacer.Replace("resource \"dodopayments_meter\" \"test\" {\n  aggregation = {\n    type = \"count\"\n  }\n  event_name       = \"tf-acc-RANDOM\"\n  measurement_unit = \"tf-acc-RANDOM\"\n  name             = \"tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_meter\" \"test\" {\n  id = dodopayments_meter.test.id\n}\n\ndata \"dodopayments_meters\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_meter.test]\n}\n")},
			{
				Query:             true,
				Config:            "provider \"dodopayments\" {}\n\nlist \"dodopayments_meter\" \"test\" {\n  provider = dodopayments\n}\n",
				QueryResultChecks: []querycheck.QueryResultCheck{querycheck.ExpectLengthAtLeast("dodopayments_meter.test", 1)},
			},
		},
	})
}
