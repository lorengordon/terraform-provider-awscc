// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

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
	registry.AddResourceTypeFactory("awscc_rds_global_cluster", globalClusterResourceType)
}

// globalClusterResourceType returns the Terraform awscc_rds_global_cluster resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RDS::GlobalCluster resource type.
func globalClusterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"deletion_protection": {
			// Property: DeletionProtection
			// CloudFormation resource type schema:
			// {
			//   "description": "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
			//   "type": "boolean"
			// }
			Description: "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"engine": {
			// Property: Engine
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			//   "enum": [
			//     "aurora",
			//     "aurora-mysql",
			//     "aurora-postgresql"
			//   ],
			//   "type": "string"
			// }
			Description: "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"aurora",
					"aurora-mysql",
					"aurora-postgresql",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"engine_version": {
			// Property: EngineVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			//   "type": "string"
			// }
			Description: "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"global_cluster_identifier": {
			// Property: GlobalClusterIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"source_db_cluster_identifier": {
			// Property: SourceDBClusterIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
			//   "oneOf": [
			//     {},
			//     {}
			//   ],
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"storage_encrypted": {
			// Property: StorageEncrypted
			// CloudFormation resource type schema:
			// {
			//   "description": " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			//   "type": "boolean"
			// }
			Description: " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
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
		Description: "Resource Type definition for AWS::RDS::GlobalCluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::GlobalCluster").WithTerraformTypeName("awscc_rds_global_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"deletion_protection":          "DeletionProtection",
		"engine":                       "Engine",
		"engine_version":               "EngineVersion",
		"global_cluster_identifier":    "GlobalClusterIdentifier",
		"source_db_cluster_identifier": "SourceDBClusterIdentifier",
		"storage_encrypted":            "StorageEncrypted",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.OneOfRequired(
		validate.Required(
			"source_db_cluster_identifier",
		),
		validate.Required(
			"engine",
		),
	),
	)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_rds_global_cluster", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
