// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAMIInvalidArchitectureRule checks the pattern is valid
type AwsAMIInvalidArchitectureRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAMIInvalidArchitectureRule returns new rule with default attributes
func NewAwsAMIInvalidArchitectureRule() *AwsAMIInvalidArchitectureRule {
	return &AwsAMIInvalidArchitectureRule{
		resourceType:  "aws_ami",
		attributeName: "architecture",
		enum: []string{
			"i386",
			"x86_64",
			"arm64",
		},
	}
}

// Name returns the rule name
func (r *AwsAMIInvalidArchitectureRule) Name() string {
	return "aws_ami_invalid_architecture"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAMIInvalidArchitectureRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAMIInvalidArchitectureRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAMIInvalidArchitectureRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAMIInvalidArchitectureRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as architecture`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
