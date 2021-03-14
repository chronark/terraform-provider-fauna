package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCollection(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCollection,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"fauna_collection.foo", "name", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccResourceCollection = `
resource "fauna_collection" "foo" {
  name = "bar"
}
`
