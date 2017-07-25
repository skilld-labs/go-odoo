package types

import (
	"time"
)

type AccountPaymentTermLine struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Days        int64     `xmlrpc:"days"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Option      string    `xmlrpc:"option"`
	PaymentId   Many2One  `xmlrpc:"payment_id"`
	Sequence    int64     `xmlrpc:"sequence"`
	Value       string    `xmlrpc:"value"`
	ValueAmount float64   `xmlrpc:"value_amount"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountPaymentTermLineNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Days        interface{} `xmlrpc:"days"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Option      interface{} `xmlrpc:"option"`
	PaymentId   interface{} `xmlrpc:"payment_id"`
	Sequence    interface{} `xmlrpc:"sequence"`
	Value       interface{} `xmlrpc:"value"`
	ValueAmount interface{} `xmlrpc:"value_amount"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountPaymentTermLineModel string = "account.payment.term.line"

type AccountPaymentTermLines []AccountPaymentTermLine

type AccountPaymentTermLinesNil []AccountPaymentTermLineNil

func (s *AccountPaymentTermLine) NilableType_() interface{} {
	return &AccountPaymentTermLineNil{}
}

func (ns *AccountPaymentTermLineNil) Type_() interface{} {
	s := &AccountPaymentTermLine{}
	return load(ns, s)
}

func (s *AccountPaymentTermLines) NilableType_() interface{} {
	return &AccountPaymentTermLinesNil{}
}

func (ns *AccountPaymentTermLinesNil) Type_() interface{} {
	s := &AccountPaymentTermLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountPaymentTermLine))
	}
	return s
}
