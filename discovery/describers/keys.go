package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/discovery/pkg/models"
	"github.com/opengovern/og-describer-tailscale/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"net/url"
	"sync"
)

func ListKeys(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processKeys(ctx, handler, tailScaleChan, &wg); err != nil {
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

func GetKey(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	key, err := processKey(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	capabilities := provider.Capabilities{
		Devices: provider.DeviceCapabilities{
			Create: provider.CreatePermissions{
				Reusable:      key.Capabilities.Devices.Create.Reusable,
				Ephemeral:     key.Capabilities.Devices.Create.Ephemeral,
				Preauthorized: key.Capabilities.Devices.Create.Preauthorized,
				Tags:          key.Capabilities.Devices.Create.Tags,
			},
		},
	}
	value := models.Resource{
		ID:   key.ID,
		Name: key.ID,
		Description: provider.KeyDescription{
			ID:           key.ID,
			Created:      key.Created,
			Expires:      key.Expires,
			Revoked:      key.Revoked,
			Capabilities: capabilities,
			Description:  key.Description,
			Invalid:      key.Invalid,
			UserID:       key.UserID,
		},
	}
	return &value, nil
}

func processKeys(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var keyListResponse provider.KeyListResponse
	baseURL := "/v2/tailnet/-/keys"

	params := url.Values{}
	params.Set("all", "true")
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailscale", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &keyListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, key := range keyListResponse.Keys {
		wg.Add(1)
		go func(key provider.KeyJSON) {
			defer wg.Done()
			capabilities := provider.Capabilities{
				Devices: provider.DeviceCapabilities{
					Create: provider.CreatePermissions{
						Reusable:      key.Capabilities.Devices.Create.Reusable,
						Ephemeral:     key.Capabilities.Devices.Create.Ephemeral,
						Preauthorized: key.Capabilities.Devices.Create.Preauthorized,
						Tags:          key.Capabilities.Devices.Create.Tags,
					},
				},
			}
			value := models.Resource{
				ID:   key.ID,
				Name: key.ID,
				Description: provider.KeyDescription{
					ID:           key.ID,
					Created:      key.Created,
					Expires:      key.Expires,
					Revoked:      key.Revoked,
					Capabilities: capabilities,
					Description:  key.Description,
					Invalid:      key.Invalid,
					UserID:       key.UserID,
				},
			}
			tailScaleChan <- value
		}(key)
	}
	return nil
}

func processKey(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.KeyJSON, error) {
	var key provider.KeyJSON
	baseURL := "/v2/tailnet/-/keys/"

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

	if err = json.Unmarshal(resp.Data, &key); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &key, nil
}
