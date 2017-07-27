package types

import (
	"time"
)

type ReportAccountReportGeneralledger struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type ReportAccountReportGeneralledgerNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var ReportAccountReportGeneralledgerModel string = "report.account.report_generalledger"

type ReportAccountReportGeneralledgers []ReportAccountReportGeneralledger

type ReportAccountReportGeneralledgersNil []ReportAccountReportGeneralledgerNil

func (s *ReportAccountReportGeneralledger) NilableType_() interface{} {
	return &ReportAccountReportGeneralledgerNil{}
}

func (ns *ReportAccountReportGeneralledgerNil) Type_() interface{} {
	s := &ReportAccountReportGeneralledger{}
	return load(ns, s)
}

func (s *ReportAccountReportGeneralledgers) NilableType_() interface{} {
	return &ReportAccountReportGeneralledgersNil{}
}

func (ns *ReportAccountReportGeneralledgersNil) Type_() interface{} {
	s := &ReportAccountReportGeneralledgers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportGeneralledger))
	}
	return s
}
