package types

import (
	"time"
)

type AccountPayment struct {
	LastUpdate                time.Time `xmlrpc:"__last_update"`
	Amount                    float64   `xmlrpc:"amount"`
	Communication             string    `xmlrpc:"communication"`
	CompanyId                 Many2One  `xmlrpc:"company_id"`
	CreateDate                time.Time `xmlrpc:"create_date"`
	CreateUid                 Many2One  `xmlrpc:"create_uid"`
	CurrencyId                Many2One  `xmlrpc:"currency_id"`
	DestinationAccountId      Many2One  `xmlrpc:"destination_account_id"`
	DestinationJournalId      Many2One  `xmlrpc:"destination_journal_id"`
	DisplayName               string    `xmlrpc:"display_name"`
	HasInvoices               bool      `xmlrpc:"has_invoices"`
	HidePaymentMethod         bool      `xmlrpc:"hide_payment_method"`
	Id                        int64     `xmlrpc:"id"`
	InvoiceIds                []int64   `xmlrpc:"invoice_ids"`
	JournalId                 Many2One  `xmlrpc:"journal_id"`
	MoveLineIds               []int64   `xmlrpc:"move_line_ids"`
	MoveName                  string    `xmlrpc:"move_name"`
	Name                      string    `xmlrpc:"name"`
	PartnerId                 Many2One  `xmlrpc:"partner_id"`
	PartnerType               string    `xmlrpc:"partner_type"`
	PaymentDate               time.Time `xmlrpc:"payment_date"`
	PaymentDifference         float64   `xmlrpc:"payment_difference"`
	PaymentDifferenceHandling string    `xmlrpc:"payment_difference_handling"`
	PaymentMethodCode         string    `xmlrpc:"payment_method_code"`
	PaymentMethodId           Many2One  `xmlrpc:"payment_method_id"`
	PaymentReference          string    `xmlrpc:"payment_reference"`
	PaymentTokenId            Many2One  `xmlrpc:"payment_token_id"`
	PaymentTransactionId      Many2One  `xmlrpc:"payment_transaction_id"`
	PaymentType               string    `xmlrpc:"payment_type"`
	State                     string    `xmlrpc:"state"`
	WriteDate                 time.Time `xmlrpc:"write_date"`
	WriteUid                  Many2One  `xmlrpc:"write_uid"`
	WriteoffAccountId         Many2One  `xmlrpc:"writeoff_account_id"`
}

type AccountPaymentNil struct {
	LastUpdate                interface{} `xmlrpc:"__last_update"`
	Amount                    interface{} `xmlrpc:"amount"`
	Communication             interface{} `xmlrpc:"communication"`
	CompanyId                 interface{} `xmlrpc:"company_id"`
	CreateDate                interface{} `xmlrpc:"create_date"`
	CreateUid                 interface{} `xmlrpc:"create_uid"`
	CurrencyId                interface{} `xmlrpc:"currency_id"`
	DestinationAccountId      interface{} `xmlrpc:"destination_account_id"`
	DestinationJournalId      interface{} `xmlrpc:"destination_journal_id"`
	DisplayName               interface{} `xmlrpc:"display_name"`
	HasInvoices               bool        `xmlrpc:"has_invoices"`
	HidePaymentMethod         bool        `xmlrpc:"hide_payment_method"`
	Id                        interface{} `xmlrpc:"id"`
	InvoiceIds                interface{} `xmlrpc:"invoice_ids"`
	JournalId                 interface{} `xmlrpc:"journal_id"`
	MoveLineIds               interface{} `xmlrpc:"move_line_ids"`
	MoveName                  interface{} `xmlrpc:"move_name"`
	Name                      interface{} `xmlrpc:"name"`
	PartnerId                 interface{} `xmlrpc:"partner_id"`
	PartnerType               interface{} `xmlrpc:"partner_type"`
	PaymentDate               interface{} `xmlrpc:"payment_date"`
	PaymentDifference         interface{} `xmlrpc:"payment_difference"`
	PaymentDifferenceHandling interface{} `xmlrpc:"payment_difference_handling"`
	PaymentMethodCode         interface{} `xmlrpc:"payment_method_code"`
	PaymentMethodId           interface{} `xmlrpc:"payment_method_id"`
	PaymentReference          interface{} `xmlrpc:"payment_reference"`
	PaymentTokenId            interface{} `xmlrpc:"payment_token_id"`
	PaymentTransactionId      interface{} `xmlrpc:"payment_transaction_id"`
	PaymentType               interface{} `xmlrpc:"payment_type"`
	State                     interface{} `xmlrpc:"state"`
	WriteDate                 interface{} `xmlrpc:"write_date"`
	WriteUid                  interface{} `xmlrpc:"write_uid"`
	WriteoffAccountId         interface{} `xmlrpc:"writeoff_account_id"`
}

var AccountPaymentModel string = "account.payment"

type AccountPayments []AccountPayment

type AccountPaymentsNil []AccountPaymentNil

func (s *AccountPayment) NilableType_() interface{} {
	return &AccountPaymentNil{}
}

func (ns *AccountPaymentNil) Type_() interface{} {
	s := &AccountPayment{}
	return load(ns, s)
}

func (s *AccountPayments) NilableType_() interface{} {
	return &AccountPaymentsNil{}
}

func (ns *AccountPaymentsNil) Type_() interface{} {
	s := &AccountPayments{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountPayment))
	}
	return s
}
