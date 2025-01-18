package provider

import (
	"errors"
	//"github.com/opengovern/og-describer-tailscale/discovery/describers"
	"github.com/opengovern/og-describer-tailscale/discovery/pkg/models"
	model "github.com/opengovern/og-describer-tailscale/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"time"
)

// DescribeListByTailScale A wrapper to pass TailScale authorization to describers functions
func DescribeListByTailScale(describe func(context.Context, *resilientbridge.ResilientBridge, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register TailScale provider
		resilientBridge.RegisterProvider("tailscale", &adapters.TailScaleAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		// Get values from describers
		var values []model.Resource
		result, err := describe(ctx, resilientBridge, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByTailScale A wrapper to pass TailScale authorization to describers functions
func DescribeSingleByTailScale(describe func(context.Context, *resilientbridge.ResilientBridge, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *model.StreamSender) (*model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register TailScale provider
		resilientBridge.RegisterProvider("tailscale", &adapters.TailScaleAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		// Get value from describers
		value, err := describe(ctx, resilientBridge, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
