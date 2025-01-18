package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/discovery/pkg/models"
	"github.com/opengovern/og-describer-tailscale/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListDevicePostures(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processDevicePostures(ctx, handler, tailScaleChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-tailScaleChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetDevicePosture(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	integration, err := processDevicePosture(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var status provider.Status
	if integration.Status != nil {
		status = provider.Status{
			LastSync:             integration.Status.LastSync,
			Error:                integration.Status.Error,
			MatchedCount:         integration.Status.MatchedCount,
			PossibleMatchedCount: integration.Status.PossibleMatchedCount,
			ProviderHostCount:    integration.Status.ProviderHostCount,
		}
	}
	value := models.Resource{
		ID:   integration.ID,
		Name: integration.ID,
		Description: provider.PostureIntegrationDescription{
			ID:            integration.ID,
			Provider:      integration.Provider,
			CloudID:       integration.CloudID,
			ClientID:      integration.ClientID,
			TenantID:      integration.TenantID,
			ClientSecret:  integration.ClientSecret,
			ConfigUpdated: integration.ConfigUpdated,
			Status:        &status,
		},
	}
	return &value, nil
}

func processDevicePostures(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var devicePostures provider.ListPostureIntegrations
	baseURL := "/v2/tailnet/-/posture/integrations"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailscale", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &devicePostures); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, integration := range devicePostures.Integrations {
		wg.Add(1)
		go func(integration provider.PostureIntegrationJSON) {
			defer wg.Done()
			var status provider.Status
			if integration.Status != nil {
				status = provider.Status{
					LastSync:             integration.Status.LastSync,
					Error:                integration.Status.Error,
					MatchedCount:         integration.Status.MatchedCount,
					PossibleMatchedCount: integration.Status.PossibleMatchedCount,
					ProviderHostCount:    integration.Status.ProviderHostCount,
				}
			}
			value := models.Resource{
				ID:   integration.ID,
				Name: integration.ID,
				Description: provider.PostureIntegrationDescription{
					ID:            integration.ID,
					Provider:      integration.Provider,
					CloudID:       integration.CloudID,
					ClientID:      integration.ClientID,
					TenantID:      integration.TenantID,
					ClientSecret:  integration.ClientSecret,
					ConfigUpdated: integration.ConfigUpdated,
					Status:        &status,
				},
			}
			tailScaleChan <- value
		}(integration)
	}
	return nil
}

func processDevicePosture(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.PostureIntegrationJSON, error) {
	var devicePosture provider.PostureIntegrationJSON
	baseURL := "/v2/posture/integrations/"

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

	if err = json.Unmarshal(resp.Data, &devicePosture); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &devicePosture, nil
}
