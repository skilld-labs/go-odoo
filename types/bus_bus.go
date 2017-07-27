package types

import (
	"time"
)

type BusBus struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Channel     string    `xmlrpc:"channel"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Message     string    `xmlrpc:"message"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BusBusNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Channel     interface{} `xmlrpc:"channel"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Message     interface{} `xmlrpc:"message"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BusBusModel string = "bus.bus"

type BusBuss []BusBus

type BusBussNil []BusBusNil

func (s *BusBus) NilableType_() interface{} {
	return &BusBusNil{}
}

func (ns *BusBusNil) Type_() interface{} {
	s := &BusBus{}
	return load(ns, s)
}

func (s *BusBuss) NilableType_() interface{} {
	return &BusBussNil{}
}

func (ns *BusBussNil) Type_() interface{} {
	s := &BusBuss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BusBus))
	}
	return s
}
