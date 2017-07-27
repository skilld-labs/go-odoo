package types

import (
	"time"
)

type AccountFinancialReport struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	AccountIds      []int64   `xmlrpc:"account_ids"`
	AccountReportId Many2One  `xmlrpc:"account_report_id"`
	AccountTypeIds  []int64   `xmlrpc:"account_type_ids"`
	ChildrenIds     []int64   `xmlrpc:"children_ids"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayDetail   string    `xmlrpc:"display_detail"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	Level           int64     `xmlrpc:"level"`
	Name            string    `xmlrpc:"name"`
	ParentId        Many2One  `xmlrpc:"parent_id"`
	Sequence        int64     `xmlrpc:"sequence"`
	Sign            string    `xmlrpc:"sign"`
	StyleOverwrite  string    `xmlrpc:"style_overwrite"`
	Type            string    `xmlrpc:"type"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountFinancialReportNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	AccountIds      interface{} `xmlrpc:"account_ids"`
	AccountReportId interface{} `xmlrpc:"account_report_id"`
	AccountTypeIds  interface{} `xmlrpc:"account_type_ids"`
	ChildrenIds     interface{} `xmlrpc:"children_ids"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayDetail   interface{} `xmlrpc:"display_detail"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	Level           interface{} `xmlrpc:"level"`
	Name            interface{} `xmlrpc:"name"`
	ParentId        interface{} `xmlrpc:"parent_id"`
	Sequence        interface{} `xmlrpc:"sequence"`
	Sign            interface{} `xmlrpc:"sign"`
	StyleOverwrite  interface{} `xmlrpc:"style_overwrite"`
	Type            interface{} `xmlrpc:"type"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountFinancialReportModel string = "account.financial.report"

type AccountFinancialReports []AccountFinancialReport

type AccountFinancialReportsNil []AccountFinancialReportNil

func (s *AccountFinancialReport) NilableType_() interface{} {
	return &AccountFinancialReportNil{}
}

func (ns *AccountFinancialReportNil) Type_() interface{} {
	s := &AccountFinancialReport{}
	return load(ns, s)
}

func (s *AccountFinancialReports) NilableType_() interface{} {
	return &AccountFinancialReportsNil{}
}

func (ns *AccountFinancialReportsNil) Type_() interface{} {
	s := &AccountFinancialReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFinancialReport))
	}
	return s
}
