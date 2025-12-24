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

	if linkedAccount.MetricCollectionMode != nil {
		_ = d.Set("metric_collection_mode", string(*linkedAccount.MetricCollectionMode))
	}

	return nil
}

// expandAwsEuSovereignLinkAccountInputForUpdate expands the schema data into a CloudUpdateCloudAccountsInput for update operations
func expandAwsEuSovereignLinkAccountInputForUpdate(d *schema.ResourceData, linkedAccountID int) cloud.CloudUpdateCloudAccountsInput {
	input := cloud.CloudUpdateCloudAccountsInput{
		Aws: cloud.CloudAwsUpdateInput{
			LinkedAccountId:      linkedAccountID,
			MetricCollectionMode: cloud.CloudMetricCollectionMode(d.Get("metric_collection_mode").(string)),
			Name:                 d.Get("name").(string),
		},
	}

	return input
}

// getLinkedAccountIDFromState extracts the linked account ID from the terraform state
func getLinkedAccountIDFromState(d *schema.ResourceData) int {
	linkedAccountID, _ := parseIDs(d.Id(), 1)
	return linkedAccountID[0]
}

// expandCloudAwsEuSovereignIntegrationsInput expands the schema data for configuring integrations
func expandCloudAwsEuSovereignIntegrationsInput(d *schema.ResourceData, linkedAccountID int) cloud.CloudConfigureIntegrationInput {
	awsInput := cloud.CloudAwsIntegrationsInput{
		LinkedAccountId: linkedAccountID,
	}

	// ALB Integration
	if attr, ok := d.GetOk("alb"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationAlb(attr.([]interface{}))
		if expanded != nil {
			awsInput.Alb = expanded
		}
	}

	// API Gateway Integration
	if attr, ok := d.GetOk("api_gateway"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationApiGateway(attr.([]interface{}))
		if expanded != nil {
			awsInput.ApiGateway = expanded
		}
	}

	// Auto Scaling Integration
	if attr, ok := d.GetOk("auto_scaling"); ok {
		expanded := expandCloudAwsEuSovereignIntegrationAutoScaling(attr.([]interface{}))
		if expanded != nil {
			awsInput.AutoScaling = expanded
		}
	}

	// Add more integrations as needed...
	// This pattern would continue for all supported integrations

	input := cloud.CloudConfigureIntegrationInput{
		Aws: awsInput,
	}

	return input
}

// expandCloudAwsEuSovereignDisableIntegrationsInput expands the schema data for disabling integrations
func expandCloudAwsEuSovereignDisableIntegrationsInput(d *schema.ResourceData, linkedAccountID int) cloud.CloudDisableIntegrationInput {
	awsInput := cloud.CloudAwsDisableIntegrationsInput{
		LinkedAccountId: linkedAccountID,
	}

	// Check which integrations exist and need to be disabled
	if _, ok := d.GetOk("alb"); ok {
		awsInput.Alb = []cloud.CloudDisableAccountIntegrationInput{{}}
	}

	if _, ok := d.GetOk("api_gateway"); ok {
		awsInput.ApiGateway = []cloud.CloudDisableAccountIntegrationInput{{}}
	}

	if _, ok := d.GetOk("auto_scaling"); ok {
		awsInput.AutoScaling = []cloud.CloudDisableAccountIntegrationInput{{}}
	}

	// Continue for all integrations...

	input := cloud.CloudDisableIntegrationInput{
		Aws: awsInput,
	}

	return input
}

// flattenCloudAwsEuSovereignIntegrations flattens the integrations data from the API into the schema
func flattenCloudAwsEuSovereignIntegrations(linkedAccount *cloud.CloudLinkedAccount, accountID int, d *schema.ResourceData) error {
	_ = d.Set("account_id", accountID)
	_ = d.Set("linked_account_id", linkedAccount.Id)

	// Flatten each integration type based on what's configured
	for _, integration := range linkedAccount.Integrations {
		switch integration.Service.Slug {
		case "alb":
			if err := d.Set("alb", flattenCloudAwsEuSovereignIntegrationAlb(integration)); err != nil {
				return err
			}
		case "api-gateway":
			if err := d.Set("api_gateway", flattenCloudAwsEuSovereignIntegrationApiGateway(integration)); err != nil {
				return err
			}
		case "auto-scaling":
			if err := d.Set("auto_scaling", flattenCloudAwsEuSovereignIntegrationAutoScaling(integration)); err != nil {
				return err
			}
		// Add cases for all supported integrations...
		}
	}

	return nil
}

