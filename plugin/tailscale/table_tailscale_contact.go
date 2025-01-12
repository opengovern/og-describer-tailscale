package tailscale

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-tailscale/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleContact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_contact",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListContact,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("account_email"),
			Hydrate:    opengovernance.GetContact,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "account", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Account"), Description: "The email address of the account contact."},
			{Name: "support", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Support"), Description: "The email address of the support contact."},
			{Name: "security", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Security"), Description: "The email address of the security contact."},
		}),
	}
}
