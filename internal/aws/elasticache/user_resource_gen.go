// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache

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
	registry.AddResourceTypeFactory("awscc_elasticache_user", userResourceType)
}

// userResourceType returns the Terraform awscc_elasticache_user resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ElastiCache::User resource type.
func userResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_string": {
			// Property: AccessString
			// CloudFormation resource type schema:
			// {
			//   "description": "Access permissions string used for this user account.",
			//   "type": "string"
			// }
			Description: "Access permissions string used for this user account.",
			Type:        types.StringType,
			Optional:    true,
			// AccessString is a write-only property.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the user account.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the user account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"engine": {
			// Property: Engine
			// CloudFormation resource type schema:
			// {
			//   "description": "Must be redis.",
			//   "enum": [
			//     "redis"
			//   ],
			//   "type": "string"
			// }
			Description: "Must be redis.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"redis",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"no_password_required": {
			// Property: NoPasswordRequired
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates a password is not required for this user account.",
			//   "type": "boolean"
			// }
			Description: "Indicates a password is not required for this user account.",
			Type:        types.BoolType,
			Optional:    true,
			// NoPasswordRequired is a write-only property.
		},
		"passwords": {
			// Property: Passwords
			// CloudFormation resource type schema:
			// {
			//   "$comment": "List of passwords.",
			//   "description": "Passwords used for this user account. You can create up to two passwords for each user.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Passwords used for this user account. You can create up to two passwords for each user.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			// Passwords is a write-only property.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			//   "type": "string"
			// }
			Description: "Indicates the user status. Can be \"active\", \"modifying\" or \"deleting\".",
			Type:        types.StringType,
			Computed:    true,
		},
		"user_id": {
			// Property: UserId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the user.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the user.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"user_name": {
			// Property: UserName
			// CloudFormation resource type schema:
			// {
			//   "description": "The username of the user.",
			//   "type": "string"
			// }
			Description: "The username of the user.",
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
		Description: "Resource Type definition for AWS::ElastiCache::User",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::User").WithTerraformTypeName("awscc_elasticache_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_string":        "AccessString",
		"arn":                  "Arn",
		"engine":               "Engine",
		"no_password_required": "NoPasswordRequired",
		"passwords":            "Passwords",
		"status":               "Status",
		"user_id":              "UserId",
		"user_name":            "UserName",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Passwords",
		"/properties/NoPasswordRequired",
		"/properties/AccessString",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_elasticache_user", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
