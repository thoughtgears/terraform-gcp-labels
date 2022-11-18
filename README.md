# labels module

![Terraform Version](https://img.shields.io/badge/tf-%3E%3D1.0.x-blue.svg)

## Usage

```hcl
module "my" {
  source  = "git::git@github.com:thoughtgears/gcp-core-module.git//modules/labels?ref=latest"

  custom = {
    label1 = "example"
  }
}
```

## What's a Module?

A Module is a canonical, reusable, best-practices definition for how to run a single piece of infrastructure, such as a database or server cluster.
Each Module is written using a combination of [Terraform](https://www.terraform.io/) and scripts (mostly bash) and include automated tests, documentation, and examples.
It is maintained both by the open source community and companies that provide commercial support.

Instead of figuring out the details of how to run a piece of infrastructure from scratch, you can reuse  existing code that has been proven in production.
And instead of maintaining all that infrastructure code yourself, you can leverage the work of the Module community to pick up infrastructure improvements through a version number bump.

## How is this Module versioned?

This Module follows the principles of [Semantic Versioning](http://semver.org/). You can find each new release, along with the changelog, in the [Releases Page](https://github.com/thoughtgears/gcp-xxx-modules/releases).

During initial development, the major version will be 0 (e.g., `0.x.y`), which indicates the code does not yet have a stable API.
Once we hit `1.0.0`, we will make every effort to maintain a backwards compatible API and use the MAJOR, MINOR, and PATCH versions on each release to indicate any incompatibilities.

## Module resource documentation

<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_external"></a> [external](#provider\_external) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [external_external.get_git_info](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_contact"></a> [contact](#input\_contact) | Contact email for main contact of project | `string` | `null` | no |
| <a name="input_custom"></a> [custom](#input\_custom) | An object of custom labels to merge with the default ones.<br>Define them by passing in an object:<br>custom = {<br>    label1 = "example1"<br>    label2 = "example2"<br>} | `any` | `null` | no |
| <a name="input_git_info"></a> [git\_info](#input\_git\_info) | If git information should be included, if true it will look up information about your repository to ensure<br>we can get repo name, organization and commit sha to use for default labels. | `bool` | `true` | no |
| <a name="input_project_name"></a> [project\_name](#input\_project\_name) | Name of the project for the labels | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_all"></a> [all](#output\_all) | n/a |
| <a name="output_custom"></a> [custom](#output\_custom) | n/a |
| <a name="output_default"></a> [default](#output\_default) | n/a |
<!-- END_TF_DOCS -->
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_external"></a> [external](#provider\_external) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [external_external.get_git_info](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_contact"></a> [contact](#input\_contact) | Contact email for main contact of project | `string` | `null` | no |
| <a name="input_custom"></a> [custom](#input\_custom) | An object of custom labels to merge with the default ones.<br>Define them by passing in an object:<br>custom = {<br>    label1 = "example1"<br>    label2 = "example2"<br>} | `any` | `null` | no |
| <a name="input_git_info"></a> [git\_info](#input\_git\_info) | If git information should be included, if true it will look up information about your repository to ensure<br>we can get repo name, organization and commit sha to use for default labels. | `bool` | `true` | no |
| <a name="input_project_name"></a> [project\_name](#input\_project\_name) | Name of the project for the labels | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_all"></a> [all](#output\_all) | n/a |
| <a name="output_custom"></a> [custom](#output\_custom) | n/a |
| <a name="output_default"></a> [default](#output\_default) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
