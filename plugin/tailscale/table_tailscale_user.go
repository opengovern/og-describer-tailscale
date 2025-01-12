package tailscale

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_user",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    nil,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the user."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.DisplayName"), Description: "The display name of the user."},
			{Name: "login_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LoginName"), Description: "The login name of the user."},
			//{Name: "profile_pic_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ProfilePicURL"), Description: "The profile picture URL of the user."},
			{Name: "tailnet_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.TailnetID"), Description: "The tailnet ID associated with the user."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Created"), Description: "The time the user was created."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "The type of user, such as member or shared."},
			{Name: "role", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Role"), Description: "The role of the user, such as owner, admin, or member."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Status"), Description: "The status of the user, such as active, idle, or suspended."},
			{Name: "device_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.DeviceCount"), Description: "The number of devices owned by the user."},
			{Name: "last_seen", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.LastSeen"), Description: "The last time the user was seen."},
			{Name: "currently_connected", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.CurrentlyConnected"), Description: "Indicates if the user is currently connected."},
		}),
	}
}
