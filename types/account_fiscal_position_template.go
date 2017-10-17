package types

import (
	"time"
)

type AccountFiscalPositionTemplate struct {
	AccountIds      []int64   `xmlrpc:"account_ids"`
	AutoApply       bool      `xmlrpc:"auto_apply"`
	ChartTemplateId Many2One  `xmlrpc:"chart_template_id"`
	CountryGroupId  Many2One  `xmlrpc:"country_group_id"`
	CountryId       Many2One  `xmlrpc:"country_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	Name            string    `xmlrpc:"name"`
	Note            string    `xmlrpc:"note"`
	Sequence        int64     `xmlrpc:"sequence"`
	StateIds        []int64   `xmlrpc:"state_ids"`
	TaxIds          []int64   `xmlrpc:"tax_ids"`
	VatRequired     bool      `xmlrpc:"vat_required"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
	ZipFrom         int64     `xmlrpc:"zip_from"`
	ZipTo           int64     `xmlrpc:"zip_to"`
}

type AccountFiscalPositionTemplateNil struct {
	AccountIds      interface{} `xmlrpc:"account_ids"`
	AutoApply       bool        `xmlrpc:"auto_apply"`
	ChartTemplateId interface{} `xmlrpc:"chart_template_id"`
	CountryGroupId  interface{} `xmlrpc:"country_group_id"`
	CountryId       interface{} `xmlrpc:"country_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Name            interface{} `xmlrpc:"name"`
	Note            interface{} `xmlrpc:"note"`
	Sequence        interface{} `xmlrpc:"sequence"`
	StateIds        interface{} `xmlrpc:"state_ids"`
	TaxIds          interface{} `xmlrpc:"tax_ids"`
	VatRequired     bool        `xmlrpc:"vat_required"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
	ZipFrom         interface{} `xmlrpc:"zip_from"`
	ZipTo           interface{} `xmlrpc:"zip_to"`
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
