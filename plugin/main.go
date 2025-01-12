package main

import (
	"github.com/opengovern/og-describer-tailscale/plugin/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: tailscale.Plugin})
}