// Integration-specific expand and flatten functions
func expandCloudAwsEuSovereignIntegrationAlb(b []interface{}) *cloud.CloudAwsAlbIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudAwsAlbIntegrationInput{}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	if v, ok := cfg["aws_regions"]; ok && len(v.([]interface{})) > 0 {
		input.AwsRegions = expandStringList(v.([]interface{}))
	}

	if v, ok := cfg["fetch_extended_inventory"]; ok {
		input.FetchExtendedInventory = v.(bool)
	}

	if v, ok := cfg["fetch_tags"]; ok {
		input.FetchTags = v.(bool)
	}

	if v, ok := cfg["load_balancer_prefixes"]; ok && len(v.([]interface{})) > 0 {
		input.LoadBalancerPrefixes = expandStringList(v.([]interface{}))
	}

	if v, ok := cfg["tag_key"]; ok && v.(string) != "" {
		input.TagKey = v.(string)
	}

	if v, ok := cfg["tag_value"]; ok && v.(string) != "" {
		input.TagValue = v.(string)
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationAlb(integration cloud.CloudIntegration) []interface{} {
	result := make(map[string]interface{})

	if integration.MetricsPollingInterval != nil {
		result["metrics_polling_interval"] = *integration.MetricsPollingInterval
	}

	if integration.AwsRegions != nil {
		result["aws_regions"] = *integration.AwsRegions
	}

	if integration.FetchExtendedInventory != nil {
		result["fetch_extended_inventory"] = *integration.FetchExtendedInventory
	}

	if integration.FetchTags != nil {
		result["fetch_tags"] = *integration.FetchTags
	}

	if integration.LoadBalancerPrefixes != nil {
		result["load_balancer_prefixes"] = *integration.LoadBalancerPrefixes
	}

	if integration.TagKey != nil {
		result["tag_key"] = *integration.TagKey
	}

	if integration.TagValue != nil {
		result["tag_value"] = *integration.TagValue
	}

	return []interface{}{result}
}

func expandCloudAwsEuSovereignIntegrationApiGateway(b []interface{}) *cloud.CloudAwsApigatewayIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudAwsApigatewayIntegrationInput{}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	if v, ok := cfg["aws_regions"]; ok && len(v.([]interface{})) > 0 {
		input.AwsRegions = expandStringList(v.([]interface{}))
	}

	if v, ok := cfg["stage_prefixes"]; ok && len(v.([]interface{})) > 0 {
		input.StagePrefixes = expandStringList(v.([]interface{}))
	}

	if v, ok := cfg["tag_key"]; ok && v.(string) != "" {
		input.TagKey = v.(string)
	}

	if v, ok := cfg["tag_value"]; ok && v.(string) != "" {
		input.TagValue = v.(string)
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationApiGateway(integration cloud.CloudIntegration) []interface{} {
	result := make(map[string]interface{})

	if integration.MetricsPollingInterval != nil {
		result["metrics_polling_interval"] = *integration.MetricsPollingInterval
	}

	if integration.AwsRegions != nil {
		result["aws_regions"] = *integration.AwsRegions
	}

	if integration.StagePrefixes != nil {
		result["stage_prefixes"] = *integration.StagePrefixes
	}

	if integration.TagKey != nil {
		result["tag_key"] = *integration.TagKey
	}

	if integration.TagValue != nil {
		result["tag_value"] = *integration.TagValue
	}

	return []interface{}{result}
}

func expandCloudAwsEuSovereignIntegrationAutoScaling(b []interface{}) *cloud.CloudAwsAutoscalingIntegrationInput {
	if len(b) == 0 || b[0] == nil {
		return nil
	}

	cfg := b[0].(map[string]interface{})
	input := &cloud.CloudAwsAutoscalingIntegrationInput{}

	if v, ok := cfg["metrics_polling_interval"]; ok && v.(int) != 0 {
		input.MetricsPollingInterval = v.(int)
	}

	if v, ok := cfg["aws_regions"]; ok && len(v.([]interface{})) > 0 {
		input.AwsRegions = expandStringList(v.([]interface{}))
	}

	return input
}

func flattenCloudAwsEuSovereignIntegrationAutoScaling(integration cloud.CloudIntegration) []interface{} {
	result := make(map[string]interface{})

	if integration.MetricsPollingInterval != nil {
		result["metrics_polling_interval"] = *integration.MetricsPollingInterval
	}

	if integration.AwsRegions != nil {
		result["aws_regions"] = *integration.AwsRegions
	}

	return []interface{}{result}
}

// Additional integration expand/flatten functions would follow the same pattern...
// For brevity, these are omitted but would be needed for a complete implementation