package steampipe

import (
	"github.com/opengovern/og-describer-tailscale/pkg/sdk/es"
)

var Map = map[string]string{
  "TailScale/Device": "tailscale_device",
  "TailScale/User": "tailscale_user",
}

var DescriptionMap = map[string]interface{}{
  "TailScale/Device": opengovernance.Device{},
  "TailScale/User": opengovernance.User{},
}

var ReverseMap = map[string]string{
  "tailscale_device": "TailScale/Device",
  "tailscale_user": "TailScale/User",
}
