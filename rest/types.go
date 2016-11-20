package rest

import "net"

// Site contains Site definition
type Site struct {
	Description string
	ID          int64
	Name        string
}

// Attribute contains Attribute definition
type Attribute struct {
	Constraints  Constraints
	Description  string
	Display      bool
	ID           int64
	Multi        bool
	Name         string
	Required     bool
	ResourceName string `json:"resource_name"`
	SiteID       int64  `json:"site_id"`
}

// Constraints contains Attribute Constraints
type Constraints struct {
	AllowEmpty  bool `json:"allow_empty"`
	Pattern     string
	ValidValues []interface{} `json:"valid_values"`
}

// Device contains Device definition
type Device struct {
	Attributes map[string]string
	Hostname   string
	ID         int64
	SiteID     int64 `json:"site_id"`
}

// Interface contains Interface definition
type Interface struct {
	Addresses   []IP
	Attributes  map[string]string
	Description string
	Device      int64
	ID          int64
	MacAddress  HardwareAddr `json:"mac_address"`
	Name        string
	Networks    []IPNet
	ParentID    int64 `json:"parent_id"`
	Speed       int64
	Type        int64
}

// IP is a wrapper around net.IP to provide (Un)Marshaling
// ffjson: skip
type IP struct{ net.IP }

// UnmarshalJSON converts from []byte into meaningful type
func (ip *IP) UnmarshalJSON(text []byte) (err error) {
	addr, _, err := net.ParseCIDR(string(text))
	if err != nil {
		return err
	}
	*ip = IP{addr}
	return nil
}

// HardwareAddr is a wrapper around net.HardwareAddr to provide (Un)Marshaling
// ffjson: skip
type HardwareAddr struct{ net.HardwareAddr }

// UnmarshalJSON converts from []byte into meaningful type
func (mac *HardwareAddr) UnmarshalJSON(text []byte) (err error) {
	m, err := net.ParseMAC(string(text))
	if err != nil {
		return err
	}
	*mac = HardwareAddr{m}
	return nil
}

// IPNet is a wrapper around net.IPNet to provide (Un)Marshaling
// ffjson: skip
type IPNet struct{ net.IPNet }

// UnmarshalJSON converts from []byte into meaningful type
func (n *IPNet) UnmarshalJSON(text []byte) (err error) {
	_, net, err := net.ParseCIDR(string(text))
	if err != nil {
		return err
	}
	*n = IPNet{*net}
	return nil
}

// Circuit contains Circuit definition
type Circuit struct {
	AEndpoint  int64 `json:"a_endpoint"`
	Attributes map[string]string
	ID         int64
	Name       string
	ZEndpoint  int64 `json:"z_endpoint"`
}

// Network contains Network definition
type Network struct {
	Attributes     map[string]string
	ID             int64
	IPVersion      string `json:"ip_version"`
	IsIP           bool   `json:"is_ip"`
	NetworkAddress IP     `json:"network_address"`
	ParentID       int64  `json:"parent_id"`
	PrefixLength   int    `json:"prefix_length"`
	SiteID         int64  `json:"site_id"`
	State          string
}

// Network returns net.IPNet
func (n *Network) Network() net.IPNet {
	var mask net.IPMask
	ip := n.NetworkAddress
	switch ip.To16() {
	default:
		mask = net.CIDRMask(n.PrefixLength, 8*net.IPv6len)
	case nil:
		mask = net.CIDRMask(n.PrefixLength, 8*net.IPv4len)
	}
	return net.IPNet{IP: ip.IP, Mask: mask}
}

// User contains User definition
type User struct {
	Email       string
	ID          int64
	Permissions Permissions
}

// Permissions contains Permissions for a User
type Permissions map[string]struct {
	Permissions []string
	SiteID      int64 `json:"site_id"`
	UserID      int64 `json:"user_id"`
}

// Change contains Change definition
type Change struct {
	ChangeAt     int64 `json:"change_at"`
	Event        string
	ID           int64
	Resource     interface{}
	ResourceID   int64  `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	Site         Site
	User         User
}

// Value contains Value definition
type Value struct {
	Attribute    int64
	ID           int64
	Name         string
	ResourceID   int64  `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	Value        string
}
