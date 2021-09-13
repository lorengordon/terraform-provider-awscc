// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_apigateway_usage_plan_key", usagePlanKeyResourceType)
}

// usagePlanKeyResourceType returns the Terraform awscc_apigateway_usage_plan_key resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ApiGateway::UsagePlanKey resource type.
func usagePlanKeyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.",
			//   "type": "string"
			// }
			Description: "An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.",
			Type:        types.StringType,
			Computed:    true,
		},
		"key_id": {
			// Property: KeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the usage plan key.",
			//   "type": "string"
			// }
			Description: "The ID of the usage plan key.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"key_type": {
			// Property: KeyType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of usage plan key. Currently, the only valid key type is API_KEY.",
			//   "enum": [
			//     "API_KEY"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of usage plan key. Currently, the only valid key type is API_KEY.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"API_KEY",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"usage_plan_id": {
			// Property: UsagePlanId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the usage plan.",
			//   "type": "string"
			// }
			Description: "The ID of the usage plan.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ApiGateway::UsagePlanKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::UsagePlanKey").WithTerraformTypeName("awscc_apigateway_usage_plan_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":            "Id",
		"key_id":        "KeyId",
		"key_type":      "KeyType",
		"usage_plan_id": "UsagePlanId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_apigateway_usage_plan_key", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
