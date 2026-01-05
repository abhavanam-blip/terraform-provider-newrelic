package newrelic

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/newrelic/newrelic-client-go/v2/pkg/cloud"
)

func resourceNewRelicCloudAwsEuSovereignIntegrations() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNewRelicCloudAwsEuSovereignIntegrationsCreate,
		ReadContext:   resourceNewRelicCloudAwsEuSovereignIntegrationsRead,
		UpdateContext: resourceNewRelicCloudAwsEuSovereignIntegrationsUpdate,
		DeleteContext: resourceNewRelicCloudAwsEuSovereignIntegrationsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The ID of the account in New Relic.",
			},
			"linked_account_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the linked AWS EU Sovereign account in New Relic.",
			},
			"cloudtrail": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "CloudTrail",
				Elem:        cloudAwsEuSovereignIntegrationsCloudtrailElem(),
			},
			"health": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "AWS Health",
				Elem:        cloudAwsEuSovereignIntegrationsHealthElem(),
			},
			"trusted_advisor": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "AWS Trusted Advisor",
				Elem:        cloudAwsEuSovereignIntegrationsTrustedAdvisorElem(),
			},
			"xray": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "AWS X-Ray",
				Elem:        cloudAwsEuSovereignIntegrationsXrayElem(),
			},
		},
	}
}

func resourceNewRelicCloudAwsEuSovereignIntegrationsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConfig := meta.(*ProviderConfig)
	client := providerConfig.NewClient

	accountID := selectAccountID(providerConfig, d)
	linkedAccountID := d.Get("linked_account_id").(int)

	configureInput := expandCloudAwsEuSovereignIntegrationsInput(d, linkedAccountID)

	log.Printf("[INFO] Creating New Relic AWS EU Sovereign integration for linked account %d", linkedAccountID)

	cloudConfigureIntegration, err := client.Cloud.CloudConfigureIntegrationWithContext(ctx, accountID, configureInput)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(cloudConfigureIntegration.Integrations) > 0 {
		d.SetId(strconv.Itoa(linkedAccountID))
	}

	return resourceNewRelicCloudAwsEuSovereignIntegrationsRead(ctx, d, meta)
}

func resourceNewRelicCloudAwsEuSovereignIntegrationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConfig := meta.(*ProviderConfig)
	client := providerConfig.NewClient

	accountID := selectAccountID(providerConfig, d)
	linkedAccountID := d.Get("linked_account_id").(int)

	log.Printf("[INFO] Reading New Relic AWS EU Sovereign integration for linked account %d", linkedAccountID)

	linkedAccount, err := client.Cloud.GetLinkedAccount(accountID, linkedAccountID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diag.FromErr(flattenCloudAwsEuSovereignIntegrations(linkedAccount, accountID, d))
}

func resourceNewRelicCloudAwsEuSovereignIntegrationsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConfig := meta.(*ProviderConfig)
	client := providerConfig.NewClient
	accountID := selectAccountID(providerConfig, d)
	linkedAccountID := d.Get("linked_account_id").(int)

	configureInput := expandCloudAwsEuSovereignIntegrationsInput(d, linkedAccountID)

	log.Printf("[INFO] Updating New Relic AWS EU Sovereign integration for linked account %d", linkedAccountID)

	_, err := client.Cloud.CloudConfigureIntegrationWithContext(ctx, accountID, configureInput)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNewRelicCloudAwsEuSovereignIntegrationsRead(ctx, d, meta)
}

func resourceNewRelicCloudAwsEuSovereignIntegrationsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConfig := meta.(*ProviderConfig)
	client := providerConfig.NewClient

	accountID := selectAccountID(providerConfig, d)
	linkedAccountID := d.Get("linked_account_id").(int)

	disableInput := expandCloudAwsEuSovereignDisableIntegrationsInput(d, linkedAccountID)

	log.Printf("[INFO] Deleting New Relic AWS EU Sovereign integration for linked account %d", linkedAccountID)

	_, err := client.Cloud.CloudDisableIntegrationWithContext(ctx, accountID, disableInput)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// CloudTrail integration schema element
func cloudAwsEuSovereignIntegrationsCloudtrailElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_polling_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The data polling interval in seconds",
			},
			"aws_regions": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Specify each AWS region that includes the resources that you want to monitor",
			},
			"fetch_extended_inventory": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine if extra inventory data be collected or not. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"fetch_tags": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specify if tags should be collected. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"tag_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
			"tag_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
		},
	}
}

