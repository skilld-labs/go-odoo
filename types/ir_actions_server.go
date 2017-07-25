package types

import (
	"time"
)

type IrActionsServer struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	ActionId            Many2One  `xmlrpc:"action_id"`
	BodyHtml            string    `xmlrpc:"body_html"`
	ChildIds            []int64   `xmlrpc:"child_ids"`
	Code                string    `xmlrpc:"code"`
	Condition           string    `xmlrpc:"condition"`
	Copyvalue           string    `xmlrpc:"copyvalue"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	CreateUid           Many2One  `xmlrpc:"create_uid"`
	CrudModelId         Many2One  `xmlrpc:"crud_model_id"`
	CrudModelName       string    `xmlrpc:"crud_model_name"`
	DisplayName         string    `xmlrpc:"display_name"`
	EmailFrom           string    `xmlrpc:"email_from"`
	EmailTo             string    `xmlrpc:"email_to"`
	FieldsLines         []int64   `xmlrpc:"fields_lines"`
	Help                string    `xmlrpc:"help"`
	Id                  int64     `xmlrpc:"id"`
	IdObject            string    `xmlrpc:"id_object"`
	IdValue             string    `xmlrpc:"id_value"`
	LinkFieldId         Many2One  `xmlrpc:"link_field_id"`
	LinkNewRecord       bool      `xmlrpc:"link_new_record"`
	MenuIrValuesId      Many2One  `xmlrpc:"menu_ir_values_id"`
	ModelId             Many2One  `xmlrpc:"model_id"`
	ModelName           string    `xmlrpc:"model_name"`
	ModelObjectField    Many2One  `xmlrpc:"model_object_field"`
	Name                string    `xmlrpc:"name"`
	PartnerTo           string    `xmlrpc:"partner_to"`
	RefObject           string    `xmlrpc:"ref_object"`
	Sequence            int64     `xmlrpc:"sequence"`
	State               string    `xmlrpc:"state"`
	SubModelObjectField Many2One  `xmlrpc:"sub_model_object_field"`
	SubObject           Many2One  `xmlrpc:"sub_object"`
	Subject             string    `xmlrpc:"subject"`
	TemplateId          Many2One  `xmlrpc:"template_id"`
	Type                string    `xmlrpc:"type"`
	Usage               string    `xmlrpc:"usage"`
	UseCreate           string    `xmlrpc:"use_create"`
	UseRelationalModel  string    `xmlrpc:"use_relational_model"`
	UseWrite            string    `xmlrpc:"use_write"`
	WkfFieldId          Many2One  `xmlrpc:"wkf_field_id"`
	WkfModelId          Many2One  `xmlrpc:"wkf_model_id"`
	WkfModelName        string    `xmlrpc:"wkf_model_name"`
	WkfTransitionId     Many2One  `xmlrpc:"wkf_transition_id"`
	WriteDate           time.Time `xmlrpc:"write_date"`
	WriteExpression     string    `xmlrpc:"write_expression"`
	WriteUid            Many2One  `xmlrpc:"write_uid"`
	XmlId               string    `xmlrpc:"xml_id"`
}

type IrActionsServerNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	ActionId            interface{} `xmlrpc:"action_id"`
	BodyHtml            interface{} `xmlrpc:"body_html"`
	ChildIds            interface{} `xmlrpc:"child_ids"`
	Code                interface{} `xmlrpc:"code"`
	Condition           interface{} `xmlrpc:"condition"`
	Copyvalue           interface{} `xmlrpc:"copyvalue"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	CreateUid           interface{} `xmlrpc:"create_uid"`
	CrudModelId         interface{} `xmlrpc:"crud_model_id"`
	CrudModelName       interface{} `xmlrpc:"crud_model_name"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	EmailFrom           interface{} `xmlrpc:"email_from"`
	EmailTo             interface{} `xmlrpc:"email_to"`
	FieldsLines         interface{} `xmlrpc:"fields_lines"`
	Help                interface{} `xmlrpc:"help"`
	Id                  interface{} `xmlrpc:"id"`
	IdObject            interface{} `xmlrpc:"id_object"`
	IdValue             interface{} `xmlrpc:"id_value"`
	LinkFieldId         interface{} `xmlrpc:"link_field_id"`
	LinkNewRecord       bool        `xmlrpc:"link_new_record"`
	MenuIrValuesId      interface{} `xmlrpc:"menu_ir_values_id"`
	ModelId             interface{} `xmlrpc:"model_id"`
	ModelName           interface{} `xmlrpc:"model_name"`
	ModelObjectField    interface{} `xmlrpc:"model_object_field"`
	Name                interface{} `xmlrpc:"name"`
	PartnerTo           interface{} `xmlrpc:"partner_to"`
	RefObject           interface{} `xmlrpc:"ref_object"`
	Sequence            interface{} `xmlrpc:"sequence"`
	State               interface{} `xmlrpc:"state"`
	SubModelObjectField interface{} `xmlrpc:"sub_model_object_field"`
	SubObject           interface{} `xmlrpc:"sub_object"`
	Subject             interface{} `xmlrpc:"subject"`
	TemplateId          interface{} `xmlrpc:"template_id"`
	Type                interface{} `xmlrpc:"type"`
	Usage               interface{} `xmlrpc:"usage"`
	UseCreate           interface{} `xmlrpc:"use_create"`
	UseRelationalModel  interface{} `xmlrpc:"use_relational_model"`
	UseWrite            interface{} `xmlrpc:"use_write"`
	WkfFieldId          interface{} `xmlrpc:"wkf_field_id"`
	WkfModelId          interface{} `xmlrpc:"wkf_model_id"`
	WkfModelName        interface{} `xmlrpc:"wkf_model_name"`
	WkfTransitionId     interface{} `xmlrpc:"wkf_transition_id"`
	WriteDate           interface{} `xmlrpc:"write_date"`
	WriteExpression     interface{} `xmlrpc:"write_expression"`
	WriteUid            interface{} `xmlrpc:"write_uid"`
	XmlId               interface{} `xmlrpc:"xml_id"`
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
