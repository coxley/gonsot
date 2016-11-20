package rest

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// Stringer implementation

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
	addr, _, err := net.ParseCIDR(string(text))
	if err != nil {
		return err
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
