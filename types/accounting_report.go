package types

import (
	"time"
)

type AccountingReport struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	AccountReportId Many2One  `xmlrpc:"account_report_id"`
	CompanyId       Many2One  `xmlrpc:"company_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DateFrom        time.Time `xmlrpc:"date_from"`
	DateFromCmp     time.Time `xmlrpc:"date_from_cmp"`
	DateTo          time.Time `xmlrpc:"date_to"`
	DateToCmp       time.Time `xmlrpc:"date_to_cmp"`
	DebitCredit     bool      `xmlrpc:"debit_credit"`
	DisplayName     string    `xmlrpc:"display_name"`
	EnableFilter    bool      `xmlrpc:"enable_filter"`
	FilterCmp       string    `xmlrpc:"filter_cmp"`
	Id              int64     `xmlrpc:"id"`
	JournalIds      []int64   `xmlrpc:"journal_ids"`
	LabelFilter     string    `xmlrpc:"label_filter"`
	TargetMove      string    `xmlrpc:"target_move"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountingReportNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	AccountReportId interface{} `xmlrpc:"account_report_id"`
	CompanyId       interface{} `xmlrpc:"company_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DateFrom        interface{} `xmlrpc:"date_from"`
	DateFromCmp     interface{} `xmlrpc:"date_from_cmp"`
	DateTo          interface{} `xmlrpc:"date_to"`
	DateToCmp       interface{} `xmlrpc:"date_to_cmp"`
	DebitCredit     bool        `xmlrpc:"debit_credit"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	EnableFilter    bool        `xmlrpc:"enable_filter"`
	FilterCmp       interface{} `xmlrpc:"filter_cmp"`
	Id              interface{} `xmlrpc:"id"`
	JournalIds      interface{} `xmlrpc:"journal_ids"`
	LabelFilter     interface{} `xmlrpc:"label_filter"`
	TargetMove      interface{} `xmlrpc:"target_move"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountingReportModel string = "accounting.report"

type AccountingReports []AccountingReport

type AccountingReportsNil []AccountingReportNil

func (s *AccountingReport) NilableType_() interface{} {
	return &AccountingReportNil{}
}

func (ns *AccountingReportNil) Type_() interface{} {
	s := &AccountingReport{}
	return load(ns, s)
}

func (s *AccountingReports) NilableType_() interface{} {
	return &AccountingReportsNil{}
}

func (ns *AccountingReportsNil) Type_() interface{} {
	s := &AccountingReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountingReport))
	}
	return s
}
