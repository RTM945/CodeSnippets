package main

import (
	"errors"
	"sync"
)

type user struct {
	username string
	password string
	sdps     map[string]string
}

var users sync.Map

func register(username, password string) error {
	if _, ok := users.Load(username); ok {
		return errors.New("already exist")
	}
	user := new(user)
	user.username = username
	user.password = password
	users.Store(username, user)
	return nil
}
