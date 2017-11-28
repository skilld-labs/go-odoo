package types

import (
	"time"
)

type IrActionsTodo struct {
	ActionId    Many2One    `xmlrpc:"action_id"`
	CreateDate  time.Time   `xmlrpc:"create_date"`
	CreateUid   Many2One    `xmlrpc:"create_uid"`
	DisplayName string      `xmlrpc:"display_name"`
	Id          int64       `xmlrpc:"id"`
	LastUpdate  time.Time   `xmlrpc:"__last_update"`
	Name        string      `xmlrpc:"name"`
	Sequence    int64       `xmlrpc:"sequence"`
	State       interface{} `xmlrpc:"state"`
	WriteDate   time.Time   `xmlrpc:"write_date"`
	WriteUid    Many2One    `xmlrpc:"write_uid"`
}

type IrActionsTodoNil struct {
	ActionId    interface{} `xmlrpc:"action_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	Sequence    interface{} `xmlrpc:"sequence"`
	State       interface{} `xmlrpc:"state"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrActionsTodoModel string = "ir.actions.todo"

type IrActionsTodos []IrActionsTodo

type IrActionsTodosNil []IrActionsTodoNil

func (s *IrActionsTodo) NilableType_() interface{} {
	return &IrActionsTodoNil{}
}

func (ns *IrActionsTodoNil) Type_() interface{} {
	s := &IrActionsTodo{}
	return load(ns, s)
}

func (s *IrActionsTodos) NilableType_() interface{} {
	return &IrActionsTodosNil{}
}

func (ns *IrActionsTodosNil) Type_() interface{} {
	s := &IrActionsTodos{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsTodo))
	}
	return s
}
