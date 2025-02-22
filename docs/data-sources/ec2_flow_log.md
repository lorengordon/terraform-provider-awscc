---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_flow_log Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::FlowLog
---

# awscc_ec2_flow_log (Data Source)

Data Source schema for AWS::EC2::FlowLog



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **deliver_logs_permission_arn** (String) The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn or LogGroupName.
- **log_destination** (String) Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group or an Amazon S3 bucket. The value specified for this parameter depends on the value specified for LogDestinationType.
- **log_destination_type** (String) Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
- **log_format** (String) The fields to include in the flow log record, in the order in which they should appear.
- **log_group_name** (String) The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn or LogGroupName.
- **max_aggregation_interval** (Number) The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
- **resource_id** (String) The ID of the subnet, network interface, or VPC for which you want to create a flow log.
- **resource_type** (String) The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
- **tags** (Attributes List) The tags to apply to the flow logs. (see [below for nested schema](#nestedatt--tags))
- **traffic_type** (String) The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)


