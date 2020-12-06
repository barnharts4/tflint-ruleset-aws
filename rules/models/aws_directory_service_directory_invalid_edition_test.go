// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDirectoryServiceDirectoryInvalidEditionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_directory_service_directory" "foo" {
	edition = "Free"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDirectoryServiceDirectoryInvalidEditionRule(),
					Message: `"Free" is an invalid value as edition`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_directory_service_directory" "foo" {
	edition = "Enterprise"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsDirectoryServiceDirectoryInvalidEditionRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
