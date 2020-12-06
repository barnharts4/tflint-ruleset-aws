// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMGroupMembershipInvalidGroupRule checks the pattern is valid
type AwsIAMGroupMembershipInvalidGroupRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMGroupMembershipInvalidGroupRule returns new rule with default attributes
func NewAwsIAMGroupMembershipInvalidGroupRule() *AwsIAMGroupMembershipInvalidGroupRule {
	return &AwsIAMGroupMembershipInvalidGroupRule{
		resourceType:  "aws_iam_group_membership",
		attributeName: "group",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMGroupMembershipInvalidGroupRule) Name() string {
	return "aws_iam_group_membership_invalid_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMGroupMembershipInvalidGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMGroupMembershipInvalidGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMGroupMembershipInvalidGroupRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMGroupMembershipInvalidGroupRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"group must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"group must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.@-]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
