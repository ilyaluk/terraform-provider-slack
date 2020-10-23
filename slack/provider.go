package slack

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/slack-go/slack"
)

// Provider returns a *schema.Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_TOKEN", nil),
				Description: "The Slack token",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"slack_conversation": resourceSlackConversation(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"slack_conversation": dataSourceConversation(),
			"slack_user":         dataSourceUser(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	token, ok := d.GetOk("token")
	if !ok {
		return nil, diag.Errorf("could not create slack client. Please provide a token.")
	}
	slackClient := slack.New(token.(string))
	return slackClient, diags
}

func updateChannelData(d *schema.ResourceData, channel *slack.Channel) diag.Diagnostics {
	if channel.ID == "" {
		return diag.Errorf("error setting id: returned channel does not have an ID")
	}
	d.SetId(channel.ID)

	if err := d.Set("name", channel.Name); err != nil {
		return diag.Errorf("error setting name: %s", err)
	}

	if err := d.Set("topic", channel.Topic.Value); err != nil {
		return diag.Errorf("error setting topic: %s", err)
	}

	if err := d.Set("purpose", channel.Purpose.Value); err != nil {
		return diag.Errorf("error setting purpose: %s", err)
	}

	if err := d.Set("is_archived", channel.IsArchived); err != nil {
		return diag.Errorf("error setting is_archived: %s", err)
	}

	if err := d.Set("is_shared", channel.IsShared); err != nil {
		return diag.Errorf("error setting is_shared: %s", err)
	}

	if err := d.Set("is_ext_shared", channel.IsExtShared); err != nil {
		return diag.Errorf("error setting is_ext_shared: %s", err)
	}

	if err := d.Set("is_org_shared", channel.IsOrgShared); err != nil {
		return diag.Errorf("error setting is_org_shared: %s", err)
	}

	if err := d.Set("created", channel.Created); err != nil {
		return diag.Errorf("error setting created: %s", err)
	}

	if err := d.Set("creator", channel.Creator); err != nil {
		return diag.Errorf("error setting creator: %s", err)
	}

	if err := d.Set("is_private", channel.IsPrivate); err != nil {
		return diag.Errorf("error setting is_private: %s", err)
	}

	if err := d.Set("is_general", channel.IsGeneral); err != nil {
		return diag.Errorf("error setting is_general: %s", err)
	}

	if err := d.Set("members", channel.Members); err != nil {
		return diag.Errorf("error setting members: %s", err)
	}

	return nil
}
