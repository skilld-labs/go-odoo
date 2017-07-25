package types

import (
	"time"
)

type ResConfig struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResConfigNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResConfigModel string = "res.config"

type ResConfigs []ResConfig

type ResConfigsNil []ResConfigNil

func (s *ResConfig) NilableType_() interface{} {
	return &ResConfigNil{}
}

func (ns *ResConfigNil) Type_() interface{} {
	s := &ResConfig{}
	return load(ns, s)
}

func (s *ResConfigs) NilableType_() interface{} {
	return &ResConfigsNil{}
}

func (ns *ResConfigsNil) Type_() interface{} {
	s := &ResConfigs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResConfig))
	}
	return s
}
