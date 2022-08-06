package main // import "src.techknowlogick.com/terraform-provider-gitea"

import (
	"git.uploadfilter24.eu/terraform-provider-gitea/gitea"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var Version = "development"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gitea.Provider})
}
