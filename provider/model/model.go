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
	Addresses []string `json:"addresses"`
	ID        string   `json:"id"`
	NodeID    string   `json:"nodeId"`
	User      string   `json:"user"`
	Name      string   `json:"name"`
	//Hostname                  string                  `json:"hostname"`
	ClientVersion string `json:"clientVersion"`
	//UpdateAvailable           bool                    `json:"updateAvailable"`
	OS string `json:"os"`
	//Created                   *time.Time              `json:"created"`
	LastSeen *time.Time `json:"lastSeen"`
	//KeyExpiryDisabled         bool                    `json:"keyExpiryDisabled"`
	//Expires                   *time.Time              `json:"expires"`
	//Authorized                bool                    `json:"authorized"`
	//IsExternal                bool                    `json:"isExternal"`
	//MachineKey                string                  `json:"machineKey"`
	//NodeKey                   string                  `json:"nodeKey"`
	//BlocksIncomingConnections bool                    `json:"blocksIncomingConnections"`
	//EnabledRoutes             []string                `json:"enabledRoutes"`
	//AdvertisedRoutes          []string                `json:"advertisedRoutes"`
	ClientConnectivity *ClientConnectivityJSON `json:"clientConnectivity"`
	Tags               []string                `json:"tags"`
	//TailnetLockError          string                  `json:"tailnetLockError"`
	//TailnetLockKey            string                  `json:"tailnetLockKey"`
	PostureIdentity *PostureIdentityJSON `json:"postureIdentity"`
}

type DeviceDescription struct {
	Addresses []string
	ID        string
	NodeID    string
	User      string
	Name      string
	//Hostname                  string
	ClientVersion string
	//UpdateAvailable           bool
	OS string
	//Created                   *time.Time
	LastSeen *time.Time
	//KeyExpiryDisabled         bool
	//Expires                   *time.Time
	//Authorized                bool
	//IsExternal                bool
	//MachineKey                string
	//NodeKey                   string
	//BlocksIncomingConnections bool
	//EnabledRoutes             []string
	//AdvertisedRoutes          []string
	ClientConnectivity *ClientConnectivity
	Tags               []string
	//TailnetLockError          string
	//TailnetLockKey            string
	PostureIdentity *PostureIdentity
}

type UserListResponse struct {
	Users []UserJSON `json:"users"`
}

type UserJSON struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	LoginName   string `json:"loginName"`
	//ProfilePicURL      string    `json:"profilePicUrl"`
	TailnetID          string    `json:"tailnetId"`
	Created            time.Time `json:"created"`
	Type               string    `json:"type"`
	Role               string    `json:"role"`
	Status             string    `json:"status"`
	DeviceCount        int       `json:"deviceCount"`
	LastSeen           time.Time `json:"lastSeen"`
	CurrentlyConnected bool      `json:"currentlyConnected"`
}

type UserDescription struct {
	ID          string
	DisplayName string
	LoginName   string
	//ProfilePicURL      string
	TailnetID          string
	Created            time.Time
	Type               string
	Role               string
	Status             string
	DeviceCount        int
	LastSeen           time.Time
	CurrentlyConnected bool
}

type ACLJSON struct {
	Action string   `json:"action"`
	Ports  []string `json:"ports"`
	Users  []string `json:"users"`
}

type ACL struct {
	Action string
	Ports  []string
	Users  []string
}

type PolicyJSON struct {
	ACLs   []ACLJSON           `json:"acls"`
	Groups map[string][]string `json:"groups"`
	Hosts  map[string]string   `json:"hosts"`
}

type PolicyDescription struct {
	ACLs   []ACL
	Groups map[string][]string
	Hosts  map[string]string
}

type CreatePermissionsJSON struct {
	Reusable      bool     `json:"reusable,omitempty"`
	Ephemeral     bool     `json:"ephemeral,omitempty"`
	Preauthorized bool     `json:"preauthorized,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

type CreatePermissions struct {
	Reusable      bool
	Ephemeral     bool
	Preauthorized bool
	Tags          []string
}

type DeviceCapabilitiesJSON struct {
	Create CreatePermissionsJSON `json:"create,omitempty"`
}

type DeviceCapabilities struct {
	Create CreatePermissions
}

type CapabilitiesJSON struct {
	Devices DeviceCapabilities `json:"devices,omitempty"`
}

type Capabilities struct {
	Devices DeviceCapabilities
}

type KeyListResponse struct {
	Keys []KeyJSON `json:"keys"`
}

type KeyJSON struct {
	ID           string           `json:"id"`
	Created      *time.Time       `json:"created,omitempty"`
	Expires      *time.Time       `json:"expires,omitempty"`
	Revoked      *time.Time       `json:"revoked,omitempty"`
	Capabilities CapabilitiesJSON `json:"capabilities,omitempty"`
	Description  string           `json:"description,omitempty"`
	Invalid      bool             `json:"invalid,omitempty"`
	UserID       string           `json:"userId,omitempty"`
}

type KeyDescription struct {
	ID           string
	Created      *time.Time
	Expires      *time.Time
	Revoked      *time.Time
	Capabilities Capabilities
	Description  string
	Invalid      bool
	UserID       string
}

type UserInviteJSON struct {
	ID              string     `json:"id"`
	Role            string     `json:"role"`
	TailnetID       int64      `json:"tailnetId"`
	InviterID       int64      `json:"inviterId"`
	Email           string     `json:"email,omitempty"`
	LastEmailSentAt *time.Time `json:"lastEmailSentAt,omitempty"`
	InviteURL       string     `json:"inviteUrl,omitempty"`
}

type UserInviteDescription struct {
	ID              string
	Role            string
	TailnetID       int64
	InviterID       int64
	Email           string
	LastEmailSentAt *time.Time
	InviteURL       string
}
