package types

import (
	"time"
)

type UtmSource struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type UtmSourceNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var UtmSourceModel string = "utm.source"

type UtmSources []UtmSource

type UtmSourcesNil []UtmSourceNil

func (s *UtmSource) NilableType_() interface{} {
	return &UtmSourceNil{}
}

func (ns *UtmSourceNil) Type_() interface{} {
	s := &UtmSource{}
	return load(ns, s)
}

func (s *UtmSources) NilableType_() interface{} {
	return &UtmSourcesNil{}
}

func (ns *UtmSourcesNil) Type_() interface{} {
	s := &UtmSources{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*UtmSource))
	}
	return s
}
