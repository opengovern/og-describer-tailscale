package tailscale

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-tailscale",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"tailscale_device":          tableTailScaleDevice(ctx),
			"tailscale_user":            tableTailScaleUser(ctx),
			"tailscale_contact":         tableTailScaleContact(ctx),
			"tailscale_device_invite":   tableTailScaleDeviceInvite(ctx),
			"tailscale_device_posture":  tableTailScaleDevicePosture(ctx),
			"tailscale_user_invite":     tableTailScaleUserInvite(ctx),
			"tailscale_key":             tableTailScaleKey(ctx),
			"tailscale_policy":          tableTailScalePolicy(ctx),
			"tailscale_tailnet_setting": tableTailScaleTailnetSetting(ctx),
			"tailscale_webhook":         tableTailScaleWebhook(ctx),
			"tailscale_dns":             tableTailScaleDNS(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_account_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
