package types

import (
	"time"
)

type AccountReportGeneralLedger struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DateFrom       time.Time `xmlrpc:"date_from"`
	DateTo         time.Time `xmlrpc:"date_to"`
	DisplayAccount string    `xmlrpc:"display_account"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	InitialBalance bool      `xmlrpc:"initial_balance"`
	JournalIds     []int64   `xmlrpc:"journal_ids"`
	Sortby         string    `xmlrpc:"sortby"`
	TargetMove     string    `xmlrpc:"target_move"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type AccountReportGeneralLedgerNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DateFrom       interface{} `xmlrpc:"date_from"`
	DateTo         interface{} `xmlrpc:"date_to"`
	DisplayAccount interface{} `xmlrpc:"display_account"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	InitialBalance bool        `xmlrpc:"initial_balance"`
	JournalIds     interface{} `xmlrpc:"journal_ids"`
	Sortby         interface{} `xmlrpc:"sortby"`
	TargetMove     interface{} `xmlrpc:"target_move"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var AccountReportGeneralLedgerModel string = "account.report.general.ledger"

type AccountReportGeneralLedgers []AccountReportGeneralLedger

type AccountReportGeneralLedgersNil []AccountReportGeneralLedgerNil

func (s *AccountReportGeneralLedger) NilableType_() interface{} {
	return &AccountReportGeneralLedgerNil{}
}

func (ns *AccountReportGeneralLedgerNil) Type_() interface{} {
	s := &AccountReportGeneralLedger{}
	return load(ns, s)
}

func (s *AccountReportGeneralLedgers) NilableType_() interface{} {
	return &AccountReportGeneralLedgersNil{}
}

func (ns *AccountReportGeneralLedgersNil) Type_() interface{} {
	s := &AccountReportGeneralLedgers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountReportGeneralLedger))
	}
	return s
}
