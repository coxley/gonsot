package rest

import (
	"net"
	"time"

	"github.com/coxley/gonsot/conf"
)

//go:generate ffjson $GOFILE

// Resource defines common methods that the resources should have
type Resource interface {
	Plural() string
	GetAll() ([]Resource, error)
}

// Resources contains multiple Resources
type Resources []Resource

// Site contains Site definition
type Site struct {
	Description string
	ID          int64
	Name        string
}

// Sites contains multiple Sites
type Sites []Site

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

// Attributes contains multiple Attributes
type Attributes []Attribute

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

// Devices contains multiple Devices
type Devices []Device

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

// Interfaces contains multiple Interfaces
type Interfaces []Interface

// IP is a wrapper around net.IP to provide (Un)Marshaling
// ffjson: skip
type IP struct{ net.IP }

// HardwareAddr is a wrapper around net.HardwareAddr to provide (Un)Marshaling
// ffjson: skip
type HardwareAddr struct{ net.HardwareAddr }

// IPNet is a wrapper around net.IPNet to provide (Un)Marshaling
// ffjson: skip
type IPNet struct{ net.IPNet }

// Circuit contains Circuit definition
type Circuit struct {
	AEndpoint  int64 `json:"a_endpoint"`
	Attributes map[string]string
	ID         int64
	Name       string
	ZEndpoint  int64 `json:"z_endpoint"`
}

// Circuits contains multiple Circuits
type Circuits []Circuit

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

// Would like to use enum for representing State but there are issues with
// jsonenums because it's Case Sensitive

// State represents Network resource state
// type State int

//goxx:generate stringer -type=State
//goxx:generate jsonenums -type=State
// const (
// 	Allocated State = iota
// 	Assigned
// 	Orphaned
// 	Reserved
// )

// Networks contains multiple Networks
type Networks []Network

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
	Email       conf.Email
	ID          int64
	Permissions Permissions
}

// Users contains multiple Users
type Users []User

// Permissions contains Permissions for a User
type Permissions map[string]struct {
	Permissions []string
	SiteID      int64 `json:"site_id"`
	UserID      int64 `json:"user_id"`
}

// Change contains Change definition
type Change struct {
	ChangeAt     time.Time `json:"change_at"`
	Event        string
	ID           int64
	Resource     interface{}
	ResourceID   int64  `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	Site         Site
	User         User
}

// Time is a wrapper around time.Time to provide (Un)Marshaling
// ffjson: skip
type Time struct{ time.Time }

// Changes contains multiple Changes
type Changes []Change

// Value contains Value definition
type Value struct {
	Attribute    int64
	ID           int64
	Name         string
	ResourceID   int64  `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	Value        string
}

// Values contains multiple Values
type Values []Value
