package types

import (
	"time"
)

type IrHttp struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrHttpNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrHttpModel string = "ir.http"

type IrHttps []IrHttp

type IrHttpsNil []IrHttpNil

func (s *IrHttp) NilableType_() interface{} {
	return &IrHttpNil{}
}

func (ns *IrHttpNil) Type_() interface{} {
	s := &IrHttp{}
	return load(ns, s)
}

func (s *IrHttps) NilableType_() interface{} {
	return &IrHttpsNil{}
}

func (ns *IrHttpsNil) Type_() interface{} {
	s := &IrHttps{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrHttp))
	}
	return s
}
