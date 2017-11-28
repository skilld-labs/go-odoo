package types

import (
	"time"
)

type IrCron struct {
	Active            bool        `xmlrpc:"active"`
	BindingModelId    Many2One    `xmlrpc:"binding_model_id"`
	BindingType       interface{} `xmlrpc:"binding_type"`
	ChildIds          []int64     `xmlrpc:"child_ids"`
	Code              string      `xmlrpc:"code"`
	CreateDate        time.Time   `xmlrpc:"create_date"`
	CreateUid         Many2One    `xmlrpc:"create_uid"`
	CronName          string      `xmlrpc:"cron_name"`
	CrudModelId       Many2One    `xmlrpc:"crud_model_id"`
	CrudModelName     string      `xmlrpc:"crud_model_name"`
	DisplayName       string      `xmlrpc:"display_name"`
	Doall             bool        `xmlrpc:"doall"`
	FieldsLines       []int64     `xmlrpc:"fields_lines"`
	Help              string      `xmlrpc:"help"`
	Id                int64       `xmlrpc:"id"`
	IntervalNumber    int64       `xmlrpc:"interval_number"`
	IntervalType      interface{} `xmlrpc:"interval_type"`
	IrActionsServerId Many2One    `xmlrpc:"ir_actions_server_id"`
	LastUpdate        time.Time   `xmlrpc:"__last_update"`
	LinkFieldId       Many2One    `xmlrpc:"link_field_id"`
	ModelId           Many2One    `xmlrpc:"model_id"`
	ModelName         string      `xmlrpc:"model_name"`
	Name              string      `xmlrpc:"name"`
	Nextcall          time.Time   `xmlrpc:"nextcall"`
	Numbercall        int64       `xmlrpc:"numbercall"`
	Priority          int64       `xmlrpc:"priority"`
	Sequence          int64       `xmlrpc:"sequence"`
	State             interface{} `xmlrpc:"state"`
	Type              string      `xmlrpc:"type"`
	Usage             interface{} `xmlrpc:"usage"`
	UserId            Many2One    `xmlrpc:"user_id"`
	WriteDate         time.Time   `xmlrpc:"write_date"`
	WriteUid          Many2One    `xmlrpc:"write_uid"`
	XmlId             string      `xmlrpc:"xml_id"`
}

type IrCronNil struct {
	Active            bool        `xmlrpc:"active"`
	BindingModelId    interface{} `xmlrpc:"binding_model_id"`
	BindingType       interface{} `xmlrpc:"binding_type"`
	ChildIds          interface{} `xmlrpc:"child_ids"`
	Code              interface{} `xmlrpc:"code"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CronName          interface{} `xmlrpc:"cron_name"`
	CrudModelId       interface{} `xmlrpc:"crud_model_id"`
	CrudModelName     interface{} `xmlrpc:"crud_model_name"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Doall             bool        `xmlrpc:"doall"`
	FieldsLines       interface{} `xmlrpc:"fields_lines"`
	Help              interface{} `xmlrpc:"help"`
	Id                interface{} `xmlrpc:"id"`
	IntervalNumber    interface{} `xmlrpc:"interval_number"`
	IntervalType      interface{} `xmlrpc:"interval_type"`
	IrActionsServerId interface{} `xmlrpc:"ir_actions_server_id"`
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	LinkFieldId       interface{} `xmlrpc:"link_field_id"`
	ModelId           interface{} `xmlrpc:"model_id"`
	ModelName         interface{} `xmlrpc:"model_name"`
	Name              interface{} `xmlrpc:"name"`
	Nextcall          interface{} `xmlrpc:"nextcall"`
	Numbercall        interface{} `xmlrpc:"numbercall"`
	Priority          interface{} `xmlrpc:"priority"`
	Sequence          interface{} `xmlrpc:"sequence"`
	State             interface{} `xmlrpc:"state"`
	Type              interface{} `xmlrpc:"type"`
	Usage             interface{} `xmlrpc:"usage"`
	UserId            interface{} `xmlrpc:"user_id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
	XmlId             interface{} `xmlrpc:"xml_id"`
}

var IrCronModel string = "ir.cron"

type IrCrons []IrCron

type IrCronsNil []IrCronNil

func (s *IrCron) NilableType_() interface{} {
	return &IrCronNil{}
}

func (ns *IrCronNil) Type_() interface{} {
	s := &IrCron{}
	return load(ns, s)
}

func (s *IrCrons) NilableType_() interface{} {
	return &IrCronsNil{}
}

func (ns *IrCronsNil) Type_() interface{} {
	s := &IrCrons{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrCron))
	}
	return s
}
