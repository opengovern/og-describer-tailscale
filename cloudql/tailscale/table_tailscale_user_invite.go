package tailscale

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-tailscale/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleUserInvite(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_user_invite",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUserInvite,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetUserInvite,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the user invite."},
			{Name: "role", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Role"), Description: "The role associated with the user invite."},
			{Name: "tailnet_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.TailnetID"), Description: "The tailnet ID associated with the user invite."},
			{Name: "inviter_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.InviterID"), Description: "The ID of the user who invited the new user."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Email"), Description: "The email address of the user being invited."},
			{Name: "last_email_sent_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.LastEmailSentAt"), Description: "The time when the last email was sent for this invite."},
			{Name: "invite_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.InviteURL"), Description: "The URL for the user invite."},
		}),
	}
}
