// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCurReportDefinitionInvalidReportNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cur_report_definition" "foo" {
	report_name = "example/cur-report-definition"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCurReportDefinitionInvalidReportNameRule(),
					Message: `"example/cur-report-definition" does not match valid pattern ^[0-9A-Za-z!\-_.*\'()]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cur_report_definition" "foo" {
	report_name = "example-cur-report-definition"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCurReportDefinitionInvalidReportNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
