package tailscale

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-tailscale/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleDevicePosture(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_device_posture",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPostureIntegration,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetPostureIntegration,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "provider", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Provider"), Description: "The provider of the posture integration."},
			{Name: "cloud_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CloudID"), Description: "The cloud ID associated with the posture integration."},
			{Name: "client_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ClientID"), Description: "The client ID associated with the posture integration."},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.TenantID"), Description: "The tenant ID associated with the posture integration."},
			{Name: "client_secret", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ClientSecret"), Description: "The client secret associated with the posture integration."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the posture integration."},
			{Name: "config_updated", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ConfigUpdated"), Description: "The time when the posture integration configuration was last updated."},
			{Name: "status", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Status"), Description: "The current status of the posture integration."},
		}),
	}
}
