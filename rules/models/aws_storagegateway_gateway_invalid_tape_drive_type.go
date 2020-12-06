// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayGatewayInvalidTapeDriveTypeRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidTapeDriveTypeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidTapeDriveTypeRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidTapeDriveTypeRule() *AwsStoragegatewayGatewayInvalidTapeDriveTypeRule {
	return &AwsStoragegatewayGatewayInvalidTapeDriveTypeRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "tape_drive_type",
		max:           50,
		min:           2,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidTapeDriveTypeRule) Name() string {
	return "aws_storagegateway_gateway_invalid_tape_drive_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidTapeDriveTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidTapeDriveTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidTapeDriveTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidTapeDriveTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"tape_drive_type must be 50 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"tape_drive_type must be 2 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
