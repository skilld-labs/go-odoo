package types

import (
	"time"
)

type IrActionsTodo struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ActionId    Many2One  `xmlrpc:"action_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	GroupsId    []int64   `xmlrpc:"groups_id"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Note        string    `xmlrpc:"note"`
	Sequence    int64     `xmlrpc:"sequence"`
	State       string    `xmlrpc:"state"`
	Type        string    `xmlrpc:"type"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrActionsTodoNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ActionId    interface{} `xmlrpc:"action_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	GroupsId    interface{} `xmlrpc:"groups_id"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Note        interface{} `xmlrpc:"note"`
	Sequence    interface{} `xmlrpc:"sequence"`
	State       interface{} `xmlrpc:"state"`
	Type        interface{} `xmlrpc:"type"`
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
