---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_healthlake_fhir_datastore Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::HealthLake::FHIRDatastore
---

# awscc_healthlake_fhir_datastore (Data Source)

Data Source schema for AWS::HealthLake::FHIRDatastore



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **created_at** (Attributes) The time that a Data Store was created. (see [below for nested schema](#nestedatt--created_at))
- **datastore_arn** (String) The Amazon Resource Name used in the creation of the Data Store.
- **datastore_endpoint** (String) The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL.
- **datastore_id** (String) The AWS-generated ID number for the Data Store.
- **datastore_name** (String) The user-generated name for the Data Store.
- **datastore_status** (String) The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'.
- **datastore_type_version** (String) The FHIR version. Only R4 version data is supported.
- **preload_data_config** (Attributes) The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported. (see [below for nested schema](#nestedatt--preload_data_config))
- **sse_configuration** (Attributes) The server-side encryption key configuration for a customer provided encryption key. (see [below for nested schema](#nestedatt--sse_configuration))
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--created_at"></a>
### Nested Schema for `created_at`

Read-Only:

- **nanos** (Number) Nanoseconds.
- **seconds** (String) Seconds since epoch.


<a id="nestedatt--preload_data_config"></a>
### Nested Schema for `preload_data_config`

Read-Only:

- **preload_data_type** (String) The type of preloaded data. Only Synthea preloaded data is supported.


<a id="nestedatt--sse_configuration"></a>
### Nested Schema for `sse_configuration`

Read-Only:

- **kms_encryption_config** (Attributes) The customer-managed-key (CMK) used when creating a Data Store. If a customer owned key is not specified, an AWS owned key will be used for encryption. (see [below for nested schema](#nestedatt--sse_configuration--kms_encryption_config))

<a id="nestedatt--sse_configuration--kms_encryption_config"></a>
### Nested Schema for `sse_configuration.kms_encryption_config`

Read-Only:

- **cmk_type** (String) The type of customer-managed-key (CMK) used for encryption. The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.
- **kms_key_id** (String) The KMS encryption key id/alias used to encrypt the Data Store contents at rest.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) The key of the tag.
- **value** (String) The value of the tag.


