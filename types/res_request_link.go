package types

import (
	"time"
)

type ResRequestLink struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Object      string    `xmlrpc:"object"`
	Priority    int64     `xmlrpc:"priority"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResRequestLinkNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Object      interface{} `xmlrpc:"object"`
	Priority    interface{} `xmlrpc:"priority"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResRequestLinkModel string = "res.request.link"

type ResRequestLinks []ResRequestLink

type ResRequestLinksNil []ResRequestLinkNil

func (s *ResRequestLink) NilableType_() interface{} {
	return &ResRequestLinkNil{}
}

func (ns *ResRequestLinkNil) Type_() interface{} {
	s := &ResRequestLink{}
	return load(ns, s)
}

func (s *ResRequestLinks) NilableType_() interface{} {
	return &ResRequestLinksNil{}
}

func (ns *ResRequestLinksNil) Type_() interface{} {
	s := &ResRequestLinks{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResRequestLink))
	}
	return s
}
