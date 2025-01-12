package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/pkg/sdk/models"
	"github.com/opengovern/og-describer-tailscale/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
)

func GetPolicyFile(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	policy, err := processPolicyFile(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var ACLs []model.ACL
	for _, ACL := range policy.ACLs {
		ACLs = append(ACLs, model.ACL{
			Action: ACL.Action,
			Ports:  ACL.Ports,
			Users:  ACL.Users,
		})
	}
	jsonBytes, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   string(jsonBytes),
		Name: "",
		Description: model.PolicyDescription{
			ACLs:   ACLs,
			Groups: policy.Groups,
			Hosts:  policy.Hosts,
		},
	}
	return &value, nil
}

func processPolicyFile(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.PolicyJSON, error) {
	var policy model.PolicyJSON
	baseURL := "/v2/tailnet/-/acl"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailscale", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &policy); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &policy, nil
}
