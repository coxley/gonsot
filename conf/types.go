package conf

import (
	"bytes"
	"errors"
	"fmt"
	"net/mail"
	"net/url"
)

// Config represents settings needed to connect to an NSoT server
type Config struct {
	APIVersion    string
	AuthHeader    string
	AuthMethod    AuthMethod
	DefaultDomain string
	DefaultSite   int
	Email         Email
	SecretKey     string
	URL           URL

	// Not parsed from file

	// Manually set Site to override DefaultSite or if you don't set
	// DefaultSite
	Site int
}

// Marshalling implementation

// AuthMethod is the type of authenticate we'll try to perform
type AuthMethod int // noqa

// Email is the email address to use to login
type Email mail.Address

// URL is the base URL for an NSoT instance
type URL url.URL

//go:generate stringer -type=AuthMethod
const (
	AuthHeader AuthMethod = iota
	AuthToken
)

// UnmarshalText validates before assigning
func (a *AuthMethod) UnmarshalText(text []byte) (err error) {
	switch {
	case bytes.Equal(text, []byte("auth_header")):
		*a = AuthHeader
	case bytes.Equal(text, []byte("auth_token")):
		*a = AuthToken
	default:
		return errors.New("AuthMethod not in acceptable values")
	}
	return nil
}

// UnmarshalText validates before assigning
func (e *Email) UnmarshalText(text []byte) (err error) {
	addr, err := mail.ParseAddress(string(text))
	if err != nil {
		return err
	}
	e.Address = addr.Address
	return
}

// MarshalText converts back to string form
func (e *Email) MarshalText() (text []byte, err error) {
	return []byte(e.Address), nil
}

// UnmarshalText validates before assigning
func (u *URL) UnmarshalText(text []byte) (err error) {
	url, err := url.Parse(string(text))
	if err != nil {
		return err
	}

	u.Scheme = url.Scheme
	u.Opaque = url.Opaque
	u.User = url.User
	u.Host = url.Host
	u.Path = url.Path
	u.RawPath = url.RawPath
	u.ForceQuery = url.ForceQuery
	u.RawQuery = url.RawQuery
	u.Fragment = url.Fragment
	return
}

// MarshalText converts back to string form
func (u *URL) MarshalText() (text []byte, err error) {
	fmt.Println("DOOT")
	var url url.URL
	url.Scheme = u.Scheme
	url.Opaque = u.Opaque
	url.User = u.User
	url.Host = u.Host
	url.Path = u.Path
	url.RawPath = u.RawPath
	url.ForceQuery = u.ForceQuery
	url.RawQuery = u.RawQuery
	url.Fragment = u.Fragment
	return []byte(url.String()), nil
}
