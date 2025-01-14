---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_detective_graph Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Detective::Graph
---

# awscc_detective_graph (Resource)

Resource schema for AWS::Detective::Graph



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) The Detective graph ARN
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
- **value** (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_detective_graph.example <resource ID>
```
