package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/discovery/pkg/models"
	"github.com/opengovern/og-describer-tailscale/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
)

func ListTailnetSettings(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	values, err := processTailnetSetting(ctx, handler, stream)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func processTailnetSetting(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var tailnetSettings provider.TailnetSettingsJSON
	baseURL := "/v2/tailnet/-/settings"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailscale", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &tailnetSettings); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	var values []models.Resource
	jsonBytes, err := json.Marshal(tailnetSettings)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   string(jsonBytes),
		Name: "",
		Description: provider.TailnetSettingsDescription{
			DevicesApprovalOn:                      tailnetSettings.DevicesApprovalOn,
			DevicesAutoUpdatesOn:                   tailnetSettings.DevicesAutoUpdatesOn,
			DevicesKeyDurationDays:                 tailnetSettings.DevicesKeyDurationDays,
			UsersApprovalOn:                        tailnetSettings.UsersApprovalOn,
			UsersRoleAllowedToJoinExternalTailnets: tailnetSettings.UsersRoleAllowedToJoinExternalTailnets,
			NetworkFlowLoggingOn:                   tailnetSettings.NetworkFlowLoggingOn,
			RegionalRoutingOn:                      tailnetSettings.RegionalRoutingOn,
			PostureIdentityCollectionOn:            tailnetSettings.PostureIdentityCollectionOn,
		},
	}
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	} else {
		values = append(values, value)
	}
	return values, nil
}
