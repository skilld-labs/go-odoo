package types

import (
	"time"
)

type AccountReconcileModelTemplate struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	AccountId        Many2One  `xmlrpc:"account_id"`
	Amount           float64   `xmlrpc:"amount"`
	AmountType       string    `xmlrpc:"amount_type"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DisplayName      string    `xmlrpc:"display_name"`
	HasSecondLine    bool      `xmlrpc:"has_second_line"`
	Id               int64     `xmlrpc:"id"`
	Label            string    `xmlrpc:"label"`
	Name             string    `xmlrpc:"name"`
	SecondAccountId  Many2One  `xmlrpc:"second_account_id"`
	SecondAmount     float64   `xmlrpc:"second_amount"`
	SecondAmountType string    `xmlrpc:"second_amount_type"`
	SecondLabel      string    `xmlrpc:"second_label"`
	SecondTaxId      Many2One  `xmlrpc:"second_tax_id"`
	Sequence         int64     `xmlrpc:"sequence"`
	TaxId            Many2One  `xmlrpc:"tax_id"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type AccountReconcileModelTemplateNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	AccountId        interface{} `xmlrpc:"account_id"`
	Amount           interface{} `xmlrpc:"amount"`
	AmountType       interface{} `xmlrpc:"amount_type"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	HasSecondLine    bool        `xmlrpc:"has_second_line"`
	Id               interface{} `xmlrpc:"id"`
	Label            interface{} `xmlrpc:"label"`
	Name             interface{} `xmlrpc:"name"`
	SecondAccountId  interface{} `xmlrpc:"second_account_id"`
	SecondAmount     interface{} `xmlrpc:"second_amount"`
	SecondAmountType interface{} `xmlrpc:"second_amount_type"`
	SecondLabel      interface{} `xmlrpc:"second_label"`
	SecondTaxId      interface{} `xmlrpc:"second_tax_id"`
	Sequence         interface{} `xmlrpc:"sequence"`
	TaxId            interface{} `xmlrpc:"tax_id"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var AccountReconcileModelTemplateModel string = "account.reconcile.model.template"

type AccountReconcileModelTemplates []AccountReconcileModelTemplate

type AccountReconcileModelTemplatesNil []AccountReconcileModelTemplateNil

func (s *AccountReconcileModelTemplate) NilableType_() interface{} {
	return &AccountReconcileModelTemplateNil{}
}

func (ns *AccountReconcileModelTemplateNil) Type_() interface{} {
	s := &AccountReconcileModelTemplate{}
	return load(ns, s)
}

func (s *AccountReconcileModelTemplates) NilableType_() interface{} {
	return &AccountReconcileModelTemplatesNil{}
}

func (ns *AccountReconcileModelTemplatesNil) Type_() interface{} {
	s := &AccountReconcileModelTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountReconcileModelTemplate))
	}
	return s
}
