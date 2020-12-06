// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCognitoUserPoolInvalidMfaConfigurationRule checks the pattern is valid
type AwsCognitoUserPoolInvalidMfaConfigurationRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCognitoUserPoolInvalidMfaConfigurationRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidMfaConfigurationRule() *AwsCognitoUserPoolInvalidMfaConfigurationRule {
	return &AwsCognitoUserPoolInvalidMfaConfigurationRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "mfa_configuration",
		enum: []string{
			"OFF",
			"ON",
			"OPTIONAL",
		},
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidMfaConfigurationRule) Name() string {
	return "aws_cognito_user_pool_invalid_mfa_configuration"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidMfaConfigurationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolInvalidMfaConfigurationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidMfaConfigurationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidMfaConfigurationRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as mfa_configuration`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
