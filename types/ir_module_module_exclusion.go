package types

import (
	"time"
)

type IrModuleModuleExclusion struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	ExclusionId Many2One  `xmlrpc:"exclusion_id"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ModuleId    Many2One  `xmlrpc:"module_id"`
	Name        string    `xmlrpc:"name"`
	State       string    `xmlrpc:"state"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrModuleModuleExclusionNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	ExclusionId interface{} `xmlrpc:"exclusion_id"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ModuleId    interface{} `xmlrpc:"module_id"`
	Name        interface{} `xmlrpc:"name"`
	State       interface{} `xmlrpc:"state"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrModuleModuleExclusionModel string = "ir.module.module.exclusion"

type IrModuleModuleExclusions []IrModuleModuleExclusion

type IrModuleModuleExclusionsNil []IrModuleModuleExclusionNil

func (s *IrModuleModuleExclusion) NilableType_() interface{} {
	return &IrModuleModuleExclusionNil{}
}

func (ns *IrModuleModuleExclusionNil) Type_() interface{} {
	s := &IrModuleModuleExclusion{}
	return load(ns, s)
}

func (s *IrModuleModuleExclusions) NilableType_() interface{} {
	return &IrModuleModuleExclusionsNil{}
}

func (ns *IrModuleModuleExclusionsNil) Type_() interface{} {
	s := &IrModuleModuleExclusions{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModuleModuleExclusion))
	}
	return s
}
