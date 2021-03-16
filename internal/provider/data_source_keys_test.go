package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceScaffolding(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("data.fauna_keys.foo", ""),
				),
			},
		},
	})
}

const testAccDataSourceScaffolding = `data "fauna_keys" "foo" {}`
