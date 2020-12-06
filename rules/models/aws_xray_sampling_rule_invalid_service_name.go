// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsXraySamplingRuleInvalidServiceNameRule checks the pattern is valid
type AwsXraySamplingRuleInvalidServiceNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsXraySamplingRuleInvalidServiceNameRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidServiceNameRule() *AwsXraySamplingRuleInvalidServiceNameRule {
	return &AwsXraySamplingRuleInvalidServiceNameRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "service_name",
		max:           64,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidServiceNameRule) Name() string {
	return "aws_xray_sampling_rule_invalid_service_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidServiceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsXraySamplingRuleInvalidServiceNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidServiceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidServiceNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"service_name must be 64 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
