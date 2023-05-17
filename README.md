# terraform-provider-gitea

Terraform Gitea Provider

[![Build Status](https://drone.uploadfilter24.eu/api/badges/lerentis/terraform-provider-gitea/status.svg)](https://drone.uploadfilter24.eu/lerentis/terraform-provider-gitea)

## History

This is a fork of https://gitea.com/gitea/terraform-provider-gitea. Many thanks for the foundation of this provider  

## Usage

This is not a 1.0 release, so usage is subject to change!

```terraform
terraform {
  required_providers {
    gitea = {
      source = "Lerentis/gitea"
      version = "0.13.0"
    }
  }
}

provider "gitea" {
  base_url = var.gitea_url # optionally use GITEA_BASE_URL env var
  token    = var.gitea_token # optionally use GITEA_TOKEN env var

  # Username/Password authentication is mutally exclusive with token authentication
  # username = var.username # optionally use GITEA_USERNAME env var
  # password = var.password # optionally use GITEA_PASSWORD env var

  # A file containing the ca certificate to use in case ssl certificate is not from a standard chain
  cacert_file = var.cacert_file 
  
  # If you are running a gitea instance with self signed TLS certificates
  # and you want to disable certificate validation you can deactivate it with this flag
  insecure = false 
}

resource "gitea_repository" "test" {
  username     = "lerentis"
  name         = "test"
  private      = true
  issue_labels = "Default"
  license      = "MIT"
  gitignores   = "Go"
}

resource "gitea_repository" "mirror" {
  username                     = "lerentis"
  name                         = "terraform-provider-gitea-mirror"
  description                  = "Mirror of Terraform Provider"
  mirror                       = true
  migration_clone_addresse     = "https://git.uploadfilter24.eu/lerentis/terraform-provider-gitea.git"
  migration_service            = "gitea"
  migration_service_auth_token = var.gitea_mirror_token
}

resource "gitea_org" "test_org" {
  name = "test-org"
}

resource "gitea_repository" "org_repo" {
  username = gitea_org.test_org.name
  name = "org-test-repo"
}

```

## Contributing

This repo is a mirror of [uploadfilter24.eu](https://git.uploadfilter24.eu/lerentis/terraform-provider-gitea), where i mostly develop. PRs will be manually merged on the gitea instance as keeping these two repositories in sync can be very error prune.
