// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_transit_gateway_connect", transitGatewayConnectResourceType)
}

// transitGatewayConnectResourceType returns the Terraform awscc_ec2_transit_gateway_connect resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGatewayConnect resource type.
func transitGatewayConnectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The creation time.",
			//   "type": "string"
			// }
			Description: "The creation time.",
			Type:        types.StringType,
			Computed:    true,
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Protocol": {
			//       "description": "The tunnel protocol.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"protocol": {
						// Property: Protocol
						Description: "The tunnel protocol.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the attachment.",
			//   "type": "string"
			// }
			Description: "The state of the attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the attachment.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The tags for the attachment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Type:        types.StringType,
						Optional:    true,
					},
					"value": {
						// Property: Value
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"transit_gateway_attachment_id": {
			// Property: TransitGatewayAttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Connect attachment.",
			//   "type": "string"
			// }
			Description: "The ID of the Connect attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transport_transit_gateway_attachment_id": {
			// Property: TransportTransitGatewayAttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the attachment from which the Connect attachment was created.",
			//   "type": "string"
			// }
			Description: "The ID of the attachment from which the Connect attachment was created.",
			Type:        types.StringType,
			Required:    true,
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
		Description: "The AWS::EC2::TransitGatewayConnect type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayConnect").WithTerraformTypeName("awscc_ec2_transit_gateway_connect")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                 "CreationTime",
		"key":                           "Key",
		"options":                       "Options",
		"protocol":                      "Protocol",
		"state":                         "State",
		"tags":                          "Tags",
		"transit_gateway_attachment_id": "TransitGatewayAttachmentId",
		"transit_gateway_id":            "TransitGatewayId",
		"transport_transit_gateway_attachment_id": "TransportTransitGatewayAttachmentId",
		"value": "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_transit_gateway_connect", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
