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

func ListUsers(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processUsers(ctx, handler, tailScaleChan, &wg); err != nil {
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

func GetUser(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	user, err := processUser(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   user.ID,
		Name: user.LoginName,
		Description: provider.UserDescription{
			ID:                 user.ID,
			DisplayName:        user.DisplayName,
			LoginName:          user.LoginName,
			TailnetID:          user.TailnetID,
			Created:            user.Created,
			Type:               user.Type,
			Role:               user.Role,
			Status:             user.Status,
			DeviceCount:        user.DeviceCount,
			LastSeen:           user.LastSeen,
			CurrentlyConnected: user.CurrentlyConnected,
		},
	}
	return &value, nil
}

func processUsers(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var userListResponse provider.UserListResponse
	baseURL := "/v2/tailnet/-/users"

	params := url.Values{}
	params.Set("type", "all")
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

	if err = json.Unmarshal(resp.Data, &userListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, user := range userListResponse.Users {
		wg.Add(1)
		go func(user provider.UserJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   user.ID,
				Name: user.LoginName,
				Description: provider.UserDescription{
					ID:                 user.ID,
					DisplayName:        user.DisplayName,
					LoginName:          user.LoginName,
					TailnetID:          user.TailnetID,
					Created:            user.Created,
					Type:               user.Type,
					Role:               user.Role,
					Status:             user.Status,
					DeviceCount:        user.DeviceCount,
					LastSeen:           user.LastSeen,
					CurrentlyConnected: user.CurrentlyConnected,
				},
			}
			tailScaleChan <- value
		}(user)
	}
	return nil
}

func processUser(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.UserJSON, error) {
	var user provider.UserJSON
	baseURL := "/v2/user/"

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

	if err = json.Unmarshal(resp.Data, &user); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &user, nil
}
