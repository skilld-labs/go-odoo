package types

import (
	"time"
)

type BaseModuleConfiguration struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseModuleConfigurationNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseModuleConfigurationModel string = "base.module.configuration"

type BaseModuleConfigurations []BaseModuleConfiguration

type BaseModuleConfigurationsNil []BaseModuleConfigurationNil

func (s *BaseModuleConfiguration) NilableType_() interface{} {
	return &BaseModuleConfigurationNil{}
}

func (ns *BaseModuleConfigurationNil) Type_() interface{} {
	s := &BaseModuleConfiguration{}
	return load(ns, s)
}

func (s *BaseModuleConfigurations) NilableType_() interface{} {
	return &BaseModuleConfigurationsNil{}
}

func (ns *BaseModuleConfigurationsNil) Type_() interface{} {
	s := &BaseModuleConfigurations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseModuleConfiguration))
	}
	return s
}
