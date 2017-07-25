package types

import (
	"time"
)

type AccountPartialReconcile struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	Amount            float64   `xmlrpc:"amount"`
	AmountCurrency    float64   `xmlrpc:"amount_currency"`
	CompanyCurrencyId Many2One  `xmlrpc:"company_currency_id"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CreditMoveId      Many2One  `xmlrpc:"credit_move_id"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	DebitMoveId       Many2One  `xmlrpc:"debit_move_id"`
	DisplayName       string    `xmlrpc:"display_name"`
	FullReconcileId   Many2One  `xmlrpc:"full_reconcile_id"`
	Id                int64     `xmlrpc:"id"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type AccountPartialReconcileNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	Amount            interface{} `xmlrpc:"amount"`
	AmountCurrency    interface{} `xmlrpc:"amount_currency"`
	CompanyCurrencyId interface{} `xmlrpc:"company_currency_id"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CreditMoveId      interface{} `xmlrpc:"credit_move_id"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	DebitMoveId       interface{} `xmlrpc:"debit_move_id"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	FullReconcileId   interface{} `xmlrpc:"full_reconcile_id"`
	Id                interface{} `xmlrpc:"id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var AccountPartialReconcileModel string = "account.partial.reconcile"

type AccountPartialReconciles []AccountPartialReconcile

type AccountPartialReconcilesNil []AccountPartialReconcileNil

func (s *AccountPartialReconcile) NilableType_() interface{} {
	return &AccountPartialReconcileNil{}
}

func (ns *AccountPartialReconcileNil) Type_() interface{} {
	s := &AccountPartialReconcile{}
	return load(ns, s)
}

func (s *AccountPartialReconciles) NilableType_() interface{} {
	return &AccountPartialReconcilesNil{}
}

func (ns *AccountPartialReconcilesNil) Type_() interface{} {
	s := &AccountPartialReconciles{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountPartialReconcile))
	}
	return s
}
