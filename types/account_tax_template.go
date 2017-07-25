package types

import (
	"time"
)

type AccountTaxTemplate struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	AccountId         Many2One  `xmlrpc:"account_id"`
	Active            bool      `xmlrpc:"active"`
	Amount            float64   `xmlrpc:"amount"`
	AmountType        string    `xmlrpc:"amount_type"`
	Analytic          bool      `xmlrpc:"analytic"`
	ChartTemplateId   Many2One  `xmlrpc:"chart_template_id"`
	ChildrenTaxIds    []int64   `xmlrpc:"children_tax_ids"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	Description       string    `xmlrpc:"description"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	IncludeBaseAmount bool      `xmlrpc:"include_base_amount"`
	Name              string    `xmlrpc:"name"`
	PriceInclude      bool      `xmlrpc:"price_include"`
	RefundAccountId   Many2One  `xmlrpc:"refund_account_id"`
	Sequence          int64     `xmlrpc:"sequence"`
	TagIds            []int64   `xmlrpc:"tag_ids"`
	TaxAdjustment     bool      `xmlrpc:"tax_adjustment"`
	TaxGroupId        Many2One  `xmlrpc:"tax_group_id"`
	TypeTaxUse        string    `xmlrpc:"type_tax_use"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type AccountTaxTemplateNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	AccountId         interface{} `xmlrpc:"account_id"`
	Active            bool        `xmlrpc:"active"`
	Amount            interface{} `xmlrpc:"amount"`
	AmountType        interface{} `xmlrpc:"amount_type"`
	Analytic          bool        `xmlrpc:"analytic"`
	ChartTemplateId   interface{} `xmlrpc:"chart_template_id"`
	ChildrenTaxIds    interface{} `xmlrpc:"children_tax_ids"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	Description       interface{} `xmlrpc:"description"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	IncludeBaseAmount bool        `xmlrpc:"include_base_amount"`
	Name              interface{} `xmlrpc:"name"`
	PriceInclude      bool        `xmlrpc:"price_include"`
	RefundAccountId   interface{} `xmlrpc:"refund_account_id"`
	Sequence          interface{} `xmlrpc:"sequence"`
	TagIds            interface{} `xmlrpc:"tag_ids"`
	TaxAdjustment     bool        `xmlrpc:"tax_adjustment"`
	TaxGroupId        interface{} `xmlrpc:"tax_group_id"`
	TypeTaxUse        interface{} `xmlrpc:"type_tax_use"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var AccountTaxTemplateModel string = "account.tax.template"

type AccountTaxTemplates []AccountTaxTemplate

type AccountTaxTemplatesNil []AccountTaxTemplateNil

func (s *AccountTaxTemplate) NilableType_() interface{} {
	return &AccountTaxTemplateNil{}
}

func (ns *AccountTaxTemplateNil) Type_() interface{} {
	s := &AccountTaxTemplate{}
	return load(ns, s)
}

func (s *AccountTaxTemplates) NilableType_() interface{} {
	return &AccountTaxTemplatesNil{}
}

func (ns *AccountTaxTemplatesNil) Type_() interface{} {
	s := &AccountTaxTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountTaxTemplate))
	}
	return s
}
