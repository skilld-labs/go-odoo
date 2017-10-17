package types

import (
	"time"
)

type IrQwebField struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldModel string = "ir.qweb.field"

type IrQwebFields []IrQwebField

type IrQwebFieldsNil []IrQwebFieldNil

func (s *IrQwebField) NilableType_() interface{} {
	return &IrQwebFieldNil{}
}

func (ns *IrQwebFieldNil) Type_() interface{} {
	s := &IrQwebField{}
	return load(ns, s)
}

func (s *IrQwebFields) NilableType_() interface{} {
	return &IrQwebFieldsNil{}
}

func (ns *IrQwebFieldsNil) Type_() interface{} {
	s := &IrQwebFields{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebField))
	}
	return s
}
