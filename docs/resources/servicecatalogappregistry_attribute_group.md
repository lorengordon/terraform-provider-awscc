---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_servicecatalogappregistry_attribute_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.
---

# awscc_servicecatalogappregistry_attribute_group (Resource)

Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **attributes** (Map of String)
- **name** (String) The name of the attribute group.

### Optional

- **description** (String) The description of the attribute group.
- **tags** (Map of String)

### Read-Only

- **arn** (String)
- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_servicecatalogappregistry_attribute_group.example <resource ID>
```
