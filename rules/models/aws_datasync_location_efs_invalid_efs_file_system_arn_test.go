// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDatasyncLocationEfsInvalidEfsFileSystemArnRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_datasync_location_efs" "foo" {
	efs_file_system_arn = "arn:aws:eks:us-east-1:123456789012:cluster/my-cluster"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDatasyncLocationEfsInvalidEfsFileSystemArnRule(),
					Message: `"arn:aws:eks:us-east-1:123456789012:cluster/my-cluster" does not match valid pattern ^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_datasync_location_efs" "foo" {
	efs_file_system_arn = "arn:aws:elasticfilesystem:us-east-1:123456789012:file-system/fs-12345678"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsDatasyncLocationEfsInvalidEfsFileSystemArnRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
