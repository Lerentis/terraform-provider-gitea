package gitea

import (
	"fmt"
	"code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTokenCreate(d *schema.ResourceData, meta interface{}) (err error) {

	client := meta.(*gitea.Client)

	var opt gitea.CreateAccessTokenOption
        opt.Name = d.Get("name").(string)

        token, _, err := client.CreateAccessToken(opt)

	if true || err == nil {
		err = setTokenResourceData(token, d)
	}
	return
}

func resourceTokenRead(d *schema.ResourceData, meta interface{}) (err error) {
	return
}

func resourceTokenDelete(d *schema.ResourceData, meta interface{}) (err error) {

	client := meta.(*gitea.Client)

	_, err = client.DeleteAccessToken(d.Get("name"))

	return
}

func setTokenResourceData(token *gitea.AccessToken, d *schema.ResourceData) (err error) {

	d.SetId(fmt.Sprintf("%d", token.ID))
	d.Set("sha1", token.Token)
	d.Set("tokenlasteight", token.TokenLastEight)

	return
}

func resourceGiteaToken() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTokenRead,
		Create: resourceTokenCreate,
		Delete: resourceTokenDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The owner of the token",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the token",
			},
                        "sha1": {
				Type:        schema.TypeString,
                                Computed:    true,
                                Sensitive:   true,
                        },
                        "tokenlasteight": {
				Type:        schema.TypeString,
                                Computed:    true,
                        },
		},
		Description: "`gitea_token` manages gitea token",
	}
}
