---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53recoverycontrol_routing_control Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53RecoveryControl::RoutingControl
---

# awscc_route53recoverycontrol_routing_control (Data Source)

Data Source schema for AWS::Route53RecoveryControl::RoutingControl



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **cluster_arn** (String) Arn associated with Control Panel
- **control_panel_arn** (String) The Amazon Resource Name (ARN) of the control panel.
- **name** (String) The name of the routing control. You can use any non-white space character in the name.
- **routing_control_arn** (String) The Amazon Resource Name (ARN) of the routing control.
- **status** (String) The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.


