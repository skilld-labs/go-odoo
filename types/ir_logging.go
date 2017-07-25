package types

import (
	"time"
)

type IrLogging struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   int64     `xmlrpc:"create_uid"`
	Dbname      string    `xmlrpc:"dbname"`
	DisplayName string    `xmlrpc:"display_name"`
	Func        string    `xmlrpc:"func"`
	Id          int64     `xmlrpc:"id"`
	Level       string    `xmlrpc:"level"`
	Line        string    `xmlrpc:"line"`
	Message     string    `xmlrpc:"message"`
	Name        string    `xmlrpc:"name"`
	Path        string    `xmlrpc:"path"`
	Type        string    `xmlrpc:"type"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrLoggingNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Dbname      interface{} `xmlrpc:"dbname"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Func        interface{} `xmlrpc:"func"`
	Id          interface{} `xmlrpc:"id"`
	Level       interface{} `xmlrpc:"level"`
	Line        interface{} `xmlrpc:"line"`
	Message     interface{} `xmlrpc:"message"`
	Name        interface{} `xmlrpc:"name"`
	Path        interface{} `xmlrpc:"path"`
	Type        interface{} `xmlrpc:"type"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrLoggingModel string = "ir.logging"

type IrLoggings []IrLogging

type IrLoggingsNil []IrLoggingNil

func (s *IrLogging) NilableType_() interface{} {
	return &IrLoggingNil{}
}

func (ns *IrLoggingNil) Type_() interface{} {
	s := &IrLogging{}
	return load(ns, s)
}

func (s *IrLoggings) NilableType_() interface{} {
	return &IrLoggingsNil{}
}

func (ns *IrLoggingsNil) Type_() interface{} {
	s := &IrLoggings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrLogging))
	}
	return s
}
