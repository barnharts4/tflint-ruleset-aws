// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverRuleInvalidRuleTypeRule checks the pattern is valid
type AwsRoute53ResolverRuleInvalidRuleTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53ResolverRuleInvalidRuleTypeRule returns new rule with default attributes
func NewAwsRoute53ResolverRuleInvalidRuleTypeRule() *AwsRoute53ResolverRuleInvalidRuleTypeRule {
	return &AwsRoute53ResolverRuleInvalidRuleTypeRule{
		resourceType:  "aws_route53_resolver_rule",
		attributeName: "rule_type",
		enum: []string{
			"FORWARD",
			"SYSTEM",
			"RECURSIVE",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverRuleInvalidRuleTypeRule) Name() string {
	return "aws_route53_resolver_rule_invalid_rule_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverRuleInvalidRuleTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverRuleInvalidRuleTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverRuleInvalidRuleTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverRuleInvalidRuleTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as rule_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
