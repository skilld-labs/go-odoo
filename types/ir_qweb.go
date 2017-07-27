package types

import (
	"time"
)

type IrQweb struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebModel string = "ir.qweb"

type IrQwebs []IrQweb

type IrQwebsNil []IrQwebNil

func (s *IrQweb) NilableType_() interface{} {
	return &IrQwebNil{}
}

func (ns *IrQwebNil) Type_() interface{} {
	s := &IrQweb{}
	return load(ns, s)
}

func (s *IrQwebs) NilableType_() interface{} {
	return &IrQwebsNil{}
}

func (ns *IrQwebsNil) Type_() interface{} {
	s := &IrQwebs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQweb))
	}
	return s
}
