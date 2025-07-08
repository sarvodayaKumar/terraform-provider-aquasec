package aquasec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAquasecresourceRegistryTypeAny(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	url := "https://docker.io"
	rtype := "HUB"
	username := ""
	password := ""
	autopull := false
	scanner_type := "any"
	description := "Terrafrom-test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("aquasec_integration_registry.new"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAquasecRegistry(name, url, rtype, username, password, autopull, scanner_type, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAquasecRegistryExists("aquasec_integration_registry.new"),
				),
			},
			{
				ResourceName:            "aquasec_integration_registry.new",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"prefixes", "scanner_name", "last_updated"}, //TODO: implement read prefixes
			},
		},
	})
}

func TestAquasecresourceRegistryTypeSpecific(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	url := "https://docker.io"
	rtype := "HUB"
	username := ""
	password := ""
	autopull := false
	scanner_type := "specific"
	scanner_group_name := "terraform-test"
	description := "Terrafrom-test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("aquasec_integration_registry.new"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAquasecRegistryTypeSpecific(name, url, rtype, username, password, autopull, scanner_type, description, scanner_group_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAquasecRegistryExists("aquasec_integration_registry.new"),
				),
			},
			{
				ResourceName:            "aquasec_integration_registry.new",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"prefixes", "scanner_name", "last_updated"}, //TODO: implement read prefixes
			},
		},
	})
}

func testAccCheckAquasecRegistry(name string, url string, rtype string, username string, password string, autopull bool, scanner_type string, description string) string {
	return fmt.Sprintf(`
	resource "aquasec_integration_registry" "new" {
		name = "%s"
		url = "%s"
		type = "%s"
		username = "%s"
		password = "%s"
		auto_pull = "%v"
		scanner_type = "%s"
		description = "%s"
	}`, name, url, rtype, username, password, autopull, scanner_type, description)

}

func testAccCheckAquasecRegistryTypeSpecific(name string, url string, rtype string, username string, password string, autopull bool, scanner_type string, description string, scanner_group_name string) string {
	return fmt.Sprintf(`
	resource "aquasec_integration_registry" "new" {
		name = "%s"
		url = "%s"
		type = "%s"
		username = "%s"
		password = "%s"
		auto_pull = "%v"
		scanner_type = "%s"
		description = "%s"
		scanner_group_name = "%s"
	}`, name, url, rtype, username, password, autopull, scanner_type, description, scanner_group_name)

}

func testAccCheckAquasecRegistryExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("ID for %s in state", n)
		}

		return nil
	}
}
