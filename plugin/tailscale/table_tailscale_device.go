package tailscale

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableTailScaleDevice(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "tailscale_device",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    nil,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "addresses", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Addresses"), Description: ""},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier of the device."},
			{Name: "node_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.NodeID"), Description: "The Node ID of the device."},
			{Name: "user", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.User"), Description: "The user associated with the device."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the device."},
			//{Name: "hostname", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Hostname"), Description: "The hostname of the device."},
			{Name: "client_version", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ClientVersion"), Description: "The client version of the device."},
			//{Name: "update_available", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.UpdateAvailable"), Description: "Indicates if an update is available for the device."},
			{Name: "os", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.OS"), Description: "The operating system of the device."},
			//{Name: "created", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Created").Transform(transform.ToString), Description: "The creation timestamp of the device."},
			{Name: "last_seen", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LastSeen").Transform(transform.ToString), Description: "The last seen timestamp of the device."},
			//{Name: "key_expiry_disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.KeyExpiryDisabled"), Description: "Indicates if the key expiry is disabled."},
			//{Name: "expires", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Expires").Transform(transform.ToString), Description: "The expiration timestamp of the device."},
			//{Name: "authorized", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Authorized"), Description: "Indicates if the device is authorized."},
			//{Name: "is_external", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.IsExternal"), Description: "Indicates if the device is external."},
			//{Name: "machine_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.MachineKey"), Description: "The machine key of the device."},
			//{Name: "node_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.NodeKey"), Description: "The node key of the device."},
			//{Name: "blocks_incoming_connections", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.BlocksIncomingConnections"), Description: "Indicates if incoming connections are blocked."},
			//{Name: "enabled_routes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.EnabledRoutes"), Description: "The routes enabled on the device."},
			//{Name: "advertised_routes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.AdvertisedRoutes"), Description: "The routes advertised by the device."},
			{Name: "client_connectivity", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.ClientConnectivity"), Description: ""},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Tags"), Description: "The tags associated with the device."},
			//{Name: "tailnet_lock_error", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.TailnetLockError"), Description: "The error related to tailnet lock, if any."},
			//{Name: "tailnet_lock_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.TailnetLockKey"), Description: "The tailnet lock key for the device."},
			{Name: "posture_identity", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.PostureIdentity"), Description: ""},
		}),
	}
}
