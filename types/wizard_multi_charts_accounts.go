package types

import (
	"time"
)

type WizardMultiChartsAccounts struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	BankAccountCodePrefix string    `xmlrpc:"bank_account_code_prefix"`
	BankAccountIds        []int64   `xmlrpc:"bank_account_ids"`
	CashAccountCodePrefix string    `xmlrpc:"cash_account_code_prefix"`
	ChartTemplateId       Many2One  `xmlrpc:"chart_template_id"`
	CodeDigits            int64     `xmlrpc:"code_digits"`
	CompanyId             Many2One  `xmlrpc:"company_id"`
	CompleteTaxSet        bool      `xmlrpc:"complete_tax_set"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	CurrencyId            Many2One  `xmlrpc:"currency_id"`
	DisplayName           string    `xmlrpc:"display_name"`
	Id                    int64     `xmlrpc:"id"`
	OnlyOneChartTemplate  bool      `xmlrpc:"only_one_chart_template"`
	PurchaseTaxId         Many2One  `xmlrpc:"purchase_tax_id"`
	PurchaseTaxRate       float64   `xmlrpc:"purchase_tax_rate"`
	SaleTaxId             Many2One  `xmlrpc:"sale_tax_id"`
	SaleTaxRate           float64   `xmlrpc:"sale_tax_rate"`
	TransferAccountId     Many2One  `xmlrpc:"transfer_account_id"`
	UseAngloSaxon         bool      `xmlrpc:"use_anglo_saxon"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type WizardMultiChartsAccountsNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	BankAccountCodePrefix interface{} `xmlrpc:"bank_account_code_prefix"`
	BankAccountIds        interface{} `xmlrpc:"bank_account_ids"`
	CashAccountCodePrefix interface{} `xmlrpc:"cash_account_code_prefix"`
	ChartTemplateId       interface{} `xmlrpc:"chart_template_id"`
	CodeDigits            interface{} `xmlrpc:"code_digits"`
	CompanyId             interface{} `xmlrpc:"company_id"`
	CompleteTaxSet        bool        `xmlrpc:"complete_tax_set"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	CurrencyId            interface{} `xmlrpc:"currency_id"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Id                    interface{} `xmlrpc:"id"`
	OnlyOneChartTemplate  bool        `xmlrpc:"only_one_chart_template"`
	PurchaseTaxId         interface{} `xmlrpc:"purchase_tax_id"`
	PurchaseTaxRate       interface{} `xmlrpc:"purchase_tax_rate"`
	SaleTaxId             interface{} `xmlrpc:"sale_tax_id"`
	SaleTaxRate           interface{} `xmlrpc:"sale_tax_rate"`
	TransferAccountId     interface{} `xmlrpc:"transfer_account_id"`
	UseAngloSaxon         bool        `xmlrpc:"use_anglo_saxon"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var WizardMultiChartsAccountsModel string = "wizard.multi.charts.accounts"

type WizardMultiChartsAccountss []WizardMultiChartsAccounts

type WizardMultiChartsAccountssNil []WizardMultiChartsAccountsNil

func (s *WizardMultiChartsAccounts) NilableType_() interface{} {
	return &WizardMultiChartsAccountsNil{}
}

func (ns *WizardMultiChartsAccountsNil) Type_() interface{} {
	s := &WizardMultiChartsAccounts{}
	return load(ns, s)
}

func (s *WizardMultiChartsAccountss) NilableType_() interface{} {
	return &WizardMultiChartsAccountssNil{}
}

func (ns *WizardMultiChartsAccountssNil) Type_() interface{} {
	s := &WizardMultiChartsAccountss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WizardMultiChartsAccounts))
	}
	return s
}
