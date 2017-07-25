package types

import (
	"time"
)

type Base struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type BaseNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var BaseModel string = "base"

type Bases []Base

type BasesNil []BaseNil

func (s *Base) NilableType_() interface{} {
	return &BaseNil{}
}

func (ns *BaseNil) Type_() interface{} {
	s := &Base{}
	return load(ns, s)
}

func (s *Bases) NilableType_() interface{} {
	return &BasesNil{}
}

func (ns *BasesNil) Type_() interface{} {
	s := &Bases{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*Base))
	}
	return s
}
