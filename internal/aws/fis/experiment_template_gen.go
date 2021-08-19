// Code generated by generators/resource/main.go; DO NOT EDIT.

package fis

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_fis_experiment_template", experimentTemplateResourceType)
}

// experimentTemplateResourceType returns the Terraform awscc_fis_experiment_template resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::FIS::ExperimentTemplate resource type.
func experimentTemplateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"actions": {
			// Property: Actions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The actions for the experiment.",
			//   "patternProperties": {
			//     "": {
			//       "additionalProperties": false,
			//       "description": "Specifies an action for the experiment template.",
			//       "properties": {
			//         "ActionId": {
			//           "description": "The ID of the action.",
			//           "maxLength": 64,
			//           "type": "string"
			//         },
			//         "Description": {
			//           "description": "A description for the action.",
			//           "maxLength": 512,
			//           "type": "string"
			//         },
			//         "Parameters": {
			//           "additionalProperties": false,
			//           "description": "The parameters for the action, if applicable.",
			//           "patternProperties": {
			//             "": {
			//               "maxLength": 1024,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "StartAfter": {
			//           "description": "The names of the actions that must be completed before the current action starts.",
			//           "items": {
			//             "maxLength": 64,
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "Targets": {
			//           "additionalProperties": false,
			//           "description": "One or more targets for the action.",
			//           "patternProperties": {
			//             "": {
			//               "maxLength": 64,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "ActionId"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The actions for the experiment.",
			// Pattern: ""
			Attributes: schema.MapNestedAttributes(
				map[string]schema.Attribute{
					"action_id": {
						// Property: ActionId
						Description: "The ID of the action.",
						Type:        types.StringType,
						Optional:    true,
					},
					"description": {
						// Property: Description
						Description: "A description for the action.",
						Type:        types.StringType,
						Optional:    true,
					},
					"parameters": {
						// Property: Parameters
						Description: "The parameters for the action, if applicable.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"start_after": {
						// Property: StartAfter
						Description: "The names of the actions that must be completed before the current action starts.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
					"targets": {
						// Property: Targets
						Description: "One or more targets for the action.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
				},
				schema.MapNestedAttributesOptions{},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the experiment template.",
			//   "maxLength": 512,
			//   "type": "string"
			// }
			Description: "A description for the experiment template.",
			Type:        types.StringType,
			Required:    true,
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
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
			//   "maxLength": 1224,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
			Type:        types.StringType,
			Required:    true,
		},
		"stop_conditions": {
			// Property: StopConditions
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more stop conditions.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Source": {
			//         "maxLength": 64,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 2048,
			//         "minLength": 20,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Source"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "One or more stop conditions.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"source": {
						// Property: Source
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Required: true,
			// Tags is a force-new attribute.
		},
		"targets": {
			// Property: Targets
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The targets for the experiment.",
			//   "patternProperties": {
			//     "": {
			//       "additionalProperties": false,
			//       "description": "Specifies a target for an experiment.",
			//       "properties": {
			//         "Filters": {
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Describes a filter used for the target resource input in an experiment template.",
			//             "properties": {
			//               "Path": {
			//                 "description": "The attribute path for the filter.",
			//                 "maxLength": 256,
			//                 "type": "string"
			//               },
			//               "Values": {
			//                 "description": "The attribute values for the filter.",
			//                 "items": {
			//                   "maxLength": 128,
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "required": [
			//               "Path",
			//               "Values"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array"
			//         },
			//         "ResourceArns": {
			//           "description": "The Amazon Resource Names (ARNs) of the target resources.",
			//           "items": {
			//             "maxLength": 2048,
			//             "minLength": 20,
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "ResourceTags": {
			//           "additionalProperties": false,
			//           "patternProperties": {
			//             "": {
			//               "maxLength": 256,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "ResourceType": {
			//           "description": "The AWS resource type. The resource type must be supported for the specified action.",
			//           "maxLength": 64,
			//           "type": "string"
			//         },
			//         "SelectionMode": {
			//           "description": "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
			//           "maxLength": 64,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "ResourceType",
			//         "SelectionMode"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The targets for the experiment.",
			// Pattern: ""
			Attributes: schema.MapNestedAttributes(
				map[string]schema.Attribute{
					"filters": {
						// Property: Filters
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"path": {
									// Property: Path
									Description: "The attribute path for the filter.",
									Type:        types.StringType,
									Required:    true,
								},
								"values": {
									// Property: Values
									Description: "The attribute values for the filter.",
									Type:        types.ListType{ElemType: types.StringType},
									Required:    true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"resource_arns": {
						// Property: ResourceArns
						Description: "The Amazon Resource Names (ARNs) of the target resources.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
					},
					"resource_tags": {
						// Property: ResourceTags
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"resource_type": {
						// Property: ResourceType
						Description: "The AWS resource type. The resource type must be supported for the specified action.",
						Type:        types.StringType,
						Optional:    true,
					},
					"selection_mode": {
						// Property: SelectionMode
						Description: "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.MapNestedAttributesOptions{},
			),
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::FIS::ExperimentTemplate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FIS::ExperimentTemplate").WithTerraformTypeName("awscc_fis_experiment_template").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_fis_experiment_template", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}