package maps

import (
	"github.com/opengovern/og-describer-tailscale/discovery/describers"
	model "github.com/opengovern/og-describer-tailscale/discovery/pkg/models"
	"github.com/opengovern/og-describer-tailscale/discovery/provider"
	"github.com/opengovern/og-describer-tailscale/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"TailScale/Device": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Device",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListDevices),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetDevice),
	},

	"TailScale/User": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListUsers),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetUser),
	},

	"TailScale/Contact": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Contact",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListContacts),
		GetDescriber:    nil,
	},

	"TailScale/Device/Invite": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Device/Invite",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListDeviceInvites),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetDeviceInvite),
	},

	"TailScale/Device/Posture": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Device/Posture",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListDevicePostures),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetDevicePosture),
	},

	"TailScale/User/Invite": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/User/Invite",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListUserInvites),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetUserInvite),
	},

	"TailScale/Key": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Key",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListKeys),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetKey),
	},

	"TailScale/Policy": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Policy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   nil,
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetPolicyFile),
	},

	"TailScale/TailnetSetting": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/TailnetSetting",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListTailnetSettings),
		GetDescriber:    nil,
	},

	"TailScale/Webhook": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/Webhook",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListWebhooks),
		GetDescriber:    provider.DescribeSingleByTailScale(describers.GetWebhook),
	},

	"TailScale/DNS": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "TailScale/DNS",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByTailScale(describers.ListDNSs),
		GetDescriber:    nil,
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"TailScale/Device": {
		Name:            "TailScale/Device",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/User": {
		Name:            "TailScale/User",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/Contact": {
		Name:            "TailScale/Contact",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/Device/Invite": {
		Name:            "TailScale/Device/Invite",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/Device/Posture": {
		Name:            "TailScale/Device/Posture",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/User/Invite": {
		Name:            "TailScale/User/Invite",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/Key": {
		Name:            "TailScale/Key",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/Policy": {
		Name:            "TailScale/Policy",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/TailnetSetting": {
		Name:            "TailScale/TailnetSetting",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/Webhook": {
		Name:            "TailScale/Webhook",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"TailScale/DNS": {
		Name:            "TailScale/DNS",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"TailScale/Device",
	"TailScale/User",
	"TailScale/Contact",
	"TailScale/Device/Invite",
	"TailScale/Device/Posture",
	"TailScale/User/Invite",
	"TailScale/Key",
	"TailScale/Policy",
	"TailScale/TailnetSetting",
	"TailScale/Webhook",
	"TailScale/DNS",
}
