package types

import (
	"time"
)

type AutosalesConfig struct {
	ConfigLine  []int64   `xmlrpc:"config_line"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	Sequence    int64     `xmlrpc:"sequence"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AutosalesConfigNil struct {
	ConfigLine  interface{} `xmlrpc:"config_line"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	Sequence    interface{} `xmlrpc:"sequence"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AutosalesConfigModel string = "autosales.config"

type AutosalesConfigs []AutosalesConfig

type AutosalesConfigsNil []AutosalesConfigNil

func (s *AutosalesConfig) NilableType_() interface{} {
	return &AutosalesConfigNil{}
}

func (ns *AutosalesConfigNil) Type_() interface{} {
	s := &AutosalesConfig{}
	return load(ns, s)
}

func (s *AutosalesConfigs) NilableType_() interface{} {
	return &AutosalesConfigsNil{}
}

func (ns *AutosalesConfigsNil) Type_() interface{} {
	s := &AutosalesConfigs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AutosalesConfig))
	}
	return s
}
