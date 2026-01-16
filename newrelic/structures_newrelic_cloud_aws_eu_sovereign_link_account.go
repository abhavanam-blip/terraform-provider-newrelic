package newrelic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/newrelic/newrelic-client-go/v2/pkg/cloud"
)

// expandAwsEuSovereignLinkAccountInputForCreate expands the schema data into a CloudLinkCloudAccountsInput for create operations
func expandAwsEuSovereignLinkAccountInputForCreate(d *schema.ResourceData) cloud.CloudLinkCloudAccountsInput {
	input := cloud.CloudLinkCloudAccountsInput{
		AwsEuSovereign: []cloud.CloudAwsEuSovereignLinkAccountInput{
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

	if linkedAccount.MetricCollectionMode != "" {
		_ = d.Set("metric_collection_mode", string(linkedAccount.MetricCollectionMode))
	}

	return nil
}

// expandAwsEuSovereignLinkAccountInputForUpdate expands the schema data into a CloudUpdateCloudAccountsInput for update operations
func expandAwsEuSovereignLinkAccountInputForUpdate(d *schema.ResourceData, linkedAccountID int) cloud.CloudUpdateCloudAccountsInput {
	input := cloud.CloudUpdateCloudAccountsInput{
		AwsEuSovereign: []cloud.CloudAwsEuSovereignUpdateAccountInput{
			{
				LinkedAccountId: linkedAccountID,
				Name:            d.Get("name").(string),
			},
		},
	}

	return input
}

// getAwsEuSovereignLinkedAccountIDFromState extracts the linked account ID from the terraform state
func getAwsEuSovereignLinkedAccountIDFromState(d *schema.ResourceData) int {
	linkedAccountID, _ := parseIDs(d.Id(), 1)
	return linkedAccountID[0]
}

// expandCloudAwsEuSovereignIntegrationsInput expands the schema data for configuring integrations
// EU Sovereign only supports: cloudtrail, xray, health, trustedadvisor
func expandCloudAwsEuSovereignIntegrationsInput(d *schema.ResourceData, linkedAccountID int) cloud.CloudIntegrationsInput {
	awsEuSovereignInput := cloud.CloudAwsEuSovereignIntegrationsInput{}

	// CloudTrail Integration
	if attr, ok := d.GetOk("cloudtrail"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationCloudtrail(attr.([]interface{}), linkedAccountID)
		if expanded != nil {
			awsEuSovereignInput.Cloudtrail = []cloud.CloudCloudtrailIntegrationInput{*expanded}
		}
	}

	// Health Integration
	if attr, ok := d.GetOk("health"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationHealth(attr.([]interface{}), linkedAccountID)
		if expanded != nil {
			awsEuSovereignInput.Health = []cloud.CloudHealthIntegrationInput{*expanded}
		}
	}

	// Trusted Advisor Integration
	if attr, ok := d.GetOk("trusted_advisor"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationTrustedAdvisor(attr.([]interface{}), linkedAccountID)
		if expanded != nil {
			awsEuSovereignInput.Trustedadvisor = []cloud.CloudTrustedadvisorIntegrationInput{*expanded}
		}
	}

	// X-Ray Integration
	if attr, ok := d.GetOk("x_ray"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationXRay(attr.([]interface{}), linkedAccountID)
		if expanded != nil {
			awsEuSovereignInput.AwsXray = []cloud.CloudAwsXrayIntegrationInput{*expanded}
		}
	}

	input := cloud.CloudIntegrationsInput{
		AwsEuSovereign: awsEuSovereignInput,
	}

	return input
}

// expandCloudAwsEuSovereignDisableIntegrationsInput expands the schema data for disabling integrations
func expandCloudAwsEuSovereignDisableIntegrationsInput(d *schema.ResourceData, linkedAccountID int) cloud.CloudDisableIntegrationsInput {
	awsEuSovereignInput := cloud.CloudAwsEuSovereignDisableIntegrationsInput{}

	if _, ok := d.GetOk("cloudtrail"); ok {
		awsEuSovereignInput.Cloudtrail = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	if _, ok := d.GetOk("health"); ok {
		awsEuSovereignInput.Health = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	if _, ok := d.GetOk("trusted_advisor"); ok {
		awsEuSovereignInput.Trustedadvisor = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	if _, ok := d.GetOk("x_ray"); ok {
		awsEuSovereignInput.AwsXray = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	input := cloud.CloudDisableIntegrationsInput{
		AwsEuSovereign: awsEuSovereignInput,
	}

	return input
}

// flattenCloudAwsEuSovereignIntegrations flattens the integrations data from the API into the schema
func flattenCloudAwsEuSovereignIntegrations(linkedAccount *cloud.CloudLinkedAccount, accountID int, d *schema.ResourceData) error {
	_ = d.Set("account_id", accountID)
	_ = d.Set("linked_account_id", linkedAccount.ID)

	for _, i := range linkedAccount.Integrations {
		switch t := i.(type) {
		case *cloud.CloudCloudtrailIntegration:
			_ = d.Set("cloudtrail", flattenCloudAwsEuSovereignIntegrationCloudtrail(t))
		case *cloud.CloudHealthIntegration:
			_ = d.Set("health", flattenCloudAwsEuSovereignIntegrationHealth(t))
		case *cloud.CloudTrustedadvisorIntegration:
			_ = d.Set("trusted_advisor", flattenCloudAwsEuSovereignIntegrationTrustedAdvisor(t))
		case *cloud.CloudAwsXrayIntegration:
			_ = d.Set("x_ray", flattenCloudAwsEuSovereignIntegrationXRay(t))
		}
	}

	return nil
}

// CloudTrail expand/flatten
func expandCloudAwsEuSovereignIntegrationCloudtrail(b []interface{}, linkedAccountID int) *cloud.CloudCloudtrailIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudCloudtrailIntegrationInput{
		LinkedAccountId: linkedAccountID,
	}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	if v, ok := cfg["aws_regions"]; ok && len(v.([]interface{})) > 0 {
		input.AwsRegions = expandStringList(v.([]interface{}))
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationCloudtrail(t *cloud.CloudCloudtrailIntegration) []interface{} {
	result := make(map[string]interface{})
	result["metrics_polling_interval"] = t.MetricsPollingInterval
	result["aws_regions"] = t.AwsRegions
	return []interface{}{result}
}

// Health expand/flatten
func expandCloudAwsEuSovereignIntegrationHealth(b []interface{}, linkedAccountID int) *cloud.CloudHealthIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudHealthIntegrationInput{
		LinkedAccountId: linkedAccountID,
	}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationHealth(t *cloud.CloudHealthIntegration) []interface{} {
	result := make(map[string]interface{})
	result["metrics_polling_interval"] = t.MetricsPollingInterval
	return []interface{}{result}
}

// Trusted Advisor expand/flatten
func expandCloudAwsEuSovereignIntegrationTrustedAdvisor(b []interface{}, linkedAccountID int) *cloud.CloudTrustedadvisorIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudTrustedadvisorIntegrationInput{
		LinkedAccountId: linkedAccountID,
	}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationTrustedAdvisor(t *cloud.CloudTrustedadvisorIntegration) []interface{} {
	result := make(map[string]interface{})
	result["metrics_polling_interval"] = t.MetricsPollingInterval
	return []interface{}{result}
}

// X-Ray expand/flatten
func expandCloudAwsEuSovereignIntegrationXRay(b []interface{}, linkedAccountID int) *cloud.CloudAwsXrayIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudAwsXrayIntegrationInput{
		LinkedAccountId: linkedAccountID,
	}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	if v, ok := cfg["aws_regions"]; ok && len(v.([]interface{})) > 0 {
		input.AwsRegions = expandStringList(v.([]interface{}))
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationXRay(t *cloud.CloudAwsXrayIntegration) []interface{} {
	result := make(map[string]interface{})
	result["metrics_polling_interval"] = t.MetricsPollingInterval
	result["aws_regions"] = t.AwsRegions
	return []interface{}{result}
}