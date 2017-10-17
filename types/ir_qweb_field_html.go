package types

import (
	"time"
)

type IrQwebFieldHtml struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldHtmlNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldHtmlModel string = "ir.qweb.field.html"

type IrQwebFieldHtmls []IrQwebFieldHtml

type IrQwebFieldHtmlsNil []IrQwebFieldHtmlNil

func (s *IrQwebFieldHtml) NilableType_() interface{} {
	return &IrQwebFieldHtmlNil{}
}

func (ns *IrQwebFieldHtmlNil) Type_() interface{} {
	s := &IrQwebFieldHtml{}
	return load(ns, s)
}

func (s *IrQwebFieldHtmls) NilableType_() interface{} {
	return &IrQwebFieldHtmlsNil{}
}

func (ns *IrQwebFieldHtmlsNil) Type_() interface{} {
	s := &IrQwebFieldHtmls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldHtml))
	}
	return s
}
