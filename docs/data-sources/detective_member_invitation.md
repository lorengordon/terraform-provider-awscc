---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_detective_member_invitation Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Detective::MemberInvitation
---

# awscc_detective_member_invitation (Data Source)

Data Source schema for AWS::Detective::MemberInvitation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **disable_email_notification** (Boolean) When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.
- **graph_arn** (String) The ARN of the graph to which the member account will be invited
- **member_email_address** (String) The root email address for the account to be invited, for validation. Updating this field has no effect.
- **member_id** (String) The AWS account ID to be invited to join the graph as a member
- **message** (String) A message to be included in the email invitation sent to the invited account. Updating this field has no effect.


