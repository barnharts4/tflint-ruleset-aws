// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigConfigurationAggregatorInvalidNameRule checks the pattern is valid
type AwsConfigConfigurationAggregatorInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsConfigConfigurationAggregatorInvalidNameRule returns new rule with default attributes
func NewAwsConfigConfigurationAggregatorInvalidNameRule() *AwsConfigConfigurationAggregatorInvalidNameRule {
	return &AwsConfigConfigurationAggregatorInvalidNameRule{
		resourceType:  "aws_config_configuration_aggregator",
		attributeName: "name",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w\-]+$`),
	}
}

// Name returns the rule name
func (r *AwsConfigConfigurationAggregatorInvalidNameRule) Name() string {
	return "aws_config_configuration_aggregator_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigConfigurationAggregatorInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigConfigurationAggregatorInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigConfigurationAggregatorInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigConfigurationAggregatorInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w\-]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
