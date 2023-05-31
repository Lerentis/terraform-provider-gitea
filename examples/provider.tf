terraform {
  required_providers {
    gitea = {
      source = "terraform.local/lerentis/gitea"
      version = "0.14.0"
    }
  }
}

provider "gitea" {
  base_url = var.gitea_url
  username = "lerentis"
  password = var.gitea_password
  #token    = var.gitea_token
}