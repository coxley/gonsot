package rest

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Resource implementations

// Plural returns string representing lowercase-plural resource name
func (s *Site) Plural() string {
	return "sites"
}

// GetAll queries NSoT for all Sites
func (s Site) GetAll() ([]Resource, error) {
	body, err := getAll(&Site{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Sites)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (a *Attribute) Plural() string {
	return "attributes"
}

// GetAll queries NSoT for all Attributes
func (a Attribute) GetAll() ([]Resource, error) {
	body, err := getAll(&Attribute{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Attributes)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (d *Device) Plural() string {
	return "devices"
}

// GetAll queries NSoT for all Devices
func (d Device) GetAll() ([]Resource, error) {
	body, err := getAll(&Device{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Devices)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (i *Interface) Plural() string {
	return "interfaces"
}

// GetAll queries NSoT for all Interfaces
func (i Interface) GetAll() ([]Resource, error) {
	body, err := getAll(&Interface{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Interfaces)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (c *Circuit) Plural() string {
	return "circuits"
}

// GetAll queries NSoT for all Circuits
func (c Circuit) GetAll() ([]Resource, error) {
	body, err := getAll(&Circuit{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Circuits)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (n *Network) Plural() string {
	return "networks"
}

// GetAll queries NSoT for all Networks
func (n Network) GetAll() ([]Resource, error) {
	body, err := getAll(&Network{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Networks)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (u *User) Plural() string {
	return "users"
}

// GetAll queries NSoT for all Users
func (u User) GetAll() ([]Resource, error) {
	body, err := getAll(&User{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Users)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (c *Change) Plural() string {
	return "changes"
}

// GetAll queries NSoT for all Changes
func (c Change) GetAll() ([]Resource, error) {
	body, err := getAll(&Change{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Changes)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Plural returns string representing lowercase-plural resource name
func (v *Value) Plural() string {
	return "values"
}

// GetAll queries NSoT for all Values
func (v Value) GetAll() ([]Resource, error) {
	body, err := getAll(&Value{})
	if err != nil {
		return nil, errors.Wrap(err, "Fetching resources failed")
	}
	var rs = new(Values)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshaling failed")
	}
	resources := make([]Resource, len(*rs))
	for i, v := range *rs {
		resources[i] = Resource(&v)
	}

	return resources, nil
}

// Stringer implementations

func (s *Site) String() string {
	return s.Name
}

func (a *Attribute) String() string {
	return a.ResourceName + ":" + a.Name
}

func (d *Device) String() string {
	return d.Hostname
}

func (i *Interface) String() string {
	return fmt.Sprintf("Device ID: %d, Name: %s", i.Device, i.Name)
}

func (c *Circuit) String() string {
	return c.Name
}

func (n *Network) String() string {
	return fmt.Sprint(n.Network())
}

func (u *User) String() string {
	return fmt.Sprint(u.Email)
}

// JSON (Un)Marshaler implementations

// UnmarshalJSON converts from []byte into meaningful type
func (ip *IP) UnmarshalJSON(text []byte) (err error) {

	// For some reason, text contains the quotes from JSON
	var s string
	if s, err = strconv.Unquote(string(text)); err != nil {
		s = string(text)
	}

	addr, _, err := net.ParseCIDR(s)
	if err != nil {
		addr = net.ParseIP(s)
		if addr == nil {
			return errors.Wrap(err, "Could not parse as CIDR or IP")
		}
	}
	*ip = IP{addr}
	return nil
}

// MarshalJSON converts from []byte into meaningful type
func (ip *IP) MarshalJSON() (text []byte, err error) {
	return []byte(fmt.Sprint(ip)), nil
}

// UnmarshalJSON converts from []byte into meaningful type
func (mac *HardwareAddr) UnmarshalJSON(text []byte) (err error) {
	m, err := net.ParseMAC(string(text))
	if err != nil {
		return err
	}
	*mac = HardwareAddr{m}
	return nil
}

// MarshalJSON converts from []byte into meaningful type
func (mac *HardwareAddr) MarshalJSON() (text []byte, err error) {
	return mac.HardwareAddr, nil
}

// UnmarshalJSON converts from []byte into meaningful type
func (n *IPNet) UnmarshalJSON(text []byte) (err error) {
	_, net, err := net.ParseCIDR(string(text))
	if err != nil {
		return err
	}
	*n = IPNet{*net}
	return nil
}

// MarshalJSON converts from []byte into meaningful type
func (n *IPNet) MarshalJSON() (text []byte, err error) {
	return []byte(fmt.Sprint(n)), nil
}

// UnmarshalJSON converts from []byte into meaningful type
func (t *Time) UnmarshalJSON(text []byte) (err error) {
	ts, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		return err
	}
	*t = Time{time.Unix(ts, 0)}
	return nil
}

// MarshalJSON converts from []byte into meaningful type
func (t *Time) MarshalJSON() (text []byte, err error) {
	return []byte(string(t.Unix())), nil
}
