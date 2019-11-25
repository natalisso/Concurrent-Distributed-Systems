package subscribermanager

import (
	"net"
)

type subscribermanager struct {
	Subscribers map[string]net.Conn
}

func NewSuscriberManager() subscribermanager {
	sm := new(subscribermanager)
	sm.Subscribers = make(map[string]net.Conn)

	return *sm
}

