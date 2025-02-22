---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_ec2_fleet Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::EC2Fleet
---

# awscc_ec2_ec2_fleet (Resource)

Resource Type definition for AWS::EC2::EC2Fleet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **launch_template_configs** (Attributes List) (see [below for nested schema](#nestedatt--launch_template_configs))
- **target_capacity_specification** (Attributes) (see [below for nested schema](#nestedatt--target_capacity_specification))

### Optional

- **context** (String)
- **excess_capacity_termination_policy** (String)
- **on_demand_options** (Attributes) (see [below for nested schema](#nestedatt--on_demand_options))
- **replace_unhealthy_instances** (Boolean)
- **spot_options** (Attributes) (see [below for nested schema](#nestedatt--spot_options))
- **tag_specifications** (Attributes List) (see [below for nested schema](#nestedatt--tag_specifications))
- **terminate_instances_with_expiration** (Boolean)
- **type** (String)
- **valid_from** (String)
- **valid_until** (String)

### Read-Only

- **fleet_id** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--launch_template_configs"></a>
### Nested Schema for `launch_template_configs`

Required:

- **launch_template_specification** (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--launch_template_specification))
- **overrides** (Attributes List) (see [below for nested schema](#nestedatt--launch_template_configs--overrides))

<a id="nestedatt--launch_template_configs--launch_template_specification"></a>
### Nested Schema for `launch_template_configs.launch_template_specification`

Required:

- **launch_template_id** (String)
- **launch_template_name** (String)
- **version** (String)


<a id="nestedatt--launch_template_configs--overrides"></a>
### Nested Schema for `launch_template_configs.overrides`

Required:

- **availability_zone** (String)
- **instance_type** (String)
- **max_price** (String)
- **placement** (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--placement))
- **priority** (Number)
- **subnet_id** (String)
- **weighted_capacity** (Number)

<a id="nestedatt--launch_template_configs--overrides--placement"></a>
### Nested Schema for `launch_template_configs.overrides.placement`

Required:

- **affinity** (String)
- **availability_zone** (String)
- **group_name** (String)
- **host_id** (String)
- **host_resource_group_arn** (String)
- **partition_number** (Number)
- **spread_domain** (String)
- **tenancy** (String)




<a id="nestedatt--target_capacity_specification"></a>
### Nested Schema for `target_capacity_specification`

Required:

- **default_target_capacity_type** (String)
- **on_demand_target_capacity** (Number)
- **spot_target_capacity** (Number)
- **total_target_capacity** (Number)


<a id="nestedatt--on_demand_options"></a>
### Nested Schema for `on_demand_options`

Optional:

- **allocation_strategy** (String)
- **capacity_reservation_options** (Attributes) (see [below for nested schema](#nestedatt--on_demand_options--capacity_reservation_options))
- **max_total_price** (String)
- **min_target_capacity** (Number)
- **single_availability_zone** (Boolean)
- **single_instance_type** (Boolean)

<a id="nestedatt--on_demand_options--capacity_reservation_options"></a>
### Nested Schema for `on_demand_options.capacity_reservation_options`

Optional:

- **usage_strategy** (String)



<a id="nestedatt--spot_options"></a>
### Nested Schema for `spot_options`

Optional:

- **allocation_strategy** (String)
- **instance_interruption_behavior** (String)
- **instance_pools_to_use_count** (Number)
- **max_total_price** (String)
- **min_target_capacity** (Number)
- **single_availability_zone** (Boolean)
- **single_instance_type** (Boolean)


<a id="nestedatt--tag_specifications"></a>
### Nested Schema for `tag_specifications`

Optional:

- **resource_type** (String)
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tag_specifications--tags))

<a id="nestedatt--tag_specifications--tags"></a>
### Nested Schema for `tag_specifications.tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_ec2_fleet.example <resource ID>
```
