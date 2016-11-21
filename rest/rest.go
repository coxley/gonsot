package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/coxley/gonsot/conf"
	"github.com/pkg/errors"
)

var cfg *conf.Config

func getClient() *http.Client {
	var client = new(http.Client)
	client = &http.Client{Timeout: time.Second * 10}
	return client
}

// addHeaders adds to req.Header the needed headers to talk to NSoT
func addHeaders(req *http.Request) error {
	req.Header.Set("Content-Type", "application/json")

	switch cfg.AuthMethod {
	default:
		err := errors.New("Unknown AuthMethod")
		return errors.Wrap(err, "Adding authentication to request failed")
	case conf.AuthHeader:
		req.Header.Set(cfg.AuthHeader, cfg.Email.Address.Address)
		return nil
	case conf.AuthToken:
		msg := fmt.Sprintf("AuthMethod: %s not yet implemented", cfg.AuthMethod)
		err := errors.New(msg)
		return errors.Wrap(err, "Adding authentication to request failed")
	}

}

// initCfg initializes the package global cfg
func initCfg() error {
	if cfg != nil {
		return nil // Already initialized
	}
	c := &conf.Config{}
	err := c.Load()
	if err != nil {
		return errors.Wrap(err, "Couldn't initialize cfg")
	}
	cfg = c
	return nil
}

// getAll returns payload from fetching all resources for given r
func getAll(r Resource) ([]byte, error) {
	initCfg()
	url := cfg.URL
	url.Path = "/api/" + r.Plural() + "/"

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create http.Request")
	}
	addHeaders(req)

	res, err := getClient().Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Request failed")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Reading body failed")
	}
	return body, nil
}
