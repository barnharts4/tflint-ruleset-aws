// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDirectoryServiceDirectoryInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_directory_service_directory" "foo" {
	name = "@example.com"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDirectoryServiceDirectoryInvalidNameRule(),
					Message: `"@example.com" does not match valid pattern ^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_directory_service_directory" "foo" {
	name = "corp.notexample.com"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsDirectoryServiceDirectoryInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
