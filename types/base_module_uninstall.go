package types

import (
	"time"
)

type BaseModuleUninstall struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ModelIds    []int64   `xmlrpc:"model_ids"`
	ModuleId    Many2One  `xmlrpc:"module_id"`
	ModuleIds   []int64   `xmlrpc:"module_ids"`
	ShowAll     bool      `xmlrpc:"show_all"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseModuleUninstallNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ModelIds    interface{} `xmlrpc:"model_ids"`
	ModuleId    interface{} `xmlrpc:"module_id"`
	ModuleIds   interface{} `xmlrpc:"module_ids"`
	ShowAll     bool        `xmlrpc:"show_all"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseModuleUninstallModel string = "base.module.uninstall"

type BaseModuleUninstalls []BaseModuleUninstall

type BaseModuleUninstallsNil []BaseModuleUninstallNil

func (s *BaseModuleUninstall) NilableType_() interface{} {
	return &BaseModuleUninstallNil{}
}

func (ns *BaseModuleUninstallNil) Type_() interface{} {
	s := &BaseModuleUninstall{}
	return load(ns, s)
}

func (s *BaseModuleUninstalls) NilableType_() interface{} {
	return &BaseModuleUninstallsNil{}
}

func (ns *BaseModuleUninstallsNil) Type_() interface{} {
	s := &BaseModuleUninstalls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseModuleUninstall))
	}
	return s
}
