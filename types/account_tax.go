package types

import (
	"time"
)

type AccountTax struct {
	AccountId          Many2One    `xmlrpc:"account_id"`
	Active             bool        `xmlrpc:"active"`
	Amount             float64     `xmlrpc:"amount"`
	AmountType         interface{} `xmlrpc:"amount_type"`
	Analytic           bool        `xmlrpc:"analytic"`
	CashBasisAccount   Many2One    `xmlrpc:"cash_basis_account"`
	ChildrenTaxIds     []int64     `xmlrpc:"children_tax_ids"`
	CompanyId          Many2One    `xmlrpc:"company_id"`
	CreateDate         time.Time   `xmlrpc:"create_date"`
	CreateUid          Many2One    `xmlrpc:"create_uid"`
	Description        string      `xmlrpc:"description"`
	DisplayName        string      `xmlrpc:"display_name"`
	HideTaxExigibility bool        `xmlrpc:"hide_tax_exigibility"`
	Id                 int64       `xmlrpc:"id"`
	IncludeBaseAmount  bool        `xmlrpc:"include_base_amount"`
	LastUpdate         time.Time   `xmlrpc:"__last_update"`
	Name               string      `xmlrpc:"name"`
	PriceInclude       bool        `xmlrpc:"price_include"`
	RefundAccountId    Many2One    `xmlrpc:"refund_account_id"`
	Sequence           int64       `xmlrpc:"sequence"`
	TagIds             []int64     `xmlrpc:"tag_ids"`
	TaxAdjustment      bool        `xmlrpc:"tax_adjustment"`
	TaxExigibility     interface{} `xmlrpc:"tax_exigibility"`
	TaxGroupId         Many2One    `xmlrpc:"tax_group_id"`
	TypeTaxUse         interface{} `xmlrpc:"type_tax_use"`
	WriteDate          time.Time   `xmlrpc:"write_date"`
	WriteUid           Many2One    `xmlrpc:"write_uid"`
}

type AccountTaxNil struct {
	AccountId          interface{} `xmlrpc:"account_id"`
	Active             bool        `xmlrpc:"active"`
	Amount             interface{} `xmlrpc:"amount"`
	AmountType         interface{} `xmlrpc:"amount_type"`
	Analytic           bool        `xmlrpc:"analytic"`
	CashBasisAccount   interface{} `xmlrpc:"cash_basis_account"`
	ChildrenTaxIds     interface{} `xmlrpc:"children_tax_ids"`
	CompanyId          interface{} `xmlrpc:"company_id"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	Description        interface{} `xmlrpc:"description"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	HideTaxExigibility bool        `xmlrpc:"hide_tax_exigibility"`
	Id                 interface{} `xmlrpc:"id"`
	IncludeBaseAmount  bool        `xmlrpc:"include_base_amount"`
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	Name               interface{} `xmlrpc:"name"`
	PriceInclude       bool        `xmlrpc:"price_include"`
	RefundAccountId    interface{} `xmlrpc:"refund_account_id"`
	Sequence           interface{} `xmlrpc:"sequence"`
	TagIds             interface{} `xmlrpc:"tag_ids"`
	TaxAdjustment      bool        `xmlrpc:"tax_adjustment"`
	TaxExigibility     interface{} `xmlrpc:"tax_exigibility"`
	TaxGroupId         interface{} `xmlrpc:"tax_group_id"`
	TypeTaxUse         interface{} `xmlrpc:"type_tax_use"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var AccountTaxModel string = "account.tax"

type AccountTaxs []AccountTax

type AccountTaxsNil []AccountTaxNil

func (s *AccountTax) NilableType_() interface{} {
	return &AccountTaxNil{}
}

func (ns *AccountTaxNil) Type_() interface{} {
	s := &AccountTax{}
	return load(ns, s)
}

func (s *AccountTaxs) NilableType_() interface{} {
	return &AccountTaxsNil{}
}

func (ns *AccountTaxsNil) Type_() interface{} {
	s := &AccountTaxs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountTax))
	}
	return s
}
