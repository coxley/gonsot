package conf

import (
	"bytes"
	"net/mail"
	"net/url"

	"github.com/pkg/errors"
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
type AuthMethod int

// Email is the email address to use to login
type Email struct{ mail.Address }

// URL is the base URL for an NSoT instance
type URL struct{ url.URL }

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
		err := errors.New("AuthMethod not in acceptable values")
		return errors.Wrap(err, "Failed Unmarshaling AuthMethod")
	}
	return nil
}

// UnmarshalText validates before assigning
func (e *Email) UnmarshalText(text []byte) (err error) {
	addr, err := mail.ParseAddress(string(text))
	if err != nil {
		return err
	}
	*e = Email{*addr}
	return nil
}

// MarshalText converts back to string form
func (e *Email) MarshalText() (text []byte, err error) {
	return []byte(e.Address.Address), nil
}

// UnmarshalText validates before assigning
func (u *URL) UnmarshalText(text []byte) (err error) {
	url, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	*u = URL{*url}
	return nil
}

// MarshalText converts back to string form
func (u *URL) MarshalText() (text []byte, err error) {
	return []byte(u.String()), nil
}
