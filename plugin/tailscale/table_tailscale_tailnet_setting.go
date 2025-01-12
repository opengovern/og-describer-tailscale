package tailscale

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleTailnetSetting(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_tailnet_setting",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    nil,
		},
		Columns: []*plugin.Column{
			{Name: "devices_approval_on", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.DevicesApprovalOn"), Description: "Whether device approval is enabled for the tailnet."},
			{Name: "devices_auto_updates_on", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.DevicesAutoUpdatesOn"), Description: "Whether auto updates are enabled for devices in the tailnet."},
			{Name: "devices_key_duration_days", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.DevicesKeyDurationDays"), Description: "The key expiry duration for devices in the tailnet (in days)."},
			{Name: "users_approval_on", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.UsersApprovalOn"), Description: "Whether user approval is enabled for the tailnet."},
			{Name: "users_role_allowed_to_join_external_tailnets", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.UsersRoleAllowedToJoinExternalTailnets"), Description: "The user role allowed to join external tailnets."},
			{Name: "network_flow_logging_on", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.NetworkFlowLoggingOn"), Description: "Whether network flow logging is enabled for the tailnet."},
			{Name: "regional_routing_on", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.RegionalRoutingOn"), Description: "Whether regional routing is enabled for the tailnet."},
			{Name: "posture_identity_collection_on", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.PostureIdentityCollectionOn"), Description: "Whether identity collection for device posture integrations is enabled for the tailnet."},
		},
	}
}
