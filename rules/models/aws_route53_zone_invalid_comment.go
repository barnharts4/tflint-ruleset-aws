// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ZoneInvalidCommentRule checks the pattern is valid
type AwsRoute53ZoneInvalidCommentRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53ZoneInvalidCommentRule returns new rule with default attributes
func NewAwsRoute53ZoneInvalidCommentRule() *AwsRoute53ZoneInvalidCommentRule {
	return &AwsRoute53ZoneInvalidCommentRule{
		resourceType:  "aws_route53_zone",
		attributeName: "comment",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsRoute53ZoneInvalidCommentRule) Name() string {
	return "aws_route53_zone_invalid_comment"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ZoneInvalidCommentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ZoneInvalidCommentRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ZoneInvalidCommentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ZoneInvalidCommentRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"comment must be 256 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
