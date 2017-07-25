package types

import (
	"time"
)

type IrModuleModuleDependency struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DependId    Many2One  `xmlrpc:"depend_id"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	ModuleId    Many2One  `xmlrpc:"module_id"`
	Name        string    `xmlrpc:"name"`
	State       string    `xmlrpc:"state"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrModuleModuleDependencyNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DependId    interface{} `xmlrpc:"depend_id"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	ModuleId    interface{} `xmlrpc:"module_id"`
	Name        interface{} `xmlrpc:"name"`
	State       interface{} `xmlrpc:"state"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrModuleModuleDependencyModel string = "ir.module.module.dependency"

type IrModuleModuleDependencys []IrModuleModuleDependency

type IrModuleModuleDependencysNil []IrModuleModuleDependencyNil

func (s *IrModuleModuleDependency) NilableType_() interface{} {
	return &IrModuleModuleDependencyNil{}
}

func (ns *IrModuleModuleDependencyNil) Type_() interface{} {
	s := &IrModuleModuleDependency{}
	return load(ns, s)
}

func (s *IrModuleModuleDependencys) NilableType_() interface{} {
	return &IrModuleModuleDependencysNil{}
}

func (ns *IrModuleModuleDependencysNil) Type_() interface{} {
	s := &IrModuleModuleDependencys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModuleModuleDependency))
	}
	return s
}
