package types

import (
	"time"
)

type ReportAccountReportOverdue struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type ReportAccountReportOverdueNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var ReportAccountReportOverdueModel string = "report.account.report_overdue"

type ReportAccountReportOverdues []ReportAccountReportOverdue

type ReportAccountReportOverduesNil []ReportAccountReportOverdueNil

func (s *ReportAccountReportOverdue) NilableType_() interface{} {
	return &ReportAccountReportOverdueNil{}
}

func (ns *ReportAccountReportOverdueNil) Type_() interface{} {
	s := &ReportAccountReportOverdue{}
	return load(ns, s)
}

func (s *ReportAccountReportOverdues) NilableType_() interface{} {
	return &ReportAccountReportOverduesNil{}
}

func (ns *ReportAccountReportOverduesNil) Type_() interface{} {
	s := &ReportAccountReportOverdues{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportOverdue))
	}
	return s
}
