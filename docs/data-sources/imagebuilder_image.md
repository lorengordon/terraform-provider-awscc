---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_imagebuilder_image Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ImageBuilder::Image
---

# awscc_imagebuilder_image (Data Source)

Data Source schema for AWS::ImageBuilder::Image



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the image.
- **container_recipe_arn** (String) The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
- **distribution_configuration_arn** (String) The Amazon Resource Name (ARN) of the distribution configuration.
- **enhanced_image_metadata_enabled** (Boolean) Collects additional information about the image being created, including the operating system (OS) version and package list.
- **image_id** (String) The AMI ID of the EC2 AMI in current region.
- **image_recipe_arn** (String) The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
- **image_tests_configuration** (Attributes) The image tests configuration used when creating this image. (see [below for nested schema](#nestedatt--image_tests_configuration))
- **infrastructure_configuration_arn** (String) The Amazon Resource Name (ARN) of the infrastructure configuration.
- **name** (String) The name of the image.
- **tags** (Map of String) The tags associated with the image.

<a id="nestedatt--image_tests_configuration"></a>
### Nested Schema for `image_tests_configuration`

Read-Only:

- **image_tests_enabled** (Boolean) ImageTestsEnabled
- **timeout_minutes** (Number) TimeoutMinutes


