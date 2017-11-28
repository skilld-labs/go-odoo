package types

import (
	"time"
)

type IrActionsServer struct {
	BindingModelId Many2One    `xmlrpc:"binding_model_id"`
	BindingType    interface{} `xmlrpc:"binding_type"`
	ChannelIds     []int64     `xmlrpc:"channel_ids"`
	ChildIds       []int64     `xmlrpc:"child_ids"`
	Code           string      `xmlrpc:"code"`
	CreateDate     time.Time   `xmlrpc:"create_date"`
	CreateUid      Many2One    `xmlrpc:"create_uid"`
	CrudModelId    Many2One    `xmlrpc:"crud_model_id"`
	CrudModelName  string      `xmlrpc:"crud_model_name"`
	DisplayName    string      `xmlrpc:"display_name"`
	FieldsLines    []int64     `xmlrpc:"fields_lines"`
	Help           string      `xmlrpc:"help"`
	Id             int64       `xmlrpc:"id"`
	LastUpdate     time.Time   `xmlrpc:"__last_update"`
	LinkFieldId    Many2One    `xmlrpc:"link_field_id"`
	ModelId        Many2One    `xmlrpc:"model_id"`
	ModelName      string      `xmlrpc:"model_name"`
	Name           string      `xmlrpc:"name"`
	PartnerIds     []int64     `xmlrpc:"partner_ids"`
	Sequence       int64       `xmlrpc:"sequence"`
	State          interface{} `xmlrpc:"state"`
	TemplateId     Many2One    `xmlrpc:"template_id"`
	Type           string      `xmlrpc:"type"`
	Usage          interface{} `xmlrpc:"usage"`
	WriteDate      time.Time   `xmlrpc:"write_date"`
	WriteUid       Many2One    `xmlrpc:"write_uid"`
	XmlId          string      `xmlrpc:"xml_id"`
}

type IrActionsServerNil struct {
	BindingModelId interface{} `xmlrpc:"binding_model_id"`
	BindingType    interface{} `xmlrpc:"binding_type"`
	ChannelIds     interface{} `xmlrpc:"channel_ids"`
	ChildIds       interface{} `xmlrpc:"child_ids"`
	Code           interface{} `xmlrpc:"code"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	CrudModelId    interface{} `xmlrpc:"crud_model_id"`
	CrudModelName  interface{} `xmlrpc:"crud_model_name"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	FieldsLines    interface{} `xmlrpc:"fields_lines"`
	Help           interface{} `xmlrpc:"help"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	LinkFieldId    interface{} `xmlrpc:"link_field_id"`
	ModelId        interface{} `xmlrpc:"model_id"`
	ModelName      interface{} `xmlrpc:"model_name"`
	Name           interface{} `xmlrpc:"name"`
	PartnerIds     interface{} `xmlrpc:"partner_ids"`
	Sequence       interface{} `xmlrpc:"sequence"`
	State          interface{} `xmlrpc:"state"`
	TemplateId     interface{} `xmlrpc:"template_id"`
	Type           interface{} `xmlrpc:"type"`
	Usage          interface{} `xmlrpc:"usage"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
	XmlId          interface{} `xmlrpc:"xml_id"`
}

var IrActionsServerModel string = "ir.actions.server"

type IrActionsServers []IrActionsServer

type IrActionsServersNil []IrActionsServerNil

func (s *IrActionsServer) NilableType_() interface{} {
	return &IrActionsServerNil{}
}

func (ns *IrActionsServerNil) Type_() interface{} {
	s := &IrActionsServer{}
	return load(ns, s)
}

func (s *IrActionsServers) NilableType_() interface{} {
	return &IrActionsServersNil{}
}

func (ns *IrActionsServersNil) Type_() interface{} {
	s := &IrActionsServers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsServer))
	}
	return s
}
