// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53ResolverResolverQueryLoggingConfig_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverQueryLoggingConfig", "awscc_route53resolver_resolver_query_logging_config", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
	})
}

func TestAccAWSRoute53ResolverResolverQueryLoggingConfig_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverQueryLoggingConfig", "awscc_route53resolver_resolver_query_logging_config", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}