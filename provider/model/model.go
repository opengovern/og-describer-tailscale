//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import "time"

type Metadata struct{}

type ClientConnectivityJSON struct {
	Endpoints             []string                `json:"endpoints"`
	MappingVariesByDestIP bool                    `json:"mappingVariesByDestIP"`
	Latency               map[string]LatencyJSON  `json:"latency"`
	ClientSupports        *ClientCapabilitiesJSON `json:"clientSupports"`
}

type ClientConnectivity struct {
	Endpoints             []string
	MappingVariesByDestIP bool
	Latency               map[string]Latency
	ClientSupports        *ClientCapabilities
}

type LatencyJSON struct {
	Preferred bool    `json:"preferred"`
	LatencyMs float64 `json:"latencyMs"`
}

type Latency struct {
	Preferred bool
	LatencyMs float64
}

type ClientCapabilitiesJSON struct {
	HairPinning bool `json:"hairPinning"`
	IPv6        bool `json:"ipv6"`
	PCP         bool `json:"pcp"`
	PMP         bool `json:"pmp"`
	UDP         bool `json:"udp"`
	UPnP        bool `json:"upnp"`
}

type ClientCapabilities struct {
	HairPinning bool
	IPv6        bool
	PCP         bool
	PMP         bool
	UDP         bool
	UPnP        bool
}

type PostureIdentityJSON struct {
	SerialNumbers []string `json:"serialNumbers"`
	Disabled      bool     `json:"disabled"`
}

type PostureIdentity struct {
	SerialNumbers []string
	Disabled      bool
}

type DeviceListResponse struct {
	Devices []DeviceJSON `json:"devices"`
}

type DeviceJSON struct {
	Addresses                 []string                `json:"addresses"`
	ID                        string                  `json:"id"`
	NodeID                    string                  `json:"nodeId"`
	User                      string                  `json:"user"`
	Name                      string                  `json:"name"`
	Hostname                  string                  `json:"hostname"`
	ClientVersion             string                  `json:"clientVersion"`
	UpdateAvailable           bool                    `json:"updateAvailable"`
	OS                        string                  `json:"os"`
	Created                   *time.Time              `json:"created"`
	LastSeen                  *time.Time              `json:"lastSeen"`
	KeyExpiryDisabled         bool                    `json:"keyExpiryDisabled"`
	Expires                   *time.Time              `json:"expires"`
	Authorized                bool                    `json:"authorized"`
	IsExternal                bool                    `json:"isExternal"`
	MachineKey                string                  `json:"machineKey"`
	NodeKey                   string                  `json:"nodeKey"`
	BlocksIncomingConnections bool                    `json:"blocksIncomingConnections"`
	EnabledRoutes             []string                `json:"enabledRoutes"`
	AdvertisedRoutes          []string                `json:"advertisedRoutes"`
	ClientConnectivity        *ClientConnectivityJSON `json:"clientConnectivity"`
	Tags                      []string                `json:"tags"`
	TailnetLockError          string                  `json:"tailnetLockError"`
	TailnetLockKey            string                  `json:"tailnetLockKey"`
	PostureIdentity           *PostureIdentityJSON    `json:"postureIdentity"`
}

type DeviceDescription struct {
	Addresses                 []string
	ID                        string
	NodeID                    string
	User                      string
	Name                      string
	Hostname                  string
	ClientVersion             string
	UpdateAvailable           bool
	OS                        string
	Created                   *time.Time
	LastSeen                  *time.Time
	KeyExpiryDisabled         bool
	Expires                   *time.Time
	Authorized                bool
	IsExternal                bool
	MachineKey                string
	NodeKey                   string
	BlocksIncomingConnections bool
	EnabledRoutes             []string
	AdvertisedRoutes          []string
	ClientConnectivity        *ClientConnectivity
	Tags                      []string
	TailnetLockError          string
	TailnetLockKey            string
	PostureIdentity           *PostureIdentity
}
