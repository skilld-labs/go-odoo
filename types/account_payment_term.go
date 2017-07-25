package types

import (
	"time"
)

type AccountPaymentTerm struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LineIds     []int64   `xmlrpc:"line_ids"`
	Name        string    `xmlrpc:"name"`
	Note        string    `xmlrpc:"note"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountPaymentTermNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LineIds     interface{} `xmlrpc:"line_ids"`
	Name        interface{} `xmlrpc:"name"`
	Note        interface{} `xmlrpc:"note"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountPaymentTermModel string = "account.payment.term"

type AccountPaymentTerms []AccountPaymentTerm

type AccountPaymentTermsNil []AccountPaymentTermNil

func (s *AccountPaymentTerm) NilableType_() interface{} {
	return &AccountPaymentTermNil{}
}

func (ns *AccountPaymentTermNil) Type_() interface{} {
	s := &AccountPaymentTerm{}
	return load(ns, s)
}

func (s *AccountPaymentTerms) NilableType_() interface{} {
	return &AccountPaymentTermsNil{}
}

func (ns *AccountPaymentTermsNil) Type_() interface{} {
	s := &AccountPaymentTerms{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountPaymentTerm))
	}
	return s
}
