---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_model Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ApiGateway::Model
---

# awscc_apigateway_model (Data Source)

Data Source schema for AWS::ApiGateway::Model



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **content_type** (String) The content type for the model.
- **description** (String) A description that identifies this model.
- **name** (String) A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
- **rest_api_id** (String) The ID of a REST API with which to associate this model.
- **schema** (String) The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.


