package types

import (
	"time"
)

type AccountAbstractPayment struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	Amount            float64   `xmlrpc:"amount"`
	Communication     string    `xmlrpc:"communication"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	DisplayName       string    `xmlrpc:"display_name"`
	HidePaymentMethod bool      `xmlrpc:"hide_payment_method"`
	Id                int64     `xmlrpc:"id"`
	JournalId         Many2One  `xmlrpc:"journal_id"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	PartnerType       string    `xmlrpc:"partner_type"`
	PaymentDate       time.Time `xmlrpc:"payment_date"`
	PaymentMethodCode string    `xmlrpc:"payment_method_code"`
	PaymentMethodId   Many2One  `xmlrpc:"payment_method_id"`
	PaymentType       string    `xmlrpc:"payment_type"`
}

type AccountAbstractPaymentNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	Amount            interface{} `xmlrpc:"amount"`
	Communication     interface{} `xmlrpc:"communication"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	HidePaymentMethod bool        `xmlrpc:"hide_payment_method"`
	Id                interface{} `xmlrpc:"id"`
	JournalId         interface{} `xmlrpc:"journal_id"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	PartnerType       interface{} `xmlrpc:"partner_type"`
	PaymentDate       interface{} `xmlrpc:"payment_date"`
	PaymentMethodCode interface{} `xmlrpc:"payment_method_code"`
	PaymentMethodId   interface{} `xmlrpc:"payment_method_id"`
	PaymentType       interface{} `xmlrpc:"payment_type"`
}

var AccountAbstractPaymentModel string = "account.abstract.payment"

type AccountAbstractPayments []AccountAbstractPayment

type AccountAbstractPaymentsNil []AccountAbstractPaymentNil

func (s *AccountAbstractPayment) NilableType_() interface{} {
	return &AccountAbstractPaymentNil{}
}

func (ns *AccountAbstractPaymentNil) Type_() interface{} {
	s := &AccountAbstractPayment{}
	return load(ns, s)
}

func (s *AccountAbstractPayments) NilableType_() interface{} {
	return &AccountAbstractPaymentsNil{}
}

func (ns *AccountAbstractPaymentsNil) Type_() interface{} {
	s := &AccountAbstractPayments{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAbstractPayment))
	}
	return s
}
