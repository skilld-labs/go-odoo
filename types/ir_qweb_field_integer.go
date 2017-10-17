package types

import (
	"time"
)

type IrQwebFieldInteger struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldIntegerNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldIntegerModel string = "ir.qweb.field.integer"

type IrQwebFieldIntegers []IrQwebFieldInteger

type IrQwebFieldIntegersNil []IrQwebFieldIntegerNil

func (s *IrQwebFieldInteger) NilableType_() interface{} {
	return &IrQwebFieldIntegerNil{}
}

func (ns *IrQwebFieldIntegerNil) Type_() interface{} {
	s := &IrQwebFieldInteger{}
	return load(ns, s)
}

func (s *IrQwebFieldIntegers) NilableType_() interface{} {
	return &IrQwebFieldIntegersNil{}
}

func (ns *IrQwebFieldIntegersNil) Type_() interface{} {
	s := &IrQwebFieldIntegers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldInteger))
	}
	return s
}
