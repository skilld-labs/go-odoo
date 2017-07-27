package types

import (
	"time"
)

type AccountRegisterPayments struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	Amount            float64   `xmlrpc:"amount"`
	Communication     string    `xmlrpc:"communication"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
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
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type AccountRegisterPaymentsNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	Amount            interface{} `xmlrpc:"amount"`
	Communication     interface{} `xmlrpc:"communication"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
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
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var AccountRegisterPaymentsModel string = "account.register.payments"

type AccountRegisterPaymentss []AccountRegisterPayments

type AccountRegisterPaymentssNil []AccountRegisterPaymentsNil

func (s *AccountRegisterPayments) NilableType_() interface{} {
	return &AccountRegisterPaymentsNil{}
}

func (ns *AccountRegisterPaymentsNil) Type_() interface{} {
	s := &AccountRegisterPayments{}
	return load(ns, s)
}

func (s *AccountRegisterPaymentss) NilableType_() interface{} {
	return &AccountRegisterPaymentssNil{}
}

func (ns *AccountRegisterPaymentssNil) Type_() interface{} {
	s := &AccountRegisterPaymentss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountRegisterPayments))
	}
	return s
}
