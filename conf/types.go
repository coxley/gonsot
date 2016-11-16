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
	ApiVersion    string
	AuthHeader    string
	AuthMethod    AuthMethod
	DefaultDomain string
	Email         Email
	SecretKey     string
	URL           URL
	DefaultSite   int

	// Not parsed from file

	// Manually set Site to override DefaultSite or if you don't set
	// DefaultSite
	Site int
}

// Marshalling implementation

type AuthMethod string
type Email mail.Address
type URL url.URL

func (a *AuthMethod) UnmarshalText(text []byte) (err error) {
	switch {
	case bytes.Equal(text, []byte("auth_header")):
		break
	case bytes.Equal(text, []byte("auth_token")):
		break
	default:
		return errors.New("AuthMethod not in acceptable values")
	}
	*a = AuthMethod(text)
	return nil
}

func (e *Email) UnmarshalText(text []byte) (err error) {
	addr, err := mail.ParseAddress(string(text))
	if err != nil {
		return err
	}
	e.Address = addr.Address
	return
}

func (e *Email) MarshalText() (text []byte, err error) {
	return []byte(e.Address), nil
}

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
