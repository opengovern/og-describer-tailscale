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

func ListDNSs(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	tailScaleChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(tailScaleChan)
		defer close(errorChan)
		if err := processDNSs(ctx, handler, tailScaleChan, &wg); err != nil {
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

func processDNSs(ctx context.Context, handler *resilientbridge.ResilientBridge, tailScaleChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var nameServers provider.ListDNSNameServerResponse
	var searchPaths provider.ListDNSSearchPathsResponse
	baseURL := "/v2/tailnet/-/dns/nameservers"

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

	if err = json.Unmarshal(resp.Data, &nameServers); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, dns := range nameServers.DNS {
		wg.Add(1)
		go func(dns string) {
			defer wg.Done()
			value := models.Resource{
				ID:   dns,
				Name: "",
				Description: provider.DNSDescription{
					DNS: dns,
				},
			}
			tailScaleChan <- value
		}(dns)
	}

	baseURL = "/v2/tailnet/-/dns/searchpaths"

	req = &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err = handler.Request("tailscale", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &searchPaths); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, searchPath := range searchPaths.SearchPaths {
		wg.Add(1)
		go func(searchPath string) {
			defer wg.Done()
			value := models.Resource{
				ID:   searchPath,
				Name: "",
				Description: provider.DNSDescription{
					DNS: searchPath,
				},
			}
			tailScaleChan <- value
		}(searchPath)
	}
	return nil
}
