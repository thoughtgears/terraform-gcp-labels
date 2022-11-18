package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestLabelsModule(t *testing.T) {
	t.Parallel()

	examplesDir := "../example"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: examplesDir,
	})
	defer terraform.Destroy(t, terraformOptions)

	t.Logf("Running in %s", examplesDir)
	terraform.InitAndApply(t, terraformOptions)

	// Do assertions below
	labels := terraform.Output(t, terraformOptions, "labels")
	assert.Equal(t, `map[project_name:test-project terraform:true]`, labels)
}
