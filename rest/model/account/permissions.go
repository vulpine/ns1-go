package account

// PermissionsMap wraps a User's "permissions" attribute
type PermissionsMap struct {
	DNS        PermissionsDNS        `json:"dns"`
	Data       PermissionsData       `json:"data"`
	Account    PermissionsAccount    `json:"account"`
	Monitoring PermissionsMonitoring `json:"monitoring"`
	Security   *PermissionsSecurity  `json:"security,omitempty"`
	DHCP       *PermissionsDHCP      `json:"dhcp,omitempty"`
	IPAM       *PermissionsIPAM      `json:"ipam,omitempty"`
}

// PermissionsDNS wraps a User's "permissions.dns" attribute
type PermissionsDNS struct {
	ViewZones           bool     `json:"view_zones"`
	ManageZones         bool     `json:"manage_zones"`
	ZonesAllowByDefault bool     `json:"zones_allow_by_default"`
	ZonesDeny           []string `json:"zones_deny"`
	ZonesAllow          []string `json:"zones_allow"`
}

// PermissionsData wraps a User's "permissions.data" attribute
type PermissionsData struct {
	PushToDatafeeds   bool `json:"push_to_datafeeds"`
	ManageDatasources bool `json:"manage_datasources"`
	ManageDatafeeds   bool `json:"manage_datafeeds"`
}

// PermissionsAccount wraps a User's "permissions.account" attribute
type PermissionsAccount struct {
	ManageUsers           bool `json:"manage_users"`
	ManagePaymentMethods  bool `json:"manage_payment_methods"`
	ManagePlan            bool `json:"manage_plan"`
	ManageTeams           bool `json:"manage_teams"`
	ManageApikeys         bool `json:"manage_apikeys"`
	ManageAccountSettings bool `json:"manage_account_settings"`
	ViewActivityLog       bool `json:"view_activity_log"`
	ViewInvoices          bool `json:"view_invoices"`
}

// PermissionsSecurity wraps a User's "permissions.security" attribute.
type PermissionsSecurity struct {
	ManageGlobal2FA       bool `json:"manage_global_2fa"`
	ManageActiveDirectory bool `json:"manage_active_directory,omitempty"`
}

// PermissionsMonitoring wraps a User's "permissions.monitoring" attribute
// Only relevant for the managed product.
type PermissionsMonitoring struct {
	ManageLists bool `json:"manage_lists"`
	ManageJobs  bool `json:"manage_jobs"`
	ViewJobs    bool `json:"view_jobs"`
}

// DDIPermissionsMap wraps a User's "permissions" attribute for DDI.
type DDIPermissionsMap struct {
	DNS      PermissionsDNS         `json:"dns"`
	Data     PermissionsData        `json:"data"`
	Account  PermissionsDDIAccount  `json:"account"`
	Security PermissionsDDISecurity `json:"security"`
	DHCP     PermissionsDHCP        `json:"dhcp"`
	IPAM     PermissionsIPAM        `json:"ipam"`
}

// PermissionsDDIAccount wraps a User's "permissions.account" attribute for DDI.
type PermissionsDDIAccount struct {
	ManageUsers           bool `json:"manage_users"`
	ManageTeams           bool `json:"manage_teams"`
	ManageApikeys         bool `json:"manage_apikeys"`
	ManageAccountSettings bool `json:"manage_account_settings"`
	ViewActivityLog       bool `json:"view_activity_log"`
}

// PermissionsDDISecurity wraps a User's "permissions.security" attribute for DDI.
type PermissionsDDISecurity struct {
	ManageGlobal2FA       bool `json:"manage_global_2fa"`
	ManageActiveDirectory bool `json:"manage_active_directory"`
}

// PermissionsDHCP wraps a User's "permissions.dhcp" attribute for DDI.
type PermissionsDHCP struct {
	ManageDHCP bool `json:"manage_dhcp"`
	ViewDHCP   bool `json:"view_dhcp"`
}

// PermissionsIPAM wraps a User's "permissions.ipam" attribute for DDI.
type PermissionsIPAM struct {
	ManageIPAM bool `json:"manage_ipam"`
	ViewIPAM   bool `json:"view_ipam"`
}
