package main

import (
	"sync"
)

var sessions sync.Map

func GetSession(sid int32) *Session {
	value, ok := sessions.Load(sid)
	if ok {
		return value.(*Session)
	}
	return nil
}

func CheckAlive() {
	sessions.Range(func(key, value interface{}) bool {
		session := value.(*Session)
		if !session.alive() {
			session.OnClose()
		}
		return true
	})
}
