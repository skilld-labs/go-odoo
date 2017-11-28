package types

import (
	"time"
)

type AccountJournal struct {
	AccountControlIds                  []int64     `xmlrpc:"account_control_ids"`
	AccountSetupBankDataDone           bool        `xmlrpc:"account_setup_bank_data_done"`
	Active                             bool        `xmlrpc:"active"`
	AtLeastOneInbound                  bool        `xmlrpc:"at_least_one_inbound"`
	AtLeastOneOutbound                 bool        `xmlrpc:"at_least_one_outbound"`
	BankAccNumber                      string      `xmlrpc:"bank_acc_number"`
	BankAccountId                      Many2One    `xmlrpc:"bank_account_id"`
	BankId                             Many2One    `xmlrpc:"bank_id"`
	BankStatementsSource               interface{} `xmlrpc:"bank_statements_source"`
	BelongsToCompany                   bool        `xmlrpc:"belongs_to_company"`
	CheckManualSequencing              bool        `xmlrpc:"check_manual_sequencing"`
	CheckNextNumber                    int64       `xmlrpc:"check_next_number"`
	CheckPrintingPaymentMethodSelected bool        `xmlrpc:"check_printing_payment_method_selected"`
	CheckSequenceId                    Many2One    `xmlrpc:"check_sequence_id"`
	Code                               string      `xmlrpc:"code"`
	Color                              int64       `xmlrpc:"color"`
	CompanyId                          Many2One    `xmlrpc:"company_id"`
	CreateDate                         time.Time   `xmlrpc:"create_date"`
	CreateUid                          Many2One    `xmlrpc:"create_uid"`
	CurrencyId                         Many2One    `xmlrpc:"currency_id"`
	DefaultCreditAccountId             Many2One    `xmlrpc:"default_credit_account_id"`
	DefaultDebitAccountId              Many2One    `xmlrpc:"default_debit_account_id"`
	DisplayName                        string      `xmlrpc:"display_name"`
	GroupInvoiceLines                  bool        `xmlrpc:"group_invoice_lines"`
	Id                                 int64       `xmlrpc:"id"`
	InboundPaymentMethodIds            []int64     `xmlrpc:"inbound_payment_method_ids"`
	KanbanDashboard                    string      `xmlrpc:"kanban_dashboard"`
	KanbanDashboardGraph               string      `xmlrpc:"kanban_dashboard_graph"`
	LastUpdate                         time.Time   `xmlrpc:"__last_update"`
	LossAccountId                      Many2One    `xmlrpc:"loss_account_id"`
	Name                               string      `xmlrpc:"name"`
	OutboundPaymentMethodIds           []int64     `xmlrpc:"outbound_payment_method_ids"`
	ProfitAccountId                    Many2One    `xmlrpc:"profit_account_id"`
	RefundSequence                     bool        `xmlrpc:"refund_sequence"`
	RefundSequenceId                   Many2One    `xmlrpc:"refund_sequence_id"`
	RefundSequenceNumberNext           int64       `xmlrpc:"refund_sequence_number_next"`
	Sequence                           int64       `xmlrpc:"sequence"`
	SequenceId                         Many2One    `xmlrpc:"sequence_id"`
	SequenceNumberNext                 int64       `xmlrpc:"sequence_number_next"`
	ShowOnDashboard                    bool        `xmlrpc:"show_on_dashboard"`
	Type                               interface{} `xmlrpc:"type"`
	TypeControlIds                     []int64     `xmlrpc:"type_control_ids"`
	UpdatePosted                       bool        `xmlrpc:"update_posted"`
	WriteDate                          time.Time   `xmlrpc:"write_date"`
	WriteUid                           Many2One    `xmlrpc:"write_uid"`
}

