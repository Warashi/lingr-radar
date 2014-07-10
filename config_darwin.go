package main

//Copyright (c) 2014 lestrrat
import (
	"fmt"
	"os"
)

func homedir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("environment variable HOME not set")
	}

	return home, nil
}
