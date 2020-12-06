// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayNfsFileShareInvalidGatewayArnRule checks the pattern is valid
type AwsStoragegatewayNfsFileShareInvalidGatewayArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayNfsFileShareInvalidGatewayArnRule returns new rule with default attributes
func NewAwsStoragegatewayNfsFileShareInvalidGatewayArnRule() *AwsStoragegatewayNfsFileShareInvalidGatewayArnRule {
	return &AwsStoragegatewayNfsFileShareInvalidGatewayArnRule{
		resourceType:  "aws_storagegateway_nfs_file_share",
		attributeName: "gateway_arn",
		max:           500,
		min:           50,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayNfsFileShareInvalidGatewayArnRule) Name() string {
	return "aws_storagegateway_nfs_file_share_invalid_gateway_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayNfsFileShareInvalidGatewayArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayNfsFileShareInvalidGatewayArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayNfsFileShareInvalidGatewayArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayNfsFileShareInvalidGatewayArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"gateway_arn must be 500 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"gateway_arn must be 50 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
