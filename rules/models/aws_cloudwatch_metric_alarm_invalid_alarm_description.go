// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule() *AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule {
	return &AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "alarm_description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_alarm_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"alarm_description must be 1024 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
