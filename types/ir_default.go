package types

import (
	"time"
)

type IrDefault struct {
	CompanyId   Many2One  `xmlrpc:"company_id"`
	Condition   string    `xmlrpc:"condition"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	FieldId     Many2One  `xmlrpc:"field_id"`
	Id          int64     `xmlrpc:"id"`
	JsonValue   string    `xmlrpc:"json_value"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	UserId      Many2One  `xmlrpc:"user_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrDefaultNil struct {
	CompanyId   interface{} `xmlrpc:"company_id"`
	Condition   interface{} `xmlrpc:"condition"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	FieldId     interface{} `xmlrpc:"field_id"`
	Id          interface{} `xmlrpc:"id"`
	JsonValue   interface{} `xmlrpc:"json_value"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	UserId      interface{} `xmlrpc:"user_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrDefaultModel string = "ir.default"

type IrDefaults []IrDefault

type IrDefaultsNil []IrDefaultNil

func (s *IrDefault) NilableType_() interface{} {
	return &IrDefaultNil{}
}

func (ns *IrDefaultNil) Type_() interface{} {
	s := &IrDefault{}
	return load(ns, s)
}

func (s *IrDefaults) NilableType_() interface{} {
	return &IrDefaultsNil{}
}

func (ns *IrDefaultsNil) Type_() interface{} {
	s := &IrDefaults{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrDefault))
	}
	return s
}