type AccountJournalNil struct {
	AccountControlIds                  interface{} `xmlrpc:"account_control_ids"`
	AccountSetupBankDataDone           bool        `xmlrpc:"account_setup_bank_data_done"`
	Active                             bool        `xmlrpc:"active"`
	AtLeastOneInbound                  bool        `xmlrpc:"at_least_one_inbound"`
	AtLeastOneOutbound                 bool        `xmlrpc:"at_least_one_outbound"`
	BankAccNumber                      interface{} `xmlrpc:"bank_acc_number"`
	BankAccountId                      interface{} `xmlrpc:"bank_account_id"`
	BankId                             interface{} `xmlrpc:"bank_id"`
	BankStatementsSource               interface{} `xmlrpc:"bank_statements_source"`
	BelongsToCompany                   bool        `xmlrpc:"belongs_to_company"`
	CheckManualSequencing              bool        `xmlrpc:"check_manual_sequencing"`
	CheckNextNumber                    interface{} `xmlrpc:"check_next_number"`
	CheckPrintingPaymentMethodSelected bool        `xmlrpc:"check_printing_payment_method_selected"`
	CheckSequenceId                    interface{} `xmlrpc:"check_sequence_id"`
	Code                               interface{} `xmlrpc:"code"`
	Color                              interface{} `xmlrpc:"color"`
	CompanyId                          interface{} `xmlrpc:"company_id"`
	CreateDate                         interface{} `xmlrpc:"create_date"`
	CreateUid                          interface{} `xmlrpc:"create_uid"`
	CurrencyId                         interface{} `xmlrpc:"currency_id"`
	DefaultCreditAccountId             interface{} `xmlrpc:"default_credit_account_id"`
	DefaultDebitAccountId              interface{} `xmlrpc:"default_debit_account_id"`
	DisplayName                        interface{} `xmlrpc:"display_name"`
	GroupInvoiceLines                  bool        `xmlrpc:"group_invoice_lines"`
	Id                                 interface{} `xmlrpc:"id"`
	InboundPaymentMethodIds            interface{} `xmlrpc:"inbound_payment_method_ids"`
	KanbanDashboard                    interface{} `xmlrpc:"kanban_dashboard"`
	KanbanDashboardGraph               interface{} `xmlrpc:"kanban_dashboard_graph"`
	LastUpdate                         interface{} `xmlrpc:"__last_update"`
	LossAccountId                      interface{} `xmlrpc:"loss_account_id"`
	Name                               interface{} `xmlrpc:"name"`
	OutboundPaymentMethodIds           interface{} `xmlrpc:"outbound_payment_method_ids"`
	ProfitAccountId                    interface{} `xmlrpc:"profit_account_id"`
	RefundSequence                     bool        `xmlrpc:"refund_sequence"`
	RefundSequenceId                   interface{} `xmlrpc:"refund_sequence_id"`
	RefundSequenceNumberNext           interface{} `xmlrpc:"refund_sequence_number_next"`
	Sequence                           interface{} `xmlrpc:"sequence"`
	SequenceId                         interface{} `xmlrpc:"sequence_id"`
	SequenceNumberNext                 interface{} `xmlrpc:"sequence_number_next"`
	ShowOnDashboard                    bool        `xmlrpc:"show_on_dashboard"`
	Type                               interface{} `xmlrpc:"type"`
	TypeControlIds                     interface{} `xmlrpc:"type_control_ids"`
	UpdatePosted                       bool        `xmlrpc:"update_posted"`
	WriteDate                          interface{} `xmlrpc:"write_date"`
	WriteUid                           interface{} `xmlrpc:"write_uid"`
}

var AccountJournalModel string = "account.journal"

type AccountJournals []AccountJournal

type AccountJournalsNil []AccountJournalNil

func (s *AccountJournal) NilableType_() interface{} {
	return &AccountJournalNil{}
}

func (ns *AccountJournalNil) Type_() interface{} {
	s := &AccountJournal{}
	return load(ns, s)
}

func (s *AccountJournals) NilableType_() interface{} {
	return &AccountJournalsNil{}
}

func (ns *AccountJournalsNil) Type_() interface{} {
	s := &AccountJournals{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountJournal))
	}
	return s
}
