package tailscale

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_key",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    nil,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the key."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Created"), Description: "The time when the key was created."},
			{Name: "expires", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Expires"), Description: "The time when the key expires."},
			{Name: "revoked", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.Revoked"), Description: "The time when the key was revoked."},
			{Name: "capabilities", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Capabilities"), Description: "The capabilities associated with the key."},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Description"), Description: "The description of the key."},
			{Name: "invalid", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Invalid"), Description: "Indicates whether the key is invalid."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.UserID"), Description: "The user ID associated with the key."},
		}),
	}
}
