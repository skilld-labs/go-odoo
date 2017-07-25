package types

import (
	"time"
)

type IrProperty struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	FieldsId       Many2One  `xmlrpc:"fields_id"`
	Id             int64     `xmlrpc:"id"`
	Name           string    `xmlrpc:"name"`
	ResId          string    `xmlrpc:"res_id"`
	Type           string    `xmlrpc:"type"`
	ValueBinary    string    `xmlrpc:"value_binary"`
	ValueDatetime  time.Time `xmlrpc:"value_datetime"`
	ValueFloat     float64   `xmlrpc:"value_float"`
	ValueInteger   int64     `xmlrpc:"value_integer"`
	ValueReference string    `xmlrpc:"value_reference"`
	ValueText      string    `xmlrpc:"value_text"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type IrPropertyNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	FieldsId       interface{} `xmlrpc:"fields_id"`
	Id             interface{} `xmlrpc:"id"`
	Name           interface{} `xmlrpc:"name"`
	ResId          interface{} `xmlrpc:"res_id"`
	Type           interface{} `xmlrpc:"type"`
	ValueBinary    interface{} `xmlrpc:"value_binary"`
	ValueDatetime  interface{} `xmlrpc:"value_datetime"`
	ValueFloat     interface{} `xmlrpc:"value_float"`
	ValueInteger   interface{} `xmlrpc:"value_integer"`
	ValueReference interface{} `xmlrpc:"value_reference"`
	ValueText      interface{} `xmlrpc:"value_text"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var IrPropertyModel string = "ir.property"

type IrPropertys []IrProperty

type IrPropertysNil []IrPropertyNil

func (s *IrProperty) NilableType_() interface{} {
	return &IrPropertyNil{}
}

func (ns *IrPropertyNil) Type_() interface{} {
	s := &IrProperty{}
	return load(ns, s)
}

func (s *IrPropertys) NilableType_() interface{} {
	return &IrPropertysNil{}
}

func (ns *IrPropertysNil) Type_() interface{} {
	s := &IrPropertys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrProperty))
	}
	return s
}
