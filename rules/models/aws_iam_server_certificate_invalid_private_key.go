// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMServerCertificateInvalidPrivateKeyRule checks the pattern is valid
type AwsIAMServerCertificateInvalidPrivateKeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMServerCertificateInvalidPrivateKeyRule returns new rule with default attributes
func NewAwsIAMServerCertificateInvalidPrivateKeyRule() *AwsIAMServerCertificateInvalidPrivateKeyRule {
	return &AwsIAMServerCertificateInvalidPrivateKeyRule{
		resourceType:  "aws_iam_server_certificate",
		attributeName: "private_key",
		max:           16384,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMServerCertificateInvalidPrivateKeyRule) Name() string {
	return "aws_iam_server_certificate_invalid_private_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMServerCertificateInvalidPrivateKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMServerCertificateInvalidPrivateKeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMServerCertificateInvalidPrivateKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMServerCertificateInvalidPrivateKeyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"private_key must be 16384 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"private_key must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					`private_key does not match valid pattern ^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`,
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
