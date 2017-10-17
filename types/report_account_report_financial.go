package types

import (
	"time"
)

type ReportAccountReportFinancial struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportAccountReportFinancialNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportAccountReportFinancialModel string = "report.account.report_financial"

type ReportAccountReportFinancials []ReportAccountReportFinancial

type ReportAccountReportFinancialsNil []ReportAccountReportFinancialNil

func (s *ReportAccountReportFinancial) NilableType_() interface{} {
	return &ReportAccountReportFinancialNil{}
}

func (ns *ReportAccountReportFinancialNil) Type_() interface{} {
	s := &ReportAccountReportFinancial{}
	return load(ns, s)
}

func (s *ReportAccountReportFinancials) NilableType_() interface{} {
	return &ReportAccountReportFinancialsNil{}
}

func (ns *ReportAccountReportFinancialsNil) Type_() interface{} {
	s := &ReportAccountReportFinancials{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportFinancial))
	}
	return s
}
