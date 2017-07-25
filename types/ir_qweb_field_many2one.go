package types

import (
	"time"
)

type IrQwebFieldMany2one struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldMany2oneNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldMany2oneModel string = "ir.qweb.field.many2one"

type IrQwebFieldMany2ones []IrQwebFieldMany2one

type IrQwebFieldMany2onesNil []IrQwebFieldMany2oneNil

func (s *IrQwebFieldMany2one) NilableType_() interface{} {
	return &IrQwebFieldMany2oneNil{}
}

func (ns *IrQwebFieldMany2oneNil) Type_() interface{} {
	s := &IrQwebFieldMany2one{}
	return load(ns, s)
}

func (s *IrQwebFieldMany2ones) NilableType_() interface{} {
	return &IrQwebFieldMany2onesNil{}
}

func (ns *IrQwebFieldMany2onesNil) Type_() interface{} {
	s := &IrQwebFieldMany2ones{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldMany2one))
	}
	return s
}
