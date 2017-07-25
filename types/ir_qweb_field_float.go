package types

import (
	"time"
)

type IrQwebFieldFloat struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldFloatNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldFloatModel string = "ir.qweb.field.float"

type IrQwebFieldFloats []IrQwebFieldFloat

type IrQwebFieldFloatsNil []IrQwebFieldFloatNil

func (s *IrQwebFieldFloat) NilableType_() interface{} {
	return &IrQwebFieldFloatNil{}
}

func (ns *IrQwebFieldFloatNil) Type_() interface{} {
	s := &IrQwebFieldFloat{}
	return load(ns, s)
}

func (s *IrQwebFieldFloats) NilableType_() interface{} {
	return &IrQwebFieldFloatsNil{}
}

func (ns *IrQwebFieldFloatsNil) Type_() interface{} {
	s := &IrQwebFieldFloats{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldFloat))
	}
	return s
}
