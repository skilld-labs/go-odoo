package types

import (
	"time"
)

type BaseModuleUpdate struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Added       int64     `xmlrpc:"added"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	State       string    `xmlrpc:"state"`
	Updated     int64     `xmlrpc:"updated"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseModuleUpdateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Added       interface{} `xmlrpc:"added"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	State       interface{} `xmlrpc:"state"`
	Updated     interface{} `xmlrpc:"updated"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseModuleUpdateModel string = "base.module.update"

type BaseModuleUpdates []BaseModuleUpdate

type BaseModuleUpdatesNil []BaseModuleUpdateNil

func (s *BaseModuleUpdate) NilableType_() interface{} {
	return &BaseModuleUpdateNil{}
}

func (ns *BaseModuleUpdateNil) Type_() interface{} {
	s := &BaseModuleUpdate{}
	return load(ns, s)
}

func (s *BaseModuleUpdates) NilableType_() interface{} {
	return &BaseModuleUpdatesNil{}
}

func (ns *BaseModuleUpdatesNil) Type_() interface{} {
	s := &BaseModuleUpdates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseModuleUpdate))
	}
	return s
}
