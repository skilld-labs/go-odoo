package types

import (
	"time"
)

type AccountInvoiceCancel struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountInvoiceCancelNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountInvoiceCancelModel string = "account.invoice.cancel"

type AccountInvoiceCancels []AccountInvoiceCancel

type AccountInvoiceCancelsNil []AccountInvoiceCancelNil

func (s *AccountInvoiceCancel) NilableType_() interface{} {
	return &AccountInvoiceCancelNil{}
}

func (ns *AccountInvoiceCancelNil) Type_() interface{} {
	s := &AccountInvoiceCancel{}
	return load(ns, s)
}

func (s *AccountInvoiceCancels) NilableType_() interface{} {
	return &AccountInvoiceCancelsNil{}
}

func (ns *AccountInvoiceCancelsNil) Type_() interface{} {
	s := &AccountInvoiceCancels{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoiceCancel))
	}
	return s
}
