package types

import (
	"time"
)

type IrQwebFieldText struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldTextNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldTextModel string = "ir.qweb.field.text"

type IrQwebFieldTexts []IrQwebFieldText

type IrQwebFieldTextsNil []IrQwebFieldTextNil

func (s *IrQwebFieldText) NilableType_() interface{} {
	return &IrQwebFieldTextNil{}
}

func (ns *IrQwebFieldTextNil) Type_() interface{} {
	s := &IrQwebFieldText{}
	return load(ns, s)
}

func (s *IrQwebFieldTexts) NilableType_() interface{} {
	return &IrQwebFieldTextsNil{}
}

func (ns *IrQwebFieldTextsNil) Type_() interface{} {
	s := &IrQwebFieldTexts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldText))
	}
	return s
}
