package types

import (
	"time"
)

type IrActionsClient struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Context     string    `xmlrpc:"context"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Help        string    `xmlrpc:"help"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Params      string    `xmlrpc:"params"`
	ParamsStore string    `xmlrpc:"params_store"`
	ResModel    string    `xmlrpc:"res_model"`
	Tag         string    `xmlrpc:"tag"`
	Target      string    `xmlrpc:"target"`
	Type        string    `xmlrpc:"type"`
	Usage       string    `xmlrpc:"usage"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
	XmlId       string    `xmlrpc:"xml_id"`
}

type IrActionsClientNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Context     interface{} `xmlrpc:"context"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Help        interface{} `xmlrpc:"help"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Params      interface{} `xmlrpc:"params"`
	ParamsStore interface{} `xmlrpc:"params_store"`
	ResModel    interface{} `xmlrpc:"res_model"`
	Tag         interface{} `xmlrpc:"tag"`
	Target      interface{} `xmlrpc:"target"`
	Type        interface{} `xmlrpc:"type"`
	Usage       interface{} `xmlrpc:"usage"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
	XmlId       interface{} `xmlrpc:"xml_id"`
}

var IrActionsClientModel string = "ir.actions.client"

type IrActionsClients []IrActionsClient

type IrActionsClientsNil []IrActionsClientNil

func (s *IrActionsClient) NilableType_() interface{} {
	return &IrActionsClientNil{}
}

func (ns *IrActionsClientNil) Type_() interface{} {
	s := &IrActionsClient{}
	return load(ns, s)
}

func (s *IrActionsClients) NilableType_() interface{} {
	return &IrActionsClientsNil{}
}

func (ns *IrActionsClientsNil) Type_() interface{} {
	s := &IrActionsClients{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsClient))
	}
	return s
}
