package types

import (
	"time"
)

type PaymentIcon struct {
	AcquirerIds      []int64   `xmlrpc:"acquirer_ids"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DisplayName      string    `xmlrpc:"display_name"`
	Id               int64     `xmlrpc:"id"`
	Image            string    `xmlrpc:"image"`
	ImagePaymentForm string    `xmlrpc:"image_payment_form"`
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	Name             string    `xmlrpc:"name"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type PaymentIconNil struct {
	AcquirerIds      interface{} `xmlrpc:"acquirer_ids"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Id               interface{} `xmlrpc:"id"`
	Image            interface{} `xmlrpc:"image"`
	ImagePaymentForm interface{} `xmlrpc:"image_payment_form"`
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	Name             interface{} `xmlrpc:"name"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var PaymentIconModel string = "payment.icon"

type PaymentIcons []PaymentIcon

type PaymentIconsNil []PaymentIconNil

func (s *PaymentIcon) NilableType_() interface{} {
	return &PaymentIconNil{}
}

func (ns *PaymentIconNil) Type_() interface{} {
	s := &PaymentIcon{}
	return load(ns, s)
}

func (s *PaymentIcons) NilableType_() interface{} {
	return &PaymentIconsNil{}
}

func (ns *PaymentIconsNil) Type_() interface{} {
	s := &PaymentIcons{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PaymentIcon))
	}
	return s
}
