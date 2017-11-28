package types

import (
	"time"
)

type AccountPayment struct {
	Amount                    float64     `xmlrpc:"amount"`
	CheckAmountInWords        string      `xmlrpc:"check_amount_in_words"`
	CheckManualSequencing     bool        `xmlrpc:"check_manual_sequencing"`
	CheckNumber               int64       `xmlrpc:"check_number"`
	Communication             string      `xmlrpc:"communication"`
	CompanyId                 Many2One    `xmlrpc:"company_id"`
	CreateDate                time.Time   `xmlrpc:"create_date"`
	CreateUid                 Many2One    `xmlrpc:"create_uid"`
	CurrencyId                Many2One    `xmlrpc:"currency_id"`
	DestinationAccountId      Many2One    `xmlrpc:"destination_account_id"`
	DestinationJournalId      Many2One    `xmlrpc:"destination_journal_id"`
	DisplayName               string      `xmlrpc:"display_name"`
	HasInvoices               bool        `xmlrpc:"has_invoices"`
	HidePaymentMethod         bool        `xmlrpc:"hide_payment_method"`
	Id                        int64       `xmlrpc:"id"`
	InvoiceIds                []int64     `xmlrpc:"invoice_ids"`
	JournalId                 Many2One    `xmlrpc:"journal_id"`
	LastUpdate                time.Time   `xmlrpc:"__last_update"`
	MessageChannelIds         []int64     `xmlrpc:"message_channel_ids"`
	MessageFollowerIds        []int64     `xmlrpc:"message_follower_ids"`
	MessageIds                []int64     `xmlrpc:"message_ids"`
	MessageIsFollower         bool        `xmlrpc:"message_is_follower"`
	MessageLastPost           time.Time   `xmlrpc:"message_last_post"`
	MessageNeedaction         bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter  int64       `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds         []int64     `xmlrpc:"message_partner_ids"`
	MessageUnread             bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter      int64       `xmlrpc:"message_unread_counter"`
	MoveLineIds               []int64     `xmlrpc:"move_line_ids"`
	MoveName                  string      `xmlrpc:"move_name"`
	MoveReconciled            bool        `xmlrpc:"move_reconciled"`
	Name                      string      `xmlrpc:"name"`
	PartnerId                 Many2One    `xmlrpc:"partner_id"`
	PartnerType               interface{} `xmlrpc:"partner_type"`
	PaymentDate               time.Time   `xmlrpc:"payment_date"`
	PaymentDifference         float64     `xmlrpc:"payment_difference"`
	PaymentDifferenceHandling interface{} `xmlrpc:"payment_difference_handling"`
	PaymentMethodCode         string      `xmlrpc:"payment_method_code"`
	PaymentMethodId           Many2One    `xmlrpc:"payment_method_id"`
	PaymentReference          string      `xmlrpc:"payment_reference"`
	PaymentTokenId            Many2One    `xmlrpc:"payment_token_id"`
	PaymentTransactionId      Many2One    `xmlrpc:"payment_transaction_id"`
	PaymentType               interface{} `xmlrpc:"payment_type"`
	State                     interface{} `xmlrpc:"state"`
	WebsiteMessageIds         []int64     `xmlrpc:"website_message_ids"`
	WriteDate                 time.Time   `xmlrpc:"write_date"`
	WriteoffAccountId         Many2One    `xmlrpc:"writeoff_account_id"`
	WriteoffLabel             string      `xmlrpc:"writeoff_label"`
	WriteUid                  Many2One    `xmlrpc:"write_uid"`
}

type AccountPaymentNil struct {
	Amount                    interface{} `xmlrpc:"amount"`
	CheckAmountInWords        interface{} `xmlrpc:"check_amount_in_words"`
	CheckManualSequencing     bool        `xmlrpc:"check_manual_sequencing"`
	CheckNumber               interface{} `xmlrpc:"check_number"`
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
	LastUpdate                interface{} `xmlrpc:"__last_update"`
	MessageChannelIds         interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds        interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                interface{} `xmlrpc:"message_ids"`
	MessageIsFollower         bool        `xmlrpc:"message_is_follower"`
	MessageLastPost           interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction         bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter  interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds         interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread             bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter      interface{} `xmlrpc:"message_unread_counter"`
	MoveLineIds               interface{} `xmlrpc:"move_line_ids"`
	MoveName                  interface{} `xmlrpc:"move_name"`
	MoveReconciled            bool        `xmlrpc:"move_reconciled"`
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
	WebsiteMessageIds         interface{} `xmlrpc:"website_message_ids"`
	WriteDate                 interface{} `xmlrpc:"write_date"`
	WriteoffAccountId         interface{} `xmlrpc:"writeoff_account_id"`
	WriteoffLabel             interface{} `xmlrpc:"writeoff_label"`
	WriteUid                  interface{} `xmlrpc:"write_uid"`
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
