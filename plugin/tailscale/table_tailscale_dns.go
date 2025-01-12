package tailscale

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-tailscale/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleDNS(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_dns",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDNS,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("dns"),
			Hydrate:    opengovernance.GetDNS,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "dns", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.DNS"), Description: ""},
		}),
	}
}
