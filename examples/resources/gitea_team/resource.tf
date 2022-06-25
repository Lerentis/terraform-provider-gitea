resource "gitea_org" "test_org" {
  name = "test-org"
}

resource "gitea_team" "test_team" {
  name = "Devs"
  organisation = gitea_org.test_org.name
  description = "Devs of Test Org"
  permission = "write"
}