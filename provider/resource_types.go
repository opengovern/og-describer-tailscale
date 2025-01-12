package provider

import (
	model "github.com/opengovern/og-describer-tailscale/pkg/sdk/models"
	"github.com/opengovern/og-describer-tailscale/provider/configs"
	"github.com/opengovern/og-describer-tailscale/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"TailScale/Device": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Device",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListDevices),
		GetDescriber:    DescribeSingleByTailScale(describer.GetDevice),
	},

	"TailScale/User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListUsers),
		GetDescriber:    DescribeSingleByTailScale(describer.GetUser),
	},

	"TailScale/Contact": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Contact",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListContacts),
		GetDescriber:    nil,
	},

	"TailScale/Device/Invite": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Device/Invite",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListDeviceInvites),
		GetDescriber:    DescribeSingleByTailScale(describer.GetDeviceInvite),
	},

	"TailScale/Device/Posture": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Device/Posture",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListDevicePostures),
		GetDescriber:    DescribeSingleByTailScale(describer.GetDevicePosture),
	},

	"TailScale/User/Invite": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/User/Invite",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListUserInvites),
		GetDescriber:    DescribeSingleByTailScale(describer.GetUserInvite),
	},

	"TailScale/Key": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Key",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListKeys),
		GetDescriber:    DescribeSingleByTailScale(describer.GetKey),
	},

	"TailScale/Policy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Policy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   nil,
		GetDescriber:    DescribeSingleByTailScale(describer.GetPolicyFile),
	},

	"TailScale/TailnetSetting": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/TailnetSetting",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListTailnetSettings),
		GetDescriber:    nil,
	},

	"TailScale/Webhook": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "TailScale/Webhook",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByTailScale(describer.ListWebhooks),
		GetDescriber:    DescribeSingleByTailScale(describer.GetWebhook),
	},
}
