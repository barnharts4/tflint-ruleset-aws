// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogPortfolioInvalidDescriptionRule checks the pattern is valid
type AwsServicecatalogPortfolioInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsServicecatalogPortfolioInvalidDescriptionRule returns new rule with default attributes
func NewAwsServicecatalogPortfolioInvalidDescriptionRule() *AwsServicecatalogPortfolioInvalidDescriptionRule {
	return &AwsServicecatalogPortfolioInvalidDescriptionRule{
		resourceType:  "aws_servicecatalog_portfolio",
		attributeName: "description",
		max:           2000,
	}
}

// Name returns the rule name
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Name() string {
	return "aws_servicecatalog_portfolio_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogPortfolioInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 2000 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
