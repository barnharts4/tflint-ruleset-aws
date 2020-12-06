// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule checks the pattern is valid
type AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsGlobalacceleratorEndpointGroupInvalidListenerArnRule returns new rule with default attributes
func NewAwsGlobalacceleratorEndpointGroupInvalidListenerArnRule() *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule {
	return &AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule{
		resourceType:  "aws_globalaccelerator_endpoint_group",
		attributeName: "listener_arn",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Name() string {
	return "aws_globalaccelerator_endpoint_group_invalid_listener_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"listener_arn must be 255 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
