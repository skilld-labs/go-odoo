package types

import (
	"time"
)

type IrQwebFieldMonetary struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldMonetaryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldMonetaryModel string = "ir.qweb.field.monetary"

type IrQwebFieldMonetarys []IrQwebFieldMonetary

type IrQwebFieldMonetarysNil []IrQwebFieldMonetaryNil

func (s *IrQwebFieldMonetary) NilableType_() interface{} {
	return &IrQwebFieldMonetaryNil{}
}

func (ns *IrQwebFieldMonetaryNil) Type_() interface{} {
	s := &IrQwebFieldMonetary{}
	return load(ns, s)
}

func (s *IrQwebFieldMonetarys) NilableType_() interface{} {
	return &IrQwebFieldMonetarysNil{}
}

func (ns *IrQwebFieldMonetarysNil) Type_() interface{} {
	s := &IrQwebFieldMonetarys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldMonetary))
	}
	return s
}
