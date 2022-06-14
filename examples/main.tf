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
  name     = "org-test-repo"
}

data "gitea_user" "me" {
  username = "lerentis"
}

resource "gitea_user" "test" {
  username             = "test"
  login_name           = "test"
  password             = "Geheim1!"
  email                = "test@user.dev"
  must_change_password = false
  admin                = true
}
