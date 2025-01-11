// Code is generated by go generate. DO NOT EDIT.
package opengovernance

import (
	"context"
	tailscale "github.com/opengovern/og-describer-tailscale/provider/model"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"
	steampipesdk "github.com/opengovern/og-util/pkg/steampipe"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"runtime"
)

type Client struct {
	essdk.Client
}

// ==========================  START: Device =============================

type Device struct {
	ResourceID      string                      `json:"resource_id"`
	PlatformID      string                      `json:"platform_id"`
	Description     tailscale.DeviceDescription `json:"Description"`
	Metadata        tailscale.Metadata          `json:"metadata"`
	DescribedBy     string                      `json:"described_by"`
	ResourceType    string                      `json:"resource_type"`
	IntegrationType string                      `json:"integration_type"`
	IntegrationID   string                      `json:"integration_id"`
}

type DeviceHit struct {
	ID      string        `json:"_id"`
	Score   float64       `json:"_score"`
	Index   string        `json:"_index"`
	Type    string        `json:"_type"`
	Version int64         `json:"_version,omitempty"`
	Source  Device        `json:"_source"`
	Sort    []interface{} `json:"sort"`
}

type DeviceHits struct {
	Total essdk.SearchTotal `json:"total"`
	Hits  []DeviceHit       `json:"hits"`
}

type DeviceSearchResponse struct {
	PitID string     `json:"pit_id"`
	Hits  DeviceHits `json:"hits"`
}

type DevicePaginator struct {
	paginator *essdk.BaseESPaginator
}

func (k Client) NewDevicePaginator(filters []essdk.BoolFilter, limit *int64) (DevicePaginator, error) {
	paginator, err := essdk.NewPaginator(k.ES(), "tailscale_device", filters, limit)
	if err != nil {
		return DevicePaginator{}, err
	}

	p := DevicePaginator{
		paginator: paginator,
	}

	return p, nil
}

func (p DevicePaginator) HasNext() bool {
	return !p.paginator.Done()
}

func (p DevicePaginator) Close(ctx context.Context) error {
	return p.paginator.Deallocate(ctx)
}

func (p DevicePaginator) NextPage(ctx context.Context) ([]Device, error) {
	var response DeviceSearchResponse
	err := p.paginator.Search(ctx, &response)
	if err != nil {
		return nil, err
	}

	var values []Device
	for _, hit := range response.Hits.Hits {
		values = append(values, hit.Source)
	}

	hits := int64(len(response.Hits.Hits))
	if hits > 0 {
		p.paginator.UpdateState(hits, response.Hits.Hits[hits-1].Sort, response.PitID)
	} else {
		p.paginator.UpdateState(hits, nil, "")
	}

	return values, nil
}

var listDeviceFilters = map[string]string{
	"addresses":                   "Description.Addresses",
	"advertised_routes":           "Description.AdvertisedRoutes",
	"authorized":                  "Description.Authorized",
	"blocks_incoming_connections": "Description.BlocksIncomingConnections",
	"client_connectivity":         "Description.ClientConnectivity",
	"client_version":              "Description.ClientVersion",
	"enabled_routes":              "Description.EnabledRoutes",
	"hostname":                    "Description.Hostname",
	"id":                          "Description.ID",
	"is_external":                 "Description.IsExternal",
	"key_expiry_disabled":         "Description.KeyExpiryDisabled",
	"machine_key":                 "Description.MachineKey",
	"name":                        "Description.Name",
	"node_id":                     "Description.NodeID",
	"node_key":                    "Description.NodeKey",
	"os":                          "Description.OS",
	"posture_identity":            "Description.PostureIdentity",
	"tags":                        "Description.Tags",
	"tailnet_lock_error":          "Description.TailnetLockError",
	"tailnet_lock_key":            "Description.TailnetLockKey",
	"update_available":            "Description.UpdateAvailable",
	"user":                        "Description.User",
}

