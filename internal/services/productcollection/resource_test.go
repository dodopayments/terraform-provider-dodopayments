// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package productcollection_test

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

func TestAccProductCollection_basic(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))
	config := replacer.Replace("resource \"dodopayments_addon\" \"test\" {\n  currency     = \"USD\"\n  name         = \"tf-acc-RANDOM\"\n  price        = 100\n  tax_category = \"digital_products\"\n}\n\nresource \"dodopayments_credit_entitlement\" \"test\" {\n  currency         = \"USD\"\n  name             = \"tf-acc-RANDOM\"\n  overage_enabled  = false\n  precision        = 2\n  price_per_unit   = \"10.5\"\n  rollover_enabled = false\n  unit             = \"credits\"\n}\n\nresource \"dodopayments_entitlement\" \"test\" {\n  integration_config = {\n    feature_id   = \"tf-acc-RANDOM\"\n    feature_type = \"boolean\"\n  }\n  integration_type = \"feature_flag\"\n  metadata = {\n    tier = \"pro\"\n  }\n  name = \"tf-acc-RANDOM\"\n}\n\nresource \"dodopayments_product\" \"test\" {\n  addons = [dodopayments_addon.test.id]\n  credit_entitlements = [{\n    credit_entitlement_id = dodopayments_credit_entitlement.test.id\n    credits_amount        = \"100.0\"\n  }]\n  entitlements = [{\n    entitlement_id = dodopayments_entitlement.test.id\n  }]\n  metadata = {\n    plan = \"pro\"\n  }\n  name = \"tf-acc-RANDOM\"\n  price = {\n    type                         = \"recurring_price\"\n    currency                     = \"USD\"\n    price                        = 1000\n    payment_frequency_count      = 1\n    payment_frequency_interval   = \"Month\"\n    subscription_period_count    = 1\n    subscription_period_interval = \"Year\"\n  }\n  pricing_mode = \"by_currency\"\n  tax_category = \"saas\"\n}\n\nresource \"dodopayments_product_collection\" \"test\" {\n  groups = [{\n    products = [{\n      product_id = dodopayments_product.test.id\n    }]\n  }]\n  name = \"tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_product_collection\" \"test\" {\n  id = dodopayments_product_collection.test.id\n}\n\ndata \"dodopayments_product_collections\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_product_collection.test]\n}\n")
	updateConfig := replacer.Replace("resource \"dodopayments_addon\" \"test\" {\n  currency     = \"USD\"\n  name         = \"tf-acc-RANDOM\"\n  price        = 100\n  tax_category = \"digital_products\"\n}\n\nresource \"dodopayments_credit_entitlement\" \"test\" {\n  currency         = \"USD\"\n  name             = \"tf-acc-RANDOM\"\n  overage_enabled  = false\n  precision        = 2\n  price_per_unit   = \"10.5\"\n  rollover_enabled = false\n  unit             = \"credits\"\n}\n\nresource \"dodopayments_entitlement\" \"test\" {\n  integration_config = {\n    feature_id   = \"tf-acc-RANDOM\"\n    feature_type = \"boolean\"\n  }\n  integration_type = \"feature_flag\"\n  metadata = {\n    tier = \"pro\"\n  }\n  name = \"tf-acc-RANDOM\"\n}\n\nresource \"dodopayments_product\" \"test\" {\n  addons = [dodopayments_addon.test.id]\n  credit_entitlements = [{\n    credit_entitlement_id = dodopayments_credit_entitlement.test.id\n    credits_amount        = \"100.0\"\n  }]\n  entitlements = [{\n    entitlement_id = dodopayments_entitlement.test.id\n  }]\n  metadata = {\n    plan = \"pro\"\n  }\n  name = \"tf-acc-RANDOM\"\n  price = {\n    type                         = \"recurring_price\"\n    currency                     = \"USD\"\n    price                        = 1000\n    payment_frequency_count      = 1\n    payment_frequency_interval   = \"Month\"\n    subscription_period_count    = 1\n    subscription_period_interval = \"Year\"\n  }\n  pricing_mode = \"by_currency\"\n  tax_category = \"saas\"\n}\n\nresource \"dodopayments_product_collection\" \"test\" {\n  groups = [{\n    products = [{\n      product_id = dodopayments_product.test.id\n    }]\n  }]\n  name = \"tf-acc-RANDOM-updated\"\n}\n\ndata \"dodopayments_product_collection\" \"test\" {\n  id = dodopayments_product_collection.test.id\n}\n\ndata \"dodopayments_product_collections\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_product_collection.test]\n}\n")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDestroy("dodopayments_product_collection", "product-collections/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("dodopayments_product_collection.test", "id"),
					resource.TestCheckResourceAttrPair("data.dodopayments_product_collection.test", "id", "dodopayments_product_collection.test", "id"),
					resource.TestCheckResourceAttr("data.dodopayments_product_collections.test", "items.#", "1"),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentityValueMatchesState("dodopayments_product_collection.test", tfjsonpath.New("id")),
				},
			},
			{
				Config: updateConfig,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{plancheck.ExpectResourceAction("dodopayments_product_collection.test", plancheck.ResourceActionUpdate)},
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				// Import by ID with the terraform import command.
				ResourceName:            "dodopayments_product_collection.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			{
				// Import by resource identity with an import block. The plan after the import must be empty.
				ResourceName:    "dodopayments_product_collection.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithResourceIdentity,
			},
			{
				// Import with an import block and -generate-config-out. The generated configuration must plan with no changes.
				ResourceName:    "dodopayments_product_collection.test",
				ImportState:     true,
				ImportStateKind: resource.ImportBlockWithID,
				GenerateConfig:  true,
			},
		},
	})
}

