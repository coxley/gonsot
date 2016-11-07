// Package conf provides means to load gonsot configuration on-disk
package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
)

var (
	HOME = os.Getenv("HOME")
)

// Config represents settings needed to connect to an NSoT server
//
// You can fill it in manually or autoload configuration by creating an empty
// instance and call Config.Load()
type Config struct {
	ApiVersion    string `ini:"api_version"`
	AuthHeader    string `ini:"auth_header"`
	AuthMethod    string `ini:"auth_method"`
	DefaultDomain string `ini:"default_domain"`
	Email         string `ini:"email"`
	SecretKey     string `ini:"secret_key"`
	Url           string `ini:"url"`
	DefaultSite   int    `ini:"default_site"`
}

type ConfigError struct {
	s string
}

func (err *ConfigError) Error() string {
	return err.s
}

// loadINI loads configs that can be at two locations, listed in merge-priority:
// 1. $HOME/.pynsotrc
// 2. /etc/pynsotrc
func (c *Config) loadINI() error {
	user := filepath.Join(HOME, ".pynsotrc")
	system := filepath.Join("/etc", "pynsotrc")
	var systemErr error
	var userErr error

	// Try opening system-level config, mapping to *Config if it works
	if systemCfg, err := ini.Load(system); err != nil {
		systemErr = err
	} else {
		systemCfg.Section("pynsot").MapTo(c)
	}

	// Try opening user-level config and map to *Config. Should overwrite
	// values added by system and this is desired to merge
	if userCfg, err := ini.Load(user); err != nil {
		userErr = err
	} else {
		userCfg.Section("pynsot").MapTo(c)
	}

	if systemErr != nil && userErr != nil {
		msg := fmt.Sprintf("Failed to open files: %s %s", user, system)
		return &ConfigError{s: msg}
	}
	return nil
}

// loadTOML loads configs that can be at two locations, listed in
// merge-priority:
// 1. $HOME/.gonsot.toml
// 2. /etc/gonsot.toml
func (c *Config) loadTOML() error {
	user := filepath.Join(HOME, ".gonsot.toml")
	system := filepath.Join("/etc", "gonsot.toml")
	var systemErr error
	var userErr error

	// Try opening system-level config, mapping to *Config if it works
	if _, err := toml.DecodeFile(system, c); err != nil {
		systemErr = err
	}

	// Try opening user-level config and map to *Config. Should overwrite
	// values added by system and this is desired to merge
	if _, err := toml.DecodeFile(user, c); err != nil {
		userErr = err
	}

	if systemErr != nil && userErr != nil {
		msg := fmt.Sprintf("Failed to open files: %s %s", user, system)
		return &ConfigError{s: msg}
	}
	return nil
}

// Load marshalls config files into the struct automatically
//
// There are two types of configuration that will be attempted:
// 1. Native gonsot configs, written in TOML (preferred)
//  * gonsot.toml
// 2. Pynsot configs, written in INI (attempted if first fails)
//  * pynsotrc
//
// There isn't a benefit over the other unless there begin to be
// configuration elements that diverge between the projects. Pynsot is the
// official NSoT client library so I wanted to support it's configuration
//
// Regardless of which type is loaded, there are two locations for the files:
//
// * System level (Under /etc)
// * User level (Under $HOME as a dotfile)
//
// User level configs will be preferred, merging if possible with system
// config. This allows "global" configuration but lets a user do overrides
func (c *Config) Load() error {

	if err := c.loadTOML(); err != nil {
		// TOML not found, attempt INI
		if err := c.loadINI(); err != nil {
			// INI not found either
			return &ConfigError{s: "Failed to open INI or TOML configuration"}
		}
		return nil
	}
	return nil
}
