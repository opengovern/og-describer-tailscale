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

type AcceptedByJSON struct {
	ID            int64  `json:"id"`
	LoginName     string `json:"loginName"`
	ProfilePicURL string `json:"profilePicUrl"`
}

type AcceptedBy struct {
	ID            int64
	LoginName     string
	ProfilePicURL string
}

type DeviceInviteJSON struct {
	ID              string          `json:"id"`
	Created         time.Time       `json:"created"`
	TailnetID       int64           `json:"tailnetId"`
	DeviceID        int64           `json:"deviceId"`
	SharerID        int64           `json:"sharerId"`
	MultiUse        bool            `json:"multiUse,omitempty"`
	AllowExitNode   bool            `json:"allowExitNode,omitempty"`
	Email           string          `json:"email,omitempty"`
	LastEmailSentAt *time.Time      `json:"lastEmailSentAt,omitempty"`
	InviteURL       string          `json:"inviteUrl,omitempty"`
	Accepted        bool            `json:"accepted"`
	AcceptedBy      *AcceptedByJSON `json:"acceptedBy,omitempty"`
}

type DeviceInviteDescription struct {
	ID              string
	Created         time.Time
	TailnetID       int64
	DeviceID        int64
	SharerID        int64
	MultiUse        bool
	AllowExitNode   bool
	Email           string
	LastEmailSentAt *time.Time
	InviteURL       string
	Accepted        bool
	AcceptedBy      *AcceptedBy
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

type StatusJSON struct {
	LastSync             time.Time `json:"lastSync"`
	Error                string    `json:"error,omitempty"`
	MatchedCount         int       `json:"matchedCount"`
	PossibleMatchedCount int       `json:"possibleMatchedCount"`
	ProviderHostCount    int       `json:"providerHostCount"`
}

type Status struct {
	LastSync             time.Time
	Error                string
	MatchedCount         int
	PossibleMatchedCount int
	ProviderHostCount    int
}

type ListPostureIntegrations struct {
	Integrations []PostureIntegrationJSON `json:"integrations"`
}

type PostureIntegrationJSON struct {
	Provider      string      `json:"provider"`
	CloudID       string      `json:"cloudId"`
	ClientID      string      `json:"clientId,omitempty"`
	TenantID      string      `json:"tenantId,omitempty"`
	ClientSecret  string      `json:"clientSecret,omitempty"`
	ID            string      `json:"id"`
	ConfigUpdated string      `json:"configUpdated"`
	Status        *StatusJSON `json:"status,omitempty"`
}

type PostureIntegrationDescription struct {
	Provider      string
	CloudID       string
	ClientID      string
	TenantID      string
	ClientSecret  string
	ID            string
	ConfigUpdated string
	Status        *Status
}

type ContactDetailJSON struct {
	Email             string `json:"email"`
	FallbackEmail     string `json:"fallbackEmail,omitempty"`
	NeedsVerification bool   `json:"needsVerification,omitempty"`
}

type ContactDetail struct {
	Email             string
	FallbackEmail     string
	NeedsVerification bool
}

type ContactJSON struct {
	Account  ContactDetailJSON `json:"account"`
	Support  ContactDetailJSON `json:"support"`
	Security ContactDetailJSON `json:"security"`
}

type ContactDescription struct {
	Account  ContactDetail
	Support  ContactDetail
	Security ContactDetail
}

type ListWebhookResponse struct {
	Webhooks []WebhookJSON `json:"webhooks"`
}

type WebhookJSON struct {
	EndpointID       string    `json:"endpointId"`
	EndpointURL      string    `json:"endpointUrl"`
	ProviderType     string    `json:"providerType"`
	CreatorLoginName string    `json:"creatorLoginName"`
	Created          time.Time `json:"created"`
	LastModified     time.Time `json:"lastModified"`
	Subscriptions    []string  `json:"subscriptions"`
	Secret           string    `json:"secret"`
}

type WebhookDescription struct {
	EndpointID       string
	EndpointURL      string
	ProviderType     string
	CreatorLoginName string
	Created          time.Time
	LastModified     time.Time
	Subscriptions    []string
	Secret           string
}

type TailnetSettingsJSON struct {
	DevicesApprovalOn                      bool   `json:"devicesApprovalOn"`
	DevicesAutoUpdatesOn                   bool   `json:"devicesAutoUpdatesOn"`
	DevicesKeyDurationDays                 int    `json:"devicesKeyDurationDays"`
	UsersApprovalOn                        bool   `json:"usersApprovalOn"`
	UsersRoleAllowedToJoinExternalTailnets string `json:"usersRoleAllowedToJoinExternalTailnets"`
	NetworkFlowLoggingOn                   bool   `json:"networkFlowLoggingOn"`
	RegionalRoutingOn                      bool   `json:"regionalRoutingOn"`
	PostureIdentityCollectionOn            bool   `json:"postureIdentityCollectionOn"`
}

type TailnetSettingsDescription struct {
	DevicesApprovalOn                      bool
	DevicesAutoUpdatesOn                   bool
	DevicesKeyDurationDays                 int
	UsersApprovalOn                        bool
	UsersRoleAllowedToJoinExternalTailnets string
	NetworkFlowLoggingOn                   bool
	RegionalRoutingOn                      bool
	PostureIdentityCollectionOn            bool
}

type ListDNSNameServerResponse struct {
	DNS []string `json:"dns"`
}

type ListDNSSearchPathsResponse struct {
	SearchPaths []string `json:"searchPaths"`
}

type DNSDescription struct {
	DNS string
}