// AWS Health integration schema element
func cloudAwsEuSovereignIntegrationsHealthElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_polling_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The data polling interval in seconds",
			},
			"fetch_extended_inventory": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine if extra inventory data be collected or not. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"fetch_tags": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specify if tags should be collected. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"tag_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
			"tag_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
		},
	}
}

// AWS Trusted Advisor integration schema element
func cloudAwsEuSovereignIntegrationsTrustedAdvisorElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_polling_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The data polling interval in seconds",
			},
			"fetch_extended_inventory": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine if extra inventory data be collected or not. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"fetch_tags": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specify if tags should be collected. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"tag_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
			"tag_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
		},
	}
}

// AWS X-Ray integration schema element
func cloudAwsEuSovereignIntegrationsXrayElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_polling_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The data polling interval in seconds",
			},
			"aws_regions": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Specify each AWS region that includes the resources that you want to monitor",
			},
			"fetch_extended_inventory": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine if extra inventory data be collected or not. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"fetch_tags": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specify if tags should be collected. May affect total data collection time and contribute to the Cloud provider API rate limit.",
			},
			"tag_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag key associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
			"tag_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a Tag value associated with the resources that you want to monitor. Filter values are case-sensitive.",
			},
		},
	}
}

// expandCloudAwsEuSovereignIntegrationsInput expands the schema data for EU SOV integrations (4 supported services)
func expandCloudAwsEuSovereignIntegrationsInput(d *schema.ResourceData, linkedAccountID int) cloud.CloudIntegrationsInput {
	awsInput := cloud.CloudAwsIntegrationsInput{}

	// EU SOV supports only 4 services
	if v, ok := d.GetOk("cloudtrail"); ok {
		awsInput.Cloudtrail = expandCloudAwsIntegrationCloudtrailInput(v.([]interface{}), linkedAccountID)
	}

	if v, ok := d.GetOk("health"); ok {
		awsInput.Health = expandCloudAwsIntegrationHealthInput(v.([]interface{}), linkedAccountID)
	}

	if v, ok := d.GetOk("trusted_advisor"); ok {
		awsInput.Trustedadvisor = expandCloudAwsIntegrationTrustedAdvisorInput(v.([]interface{}), linkedAccountID)
	}

	if v, ok := d.GetOk("xray"); ok {
		awsInput.AwsXray = expandCloudAwsIntegrationXRayInput(v.([]interface{}), linkedAccountID)
	}

	input := cloud.CloudIntegrationsInput{
		Aws: awsInput,
	}

	return input
}

// expandCloudAwsEuSovereignDisableIntegrationsInput expands the schema data for disabling EU SOV integrations
func expandCloudAwsEuSovereignDisableIntegrationsInput(d *schema.ResourceData, linkedAccountID int) cloud.CloudDisableIntegrationsInput {
	awsInput := cloud.CloudAwsDisableIntegrationsInput{}

	// EU SOV supports only 4 services
	if _, ok := d.GetOk("cloudtrail"); ok {
		awsInput.Cloudtrail = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	if _, ok := d.GetOk("health"); ok {
		awsInput.Health = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	if _, ok := d.GetOk("trusted_advisor"); ok {
		awsInput.Trustedadvisor = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	if _, ok := d.GetOk("xray"); ok {
		awsInput.AwsXray = []cloud.CloudDisableAccountIntegrationInput{{LinkedAccountId: linkedAccountID}}
	}

	input := cloud.CloudDisableIntegrationsInput{
		Aws: awsInput,
	}

	return input
}

// flattenCloudAwsEuSovereignIntegrations flattens EU SOV integrations data from the API into the schema
func flattenCloudAwsEuSovereignIntegrations(linkedAccount *cloud.CloudLinkedAccount, accountID int, d *schema.ResourceData) error {
	_ = d.Set("account_id", accountID)
	_ = d.Set("linked_account_id", linkedAccount.ID)

	// EU SOV supports only 4 services - use regular AWS integration flatten functions
	for _, integration := range linkedAccount.Integrations {
		switch t := integration.(type) {
		case *cloud.CloudCloudtrailIntegration:
			_ = d.Set("cloudtrail", flattenCloudAwsCloudTrailIntegration(t))
		case *cloud.CloudHealthIntegration:
			_ = d.Set("health", flattenCloudAwsHealthIntegration(t))
		case *cloud.CloudTrustedadvisorIntegration:
			_ = d.Set("trusted_advisor", flattenCloudAwsTrustedAdvisorIntegration(t))
		case *cloud.CloudAwsXrayIntegration:
			_ = d.Set("xray", flattenCloudAwsXRayIntegration(t))
		}
	}

	return nil
}