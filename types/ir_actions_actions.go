package types

import (
	"time"
)

type IrActionsActions struct {
	BindingModelId Many2One    `xmlrpc:"binding_model_id"`
	BindingType    interface{} `xmlrpc:"binding_type"`
	CreateDate     time.Time   `xmlrpc:"create_date"`
	CreateUid      Many2One    `xmlrpc:"create_uid"`
	DisplayName    string      `xmlrpc:"display_name"`
	Help           string      `xmlrpc:"help"`
	Id             int64       `xmlrpc:"id"`
	LastUpdate     time.Time   `xmlrpc:"__last_update"`
	Name           string      `xmlrpc:"name"`
	Type           string      `xmlrpc:"type"`
	WriteDate      time.Time   `xmlrpc:"write_date"`
	WriteUid       Many2One    `xmlrpc:"write_uid"`
	XmlId          string      `xmlrpc:"xml_id"`
}

type IrActionsActionsNil struct {
	BindingModelId interface{} `xmlrpc:"binding_model_id"`
	BindingType    interface{} `xmlrpc:"binding_type"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Help           interface{} `xmlrpc:"help"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Name           interface{} `xmlrpc:"name"`
	Type           interface{} `xmlrpc:"type"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
	XmlId          interface{} `xmlrpc:"xml_id"`
}

var IrActionsActionsModel string = "ir.actions.actions"

type IrActionsActionss []IrActionsActions

type IrActionsActionssNil []IrActionsActionsNil

func (s *IrActionsActions) NilableType_() interface{} {
	return &IrActionsActionsNil{}
}

func (ns *IrActionsActionsNil) Type_() interface{} {
	s := &IrActionsActions{}
	return load(ns, s)
}

func (s *IrActionsActionss) NilableType_() interface{} {
	return &IrActionsActionssNil{}
}

func (ns *IrActionsActionssNil) Type_() interface{} {
	s := &IrActionsActionss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsActions))
	}
	return s
}
