// Code generated by generators/resource/main.go; DO NOT EDIT.

package customerprofiles

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
	registry.AddResourceTypeFactory("awscc_customerprofiles_domain", domainResourceType)
}

// domainResourceType returns the Terraform awscc_customerprofiles_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CustomerProfiles::Domain resource type.
func domainResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got created",
			//   "type": "string"
			// }
			Description: "The time of this integration got created",
			Type:        types.StringType,
			Computed:    true,
		},
		"dead_letter_queue_url": {
			// Property: DeadLetterQueueUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the SQS dead letter queue",
			//   "maxLength": 255,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The URL of the SQS dead letter queue",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 255),
			},
		},
		"default_encryption_key": {
			// Property: DefaultEncryptionKey
			// CloudFormation resource type schema:
			// {
			//   "description": "The default encryption key",
			//   "maxLength": 255,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The default encryption key",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 255),
			},
		},
		"default_expiration_days": {
			// Property: DefaultExpirationDays
			// CloudFormation resource type schema:
			// {
			//   "description": "The default number of days until the data within the domain expires.",
			//   "maximum": 1098,
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Description: "The default number of days until the data within the domain expires.",
			Type:        types.NumberType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(1, 1098),
			},
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of the domain.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique name of the domain.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"last_updated_at": {
			// Property: LastUpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got last updated at",
			//   "type": "string"
			// }
			Description: "The time of this integration got last updated at",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the domain",
			//   "items": {
			//     "additionalProperties": false,
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
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the domain",
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
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
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
		Description: "A domain defined for 3rd party data source in Profile Service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::Domain").WithTerraformTypeName("awscc_customerprofiles_domain")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":              "CreatedAt",
		"dead_letter_queue_url":   "DeadLetterQueueUrl",
		"default_encryption_key":  "DefaultEncryptionKey",
		"default_expiration_days": "DefaultExpirationDays",
		"domain_name":             "DomainName",
		"key":                     "Key",
		"last_updated_at":         "LastUpdatedAt",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_customerprofiles_domain", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
