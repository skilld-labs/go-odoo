package types

import (
	"time"
)

type AccountInvoiceConfirm struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountInvoiceConfirmNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountInvoiceConfirmModel string = "account.invoice.confirm"

type AccountInvoiceConfirms []AccountInvoiceConfirm

type AccountInvoiceConfirmsNil []AccountInvoiceConfirmNil

func (s *AccountInvoiceConfirm) NilableType_() interface{} {
	return &AccountInvoiceConfirmNil{}
}

func (ns *AccountInvoiceConfirmNil) Type_() interface{} {
	s := &AccountInvoiceConfirm{}
	return load(ns, s)
}

func (s *AccountInvoiceConfirms) NilableType_() interface{} {
	return &AccountInvoiceConfirmsNil{}
}

func (ns *AccountInvoiceConfirmsNil) Type_() interface{} {
	s := &AccountInvoiceConfirms{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoiceConfirm))
	}
	return s
}
