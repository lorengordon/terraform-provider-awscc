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
	registry.AddResourceTypeFactory("awscc_ec2_transit_gateway", transitGatewayResourceType)
}

// transitGatewayResourceType returns the Terraform awscc_ec2_transit_gateway resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGateway resource type.
func transitGatewayResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"amazon_side_asn": {
			// Property: AmazonSideAsn
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"association_default_route_table_id": {
			// Property: AssociationDefaultRouteTableId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"auto_accept_shared_attachments": {
			// Property: AutoAcceptSharedAttachments
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"default_route_table_association": {
			// Property: DefaultRouteTableAssociation
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"default_route_table_propagation": {
			// Property: DefaultRouteTablePropagation
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"dns_support": {
			// Property: DnsSupport
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"multicast_support": {
			// Property: MulticastSupport
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"propagation_default_route_table_id": {
			// Property: PropagationDefaultRouteTableId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"transit_gateway_cidr_blocks": {
			// Property: TransitGatewayCidrBlocks
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"vpn_ecmp_support": {
			// Property: VpnEcmpSupport
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::TransitGateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGateway").WithTerraformTypeName("awscc_ec2_transit_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amazon_side_asn":                    "AmazonSideAsn",
		"association_default_route_table_id": "AssociationDefaultRouteTableId",
		"auto_accept_shared_attachments":     "AutoAcceptSharedAttachments",
		"default_route_table_association":    "DefaultRouteTableAssociation",
		"default_route_table_propagation":    "DefaultRouteTablePropagation",
		"description":                        "Description",
		"dns_support":                        "DnsSupport",
		"id":                                 "Id",
		"key":                                "Key",
		"multicast_support":                  "MulticastSupport",
		"propagation_default_route_table_id": "PropagationDefaultRouteTableId",
		"tags":                               "Tags",
		"transit_gateway_cidr_blocks":        "TransitGatewayCidrBlocks",
		"value":                              "Value",
		"vpn_ecmp_support":                   "VpnEcmpSupport",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_transit_gateway", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
