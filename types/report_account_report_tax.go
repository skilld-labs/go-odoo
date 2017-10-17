package types

import (
	"time"
)

type ReportAccountReportTax struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportAccountReportTaxNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportAccountReportTaxModel string = "report.account.report_tax"

type ReportAccountReportTaxs []ReportAccountReportTax

type ReportAccountReportTaxsNil []ReportAccountReportTaxNil

func (s *ReportAccountReportTax) NilableType_() interface{} {
	return &ReportAccountReportTaxNil{}
}

func (ns *ReportAccountReportTaxNil) Type_() interface{} {
	s := &ReportAccountReportTax{}
	return load(ns, s)
}

func (s *ReportAccountReportTaxs) NilableType_() interface{} {
	return &ReportAccountReportTaxsNil{}
}

func (ns *ReportAccountReportTaxsNil) Type_() interface{} {
	s := &ReportAccountReportTaxs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportTax))
	}
	return s
}
