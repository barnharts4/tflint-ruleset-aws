// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppsyncFunctionInvalidResponseMappingTemplateRule checks the pattern is valid
type AwsAppsyncFunctionInvalidResponseMappingTemplateRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppsyncFunctionInvalidResponseMappingTemplateRule returns new rule with default attributes
func NewAwsAppsyncFunctionInvalidResponseMappingTemplateRule() *AwsAppsyncFunctionInvalidResponseMappingTemplateRule {
	return &AwsAppsyncFunctionInvalidResponseMappingTemplateRule{
		resourceType:  "aws_appsync_function",
		attributeName: "response_mapping_template",
		max:           65536,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppsyncFunctionInvalidResponseMappingTemplateRule) Name() string {
	return "aws_appsync_function_invalid_response_mapping_template"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncFunctionInvalidResponseMappingTemplateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncFunctionInvalidResponseMappingTemplateRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncFunctionInvalidResponseMappingTemplateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncFunctionInvalidResponseMappingTemplateRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"response_mapping_template must be 65536 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"response_mapping_template must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