// TestAccProductCollection_query finds the object with terraform query. List resources need Terraform 1.14 or later.
func TestAccProductCollection_query(t *testing.T) {
	suffix := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	replacer := strings.NewReplacer("RANDOM", strings.ToUpper(suffix))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_14_0)},
		CheckDestroy:             acctest.CheckDestroy("dodopayments_product_collection", "product-collections/%s", []string{}, ""),
		Steps: []resource.TestStep{
			{Config: replacer.Replace("resource \"dodopayments_addon\" \"test\" {\n  currency     = \"USD\"\n  name         = \"tf-acc-RANDOM\"\n  price        = 100\n  tax_category = \"digital_products\"\n}\n\nresource \"dodopayments_credit_entitlement\" \"test\" {\n  currency         = \"USD\"\n  name             = \"tf-acc-RANDOM\"\n  overage_enabled  = false\n  precision        = 2\n  price_per_unit   = \"10.5\"\n  rollover_enabled = false\n  unit             = \"credits\"\n}\n\nresource \"dodopayments_entitlement\" \"test\" {\n  integration_config = {\n    feature_id   = \"tf-acc-RANDOM\"\n    feature_type = \"boolean\"\n  }\n  integration_type = \"feature_flag\"\n  metadata = {\n    tier = \"pro\"\n  }\n  name = \"tf-acc-RANDOM\"\n}\n\nresource \"dodopayments_product\" \"test\" {\n  addons = [dodopayments_addon.test.id]\n  credit_entitlements = [{\n    credit_entitlement_id = dodopayments_credit_entitlement.test.id\n    credits_amount        = \"100.0\"\n  }]\n  entitlements = [{\n    entitlement_id = dodopayments_entitlement.test.id\n  }]\n  metadata = {\n    plan = \"pro\"\n  }\n  name = \"tf-acc-RANDOM\"\n  price = {\n    type                         = \"recurring_price\"\n    currency                     = \"USD\"\n    price                        = 1000\n    payment_frequency_count      = 1\n    payment_frequency_interval   = \"Month\"\n    subscription_period_count    = 1\n    subscription_period_interval = \"Year\"\n  }\n  pricing_mode = \"by_currency\"\n  tax_category = \"saas\"\n}\n\nresource \"dodopayments_product_collection\" \"test\" {\n  groups = [{\n    products = [{\n      product_id = dodopayments_product.test.id\n    }]\n  }]\n  name = \"tf-acc-RANDOM\"\n}\n\ndata \"dodopayments_product_collection\" \"test\" {\n  id = dodopayments_product_collection.test.id\n}\n\ndata \"dodopayments_product_collections\" \"test\" {\n  max_items  = 1\n  depends_on = [dodopayments_product_collection.test]\n}\n")},
			{
				Query:             true,
				Config:            "provider \"dodopayments\" {}\n\nlist \"dodopayments_product_collection\" \"test\" {\n  provider = dodopayments\n}\n",
				QueryResultChecks: []querycheck.QueryResultCheck{querycheck.ExpectLengthAtLeast("dodopayments_product_collection.test", 1)},
			},
		},
	})
}
