package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-tailscale/pkg/sdk/models"
	"github.com/opengovern/og-describer-tailscale/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
)

func ListContacts(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	values, err := processContacts(ctx, handler, stream)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func processContacts(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var contact model.ContactJSON
	baseURL := "/v2/tailnet/-/contacts"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("tailsclae", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &contact); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	account := model.ContactDetail{
		Email:             contact.Account.Email,
		FallbackEmail:     contact.Account.FallbackEmail,
		NeedsVerification: contact.Account.NeedsVerification,
	}
	support := model.ContactDetail{
		Email:             contact.Support.Email,
		FallbackEmail:     contact.Support.FallbackEmail,
		NeedsVerification: contact.Support.NeedsVerification,
	}
	security := model.ContactDetail{
		Email:             contact.Security.Email,
		FallbackEmail:     contact.Security.FallbackEmail,
		NeedsVerification: contact.Security.NeedsVerification,
	}
	var values []models.Resource
	value := models.Resource{
		ID:   contact.Account.Email,
		Name: "",
		Description: model.ContactDescription{
			Account:  account,
			Security: security,
			Support:  support,
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
