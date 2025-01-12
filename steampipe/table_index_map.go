package steampipe

import (
	"github.com/opengovern/og-describer-tailscale/pkg/sdk/es"
)

var Map = map[string]string{
  "TailScale/Device": "tailscale_device",
  "TailScale/User": "tailscale_user",
  "TailScale/Contact": "tailscale_contact",
  "TailScale/Device/Invite": "tailscale_device_invite",
  "TailScale/Device/Posture": "tailscale_device_posture",
  "TailScale/User/Invite": "tailscale_user_invite",
  "TailScale/Key": "tailscale_key",
  "TailScale/Policy": "tailscale_policy",
  "TailScale/TailnetSetting": "tailscale_tailnet_setting",
  "TailScale/Webhook": "tailscale_webhook",
  "TailScale/DNS": "tailscale_dns",
}

var DescriptionMap = map[string]interface{}{
  "TailScale/Device": opengovernance.Device{},
  "TailScale/User": opengovernance.User{},
  "TailScale/Contact": opengovernance.Contact{},
  "TailScale/Device/Invite": opengovernance.DeviceInvite{},
  "TailScale/Device/Posture": opengovernance.PostureIntegration{},
  "TailScale/User/Invite": opengovernance.UserInvite{},
  "TailScale/Key": opengovernance.Key{},
  "TailScale/Policy": opengovernance.Policy{},
  "TailScale/TailnetSetting": opengovernance.TailnetSettings{},
  "TailScale/Webhook": opengovernance.Webhook{},
  "TailScale/DNS": opengovernance.DNS{},
}

var ReverseMap = map[string]string{
  "tailscale_device": "TailScale/Device",
  "tailscale_user": "TailScale/User",
  "tailscale_contact": "TailScale/Contact",
  "tailscale_device_invite": "TailScale/Device/Invite",
  "tailscale_device_posture": "TailScale/Device/Posture",
  "tailscale_user_invite": "TailScale/User/Invite",
  "tailscale_key": "TailScale/Key",
  "tailscale_policy": "TailScale/Policy",
  "tailscale_tailnet_setting": "TailScale/TailnetSetting",
  "tailscale_webhook": "TailScale/Webhook",
  "tailscale_dns": "TailScale/DNS",
}
