package types

import (
	"time"
)

type AccountBalanceReport struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DateFrom       time.Time `xmlrpc:"date_from"`
	DateTo         time.Time `xmlrpc:"date_to"`
	DisplayAccount string    `xmlrpc:"display_account"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	JournalIds     []int64   `xmlrpc:"journal_ids"`
	TargetMove     string    `xmlrpc:"target_move"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type AccountBalanceReportNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DateFrom       interface{} `xmlrpc:"date_from"`
	DateTo         interface{} `xmlrpc:"date_to"`
	DisplayAccount interface{} `xmlrpc:"display_account"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	JournalIds     interface{} `xmlrpc:"journal_ids"`
	TargetMove     interface{} `xmlrpc:"target_move"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var AccountBalanceReportModel string = "account.balance.report"

type AccountBalanceReports []AccountBalanceReport

type AccountBalanceReportsNil []AccountBalanceReportNil

func (s *AccountBalanceReport) NilableType_() interface{} {
	return &AccountBalanceReportNil{}
}

func (ns *AccountBalanceReportNil) Type_() interface{} {
	s := &AccountBalanceReport{}
	return load(ns, s)
}

func (s *AccountBalanceReports) NilableType_() interface{} {
	return &AccountBalanceReportsNil{}
}

func (ns *AccountBalanceReportsNil) Type_() interface{} {
	s := &AccountBalanceReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBalanceReport))
	}
	return s
}
