// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

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
	registry.AddResourceTypeFactory("awscc_amplify_branch", branchResourceType)
}

// branchResourceType returns the Terraform awscc_amplify_branch resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Amplify::Branch resource type.
func branchResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 20),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"basic_auth_config": {
			// Property: BasicAuthConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EnableBasicAuth": {
			//       "type": "boolean"
			//     },
			//     "Password": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Username": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Username",
			//     "Password"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enable_basic_auth": {
						// Property: EnableBasicAuth
						Type:     types.BoolType,
						Optional: true,
					},
					"password": {
						// Property: Password
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"username": {
						// Property: Username
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			// BasicAuthConfig is a write-only property.
		},
		"branch_name": {
			// Property: BranchName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"build_spec": {
			// Property: BuildSpec
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 25000,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 25000),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1000),
			},
		},
		"enable_auto_build": {
			// Property: EnableAutoBuild
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"enable_performance_mode": {
			// Property: EnablePerformanceMode
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"enable_pull_request_preview": {
			// Property: EnablePullRequestPreview
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"environment_variables": {
			// Property: EnvironmentVariables
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 255,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 5500,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 255),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 5500),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"pull_request_environment_name": {
			// Property: PullRequestEnvironmentName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 20),
			},
		},
		"stage": {
			// Property: Stage
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "EXPERIMENTAL",
			//     "BETA",
			//     "PULL_REQUEST",
			//     "PRODUCTION",
			//     "DEVELOPMENT"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"EXPERIMENTAL",
					"BETA",
					"PULL_REQUEST",
					"PRODUCTION",
					"DEVELOPMENT",
				}),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "insertionOrder": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
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
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::Amplify::Branch resource creates a new branch within an app.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Branch").WithTerraformTypeName("awscc_amplify_branch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                        "AppId",
		"arn":                           "Arn",
		"basic_auth_config":             "BasicAuthConfig",
		"branch_name":                   "BranchName",
		"build_spec":                    "BuildSpec",
		"description":                   "Description",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"key":                           "Key",
		"name":                          "Name",
		"password":                      "Password",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"stage":                         "Stage",
		"tags":                          "Tags",
		"username":                      "Username",
		"value":                         "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BasicAuthConfig",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_amplify_branch", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
