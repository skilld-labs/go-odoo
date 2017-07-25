package types

import (
	"time"
)

type BaseModuleUpgrade struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	ModuleInfo  string    `xmlrpc:"module_info"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseModuleUpgradeNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	ModuleInfo  interface{} `xmlrpc:"module_info"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseModuleUpgradeModel string = "base.module.upgrade"

type BaseModuleUpgrades []BaseModuleUpgrade

type BaseModuleUpgradesNil []BaseModuleUpgradeNil

func (s *BaseModuleUpgrade) NilableType_() interface{} {
	return &BaseModuleUpgradeNil{}
}

func (ns *BaseModuleUpgradeNil) Type_() interface{} {
	s := &BaseModuleUpgrade{}
	return load(ns, s)
}

func (s *BaseModuleUpgrades) NilableType_() interface{} {
	return &BaseModuleUpgradesNil{}
}

func (ns *BaseModuleUpgradesNil) Type_() interface{} {
	s := &BaseModuleUpgrades{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseModuleUpgrade))
	}
	return s
}
