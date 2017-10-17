package types

import (
	"time"
)

type IrActionsActUrl struct {
	BindingModelId Many2One  `xmlrpc:"binding_model_id"`
	BindingType    string    `xmlrpc:"binding_type"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Help           string    `xmlrpc:"help"`
	Id             int64     `xmlrpc:"id"`
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Name           string    `xmlrpc:"name"`
	Target         string    `xmlrpc:"target"`
	Type           string    `xmlrpc:"type"`
	Url            string    `xmlrpc:"url"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
	XmlId          string    `xmlrpc:"xml_id"`
}

type IrActionsActUrlNil struct {
	BindingModelId interface{} `xmlrpc:"binding_model_id"`
	BindingType    interface{} `xmlrpc:"binding_type"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Help           interface{} `xmlrpc:"help"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Name           interface{} `xmlrpc:"name"`
	Target         interface{} `xmlrpc:"target"`
	Type           interface{} `xmlrpc:"type"`
	Url            interface{} `xmlrpc:"url"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
	XmlId          interface{} `xmlrpc:"xml_id"`
}

var IrActionsActUrlModel string = "ir.actions.act_url"

type IrActionsActUrls []IrActionsActUrl

type IrActionsActUrlsNil []IrActionsActUrlNil

func (s *IrActionsActUrl) NilableType_() interface{} {
	return &IrActionsActUrlNil{}
}

func (ns *IrActionsActUrlNil) Type_() interface{} {
	s := &IrActionsActUrl{}
	return load(ns, s)
}

func (s *IrActionsActUrls) NilableType_() interface{} {
	return &IrActionsActUrlsNil{}
}

func (ns *IrActionsActUrlsNil) Type_() interface{} {
	s := &IrActionsActUrls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsActUrl))
	}
	return s
}
