// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCognitoIdentityPoolRolesAttachmentInvalidIdentityPoolIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_identity_pool_roles_attachment" "foo" {
	identity_pool_id = "0123456789"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCognitoIdentityPoolRolesAttachmentInvalidIdentityPoolIDRule(),
					Message: `"0123456789" does not match valid pattern ^[\w-]+:[0-9a-f-]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_identity_pool_roles_attachment" "foo" {
	identity_pool_id = "us-east-1:0123456789"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCognitoIdentityPoolRolesAttachmentInvalidIdentityPoolIDRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
