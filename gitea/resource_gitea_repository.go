package gitea

import (
	"fmt"
	"strconv"

	"code.gitea.io/sdk/gitea"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	repoOwner                 string = "username"
	repoName                  string = "name"
	repoDescription           string = "description"
	repoPrivateFlag           string = "private"
	repoIssueLabels           string = "issue_labels"
	repoAutoInit              string = "auto_init"
	repoTemplate              string = "repo_template"
	repoGitignores            string = "gitignores"
	repoLicense               string = "license"
	repoReadme                string = "readme"
	repoDefaultBranch         string = "default_branch"
	repoWebsite               string = "website"
	repoIssues                string = "has_issues"
	repoWiki                  string = "has_wiki"
	repoPrs                   string = "has_pull_requests"
	repoProjects              string = "has_projects"
	repoIgnoreWhitespace      string = "ignore_whitespace_conflicts"
	repoAllowMerge            string = "allow_merge_commits"
	repoAllowRebase           string = "allow_rebase"
	repoAllowRebaseMerge      string = "allow_rebase_explicit"
	repoAllowSquash           string = "allow_squash_merge"
	repoAchived               string = "archived"
	repoAllowManualMerge      string = "allow_manual_merge"
	repoAutodetectManualMerge string = "autodetect_manual_merge"
)

func resourceRepoRead(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*gitea.Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)

	if err != nil {
		return err
	}

	repo, _, err := client.GetRepoByID(id)

	if err != nil {
		return err
	}

	err = setRepoResourceData(repo, d)

	return
}

func resourceRepoCreate(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*gitea.Client)

	var repo *gitea.Repository

	opts := gitea.CreateRepoOption{
		Name:          d.Get(repoName).(string),
		Description:   d.Get(repoDescription).(string),
		Private:       d.Get(repoPrivateFlag).(bool),
		IssueLabels:   d.Get(repoIssueLabels).(string),
		AutoInit:      d.Get(repoAutoInit).(bool),
		Template:      d.Get(repoTemplate).(bool),
		Gitignores:    d.Get(repoGitignores).(string),
		License:       d.Get(repoLicense).(string),
		Readme:        d.Get(repoReadme).(string),
		DefaultBranch: d.Get(repoDefaultBranch).(string),
		TrustModel:    "default",
	}
	repo, _, err = client.CreateRepo(opts)

	if err != nil {
		return
	}

	err = setRepoResourceData(repo, d)

	return
}

func resourceRepoUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*gitea.Client)

	var repo *gitea.Repository

	var name string = d.Get(repoName).(string)
	var description string = d.Get(repoDescription).(string)
	var website string = d.Get(repoDescription).(string)
	var private bool = d.Get(repoPrivateFlag).(bool)
	var template bool = d.Get(repoTemplate).(bool)
	var hasIssues bool = d.Get(repoIssues).(bool)
	var hasWiki bool = d.Get(repoWiki).(bool)
	var defaultBranch string = d.Get(repoDefaultBranch).(string)
	var hasPRs bool = d.Get(repoPrs).(bool)
	var hasProjects bool = d.Get(repoProjects).(bool)
	var ignoreWhitespaceConflicts bool = d.Get(repoIgnoreWhitespace).(bool)
	var allowMerge bool = d.Get(repoAllowMerge).(bool)
	var allowRebase bool = d.Get(repoAllowRebase).(bool)
	var allowRebaseMerge bool = d.Get(repoAllowRebaseMerge).(bool)
	var allowSquash bool = d.Get(repoAllowSquash).(bool)
	var archived bool = d.Get(repoAchived).(bool)
	var allowManualMerge bool = d.Get(repoAllowManualMerge).(bool)
	var autodetectManualMerge bool = d.Get(repoAutodetectManualMerge).(bool)

	opts := gitea.EditRepoOption{
		Name:                      &name,
		Description:               &description,
		Website:                   &website,
		Private:                   &private,
		Template:                  &template,
		HasIssues:                 &hasIssues,
		HasWiki:                   &hasWiki,
		DefaultBranch:             &defaultBranch,
		HasPullRequests:           &hasPRs,
		HasProjects:               &hasProjects,
		IgnoreWhitespaceConflicts: &ignoreWhitespaceConflicts,
		AllowMerge:                &allowMerge,
		AllowRebase:               &allowRebase,
		AllowRebaseMerge:          &allowRebaseMerge,
		AllowSquash:               &allowSquash,
		Archived:                  &archived,
		AllowManualMerge:          &allowManualMerge,
		AutodetectManualMerge:     &autodetectManualMerge,
	}

	repo, _, err = client.EditRepo(d.Get(repoOwner).(string), d.Get(repoName).(string), opts)

	if err != nil {
		return err
	}
	err = setRepoResourceData(repo, d)

	return

}

func respurceRepoDelete(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*gitea.Client)

	client.DeleteRepo(d.Get(repoOwner).(string), d.Get(repoName).(string))

	return
}

func setRepoResourceData(repo *gitea.Repository, d *schema.ResourceData) (err error) {
	d.SetId(fmt.Sprintf("%d", repo.ID))
	d.Set("name", repo.Name)
	d.Set("description", repo.Description)
	d.Set("full_name", repo.FullName)
	d.Set("description", repo.Description)
	d.Set("private", repo.Private)
	d.Set("fork", repo.Fork)
	d.Set("mirror", repo.Mirror)
	d.Set("size", repo.Size)
	d.Set("html_url", repo.HTMLURL)
	d.Set("ssh_url", repo.SSHURL)
	d.Set("clone_url", repo.CloneURL)
	d.Set("website", repo.Website)
	d.Set("stars", repo.Stars)
	d.Set("forks", repo.Forks)
	d.Set("watchers", repo.Watchers)
	d.Set("open_issue_count", repo.OpenIssues)
	d.Set("default_branch", repo.DefaultBranch)
	d.Set("created", repo.Created)
	d.Set("updated", repo.Updated)
	d.Set("permission_admin", repo.Permissions.Admin)
	d.Set("permission_push", repo.Permissions.Push)
	d.Set("permission_pull", repo.Permissions.Pull)

	return
}

func resourceGiteaRepository() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRepoRead,
		Create: resourceRepoCreate,
		Update: resourceRepoUpdate,
		Delete: respurceRepoDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_init": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"repo_template": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  false,
			},
			"issue_labels": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "Default",
			},
			"gitignores": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"license": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"readme": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"private": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"default_branch": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "main",
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission_admin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permission_push": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permission_pull": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"website": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"has_issues": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"has_wiki": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"has_pull_requests": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"has_projects": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"ignore_whitespace_conflicts": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"allow_merge_commits": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"allow_rebase": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"allow_rebase_explicit": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"allow_squash_merge": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"archived": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  false,
			},
			"allow_manual_merge": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"autodetect_manual_merge": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
		},
		Description: "Handling Repository resources",
	}
}
