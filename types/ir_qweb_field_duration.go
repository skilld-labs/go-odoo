package types

import (
	"time"
)

type IrQwebFieldDuration struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldDurationNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldDurationModel string = "ir.qweb.field.duration"

type IrQwebFieldDurations []IrQwebFieldDuration

type IrQwebFieldDurationsNil []IrQwebFieldDurationNil

func (s *IrQwebFieldDuration) NilableType_() interface{} {
	return &IrQwebFieldDurationNil{}
}

func (ns *IrQwebFieldDurationNil) Type_() interface{} {
	s := &IrQwebFieldDuration{}
	return load(ns, s)
}

func (s *IrQwebFieldDurations) NilableType_() interface{} {
	return &IrQwebFieldDurationsNil{}
}

func (ns *IrQwebFieldDurationsNil) Type_() interface{} {
	s := &IrQwebFieldDurations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldDuration))
	}
	return s
}
