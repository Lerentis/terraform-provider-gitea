terraform {
  required_providers {
    gitea = {
      source = "terraform.local/lerentis/gitea"
      version = "0.5.0"
    }
  }
}

provider "gitea" {
  base_url = var.gitea_url
  token    = var.gitea_token
}