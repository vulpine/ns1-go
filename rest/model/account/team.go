package account

// Team wraps an NS1 /accounts/teams resource
type Team struct {
	ID          string         `json:"id,omitempty"`
	Name        string         `json:"name"`
	Permissions PermissionsMap `json:"permissions"`
}

// DDITeam wraps an NS1 /accounts/teams resource for DDI.
type DDITeam struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Permissions DDIPermissionsMap `json:"permissions"`
}
