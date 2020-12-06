// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoResourceServerInvalidNameRule checks the pattern is valid
type AwsCognitoResourceServerInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoResourceServerInvalidNameRule returns new rule with default attributes
func NewAwsCognitoResourceServerInvalidNameRule() *AwsCognitoResourceServerInvalidNameRule {
	return &AwsCognitoResourceServerInvalidNameRule{
		resourceType:  "aws_cognito_resource_server",
		attributeName: "name",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w\s+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoResourceServerInvalidNameRule) Name() string {
	return "aws_cognito_resource_server_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoResourceServerInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoResourceServerInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoResourceServerInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoResourceServerInvalidNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w\s+=,.@-]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
