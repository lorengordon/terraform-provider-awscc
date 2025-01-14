---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_frauddetector_label Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::FraudDetector::Label
---

# awscc_frauddetector_label (Data Source)

Data Source schema for AWS::FraudDetector::Label



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The label ARN.
- **created_time** (String) The timestamp when the label was created.
- **description** (String) The label description.
- **last_updated_time** (String) The timestamp when the label was last updated.
- **name** (String) The name of the label.
- **tags** (Attributes List) Tags associated with this label. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)


