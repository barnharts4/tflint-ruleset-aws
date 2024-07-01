// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAcmpcaCertificateInvalidSigningAlgorithmRule checks the pattern is valid
type AwsAcmpcaCertificateInvalidSigningAlgorithmRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAcmpcaCertificateInvalidSigningAlgorithmRule returns new rule with default attributes
func NewAwsAcmpcaCertificateInvalidSigningAlgorithmRule() *AwsAcmpcaCertificateInvalidSigningAlgorithmRule {
	return &AwsAcmpcaCertificateInvalidSigningAlgorithmRule{
		resourceType:  "aws_acmpca_certificate",
		attributeName: "signing_algorithm",
		enum: []string{
			"SHA256WITHECDSA",
			"SHA384WITHECDSA",
			"SHA512WITHECDSA",
			"SHA256WITHRSA",
			"SHA384WITHRSA",
			"SHA512WITHRSA",
			"SM3WITHSM2",
		},
	}
}

// Name returns the rule name
func (r *AwsAcmpcaCertificateInvalidSigningAlgorithmRule) Name() string {
	return "aws_acmpca_certificate_invalid_signing_algorithm"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAcmpcaCertificateInvalidSigningAlgorithmRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAcmpcaCertificateInvalidSigningAlgorithmRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAcmpcaCertificateInvalidSigningAlgorithmRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAcmpcaCertificateInvalidSigningAlgorithmRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as signing_algorithm`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
