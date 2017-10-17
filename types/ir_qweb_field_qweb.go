package types

import (
	"time"
)

type IrQwebFieldQweb struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldQwebNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldQwebModel string = "ir.qweb.field.qweb"

type IrQwebFieldQwebs []IrQwebFieldQweb

type IrQwebFieldQwebsNil []IrQwebFieldQwebNil

func (s *IrQwebFieldQweb) NilableType_() interface{} {
	return &IrQwebFieldQwebNil{}
}

func (ns *IrQwebFieldQwebNil) Type_() interface{} {
	s := &IrQwebFieldQweb{}
	return load(ns, s)
}

func (s *IrQwebFieldQwebs) NilableType_() interface{} {
	return &IrQwebFieldQwebsNil{}
}

func (ns *IrQwebFieldQwebsNil) Type_() interface{} {
	s := &IrQwebFieldQwebs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldQweb))
	}
	return s
}
