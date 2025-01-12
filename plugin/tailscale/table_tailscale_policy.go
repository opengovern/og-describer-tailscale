package tailscale

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScalePolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_policy",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    nil,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "acls", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.ACLs"), Description: "The list of ACLs (Access Control Lists) associated with the policy."},
			{Name: "groups", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Groups"), Description: "The mapping of groups to the corresponding list of users."},
			{Name: "hosts", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Hosts"), Description: "The mapping of hosts to their corresponding IPs."},
		}),
	}
}
