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

func ListUserInvites(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processUserInvites(ctx, handler, tailScaleChan, &wg); err != nil {
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

func GetUserInvite(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	userInvite, err := processUserInvite(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   userInvite.ID,
		Name: userInvite.InviteURL,
		Description: model.UserInviteDescription{
			ID:              userInvite.ID,
			Role:            userInvite.Role,
			TailnetID:       userInvite.TailnetID,
			InviterID:       userInvite.InviterID,
			Email:           userInvite.Email,
			LastEmailSentAt: userInvite.LastEmailSentAt,
			InviteURL:       userInvite.InviteURL,
		},
	}
	return &value, nil
}

func processUserInvites(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var userInvites []model.UserInviteJSON
	baseURL := "/v2/tailnet/-/user-invites"

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

	if err = json.Unmarshal(resp.Data, &userInvites); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, userInvite := range userInvites {
		wg.Add(1)
		go func(userInvite model.UserInviteJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   userInvite.ID,
				Name: userInvite.InviteURL,
				Description: model.UserInviteDescription{
					ID:              userInvite.ID,
					Role:            userInvite.Role,
					TailnetID:       userInvite.TailnetID,
					InviterID:       userInvite.InviterID,
					Email:           userInvite.Email,
					LastEmailSentAt: userInvite.LastEmailSentAt,
					InviteURL:       userInvite.InviteURL,
				},
			}
			tailScaleChan <- value
		}(userInvite)
	}
	return nil
}

func processUserInvite(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.UserInviteJSON, error) {
	var userInvite model.UserInviteJSON
	baseURL := "/v2/user-invites/"

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

	if err = json.Unmarshal(resp.Data, &userInvite); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &userInvite, nil
}
