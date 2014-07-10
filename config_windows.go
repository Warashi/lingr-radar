package peco

//Copyright (c) 2014 lestrrat
import "os/user"

func homedir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}
