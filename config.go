package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var c *Config
var homedirFunc = homedir

func init() {
	c = NewConfig()
	filename, err := LocateRcfile()
	if err != nil {
		log.Fatal(err)
	}
	err = c.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	User     string
	Password string
	APIKey   string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ReadFile(filename string) error {
	// Copyright (c) 2014 lestrrat

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return err
	}

	return nil
}

var _locateRcfileIn = locateRcfileIn

func locateRcfileIn(dir string) (string, error) {
	// Copyright (c) 2014 lestrrat

	const basename = "config.json"
	file := filepath.Join(dir, basename)
	if _, err := os.Stat(file); err != nil {
		return "", err
	}
	return file, nil
}

func LocateRcfile() (string, error) {
	// Copyright (c) 2014 lestrrat

	// http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
	//
	// Try in this order:
	// $XDG_CONFIG_HOME/lingr-radar/config.json
	// $XDG_CONFIG_DIR/lingr-radar/config.json (where XDG_CONFIG_DIR is listed in $XDG_CONFIG_DIRS)
	// ~/.lingr-radar/config.json

	home, uErr := homedirFunc()

	// Try dir supplied via env var
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		file, err := _locateRcfileIn(filepath.Join(dir, "lingr-radar"))
		if err == nil {
			return file, nil
		}
	} else if uErr == nil { // silently ignore failure for homedir()
		// Try "default" XDG location, is user is available
		file, err := _locateRcfileIn(filepath.Join(home, ".config", "lingr-radar"))
		if err == nil {
			return file, nil
		}
	}

	// this standard does not take into consideration windows (duh)
	// while the spec says use ":" as the separator, Go provides us
	// with filepath.ListSeparator, so use it
	if dirs := os.Getenv("XDG_CONFIG_DIRS"); dirs != "" {
		for _, dir := range strings.Split(dirs, fmt.Sprintf("%c", filepath.ListSeparator)) {
			file, err := _locateRcfileIn(filepath.Join(dir, "lingr-radar"))
			if err == nil {
				return file, nil
			}
		}
	}

	if uErr == nil { // silently ignore failure for homedir()
		file, err := _locateRcfileIn(filepath.Join(home, ".lingr-radar"))
		if err == nil {
			return file, nil
		}
	}

	return "", fmt.Errorf("config file not found")
}
