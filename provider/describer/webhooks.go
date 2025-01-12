package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/pkg/sdk/models"
	"github.com/opengovern/og-describer-tailscale/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListWebhooks(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processWebhooks(ctx, handler, tailScaleChan, &wg); err != nil {
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

func GetWebhook(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	webhook, err := processWebhook(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   webhook.EndpointID,
		Name: webhook.EndpointURL,
		Description: model.WebhookDescription{
			EndpointID:       webhook.EndpointID,
			EndpointURL:      webhook.EndpointURL,
			ProviderType:     webhook.ProviderType,
			Created:          webhook.Created,
			CreatorLoginName: webhook.CreatorLoginName,
			LastModified:     webhook.LastModified,
			Secret:           webhook.Secret,
			Subscriptions:    webhook.Subscriptions,
		},
	}
	return &value, nil
}

func processWebhooks(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var webhookListResponse model.ListWebhookResponse
	baseURL := "https://api.tailscale.com/api/v2/tailnet/-/webhooks"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailsclae", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &webhookListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, webhook := range webhookListResponse.Webhooks {
		wg.Add(1)
		go func(webhook model.WebhookJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   webhook.EndpointID,
				Name: webhook.EndpointURL,
				Description: model.WebhookDescription{
					EndpointID:       webhook.EndpointID,
					EndpointURL:      webhook.EndpointURL,
					ProviderType:     webhook.ProviderType,
					Created:          webhook.Created,
					CreatorLoginName: webhook.CreatorLoginName,
					LastModified:     webhook.LastModified,
					Secret:           webhook.Secret,
					Subscriptions:    webhook.Subscriptions,
				},
			}
			tailScaleChan <- value
		}(webhook)
	}
	return nil
}

func processWebhook(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.WebhookJSON, error) {
	var webhook model.WebhookJSON
	baseURL := "/v2/webhooks/"

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

	if err = json.Unmarshal(resp.Data, &webhook); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &webhook, nil
}
