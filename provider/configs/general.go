package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "tailscale"                                    // example: aws, azure
	IntegrationName      = integration.Type("tailscale_account")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-tailscale" // example: github.com/opengovern/og-describer-aws
)
