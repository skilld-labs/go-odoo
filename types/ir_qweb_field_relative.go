package types

import (
	"time"
)

type IrQwebFieldRelative struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldRelativeNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldRelativeModel string = "ir.qweb.field.relative"

type IrQwebFieldRelatives []IrQwebFieldRelative

type IrQwebFieldRelativesNil []IrQwebFieldRelativeNil

func (s *IrQwebFieldRelative) NilableType_() interface{} {
	return &IrQwebFieldRelativeNil{}
}

func (ns *IrQwebFieldRelativeNil) Type_() interface{} {
	s := &IrQwebFieldRelative{}
	return load(ns, s)
}

func (s *IrQwebFieldRelatives) NilableType_() interface{} {
	return &IrQwebFieldRelativesNil{}
}

func (ns *IrQwebFieldRelativesNil) Type_() interface{} {
	s := &IrQwebFieldRelatives{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldRelative))
	}
	return s
}
