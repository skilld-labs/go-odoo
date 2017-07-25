package types

import (
	"time"
)

type IrQwebFieldInteger struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldIntegerNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
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
