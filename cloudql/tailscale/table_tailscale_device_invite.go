package tailscale

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-tailscale/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleDeviceInvite(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_device_invite",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDeviceInvite,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetDeviceInvite,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the device invite."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Created"), Description: "The time when the device invite was created."},
			{Name: "tailnet_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.TailnetID"), Description: "The tailnet ID associated with the device invite."},
			{Name: "device_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.DeviceID"), Description: "The device ID associated with the invite."},
			{Name: "sharer_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.SharerID"), Description: "The ID of the user who shared the device invite."},
			{Name: "multi_use", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.MultiUse"), Description: "Whether the invite is for multiple uses."},
			{Name: "allow_exit_node", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.AllowExitNode"), Description: "Indicates whether the invite allows the device to be an exit node."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Email"), Description: "The email address of the invite recipient."},
			{Name: "last_email_sent_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.LastEmailSentAt"), Description: "The time when the last email was sent for this invite."},
			{Name: "invite_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.InviteURL"), Description: "The URL for the device invite."},
			{Name: "accepted", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Accepted"), Description: "Whether the invite has been accepted."},
			{Name: "accepted_by", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.AcceptedBy"), Description: "Details of the user who accepted the invite, if applicable."},
		}),
	}
}
