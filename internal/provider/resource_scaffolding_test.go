package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"mime.foo", "input", regexp.MustCompile("^あああ$")),
				),
			},
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"mime.foo", "value", regexp.MustCompile("^=\\?utf-8\\?b\\?44GC44GC44GC\\?=$")),
				),
			},
		},
	})
}

const testAccResourceScaffolding = `
resource "mime" "foo" {
  input = "あああ"
}
`
