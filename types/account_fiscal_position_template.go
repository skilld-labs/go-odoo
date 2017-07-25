package types

import (
	"time"
)

type AccountFiscalPositionTemplate struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	AccountIds      []int64   `xmlrpc:"account_ids"`
	ChartTemplateId Many2One  `xmlrpc:"chart_template_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	Name            string    `xmlrpc:"name"`
	Note            string    `xmlrpc:"note"`
	TaxIds          []int64   `xmlrpc:"tax_ids"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountFiscalPositionTemplateNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	AccountIds      interface{} `xmlrpc:"account_ids"`
	ChartTemplateId interface{} `xmlrpc:"chart_template_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	Name            interface{} `xmlrpc:"name"`
	Note            interface{} `xmlrpc:"note"`
	TaxIds          interface{} `xmlrpc:"tax_ids"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountFiscalPositionTemplateModel string = "account.fiscal.position.template"

type AccountFiscalPositionTemplates []AccountFiscalPositionTemplate

type AccountFiscalPositionTemplatesNil []AccountFiscalPositionTemplateNil

func (s *AccountFiscalPositionTemplate) NilableType_() interface{} {
	return &AccountFiscalPositionTemplateNil{}
}

func (ns *AccountFiscalPositionTemplateNil) Type_() interface{} {
	s := &AccountFiscalPositionTemplate{}
	return load(ns, s)
}

func (s *AccountFiscalPositionTemplates) NilableType_() interface{} {
	return &AccountFiscalPositionTemplatesNil{}
}

func (ns *AccountFiscalPositionTemplatesNil) Type_() interface{} {
	s := &AccountFiscalPositionTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFiscalPositionTemplate))
	}
	return s
}
