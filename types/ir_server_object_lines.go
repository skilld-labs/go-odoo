package types

import (
	"time"
)

type IrServerObjectLines struct {
	Col1        Many2One    `xmlrpc:"col1"`
	CreateDate  time.Time   `xmlrpc:"create_date"`
	CreateUid   Many2One    `xmlrpc:"create_uid"`
	DisplayName string      `xmlrpc:"display_name"`
	Id          int64       `xmlrpc:"id"`
	LastUpdate  time.Time   `xmlrpc:"__last_update"`
	ServerId    Many2One    `xmlrpc:"server_id"`
	Type        interface{} `xmlrpc:"type"`
	Value       string      `xmlrpc:"value"`
	WriteDate   time.Time   `xmlrpc:"write_date"`
	WriteUid    Many2One    `xmlrpc:"write_uid"`
}

type IrServerObjectLinesNil struct {
	Col1        interface{} `xmlrpc:"col1"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ServerId    interface{} `xmlrpc:"server_id"`
	Type        interface{} `xmlrpc:"type"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrServerObjectLinesModel string = "ir.server.object.lines"

type IrServerObjectLiness []IrServerObjectLines

type IrServerObjectLinessNil []IrServerObjectLinesNil

func (s *IrServerObjectLines) NilableType_() interface{} {
	return &IrServerObjectLinesNil{}
}

func (ns *IrServerObjectLinesNil) Type_() interface{} {
	s := &IrServerObjectLines{}
	return load(ns, s)
}

func (s *IrServerObjectLiness) NilableType_() interface{} {
	return &IrServerObjectLinessNil{}
}

func (ns *IrServerObjectLinessNil) Type_() interface{} {
	s := &IrServerObjectLiness{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrServerObjectLines))
	}
	return s
}
