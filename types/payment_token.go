package types

import (
	"time"
)

type PaymentToken struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	AcquirerId  Many2One  `xmlrpc:"acquirer_id"`
	AcquirerRef string    `xmlrpc:"acquirer_ref"`
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	PartnerId   Many2One  `xmlrpc:"partner_id"`
	PaymentIds  []int64   `xmlrpc:"payment_ids"`
	ShortName   string    `xmlrpc:"short_name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type PaymentTokenNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	AcquirerId  interface{} `xmlrpc:"acquirer_id"`
	AcquirerRef interface{} `xmlrpc:"acquirer_ref"`
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	PaymentIds  interface{} `xmlrpc:"payment_ids"`
	ShortName   interface{} `xmlrpc:"short_name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var PaymentTokenModel string = "payment.token"

type PaymentTokens []PaymentToken

type PaymentTokensNil []PaymentTokenNil

func (s *PaymentToken) NilableType_() interface{} {
	return &PaymentTokenNil{}
}

func (ns *PaymentTokenNil) Type_() interface{} {
	s := &PaymentToken{}
	return load(ns, s)
}

func (s *PaymentTokens) NilableType_() interface{} {
	return &PaymentTokensNil{}
}

func (ns *PaymentTokensNil) Type_() interface{} {
	s := &PaymentTokens{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PaymentToken))
	}
	return s
}
