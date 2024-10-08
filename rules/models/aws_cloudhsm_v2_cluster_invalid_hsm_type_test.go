// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudhsmV2ClusterInvalidHsmTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	hsm_type = "t2.medium"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudhsmV2ClusterInvalidHsmTypeRule(),
					Message: fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage("t2.medium"), `^((p|)hsm[0-9][a-z.]*\.[a-zA-Z]+)$`),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	hsm_type = "hsm1.medium"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCloudhsmV2ClusterInvalidHsmTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
