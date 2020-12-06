// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOrganizationsPolicyInvalidContentRule checks the pattern is valid
type AwsOrganizationsPolicyInvalidContentRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsPolicyInvalidContentRule returns new rule with default attributes
func NewAwsOrganizationsPolicyInvalidContentRule() *AwsOrganizationsPolicyInvalidContentRule {
	return &AwsOrganizationsPolicyInvalidContentRule{
		resourceType:  "aws_organizations_policy",
		attributeName: "content",
		max:           1000000,
		min:           1,
		pattern:       regexp.MustCompile(`^[\s\S]*$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsPolicyInvalidContentRule) Name() string {
	return "aws_organizations_policy_invalid_content"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsPolicyInvalidContentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsPolicyInvalidContentRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsPolicyInvalidContentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsPolicyInvalidContentRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"content must be 1000000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"content must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\s\S]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
