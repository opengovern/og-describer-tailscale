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

func ListDeviceInvites(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	devices, err := ListDevices(ctx, handler, stream)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		for _, device := range devices {
			if err := processDeviceInvites(ctx, handler, device.ID, tailScaleChan, &wg); err != nil {
				errorChan <- err // Send error to the error channel
			}
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

func GetDeviceInvite(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	deviceInvite, err := processDeviceInvite(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var acceptedBy provider.AcceptedBy
	if deviceInvite.AcceptedBy != nil {
		acceptedBy = provider.AcceptedBy{
			ID:            deviceInvite.AcceptedBy.ID,
			LoginName:     deviceInvite.AcceptedBy.LoginName,
			ProfilePicURL: deviceInvite.AcceptedBy.ProfilePicURL,
		}
	}
	value := models.Resource{
		ID:   deviceInvite.ID,
		Name: deviceInvite.InviteURL,
		Description: provider.DeviceInviteDescription{
			ID:              deviceInvite.ID,
			Created:         deviceInvite.Created,
			TailnetID:       deviceInvite.TailnetID,
			DeviceID:        deviceInvite.DeviceID,
			SharerID:        deviceInvite.SharerID,
			MultiUse:        deviceInvite.MultiUse,
			AllowExitNode:   deviceInvite.AllowExitNode,
			Email:           deviceInvite.Email,
			LastEmailSentAt: deviceInvite.LastEmailSentAt,
			InviteURL:       deviceInvite.InviteURL,
			Accepted:        deviceInvite.Accepted,
			AcceptedBy:      &acceptedBy,
		},
	}
	return &value, nil
}

func processDeviceInvites(ctx context.Context, handler *resilientbridge.ResilientBridge, deviceID string, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var deviceInvites []provider.DeviceInviteJSON
	baseURL := "/v2/device/"

	finalURL := fmt.Sprintf("%s%s/device-invites", baseURL, deviceID)

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

	if err = json.Unmarshal(resp.Data, &deviceInvites); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, deviceInvite := range deviceInvites {
		wg.Add(1)
		go func(deviceInvite provider.DeviceInviteJSON) {
			defer wg.Done()
			var acceptedBy provider.AcceptedBy
			if deviceInvite.AcceptedBy != nil {
				acceptedBy = provider.AcceptedBy{
					ID:            deviceInvite.AcceptedBy.ID,
					LoginName:     deviceInvite.AcceptedBy.LoginName,
					ProfilePicURL: deviceInvite.AcceptedBy.ProfilePicURL,
				}
			}
			value := models.Resource{
				ID:   deviceInvite.ID,
				Name: deviceInvite.InviteURL,
				Description: provider.DeviceInviteDescription{
					ID:              deviceInvite.ID,
					Created:         deviceInvite.Created,
					TailnetID:       deviceInvite.TailnetID,
					DeviceID:        deviceInvite.DeviceID,
					SharerID:        deviceInvite.SharerID,
					MultiUse:        deviceInvite.MultiUse,
					AllowExitNode:   deviceInvite.AllowExitNode,
					Email:           deviceInvite.Email,
					LastEmailSentAt: deviceInvite.LastEmailSentAt,
					InviteURL:       deviceInvite.InviteURL,
					Accepted:        deviceInvite.Accepted,
					AcceptedBy:      &acceptedBy,
				},
			}
			tailScaleChan <- value
		}(deviceInvite)
	}
	return nil
}

func processDeviceInvite(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.DeviceInviteJSON, error) {
	var deviceInvite provider.DeviceInviteJSON
	baseURL := "/v2/device-invites/"

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

	if err = json.Unmarshal(resp.Data, &deviceInvite); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &deviceInvite, nil
}
