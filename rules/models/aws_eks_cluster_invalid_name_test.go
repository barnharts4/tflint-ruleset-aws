// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEksClusterInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_eks_cluster" "foo" {
	name = "@example"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEksClusterInvalidNameRule(),
					Message: `"@example" does not match valid pattern ^[0-9A-Za-z][A-Za-z0-9\-_]*`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_eks_cluster" "foo" {
	name = "example"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEksClusterInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
