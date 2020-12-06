// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoUserPoolInvalidEmailVerificationMessageRule checks the pattern is valid
type AwsCognitoUserPoolInvalidEmailVerificationMessageRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule() *AwsCognitoUserPoolInvalidEmailVerificationMessageRule {
	return &AwsCognitoUserPoolInvalidEmailVerificationMessageRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "email_verification_message",
		max:           20000,
		min:           6,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Name() string {
	return "aws_cognito_user_pool_invalid_email_verification_message"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"email_verification_message must be 20000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"email_verification_message must be 6 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
