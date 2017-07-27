package types

import (
	"time"
)

type AccountCommonJournalReport struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	AmountCurrency bool      `xmlrpc:"amount_currency"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DateFrom       time.Time `xmlrpc:"date_from"`
	DateTo         time.Time `xmlrpc:"date_to"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	JournalIds     []int64   `xmlrpc:"journal_ids"`
	TargetMove     string    `xmlrpc:"target_move"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type AccountCommonJournalReportNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	AmountCurrency bool        `xmlrpc:"amount_currency"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DateFrom       interface{} `xmlrpc:"date_from"`
	DateTo         interface{} `xmlrpc:"date_to"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	JournalIds     interface{} `xmlrpc:"journal_ids"`
	TargetMove     interface{} `xmlrpc:"target_move"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var AccountCommonJournalReportModel string = "account.common.journal.report"

type AccountCommonJournalReports []AccountCommonJournalReport

type AccountCommonJournalReportsNil []AccountCommonJournalReportNil

func (s *AccountCommonJournalReport) NilableType_() interface{} {
	return &AccountCommonJournalReportNil{}
}

func (ns *AccountCommonJournalReportNil) Type_() interface{} {
	s := &AccountCommonJournalReport{}
	return load(ns, s)
}

func (s *AccountCommonJournalReports) NilableType_() interface{} {
	return &AccountCommonJournalReportsNil{}
}

func (ns *AccountCommonJournalReportsNil) Type_() interface{} {
	s := &AccountCommonJournalReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountCommonJournalReport))
	}
	return s
}
