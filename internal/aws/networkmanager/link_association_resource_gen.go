// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

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
	registry.AddResourceTypeFactory("awscc_networkmanager_link_association", linkAssociationResourceType)
}

// linkAssociationResourceType returns the Terraform awscc_networkmanager_link_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::LinkAssociation resource type.
func linkAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"device_id": {
			// Property: DeviceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the device",
			//   "type": "string"
			// }
			Description: "The ID of the device",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the global network.",
			//   "type": "string"
			// }
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"link_id": {
			// Property: LinkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the link",
			//   "type": "string"
			// }
			Description: "The ID of the link",
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
		Description: "The AWS::NetworkManager::LinkAssociation type associates a link to a device. The device and link must be in the same global network and the same site.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::LinkAssociation").WithTerraformTypeName("awscc_networkmanager_link_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"device_id":         "DeviceId",
		"global_network_id": "GlobalNetworkId",
		"link_id":           "LinkId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkmanager_link_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
