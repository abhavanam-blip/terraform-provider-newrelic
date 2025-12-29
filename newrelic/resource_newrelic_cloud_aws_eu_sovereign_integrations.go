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
	client := meta.(*ProviderConfig).NewClient

	accountID := selectAccountID(meta, d)
	linkedAccountID := d.Get("linked_account_id").(int)

	configureInput := expandCloudAwsEuSovereignIntegrationsInput(d, linkedAccountID)

	log.Printf("[INFO] Creating New Relic AWS EU Sovereign integration for linked account %d", linkedAccountID)

	cloudConfigureIntegration, err := client.Cloud.CloudConfigureIntegrationWithContext(ctx, accountID, configureInput)
	if err != nil {
		return diag.FromErr(err)
	}

	var integrationIds []string
	for _, integration := range cloudConfigureIntegration.Integrations {
		integrationIds = append(integrationIds, strconv.Itoa(integration.Id))
	}

	d.SetId(buildCompositeID(integrationIds))

	return resourceNewRelicCloudAwsEuSovereignIntegrationsRead(ctx, d, meta)
}

func resourceNewRelicCloudAwsEuSovereignIntegrationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).NewClient

	accountID := selectAccountID(meta, d)
	linkedAccountID := d.Get("linked_account_id").(int)

	log.Printf("[INFO] Reading New Relic AWS EU Sovereign integration for linked account %d", linkedAccountID)

	linkedAccount, err := client.Cloud.GetLinkedAccount(accountID, linkedAccountID)
	if err != nil {
		if _, ok := err.(*cloud.NotFoundError); ok {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	return diag.FromErr(flattenCloudAwsEuSovereignIntegrations(linkedAccount, accountID, d))
}

func resourceNewRelicCloudAwsEuSovereignIntegrationsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).NewClient
	accountID := selectAccountID(meta, d)
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
	client := meta.(*ProviderConfig).NewClient

	accountID := selectAccountID(meta, d)
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