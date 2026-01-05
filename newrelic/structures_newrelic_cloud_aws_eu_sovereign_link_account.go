package newrelic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/newrelic/newrelic-client-go/v2/pkg/cloud"
)

// expandAwsEuSovereignLinkAccountInputForCreate expands the schema data into a CloudLinkCloudAccountsInput for create operations
func expandAwsEuSovereignLinkAccountInputForCreate(d *schema.ResourceData) cloud.CloudLinkCloudAccountsInput {
	input := cloud.CloudLinkCloudAccountsInput{
		Aws: []cloud.CloudAwsLinkAccountInput{
			{
				Arn:                  d.Get("arn").(string),
				MetricCollectionMode: cloud.CloudMetricCollectionMode(d.Get("metric_collection_mode").(string)),
				Name:                 d.Get("name").(string),
			},
		},
	}

	return input
}

// flattenAwsEuSovereignLinkAccountForRead flattens the linked account data into the schema for read operations
func flattenAwsEuSovereignLinkAccountForRead(linkedAccount *cloud.CloudLinkedAccount, d *schema.ResourceData, accountID int) error {
	_ = d.Set("account_id", accountID)
	_ = d.Set("name", linkedAccount.Name)
	_ = d.Set("arn", linkedAccount.AuthLabel)

	_ = d.Set("metric_collection_mode", string(linkedAccount.MetricCollectionMode))

	return nil
}

// expandAwsEuSovereignLinkAccountInputForUpdate expands the schema data into a CloudRenameAccountsInput for update operations
func expandAwsEuSovereignLinkAccountInputForUpdate(d *schema.ResourceData, linkedAccountID int) []cloud.CloudRenameAccountsInput {
	input := []cloud.CloudRenameAccountsInput{
		{
			LinkedAccountId: linkedAccountID,
			Name:           d.Get("name").(string),
		},
	}

	return input
}

// getEuSovereignLinkedAccountIDFromState extracts the linked account ID from the terraform state
func getEuSovereignLinkedAccountIDFromState(d *schema.ResourceData) int {
	linkedAccountID, _ := parseIDs(d.Id(), 1)
	return linkedAccountID[0]
}
