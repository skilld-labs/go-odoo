package types

import (
	"time"
)

type AccountReportPartnerLedger struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	AmountCurrency  bool      `xmlrpc:"amount_currency"`
	CompanyId       Many2One  `xmlrpc:"company_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DateFrom        time.Time `xmlrpc:"date_from"`
	DateTo          time.Time `xmlrpc:"date_to"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	JournalIds      []int64   `xmlrpc:"journal_ids"`
	Reconciled      bool      `xmlrpc:"reconciled"`
	ResultSelection string    `xmlrpc:"result_selection"`
	TargetMove      string    `xmlrpc:"target_move"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountReportPartnerLedgerNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	AmountCurrency  bool        `xmlrpc:"amount_currency"`
	CompanyId       interface{} `xmlrpc:"company_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DateFrom        interface{} `xmlrpc:"date_from"`
	DateTo          interface{} `xmlrpc:"date_to"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	JournalIds      interface{} `xmlrpc:"journal_ids"`
	Reconciled      bool        `xmlrpc:"reconciled"`
	ResultSelection interface{} `xmlrpc:"result_selection"`
	TargetMove      interface{} `xmlrpc:"target_move"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountReportPartnerLedgerModel string = "account.report.partner.ledger"

type AccountReportPartnerLedgers []AccountReportPartnerLedger

type AccountReportPartnerLedgersNil []AccountReportPartnerLedgerNil

func (s *AccountReportPartnerLedger) NilableType_() interface{} {
	return &AccountReportPartnerLedgerNil{}
}

func (ns *AccountReportPartnerLedgerNil) Type_() interface{} {
	s := &AccountReportPartnerLedger{}
	return load(ns, s)
}

func (s *AccountReportPartnerLedgers) NilableType_() interface{} {
	return &AccountReportPartnerLedgersNil{}
}

func (ns *AccountReportPartnerLedgersNil) Type_() interface{} {
	s := &AccountReportPartnerLedgers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountReportPartnerLedger))
	}
	return s
}
