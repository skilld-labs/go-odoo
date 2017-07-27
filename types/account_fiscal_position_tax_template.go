package types

import (
	"time"
)

type AccountFiscalPositionTaxTemplate struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PositionId  Many2One  `xmlrpc:"position_id"`
	TaxDestId   Many2One  `xmlrpc:"tax_dest_id"`
	TaxSrcId    Many2One  `xmlrpc:"tax_src_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountFiscalPositionTaxTemplateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PositionId  interface{} `xmlrpc:"position_id"`
	TaxDestId   interface{} `xmlrpc:"tax_dest_id"`
	TaxSrcId    interface{} `xmlrpc:"tax_src_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountFiscalPositionTaxTemplateModel string = "account.fiscal.position.tax.template"

type AccountFiscalPositionTaxTemplates []AccountFiscalPositionTaxTemplate

type AccountFiscalPositionTaxTemplatesNil []AccountFiscalPositionTaxTemplateNil

func (s *AccountFiscalPositionTaxTemplate) NilableType_() interface{} {
	return &AccountFiscalPositionTaxTemplateNil{}
}

func (ns *AccountFiscalPositionTaxTemplateNil) Type_() interface{} {
	s := &AccountFiscalPositionTaxTemplate{}
	return load(ns, s)
}

func (s *AccountFiscalPositionTaxTemplates) NilableType_() interface{} {
	return &AccountFiscalPositionTaxTemplatesNil{}
}

func (ns *AccountFiscalPositionTaxTemplatesNil) Type_() interface{} {
	s := &AccountFiscalPositionTaxTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFiscalPositionTaxTemplate))
	}
	return s
}
