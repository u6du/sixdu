package config

import (
	osUser "os/user"
	"path"
)

type user struct {
	Name string
	Root string
}

type toml struct {
	User user
}

var Toml toml

func init() {
	_user, err := osUser.Current()
	var name string
	if err != nil {
		name = "root"
	} else {
		name = _user.Username
	}
	Toml.User = user{name, path.Join("home", name) + "/"}
}
