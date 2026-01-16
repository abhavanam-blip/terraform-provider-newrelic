package newrelic

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			// EU Sovereign only supports 4 integrations: cloudtrail, xray, health, trustedadvisor
			"cloudtrail": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "CloudTrail integration",
				Elem:        cloudAwsEuSovereignIntegrationsCloudtrailElem(),
			},
			"health": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Health integration",
				Elem:        cloudAwsEuSovereignIntegrationsHealthElem(),
			},
			"trusted_advisor": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Trusted Advisor integration",
				Elem:        cloudAwsEuSovereignIntegrationsTrustedAdvisorElem(),
			},
			"x_ray": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "X-Ray integration",
				Elem:        cloudAwsEuSovereignIntegrationsXRayElem(),
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

	payload, err := client.Cloud.CloudConfigureIntegrationWithContext(ctx, accountID, configureInput)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	if len(payload.Errors) > 0 {
		for _, err := range payload.Errors {
			log.Printf("[ERROR] CloudConfigureIntegration error: Type=%s, Message=%s", err.Type, err.Message)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Type + " " + err.Message,
			})
		}
		return diags
	}

	log.Printf("[DEBUG] CloudConfigureIntegration response: Integrations=%d", len(payload.Integrations))

	// Set ID using linked account ID
	d.SetId(strconv.Itoa(linkedAccountID))

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
		d.SetId("")
		return nil
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

	payload, err := client.Cloud.CloudConfigureIntegrationWithContext(ctx, accountID, configureInput)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	if len(payload.Errors) > 0 {
		for _, err := range payload.Errors {
			log.Printf("[ERROR] CloudConfigureIntegration error: Type=%s, Message=%s", err.Type, err.Message)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Type + " " + err.Message,
			})
		}
		return diags
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

	payload, err := client.Cloud.CloudDisableIntegrationWithContext(ctx, accountID, disableInput)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	if len(payload.Errors) > 0 {
		for _, err := range payload.Errors {
			log.Printf("[ERROR] CloudDisableIntegration error: Type=%s, Message=%s", err.Type, err.Message)
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Type + " " + err.Message,
			})
		}
		return diags
	}

	return nil
}

// CloudTrail integration schema
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
		},
	}
}

// Health integration schema
func cloudAwsEuSovereignIntegrationsHealthElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_polling_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The data polling interval in seconds",
			},
		},
	}
}

// Trusted Advisor integration schema
func cloudAwsEuSovereignIntegrationsTrustedAdvisorElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_polling_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The data polling interval in seconds",
			},
		},
	}
}

// X-Ray integration schema
func cloudAwsEuSovereignIntegrationsXRayElem() *schema.Resource {
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
		},
	}
}