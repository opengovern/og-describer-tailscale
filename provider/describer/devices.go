package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/pkg/sdk/models"
	"github.com/opengovern/og-describer-tailscale/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"net/url"
	"sync"
)

func ListDevices(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processDevices(ctx, handler, tailScaleChan, &wg); err != nil {
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

func GetDevice(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	device, err := processDevice(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	latencyMap := make(map[string]model.Latency)
	for key, value := range device.ClientConnectivity.Latency {
		latency := model.Latency{
			Preferred: value.Preferred,
			LatencyMs: value.LatencyMs,
		}
		latencyMap[key] = latency
	}
	clientSupports := model.ClientCapabilities{
		HairPinning: device.ClientConnectivity.ClientSupports.HairPinning,
		IPv6:        device.ClientConnectivity.ClientSupports.IPv6,
		PCP:         device.ClientConnectivity.ClientSupports.PCP,
		PMP:         device.ClientConnectivity.ClientSupports.PMP,
		UDP:         device.ClientConnectivity.ClientSupports.UDP,
		UPnP:        device.ClientConnectivity.ClientSupports.UPnP,
	}
	clientConnectivity := model.ClientConnectivity{
		Endpoints:             device.ClientConnectivity.Endpoints,
		MappingVariesByDestIP: device.ClientConnectivity.MappingVariesByDestIP,
		Latency:               latencyMap,
		ClientSupports:        &clientSupports,
	}
	postureIdentity := model.PostureIdentity{
		SerialNumbers: device.PostureIdentity.SerialNumbers,
		Disabled:      device.PostureIdentity.Disabled,
	}
	value := models.Resource{
		ID:   device.ID,
		Name: device.Name,
		Description: model.DeviceDescription{
			Addresses:          device.Addresses,
			ID:                 device.ID,
			NodeID:             device.NodeID,
			User:               device.User,
			Name:               device.Name,
			ClientVersion:      device.ClientVersion,
			OS:                 device.OS,
			LastSeen:           device.LastSeen,
			ClientConnectivity: &clientConnectivity,
			Tags:               device.Tags,
			PostureIdentity:    &postureIdentity,
		},
	}
	return &value, nil
}

func processDevices(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var deviceListResponse model.DeviceListResponse
	baseURL := "/v2/tailnet/-/devices"

	params := url.Values{}
	params.Set("fields", "all")
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailsclae", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &deviceListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, device := range deviceListResponse.Devices {
		wg.Add(1)
		go func(device model.DeviceJSON) {
			defer wg.Done()
			latencyMap := make(map[string]model.Latency)
			for key, value := range device.ClientConnectivity.Latency {
				latency := model.Latency{
					Preferred: value.Preferred,
					LatencyMs: value.LatencyMs,
				}
				latencyMap[key] = latency
			}
			clientSupports := model.ClientCapabilities{
				HairPinning: device.ClientConnectivity.ClientSupports.HairPinning,
				IPv6:        device.ClientConnectivity.ClientSupports.IPv6,
				PCP:         device.ClientConnectivity.ClientSupports.PCP,
				PMP:         device.ClientConnectivity.ClientSupports.PMP,
				UDP:         device.ClientConnectivity.ClientSupports.UDP,
				UPnP:        device.ClientConnectivity.ClientSupports.UPnP,
			}
			clientConnectivity := model.ClientConnectivity{
				Endpoints:             device.ClientConnectivity.Endpoints,
				MappingVariesByDestIP: device.ClientConnectivity.MappingVariesByDestIP,
				Latency:               latencyMap,
				ClientSupports:        &clientSupports,
			}
			postureIdentity := model.PostureIdentity{
				SerialNumbers: device.PostureIdentity.SerialNumbers,
				Disabled:      device.PostureIdentity.Disabled,
			}
			value := models.Resource{
				ID:   device.ID,
				Name: device.Name,
				Description: model.DeviceDescription{
					Addresses:          device.Addresses,
					ID:                 device.ID,
					NodeID:             device.NodeID,
					User:               device.User,
					Name:               device.Name,
					ClientVersion:      device.ClientVersion,
					OS:                 device.OS,
					LastSeen:           device.LastSeen,
					ClientConnectivity: &clientConnectivity,
					Tags:               device.Tags,
					PostureIdentity:    &postureIdentity,
				},
			}
			tailScaleChan <- value
		}(device)
	}
	return nil
}

func processDevice(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.DeviceJSON, error) {
	var device model.DeviceJSON
	baseURL := "/v2/device/"

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

	if err = json.Unmarshal(resp.Data, &device); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &device, nil
}
