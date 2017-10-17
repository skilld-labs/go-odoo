package types

import (
	"time"
)

type AccountAccountTemplate struct {
	ChartTemplateId Many2One  `xmlrpc:"chart_template_id"`
	Code            string    `xmlrpc:"code"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	CurrencyId      Many2One  `xmlrpc:"currency_id"`
	DisplayName     string    `xmlrpc:"display_name"`
	GroupId         Many2One  `xmlrpc:"group_id"`
	Id              int64     `xmlrpc:"id"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	Name            string    `xmlrpc:"name"`
	Nocreate        bool      `xmlrpc:"nocreate"`
	Note            string    `xmlrpc:"note"`
	Reconcile       bool      `xmlrpc:"reconcile"`
	TagIds          []int64   `xmlrpc:"tag_ids"`
	TaxIds          []int64   `xmlrpc:"tax_ids"`
	UserTypeId      Many2One  `xmlrpc:"user_type_id"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountAccountTemplateNil struct {
	ChartTemplateId interface{} `xmlrpc:"chart_template_id"`
	Code            interface{} `xmlrpc:"code"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	CurrencyId      interface{} `xmlrpc:"currency_id"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	GroupId         interface{} `xmlrpc:"group_id"`
	Id              interface{} `xmlrpc:"id"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Name            interface{} `xmlrpc:"name"`
	Nocreate        bool        `xmlrpc:"nocreate"`
	Note            interface{} `xmlrpc:"note"`
	Reconcile       bool        `xmlrpc:"reconcile"`
	TagIds          interface{} `xmlrpc:"tag_ids"`
	TaxIds          interface{} `xmlrpc:"tax_ids"`
	UserTypeId      interface{} `xmlrpc:"user_type_id"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountAccountTemplateModel string = "account.account.template"

type AccountAccountTemplates []AccountAccountTemplate

type AccountAccountTemplatesNil []AccountAccountTemplateNil

func (s *AccountAccountTemplate) NilableType_() interface{} {
	return &AccountAccountTemplateNil{}
}

func (ns *AccountAccountTemplateNil) Type_() interface{} {
	s := &AccountAccountTemplate{}
	return load(ns, s)
}

func (s *AccountAccountTemplates) NilableType_() interface{} {
	return &AccountAccountTemplatesNil{}
}

func (ns *AccountAccountTemplatesNil) Type_() interface{} {
	s := &AccountAccountTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAccountTemplate))
	}
	return s
}
