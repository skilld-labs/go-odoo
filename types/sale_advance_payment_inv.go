package types

import (
	"time"
)

type SaleAdvancePaymentInv struct {
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	AdvancePaymentMethod string    `xmlrpc:"advance_payment_method"`
	Amount               float64   `xmlrpc:"amount"`
	Count                int64     `xmlrpc:"count"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	DepositAccountId     Many2One  `xmlrpc:"deposit_account_id"`
	DepositTaxesId       []int64   `xmlrpc:"deposit_taxes_id"`
	DisplayName          string    `xmlrpc:"display_name"`
	Id                   int64     `xmlrpc:"id"`
	ProductId            Many2One  `xmlrpc:"product_id"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
}

type SaleAdvancePaymentInvNil struct {
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	AdvancePaymentMethod interface{} `xmlrpc:"advance_payment_method"`
	Amount               interface{} `xmlrpc:"amount"`
	Count                interface{} `xmlrpc:"count"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	DepositAccountId     interface{} `xmlrpc:"deposit_account_id"`
	DepositTaxesId       interface{} `xmlrpc:"deposit_taxes_id"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	Id                   interface{} `xmlrpc:"id"`
	ProductId            interface{} `xmlrpc:"product_id"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
}

var SaleAdvancePaymentInvModel string = "sale.advance.payment.inv"

type SaleAdvancePaymentInvs []SaleAdvancePaymentInv

type SaleAdvancePaymentInvsNil []SaleAdvancePaymentInvNil

func (s *SaleAdvancePaymentInv) NilableType_() interface{} {
	return &SaleAdvancePaymentInvNil{}
}

func (ns *SaleAdvancePaymentInvNil) Type_() interface{} {
	s := &SaleAdvancePaymentInv{}
	return load(ns, s)
}

func (s *SaleAdvancePaymentInvs) NilableType_() interface{} {
	return &SaleAdvancePaymentInvsNil{}
}

func (ns *SaleAdvancePaymentInvsNil) Type_() interface{} {
	s := &SaleAdvancePaymentInvs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SaleAdvancePaymentInv))
	}
	return s
}
