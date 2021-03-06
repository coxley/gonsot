// Package conf provides means to load gonsot configuration on-disk
package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
	"github.com/serenize/snaker"
)

var (
	home    = os.Getenv("HOME")
	sysCFG  = filepath.Join("/etc", "gonsot.toml")
	userCFG = filepath.Join(home, ".gonsot.toml")
)

// loadTOML loads configs that can be at two locations, listed in
// merge-priority:
// 1. $HOME/.gonsot.toml
// 2. /etc/gonsot.toml
func (c *Config) loadTOML() error {
	var systemErr error
	var userErr error

	// Try opening system-level config, mapping to *Config if it works
	if _, err := toml.DecodeFile(sysCFG, c); err != nil {
		systemErr = err
	}

	// Try opening user-level config and map to *Config. Should overwrite
	// values added by system and this is desired to merge
	if _, err := toml.DecodeFile(userCFG, c); err != nil {
		userErr = err
	}

	if systemErr != nil && userErr != nil {
		msg := fmt.Sprintf("Failed to open and validate files: %s %s", userCFG, sysCFG)
		err := errors.New(msg)
		return errors.Wrap(err, "Failed to load TOML")
	}
	return nil
}

// convertINI converts pynsot config to gonsot
func convertINI() (err error) {
	p := filepath.Join("/etc", "pynsotrc") // Parent
	c := filepath.Join(home, ".pynsotrc")  // Child

	cfg, err := ini.LooseLoad(p, c) // `c` has higher merge pri over `p`
	if err != nil {
		return err
	}

	// Encode TOML to file
	f, err := os.Create("/home/coxley/.gonsot.toml")
	if err != nil {
		return err
	}
	defer f.Close()

	// INI parses values as all strings, we'll convert to int if possible
	// In addition to that, we'll map snake_case to PascalCase
	iniConfig := map[string]interface{}{}
	for k, v := range cfg.Section("pynsot").KeysHash() {
		num, err := strconv.Atoi(v)
		nk := snaker.SnakeToCamel(k) // New key
		if err != nil {
			iniConfig[nk] = v
			continue
		}
		iniConfig[nk] = num
	}
	if err := toml.NewEncoder(f).Encode(iniConfig); err != nil {
		return err
	}
	return nil
}

// Load marshalls config files into the struct automatically
//
// There are two locations for gonsot.toml
//
// * System level (Under /etc)
// * User level (Under $HOME as a dotfile)
//
// User level configs will be preferred, merging if possible with system
// config. This allows "global" configuration but lets a user do overrides
//
// If existing INI config from pynsot exists and no TOML found, will create
// TOML config with the same settings
func (c *Config) Load() error {

	if err := c.loadTOML(); err != nil {
		// TOML not found, attempt to convert pynsotrc INI to TOML if it exists
		_ = convertINI()

		// Attempt reload of TOML
		if err := c.loadTOML(); err != nil {
			return err
		}
		return nil
	}

	return nil
}

// Dump writes the Config instance to provided filename
func (c *Config) Dump(fn string) error {

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := toml.NewEncoder(f).Encode(*c); err != nil {
		return err
	}
	return nil

}
