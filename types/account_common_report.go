package types

import (
	"time"
)

type AccountCommonReport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	DateTo      time.Time `xmlrpc:"date_to"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	JournalIds  []int64   `xmlrpc:"journal_ids"`
	TargetMove  string    `xmlrpc:"target_move"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountCommonReportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	DateTo      interface{} `xmlrpc:"date_to"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	JournalIds  interface{} `xmlrpc:"journal_ids"`
	TargetMove  interface{} `xmlrpc:"target_move"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountCommonReportModel string = "account.common.report"

type AccountCommonReports []AccountCommonReport

type AccountCommonReportsNil []AccountCommonReportNil

func (s *AccountCommonReport) NilableType_() interface{} {
	return &AccountCommonReportNil{}
}

func (ns *AccountCommonReportNil) Type_() interface{} {
	s := &AccountCommonReport{}
	return load(ns, s)
}

func (s *AccountCommonReports) NilableType_() interface{} {
	return &AccountCommonReportsNil{}
}

func (ns *AccountCommonReportsNil) Type_() interface{} {
	s := &AccountCommonReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountCommonReport))
	}
	return s
}
