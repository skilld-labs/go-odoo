package types

import (
	"time"
)

type AccountInvoiceTax struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	AccountAnalyticId Many2One  `xmlrpc:"account_analytic_id"`
	AccountId         Many2One  `xmlrpc:"account_id"`
	Amount            float64   `xmlrpc:"amount"`
	Base              float64   `xmlrpc:"base"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	InvoiceId         Many2One  `xmlrpc:"invoice_id"`
	Manual            bool      `xmlrpc:"manual"`
	Name              string    `xmlrpc:"name"`
	Sequence          int64     `xmlrpc:"sequence"`
	TaxId             Many2One  `xmlrpc:"tax_id"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type AccountInvoiceTaxNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	AccountAnalyticId interface{} `xmlrpc:"account_analytic_id"`
	AccountId         interface{} `xmlrpc:"account_id"`
	Amount            interface{} `xmlrpc:"amount"`
	Base              interface{} `xmlrpc:"base"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	InvoiceId         interface{} `xmlrpc:"invoice_id"`
	Manual            bool        `xmlrpc:"manual"`
	Name              interface{} `xmlrpc:"name"`
	Sequence          interface{} `xmlrpc:"sequence"`
	TaxId             interface{} `xmlrpc:"tax_id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var AccountInvoiceTaxModel string = "account.invoice.tax"

type AccountInvoiceTaxs []AccountInvoiceTax

type AccountInvoiceTaxsNil []AccountInvoiceTaxNil

func (s *AccountInvoiceTax) NilableType_() interface{} {
	return &AccountInvoiceTaxNil{}
}

func (ns *AccountInvoiceTaxNil) Type_() interface{} {
	s := &AccountInvoiceTax{}
	return load(ns, s)
}

func (s *AccountInvoiceTaxs) NilableType_() interface{} {
	return &AccountInvoiceTaxsNil{}
}

func (ns *AccountInvoiceTaxsNil) Type_() interface{} {
	s := &AccountInvoiceTaxs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoiceTax))
	}
	return s
}
