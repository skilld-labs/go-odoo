package types

import (
	"time"
)

type ReportAccountReportAgedpartnerbalance struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportAccountReportAgedpartnerbalanceNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportAccountReportAgedpartnerbalanceModel string = "report.account.report_agedpartnerbalance"

type ReportAccountReportAgedpartnerbalances []ReportAccountReportAgedpartnerbalance

type ReportAccountReportAgedpartnerbalancesNil []ReportAccountReportAgedpartnerbalanceNil

func (s *ReportAccountReportAgedpartnerbalance) NilableType_() interface{} {
	return &ReportAccountReportAgedpartnerbalanceNil{}
}

func (ns *ReportAccountReportAgedpartnerbalanceNil) Type_() interface{} {
	s := &ReportAccountReportAgedpartnerbalance{}
	return load(ns, s)
}

func (s *ReportAccountReportAgedpartnerbalances) NilableType_() interface{} {
	return &ReportAccountReportAgedpartnerbalancesNil{}
}

func (ns *ReportAccountReportAgedpartnerbalancesNil) Type_() interface{} {
	s := &ReportAccountReportAgedpartnerbalances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportAgedpartnerbalance))
	}
	return s
}
