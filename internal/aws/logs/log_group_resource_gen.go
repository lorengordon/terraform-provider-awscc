// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_logs_log_group", logGroupResourceType)
}

// logGroupResourceType returns the Terraform awscc_logs_log_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Logs::LogGroup resource type.
func logGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The CloudWatch log group ARN.",
			//   "type": "string"
			// }
			Description: "The CloudWatch log group ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
			},
		},
		"log_group_name": {
			// Property: LogGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 512),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"retention_in_days": {
			// Property: RetentionInDays
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			//   "enum": [
			//     1,
			//     3,
			//     5,
			//     7,
			//     14,
			//     30,
			//     60,
			//     90,
			//     120,
			//     150,
			//     180,
			//     365,
			//     400,
			//     545,
			//     731,
			//     1827,
			//     3653
			//   ],
			//   "type": "integer"
			// }
			Description: "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			Type:        types.NumberType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntInSlice([]int{
					1,
					3,
					5,
					7,
					14,
					30,
					60,
					90,
					120,
					150,
					180,
					365,
					400,
					545,
					731,
					1827,
					3653,
				}),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Logs::LogGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogGroup").WithTerraformTypeName("awscc_logs_log_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"kms_key_id":        "KmsKeyId",
		"log_group_name":    "LogGroupName",
		"retention_in_days": "RetentionInDays",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
