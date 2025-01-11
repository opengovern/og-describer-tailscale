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
}
