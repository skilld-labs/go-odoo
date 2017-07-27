package types

import (
	"time"
)

type IrQwebFieldDate struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldDateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldDateModel string = "ir.qweb.field.date"

type IrQwebFieldDates []IrQwebFieldDate

type IrQwebFieldDatesNil []IrQwebFieldDateNil

func (s *IrQwebFieldDate) NilableType_() interface{} {
	return &IrQwebFieldDateNil{}
}

func (ns *IrQwebFieldDateNil) Type_() interface{} {
	s := &IrQwebFieldDate{}
	return load(ns, s)
}

func (s *IrQwebFieldDates) NilableType_() interface{} {
	return &IrQwebFieldDatesNil{}
}

func (ns *IrQwebFieldDatesNil) Type_() interface{} {
	s := &IrQwebFieldDates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldDate))
	}
	return s
}
