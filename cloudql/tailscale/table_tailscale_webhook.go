package tailscale

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-tailscale/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleWebhook(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_webhook",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWebhook,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("endpoint_id"),
			Hydrate:    opengovernance.GetWebhook,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "endpoint_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.EndpointID"), Description: "The unique identifier for the webhook endpoint."},
			{Name: "endpoint_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.EndpointURL"), Description: "The URL where events are sent to from Tailscale via POST requests."},
			{Name: "provider_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ProviderType"), Description: "The provider type for the webhook destination (e.g., slack, discord, etc.)."},
			{Name: "creator_login_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatorLoginName"), Description: "The login name of the user who created the webhook."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Created"), Description: "The time when the webhook endpoint was created."},
			{Name: "last_modified", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.LastModified"), Description: "The time when the webhook endpoint was last modified."},
			{Name: "subscriptions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Subscriptions"), Description: "The list of events that trigger POST requests to the webhook endpoint."},
			{Name: "secret", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Secret"), Description: "The secret associated with the webhook endpoint for signature verification."},
		}),
	}
}
