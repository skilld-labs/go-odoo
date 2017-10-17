package types

import (
	"time"
)

type IrQwebFieldDate struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldDateNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
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
