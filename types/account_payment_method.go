package types

import (
	"time"
)

type AccountPaymentMethod struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Code        string    `xmlrpc:"code"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	PaymentType string    `xmlrpc:"payment_type"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountPaymentMethodNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Code        interface{} `xmlrpc:"code"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	PaymentType interface{} `xmlrpc:"payment_type"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountPaymentMethodModel string = "account.payment.method"

type AccountPaymentMethods []AccountPaymentMethod

type AccountPaymentMethodsNil []AccountPaymentMethodNil

func (s *AccountPaymentMethod) NilableType_() interface{} {
	return &AccountPaymentMethodNil{}
}

func (ns *AccountPaymentMethodNil) Type_() interface{} {
	s := &AccountPaymentMethod{}
	return load(ns, s)
}

func (s *AccountPaymentMethods) NilableType_() interface{} {
	return &AccountPaymentMethodsNil{}
}

func (ns *AccountPaymentMethodsNil) Type_() interface{} {
	s := &AccountPaymentMethods{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountPaymentMethod))
	}
	return s
}
