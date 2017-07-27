package types

import (
	"time"
)

type IrModuleCategory struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ChildIds    []int64   `xmlrpc:"child_ids"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Description string    `xmlrpc:"description"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	ModuleIds   []int64   `xmlrpc:"module_ids"`
	ModuleNr    int64     `xmlrpc:"module_nr"`
	Name        string    `xmlrpc:"name"`
	ParentId    Many2One  `xmlrpc:"parent_id"`
	Sequence    int64     `xmlrpc:"sequence"`
	Visible     bool      `xmlrpc:"visible"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
	XmlId       string    `xmlrpc:"xml_id"`
}

type IrModuleCategoryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ChildIds    interface{} `xmlrpc:"child_ids"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Description interface{} `xmlrpc:"description"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	ModuleIds   interface{} `xmlrpc:"module_ids"`
	ModuleNr    interface{} `xmlrpc:"module_nr"`
	Name        interface{} `xmlrpc:"name"`
	ParentId    interface{} `xmlrpc:"parent_id"`
	Sequence    interface{} `xmlrpc:"sequence"`
	Visible     bool        `xmlrpc:"visible"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
	XmlId       interface{} `xmlrpc:"xml_id"`
}

var IrModuleCategoryModel string = "ir.module.category"

type IrModuleCategorys []IrModuleCategory

type IrModuleCategorysNil []IrModuleCategoryNil

func (s *IrModuleCategory) NilableType_() interface{} {
	return &IrModuleCategoryNil{}
}

func (ns *IrModuleCategoryNil) Type_() interface{} {
	s := &IrModuleCategory{}
	return load(ns, s)
}

func (s *IrModuleCategorys) NilableType_() interface{} {
	return &IrModuleCategorysNil{}
}

func (ns *IrModuleCategorysNil) Type_() interface{} {
	s := &IrModuleCategorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModuleCategory))
	}
	return s
}
