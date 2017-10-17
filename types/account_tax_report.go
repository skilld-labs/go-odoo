package types

import (
	"time"
)

type AccountTaxReport struct {
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	DateTo      time.Time `xmlrpc:"date_to"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	JournalIds  []int64   `xmlrpc:"journal_ids"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	TargetMove  string    `xmlrpc:"target_move"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountTaxReportNil struct {
	CompanyId   interface{} `xmlrpc:"company_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	DateTo      interface{} `xmlrpc:"date_to"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	JournalIds  interface{} `xmlrpc:"journal_ids"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	TargetMove  interface{} `xmlrpc:"target_move"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountTaxReportModel string = "account.tax.report"

type AccountTaxReports []AccountTaxReport

type AccountTaxReportsNil []AccountTaxReportNil

func (s *AccountTaxReport) NilableType_() interface{} {
	return &AccountTaxReportNil{}
}

func (ns *AccountTaxReportNil) Type_() interface{} {
	s := &AccountTaxReport{}
	return load(ns, s)
}

func (s *AccountTaxReports) NilableType_() interface{} {
	return &AccountTaxReportsNil{}
}

func (ns *AccountTaxReportsNil) Type_() interface{} {
	s := &AccountTaxReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountTaxReport))
	}
	return s
}
