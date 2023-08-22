package gitea

import (
	"fmt"
	"strconv"

	"code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	TokenUsername  string = "username"
	TokenName      string = "name"
	TokenHash      string = "token"
	TokenLastEight string = "last_eight"
	TokenScope     string = "scope"
)

func searchTokenById(c *gitea.Client, id int64) (res *gitea.AccessToken, err error) {
	page := 1

	for {
		tokens, _, err := c.ListAccessTokens(gitea.ListAccessTokensOptions{
			ListOptions: gitea.ListOptions{
				Page:     page,
				PageSize: 50,
			},
		})
		if err != nil {
			return nil, err
		}

		if len(tokens) == 0 {
			return nil, fmt.Errorf("Token with ID %d could not be found", id)
		}

		for _, token := range tokens {
			if token.ID == id {
				return token, nil
			}
		}

		page += 1
	}
}

func resourceTokenCreate(d *schema.ResourceData, meta interface{}) (err error) {

	client := meta.(*gitea.Client)

	var opt gitea.CreateAccessTokenOption
	opt.Name = d.Get(TokenName).(string)

	if err := client.CheckServerVersionConstraint(">= 1.20.0"); err != nil {
		opt.Scopes = []gitea.AccessTokenScope{}
		for _, scope := range d.Get(TokenScope).([]interface{}) {
			opt.Scopes = append(opt.Scopes, gitea.AccessTokenScope(scope.(string)))
		}
	}

	token, _, err := client.CreateAccessToken(opt)

	if err != nil {
		return err
	}

	err = setTokenResourceData(token, d, meta)

	return
}

func resourceTokenRead(d *schema.ResourceData, meta interface{}) (err error) {

	client := meta.(*gitea.Client)

	var token *gitea.AccessToken

	id, err := strconv.ParseInt(d.Id(), 10, 64)

	token, err = searchTokenById(client, id)

	if err != nil {
		return err
	}

	err = setTokenResourceData(token, d, meta)

	return
}

func resourceTokenDelete(d *schema.ResourceData, meta interface{}) (err error) {

	client := meta.(*gitea.Client)
	var resp *gitea.Response

	resp, err = client.DeleteAccessToken(d.Get(TokenName).(string))

	if err != nil {
		if resp.StatusCode == 404 {
			return
		} else {
			return err
		}
	}

	return
}

func setTokenResourceData(token *gitea.AccessToken, d *schema.ResourceData, meta interface{}) (err error) {

	client := meta.(*gitea.Client)

	d.SetId(fmt.Sprintf("%d", token.ID))
	d.Set(TokenName, token.Name)

	if err := client.CheckServerVersionConstraint(">= 1.20.0"); err != nil {
		var scopeList []string
		for _, scope := range token.Scopes {
			scopeList = append(scopeList, string(scope))
		}
		d.Set(TokenScope, scopeList)
	}

	if token.Token != "" {
		d.Set(TokenHash, token.Token)
	}
	d.Set(TokenLastEight, token.TokenLastEight)

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
				Description: "The owner of the Access Token",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Access Token",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The actual Access Token",
			},
			"scope": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				ForceNew:    true,
				Description: "The Access Token scope (all,public-only,read:activitypub,write:activitypub,read:admin,write:admin,read:misc,write:misc,read:notification,write:notification,read:organization,write:organization,read:package,write:package,read:issue,write:issue,read:repository,write:repository,read:user,write:user)",
			},
			"last_eight": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Description: "`gitea_token` manages gitea Access Tokens.\n\n" +
			"Due to upstream limitations (see https://gitea.com/gitea/go-sdk/issues/610) this resource\n" +
			"can only be used with username/password provider configuration.\n\n" +
			"WARNING:\n" +
			"Tokens will be stored in the terraform state!",
	}
}
