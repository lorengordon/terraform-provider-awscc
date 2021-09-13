// Code generated by generators/resource/main.go; DO NOT EDIT.

package autoscaling

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
	registry.AddResourceTypeFactory("awscc_autoscaling_warm_pool", warmPoolResourceType)
}

// warmPoolResourceType returns the Terraform awscc_autoscaling_warm_pool resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AutoScaling::WarmPool resource type.
func warmPoolResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_group_name": {
			// Property: AutoScalingGroupName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"max_group_prepared_capacity": {
			// Property: MaxGroupPreparedCapacity
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
		},
		"min_size": {
			// Property: MinSize
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
		},
		"pool_state": {
			// Property: PoolState
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::AutoScaling::WarmPool.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AutoScaling::WarmPool").WithTerraformTypeName("awscc_autoscaling_warm_pool")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_name":     "AutoScalingGroupName",
		"max_group_prepared_capacity": "MaxGroupPreparedCapacity",
		"min_size":                    "MinSize",
		"pool_state":                  "PoolState",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_autoscaling_warm_pool", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
