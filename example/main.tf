module "labels" {
  source       = "../"
  project_name = "test-project"
  git_info     = false
}

output "labels" {
  value = module.labels.all
}
