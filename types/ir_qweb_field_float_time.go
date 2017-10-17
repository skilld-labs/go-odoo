package types

import (
	"time"
)

type IrQwebFieldFloatTime struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldFloatTimeNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldFloatTimeModel string = "ir.qweb.field.float_time"

type IrQwebFieldFloatTimes []IrQwebFieldFloatTime

type IrQwebFieldFloatTimesNil []IrQwebFieldFloatTimeNil

func (s *IrQwebFieldFloatTime) NilableType_() interface{} {
	return &IrQwebFieldFloatTimeNil{}
}

func (ns *IrQwebFieldFloatTimeNil) Type_() interface{} {
	s := &IrQwebFieldFloatTime{}
	return load(ns, s)
}

func (s *IrQwebFieldFloatTimes) NilableType_() interface{} {
	return &IrQwebFieldFloatTimesNil{}
}

func (ns *IrQwebFieldFloatTimesNil) Type_() interface{} {
	s := &IrQwebFieldFloatTimes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldFloatTime))
	}
	return s
}
