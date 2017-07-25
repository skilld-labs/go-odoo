package types

import (
	"time"
)

type FetchmailServer struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	ActionId      Many2One  `xmlrpc:"action_id"`
	Active        bool      `xmlrpc:"active"`
	Attach        bool      `xmlrpc:"attach"`
	Configuration string    `xmlrpc:"configuration"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	Date          time.Time `xmlrpc:"date"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	IsSsl         bool      `xmlrpc:"is_ssl"`
	MessageIds    []int64   `xmlrpc:"message_ids"`
	Name          string    `xmlrpc:"name"`
	ObjectId      Many2One  `xmlrpc:"object_id"`
	Original      bool      `xmlrpc:"original"`
	Password      string    `xmlrpc:"password"`
	Port          int64     `xmlrpc:"port"`
	Priority      int64     `xmlrpc:"priority"`
	Script        string    `xmlrpc:"script"`
	Server        string    `xmlrpc:"server"`
	State         string    `xmlrpc:"state"`
	Type          string    `xmlrpc:"type"`
	User          string    `xmlrpc:"user"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type FetchmailServerNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	ActionId      interface{} `xmlrpc:"action_id"`
	Active        bool        `xmlrpc:"active"`
	Attach        bool        `xmlrpc:"attach"`
	Configuration interface{} `xmlrpc:"configuration"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	Date          interface{} `xmlrpc:"date"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	IsSsl         bool        `xmlrpc:"is_ssl"`
	MessageIds    interface{} `xmlrpc:"message_ids"`
	Name          interface{} `xmlrpc:"name"`
	ObjectId      interface{} `xmlrpc:"object_id"`
	Original      bool        `xmlrpc:"original"`
	Password      interface{} `xmlrpc:"password"`
	Port          interface{} `xmlrpc:"port"`
	Priority      interface{} `xmlrpc:"priority"`
	Script        interface{} `xmlrpc:"script"`
	Server        interface{} `xmlrpc:"server"`
	State         interface{} `xmlrpc:"state"`
	Type          interface{} `xmlrpc:"type"`
	User          interface{} `xmlrpc:"user"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var FetchmailServerModel string = "fetchmail.server"

type FetchmailServers []FetchmailServer

type FetchmailServersNil []FetchmailServerNil

func (s *FetchmailServer) NilableType_() interface{} {
	return &FetchmailServerNil{}
}

func (ns *FetchmailServerNil) Type_() interface{} {
	s := &FetchmailServer{}
	return load(ns, s)
}

func (s *FetchmailServers) NilableType_() interface{} {
	return &FetchmailServersNil{}
}

func (ns *FetchmailServersNil) Type_() interface{} {
	s := &FetchmailServers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*FetchmailServer))
	}
	return s
}
