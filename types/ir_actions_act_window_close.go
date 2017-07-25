package types

import (
	"time"
)

type IrActionsActWindowClose struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Help        string    `xmlrpc:"help"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Type        string    `xmlrpc:"type"`
	Usage       string    `xmlrpc:"usage"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
	XmlId       string    `xmlrpc:"xml_id"`
}

type IrActionsActWindowCloseNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Help        interface{} `xmlrpc:"help"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Type        interface{} `xmlrpc:"type"`
	Usage       interface{} `xmlrpc:"usage"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
	XmlId       interface{} `xmlrpc:"xml_id"`
}

var IrActionsActWindowCloseModel string = "ir.actions.act_window_close"

type IrActionsActWindowCloses []IrActionsActWindowClose

type IrActionsActWindowClosesNil []IrActionsActWindowCloseNil

func (s *IrActionsActWindowClose) NilableType_() interface{} {
	return &IrActionsActWindowCloseNil{}
}

func (ns *IrActionsActWindowCloseNil) Type_() interface{} {
	s := &IrActionsActWindowClose{}
	return load(ns, s)
}

func (s *IrActionsActWindowCloses) NilableType_() interface{} {
	return &IrActionsActWindowClosesNil{}
}

func (ns *IrActionsActWindowClosesNil) Type_() interface{} {
	s := &IrActionsActWindowCloses{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsActWindowClose))
	}
	return s
}
