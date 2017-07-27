package types

import (
	"time"
)

type AccountReconcileModel struct {
	LastUpdate              time.Time `xmlrpc:"__last_update"`
	AccountId               Many2One  `xmlrpc:"account_id"`
	Amount                  float64   `xmlrpc:"amount"`
	AmountType              string    `xmlrpc:"amount_type"`
	AnalyticAccountId       Many2One  `xmlrpc:"analytic_account_id"`
	CompanyId               Many2One  `xmlrpc:"company_id"`
	CreateDate              time.Time `xmlrpc:"create_date"`
	CreateUid               Many2One  `xmlrpc:"create_uid"`
	DisplayName             string    `xmlrpc:"display_name"`
	HasSecondLine           bool      `xmlrpc:"has_second_line"`
	Id                      int64     `xmlrpc:"id"`
	JournalId               Many2One  `xmlrpc:"journal_id"`
	Label                   string    `xmlrpc:"label"`
	Name                    string    `xmlrpc:"name"`
	SecondAccountId         Many2One  `xmlrpc:"second_account_id"`
	SecondAmount            float64   `xmlrpc:"second_amount"`
	SecondAmountType        string    `xmlrpc:"second_amount_type"`
	SecondAnalyticAccountId Many2One  `xmlrpc:"second_analytic_account_id"`
	SecondJournalId         Many2One  `xmlrpc:"second_journal_id"`
	SecondLabel             string    `xmlrpc:"second_label"`
	SecondTaxId             Many2One  `xmlrpc:"second_tax_id"`
	Sequence                int64     `xmlrpc:"sequence"`
	TaxId                   Many2One  `xmlrpc:"tax_id"`
	WriteDate               time.Time `xmlrpc:"write_date"`
	WriteUid                Many2One  `xmlrpc:"write_uid"`
}

type AccountReconcileModelNil struct {
	LastUpdate              interface{} `xmlrpc:"__last_update"`
	AccountId               interface{} `xmlrpc:"account_id"`
	Amount                  interface{} `xmlrpc:"amount"`
	AmountType              interface{} `xmlrpc:"amount_type"`
	AnalyticAccountId       interface{} `xmlrpc:"analytic_account_id"`
	CompanyId               interface{} `xmlrpc:"company_id"`
	CreateDate              interface{} `xmlrpc:"create_date"`
	CreateUid               interface{} `xmlrpc:"create_uid"`
	DisplayName             interface{} `xmlrpc:"display_name"`
	HasSecondLine           bool        `xmlrpc:"has_second_line"`
	Id                      interface{} `xmlrpc:"id"`
	JournalId               interface{} `xmlrpc:"journal_id"`
	Label                   interface{} `xmlrpc:"label"`
	Name                    interface{} `xmlrpc:"name"`
	SecondAccountId         interface{} `xmlrpc:"second_account_id"`
	SecondAmount            interface{} `xmlrpc:"second_amount"`
	SecondAmountType        interface{} `xmlrpc:"second_amount_type"`
	SecondAnalyticAccountId interface{} `xmlrpc:"second_analytic_account_id"`
	SecondJournalId         interface{} `xmlrpc:"second_journal_id"`
	SecondLabel             interface{} `xmlrpc:"second_label"`
	SecondTaxId             interface{} `xmlrpc:"second_tax_id"`
	Sequence                interface{} `xmlrpc:"sequence"`
	TaxId                   interface{} `xmlrpc:"tax_id"`
	WriteDate               interface{} `xmlrpc:"write_date"`
	WriteUid                interface{} `xmlrpc:"write_uid"`
}

var AccountReconcileModelModel string = "account.reconcile.model"

type AccountReconcileModels []AccountReconcileModel

type AccountReconcileModelsNil []AccountReconcileModelNil

func (s *AccountReconcileModel) NilableType_() interface{} {
	return &AccountReconcileModelNil{}
}

func (ns *AccountReconcileModelNil) Type_() interface{} {
	s := &AccountReconcileModel{}
	return load(ns, s)
}

func (s *AccountReconcileModels) NilableType_() interface{} {
	return &AccountReconcileModelsNil{}
}

func (ns *AccountReconcileModelsNil) Type_() interface{} {
	s := &AccountReconcileModels{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountReconcileModel))
	}
	return s
}
