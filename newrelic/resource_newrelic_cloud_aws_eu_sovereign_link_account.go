package newrelic

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/newrelic/newrelic-client-go/v2/pkg/cloud"
)

func resourceNewRelicCloudAwsEuSovereignLinkAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNewRelicCloudAwsEuSovereignLinkAccountCreate,
		ReadContext:   resourceNewRelicCloudAwsEuSovereignLinkAccountRead,
		UpdateContext: resourceNewRelicCloudAwsEuSovereignLinkAccountUpdate,
		DeleteContext: resourceNewRelicCloudAwsEuSovereignLinkAccountDelete,
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the AWS EU Sovereign account in New Relic.",
			},
			"metric_collection_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PULL",
				ForceNew:     true,
				Description:  "How metrics are collected. PULL or PUSH.",
				ValidateFunc: validation.StringInSlice([]string{"PULL", "PUSH"}, false),
			},
			"arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ARN of the IAM role.",
			},
		},
	}
}

func resourceNewRelicCloudAwsEuSovereignLinkAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).NewClient
	accountID := selectAccountID(meta, d)

	createInput := expandAwsEuSovereignLinkAccountInputForCreate(d)

	log.Printf("[INFO] Creating New Relic AWS EU Sovereign link account %s", d.Get("name"))

	cloudLinkedAccount, err := client.Cloud.CloudLinkAccountWithContext(ctx, accountID, createInput)
	if err != nil {
		return diag.FromErr(err)
	}

	var linkedAccountID int
	for _, linkedAccount := range cloudLinkedAccount.LinkedAccounts {
		linkedAccountID = linkedAccount.Id
		break
	}

	d.SetId(strconv.Itoa(linkedAccountID))

	return resourceNewRelicCloudAwsEuSovereignLinkAccountRead(ctx, d, meta)
}

func resourceNewRelicCloudAwsEuSovereignLinkAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).NewClient

	accountID := selectAccountID(meta, d)

	linkedAccountID, convErr := strconv.Atoi(d.Id())
	if convErr != nil {
		return diag.FromErr(convErr)
	}

	log.Printf("[INFO] Reading New Relic AWS EU Sovereign link account %d", linkedAccountID)

	linkedAccount, err := client.Cloud.GetLinkedAccount(accountID, linkedAccountID)
	if err != nil {
		if _, ok := err.(*cloud.NotFoundError); ok {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	return diag.FromErr(flattenAwsEuSovereignLinkAccountForRead(linkedAccount, d, accountID))
}

func resourceNewRelicCloudAwsEuSovereignLinkAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).NewClient
	accountID := selectAccountID(meta, d)

	linkedAccountID, convErr := strconv.Atoi(d.Id())
	if convErr != nil {
		return diag.FromErr(convErr)
	}

	updateInput := expandAwsEuSovereignLinkAccountInputForUpdate(d, linkedAccountID)

	log.Printf("[INFO] Updating New Relic AWS EU Sovereign link account %d", linkedAccountID)

	_, err := client.Cloud.CloudUpdateAccountWithContext(ctx, accountID, updateInput)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNewRelicCloudAwsEuSovereignLinkAccountRead(ctx, d, meta)
}

func resourceNewRelicCloudAwsEuSovereignLinkAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).NewClient

	accountID := selectAccountID(meta, d)

	linkedAccountID := getLinkedAccountIDFromState(d)

	unlinkInput := cloud.CloudUnlinkAccountsInput{
		LinkedAccountIds: []int{linkedAccountID},
	}

	log.Printf("[INFO] Unlinking New Relic AWS EU Sovereign link account %d", linkedAccountID)

	_, err := client.Cloud.CloudUnlinkAccountWithContext(ctx, accountID, unlinkInput)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}