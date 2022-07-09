package gitea

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccResourceGiteaUser_basic(t *testing.T) {
	name := fmt.Sprintf("user-%d", 1)
	mail := fmt.Sprintf("%s@test.org", name)
	fqrn := fmt.Sprintf("gitea_user.%s", name)

	userSimple := fmt.Sprintf(`
		resource "gitea_user" "%s" {
		    username = "%s"
			login_name = "%s"
			email = "%s"
			password = "Geheim1!"

		}
		`, name, name, name, mail)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:       userSimple,
				ResourceName: fqrn,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fqrn, "username", name),
				),
			},
		},
	})
}

func testAccCheckExampleResourceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	//conn := testAccProvider.Meta().(*ExampleClient)

	// loop through the resources in state, verifying each widget
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "example_widget" {
			continue
		}

		// Retrieve our widget by referencing it's state ID for API lookup
		//request := &example.DescribeWidgets{
		//	IDs: []string{rs.Primary.ID},
		//}

		//response, err := conn.DescribeWidgets(request)
		//if err == nil {
		//	if len(response.Widgets) > 0 && *response.Widgets[0].ID == rs.Primary.ID {
		//		return fmt.Errorf("Widget (%s) still exists.", rs.Primary.ID)
		//	}
		//	return nil
		//}

		// If the error is equivalent to 404 not found, the widget is destroyed.
		// Otherwise return the error
		//if !strings.Contains(err.Error(), "Widget not found") {
		//	return err
		//}
	}

	return nil
}

func testAccResourceGiteaUserSimple(fqrn string, name string, mail string) string {
	return fmt.Sprintf(`
	resource "gitea_user" "%s" {
		username = "%s"
		login_name = "%s"
		email = "%s"
		password = "Geheim1!"

	}
	`, fqrn, name, name, mail)
}