func ListDevice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ListDevice")
	runtime.GC()

	// create service
	cfg := essdk.GetConfig(d.Connection)
	ke, err := essdk.NewClientCached(cfg, d.ConnectionCache, ctx)
	if err != nil {
		plugin.Logger(ctx).Error("ListDevice NewClientCached", "error", err)
		return nil, err
	}
	k := Client{Client: ke}

	sc, err := steampipesdk.NewSelfClientCached(ctx, d.ConnectionCache)
	if err != nil {
		plugin.Logger(ctx).Error("ListDevice NewSelfClientCached", "error", err)
		return nil, err
	}
	integrationId, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyIntegrationID)
	if err != nil {
		plugin.Logger(ctx).Error("ListDevice GetConfigTableValueOrNil for OpenGovernanceConfigKeyIntegrationID", "error", err)
		return nil, err
	}
	encodedResourceCollectionFilters, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyResourceCollectionFilters)
	if err != nil {
		plugin.Logger(ctx).Error("ListDevice GetConfigTableValueOrNil for OpenGovernanceConfigKeyResourceCollectionFilters", "error", err)
		return nil, err
	}
	clientType, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyClientType)
	if err != nil {
		plugin.Logger(ctx).Error("ListDevice GetConfigTableValueOrNil for OpenGovernanceConfigKeyClientType", "error", err)
		return nil, err
	}

	paginator, err := k.NewDevicePaginator(essdk.BuildFilter(ctx, d.QueryContext, listDeviceFilters, integrationId, encodedResourceCollectionFilters, clientType), d.QueryContext.Limit)
	if err != nil {
		plugin.Logger(ctx).Error("ListDevice NewDevicePaginator", "error", err)
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("ListDevice paginator.NextPage", "error", err)
			return nil, err
		}

		for _, v := range page {
			d.StreamListItem(ctx, v)
		}
	}

	err = paginator.Close(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

var getDeviceFilters = map[string]string{
	"addresses":                   "Description.Addresses",
	"advertised_routes":           "Description.AdvertisedRoutes",
	"authorized":                  "Description.Authorized",
	"blocks_incoming_connections": "Description.BlocksIncomingConnections",
	"client_connectivity":         "Description.ClientConnectivity",
	"client_version":              "Description.ClientVersion",
	"enabled_routes":              "Description.EnabledRoutes",
	"hostname":                    "Description.Hostname",
	"id":                          "Description.ID",
	"is_external":                 "Description.IsExternal",
	"key_expiry_disabled":         "Description.KeyExpiryDisabled",
	"machine_key":                 "Description.MachineKey",
	"name":                        "Description.Name",
	"node_id":                     "Description.NodeID",
	"node_key":                    "Description.NodeKey",
	"os":                          "Description.OS",
	"posture_identity":            "Description.PostureIdentity",
	"tags":                        "Description.Tags",
	"tailnet_lock_error":          "Description.TailnetLockError",
	"tailnet_lock_key":            "Description.TailnetLockKey",
	"update_available":            "Description.UpdateAvailable",
	"user":                        "Description.User",
}

func GetDevice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("GetDevice")
	runtime.GC()
	// create service
	cfg := essdk.GetConfig(d.Connection)
	ke, err := essdk.NewClientCached(cfg, d.ConnectionCache, ctx)
	if err != nil {
		return nil, err
	}
	k := Client{Client: ke}

	sc, err := steampipesdk.NewSelfClientCached(ctx, d.ConnectionCache)
	if err != nil {
		return nil, err
	}
	integrationId, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyIntegrationID)
	if err != nil {
		return nil, err
	}
	encodedResourceCollectionFilters, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyResourceCollectionFilters)
	if err != nil {
		return nil, err
	}
	clientType, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyClientType)
	if err != nil {
		return nil, err
	}

	limit := int64(1)
	paginator, err := k.NewDevicePaginator(essdk.BuildFilter(ctx, d.QueryContext, getDeviceFilters, integrationId, encodedResourceCollectionFilters, clientType), &limit)
	if err != nil {
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page {
			return v, nil
		}
	}

	err = paginator.Close(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ==========================  END: Device =============================

// ==========================  START: User =============================

type User struct {
	ResourceID      string                    `json:"resource_id"`
	PlatformID      string                    `json:"platform_id"`
	Description     tailscale.UserDescription `json:"Description"`
	Metadata        tailscale.Metadata        `json:"metadata"`
	DescribedBy     string                    `json:"described_by"`
	ResourceType    string                    `json:"resource_type"`
	IntegrationType string                    `json:"integration_type"`
	IntegrationID   string                    `json:"integration_id"`
}

type UserHit struct {
	ID      string        `json:"_id"`
	Score   float64       `json:"_score"`
	Index   string        `json:"_index"`
	Type    string        `json:"_type"`
	Version int64         `json:"_version,omitempty"`
	Source  User          `json:"_source"`
	Sort    []interface{} `json:"sort"`
}

type UserHits struct {
	Total essdk.SearchTotal `json:"total"`
	Hits  []UserHit         `json:"hits"`
}

type UserSearchResponse struct {
	PitID string   `json:"pit_id"`
	Hits  UserHits `json:"hits"`
}

type UserPaginator struct {
	paginator *essdk.BaseESPaginator
}

func (k Client) NewUserPaginator(filters []essdk.BoolFilter, limit *int64) (UserPaginator, error) {
	paginator, err := essdk.NewPaginator(k.ES(), "tailscale_user", filters, limit)
	if err != nil {
		return UserPaginator{}, err
	}

	p := UserPaginator{
		paginator: paginator,
	}

	return p, nil
}

func (p UserPaginator) HasNext() bool {
	return !p.paginator.Done()
}

func (p UserPaginator) Close(ctx context.Context) error {
	return p.paginator.Deallocate(ctx)
}

func (p UserPaginator) NextPage(ctx context.Context) ([]User, error) {
	var response UserSearchResponse
	err := p.paginator.Search(ctx, &response)
	if err != nil {
		return nil, err
	}

	var values []User
	for _, hit := range response.Hits.Hits {
		values = append(values, hit.Source)
	}

	hits := int64(len(response.Hits.Hits))
	if hits > 0 {
		p.paginator.UpdateState(hits, response.Hits.Hits[hits-1].Sort, response.PitID)
	} else {
		p.paginator.UpdateState(hits, nil, "")
	}

	return values, nil
}

var listUserFilters = map[string]string{
	"created":             "Created",
	"currently_connected": "CurrentlyConnected",
	"device_count":        "DeviceCount",
	"display_name":        "DisplayName",
	"id":                  "ID",
	"last_seen":           "LastSeen",
	"login_name":          "LoginName",
	"profile_pic_url":     "ProfilePicURL",
	"role":                "Role",
	"status":              "Status",
	"tailnet_id":          "TailnetID",
	"type":                "Type",
}

func ListUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("ListUser")
	runtime.GC()

	// create service
	cfg := essdk.GetConfig(d.Connection)
	ke, err := essdk.NewClientCached(cfg, d.ConnectionCache, ctx)
	if err != nil {
		plugin.Logger(ctx).Error("ListUser NewClientCached", "error", err)
		return nil, err
	}
	k := Client{Client: ke}

	sc, err := steampipesdk.NewSelfClientCached(ctx, d.ConnectionCache)
	if err != nil {
		plugin.Logger(ctx).Error("ListUser NewSelfClientCached", "error", err)
		return nil, err
	}
	integrationId, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyIntegrationID)
	if err != nil {
		plugin.Logger(ctx).Error("ListUser GetConfigTableValueOrNil for OpenGovernanceConfigKeyIntegrationID", "error", err)
		return nil, err
	}
	encodedResourceCollectionFilters, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyResourceCollectionFilters)
	if err != nil {
		plugin.Logger(ctx).Error("ListUser GetConfigTableValueOrNil for OpenGovernanceConfigKeyResourceCollectionFilters", "error", err)
		return nil, err
	}
	clientType, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyClientType)
	if err != nil {
		plugin.Logger(ctx).Error("ListUser GetConfigTableValueOrNil for OpenGovernanceConfigKeyClientType", "error", err)
		return nil, err
	}

	paginator, err := k.NewUserPaginator(essdk.BuildFilter(ctx, d.QueryContext, listUserFilters, integrationId, encodedResourceCollectionFilters, clientType), d.QueryContext.Limit)
	if err != nil {
		plugin.Logger(ctx).Error("ListUser NewUserPaginator", "error", err)
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("ListUser paginator.NextPage", "error", err)
			return nil, err
		}

		for _, v := range page {
			d.StreamListItem(ctx, v)
		}
	}

	err = paginator.Close(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

var getUserFilters = map[string]string{
	"created":             "Created",
	"currently_connected": "CurrentlyConnected",
	"device_count":        "DeviceCount",
	"display_name":        "DisplayName",
	"id":                  "ID",
	"last_seen":           "LastSeen",
	"login_name":          "LoginName",
	"profile_pic_url":     "ProfilePicURL",
	"role":                "Role",
	"status":              "Status",
	"tailnet_id":          "TailnetID",
	"type":                "Type",
}

func GetUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("GetUser")
	runtime.GC()
	// create service
	cfg := essdk.GetConfig(d.Connection)
	ke, err := essdk.NewClientCached(cfg, d.ConnectionCache, ctx)
	if err != nil {
		return nil, err
	}
	k := Client{Client: ke}

	sc, err := steampipesdk.NewSelfClientCached(ctx, d.ConnectionCache)
	if err != nil {
		return nil, err
	}
	integrationId, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyIntegrationID)
	if err != nil {
		return nil, err
	}
	encodedResourceCollectionFilters, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyResourceCollectionFilters)
	if err != nil {
		return nil, err
	}
	clientType, err := sc.GetConfigTableValueOrNil(ctx, steampipesdk.OpenGovernanceConfigKeyClientType)
	if err != nil {
		return nil, err
	}

	limit := int64(1)
	paginator, err := k.NewUserPaginator(essdk.BuildFilter(ctx, d.QueryContext, getUserFilters, integrationId, encodedResourceCollectionFilters, clientType), &limit)
	if err != nil {
		return nil, err
	}

	for paginator.HasNext() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page {
			return v, nil
		}
	}

	err = paginator.Close(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ==========================  END: User =============================
