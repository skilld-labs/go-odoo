package types

import (
	"time"
)

type IrQwebFieldHtml struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldHtmlNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
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
