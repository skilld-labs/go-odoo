package types

import (
	"time"
)

type BusPresence struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	DisplayName  string    `xmlrpc:"display_name"`
	Id           int64     `xmlrpc:"id"`
	LastPoll     time.Time `xmlrpc:"last_poll"`
	LastPresence time.Time `xmlrpc:"last_presence"`
	Status       string    `xmlrpc:"status"`
	UserId       Many2One  `xmlrpc:"user_id"`
}

type BusPresenceNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Id           interface{} `xmlrpc:"id"`
	LastPoll     interface{} `xmlrpc:"last_poll"`
	LastPresence interface{} `xmlrpc:"last_presence"`
	Status       interface{} `xmlrpc:"status"`
	UserId       interface{} `xmlrpc:"user_id"`
}

var BusPresenceModel string = "bus.presence"

type BusPresences []BusPresence

type BusPresencesNil []BusPresenceNil

func (s *BusPresence) NilableType_() interface{} {
	return &BusPresenceNil{}
}

func (ns *BusPresenceNil) Type_() interface{} {
	s := &BusPresence{}
	return load(ns, s)
}

func (s *BusPresences) NilableType_() interface{} {
	return &BusPresencesNil{}
}

func (ns *BusPresencesNil) Type_() interface{} {
	s := &BusPresences{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BusPresence))
	}
	return s
}
