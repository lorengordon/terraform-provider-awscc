// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

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
	registry.AddResourceTypeFactory("awscc_cloudfront_origin_request_policy", originRequestPolicyResourceType)
}

// originRequestPolicyResourceType returns the Terraform awscc_cloudfront_origin_request_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFront::OriginRequestPolicy resource type.
func originRequestPolicyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"origin_request_policy_config": {
			// Property: OriginRequestPolicyConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Comment": {
			//       "type": "string"
			//     },
			//     "CookiesConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CookieBehavior": {
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "Cookies": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "CookieBehavior"
			//       ],
			//       "type": "object"
			//     },
			//     "HeadersConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "HeaderBehavior": {
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "Headers": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "HeaderBehavior"
			//       ],
			//       "type": "object"
			//     },
			//     "Name": {
			//       "type": "string"
			//     },
			//     "QueryStringsConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "QueryStringBehavior": {
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "QueryStrings": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "QueryStringBehavior"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "HeadersConfig",
			//     "CookiesConfig",
			//     "QueryStringsConfig"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Optional: true,
					},
					"cookies_config": {
						// Property: CookiesConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cookie_behavior": {
									// Property: CookieBehavior
									Type:     types.StringType,
									Required: true,
								},
								"cookies": {
									// Property: Cookies
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Required: true,
					},
					"headers_config": {
						// Property: HeadersConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"header_behavior": {
									// Property: HeaderBehavior
									Type:     types.StringType,
									Required: true,
								},
								"headers": {
									// Property: Headers
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Required: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
					},
					"query_strings_config": {
						// Property: QueryStringsConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"query_string_behavior": {
									// Property: QueryStringBehavior
									Type:     types.StringType,
									Required: true,
								},
								"query_strings": {
									// Property: QueryStrings
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CloudFront::OriginRequestPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::OriginRequestPolicy").WithTerraformTypeName("awscc_cloudfront_origin_request_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"comment":                      "Comment",
		"cookie_behavior":              "CookieBehavior",
		"cookies":                      "Cookies",
		"cookies_config":               "CookiesConfig",
		"header_behavior":              "HeaderBehavior",
		"headers":                      "Headers",
		"headers_config":               "HeadersConfig",
		"id":                           "Id",
		"last_modified_time":           "LastModifiedTime",
		"name":                         "Name",
		"origin_request_policy_config": "OriginRequestPolicyConfig",
		"query_string_behavior":        "QueryStringBehavior",
		"query_strings":                "QueryStrings",
		"query_strings_config":         "QueryStringsConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cloudfront_origin_request_policy", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
