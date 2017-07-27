package types

import (
	"time"
)

type AccountInvoiceRefund struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	Date         time.Time `xmlrpc:"date"`
	DateInvoice  time.Time `xmlrpc:"date_invoice"`
	Description  string    `xmlrpc:"description"`
	DisplayName  string    `xmlrpc:"display_name"`
	FilterRefund string    `xmlrpc:"filter_refund"`
	Id           int64     `xmlrpc:"id"`
	RefundOnly   bool      `xmlrpc:"refund_only"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type AccountInvoiceRefundNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	Date         interface{} `xmlrpc:"date"`
	DateInvoice  interface{} `xmlrpc:"date_invoice"`
	Description  interface{} `xmlrpc:"description"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	FilterRefund interface{} `xmlrpc:"filter_refund"`
	Id           interface{} `xmlrpc:"id"`
	RefundOnly   bool        `xmlrpc:"refund_only"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var AccountInvoiceRefundModel string = "account.invoice.refund"

type AccountInvoiceRefunds []AccountInvoiceRefund

type AccountInvoiceRefundsNil []AccountInvoiceRefundNil

func (s *AccountInvoiceRefund) NilableType_() interface{} {
	return &AccountInvoiceRefundNil{}
}

func (ns *AccountInvoiceRefundNil) Type_() interface{} {
	s := &AccountInvoiceRefund{}
	return load(ns, s)
}

func (s *AccountInvoiceRefunds) NilableType_() interface{} {
	return &AccountInvoiceRefundsNil{}
}

func (ns *AccountInvoiceRefundsNil) Type_() interface{} {
	s := &AccountInvoiceRefunds{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoiceRefund))
	}
	return s
}
