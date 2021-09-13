// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

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
	registry.AddResourceTypeFactory("awscc_iot_topic_rule_destination", topicRuleDestinationResourceType)
}

// topicRuleDestinationResourceType returns the Terraform awscc_iot_topic_rule_destination resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::TopicRuleDestination resource type.
func topicRuleDestinationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN).",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN).",
			Type:        types.StringType,
			Computed:    true,
		},
		"http_url_properties": {
			// Property: HttpUrlProperties
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ConfirmationUrl": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"confirmation_url": {
						// Property: ConfirmationUrl
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ENABLED",
			//     "IN_PROGRESS",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ENABLED",
					"IN_PROGRESS",
					"DISABLED",
				}),
			},
		},
		"status_reason": {
			// Property: StatusReason
			// CloudFormation resource type schema:
			// {
			//   "description": "The reasoning for the current status of the TopicRuleDestination.",
			//   "type": "string"
			// }
			Description: "The reasoning for the current status of the TopicRuleDestination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_properties": {
			// Property: VpcProperties
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "RoleArn": {
			//       "type": "string"
			//     },
			//     "SecurityGroups": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "SubnetIds": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "VpcId": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"role_arn": {
						// Property: RoleArn
						Type:     types.StringType,
						Optional: true,
					},
					"security_groups": {
						// Property: SecurityGroups
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
					},
					"subnet_ids": {
						// Property: SubnetIds
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
					},
					"vpc_id": {
						// Property: VpcId
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::IoT::TopicRuleDestination",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::TopicRuleDestination").WithTerraformTypeName("awscc_iot_topic_rule_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"confirmation_url":    "ConfirmationUrl",
		"http_url_properties": "HttpUrlProperties",
		"role_arn":            "RoleArn",
		"security_groups":     "SecurityGroups",
		"status":              "Status",
		"status_reason":       "StatusReason",
		"subnet_ids":          "SubnetIds",
		"vpc_id":              "VpcId",
		"vpc_properties":      "VpcProperties",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iot_topic_rule_destination", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
