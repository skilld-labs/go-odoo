package types

import (
	"time"
)

type ReportAbstractReport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type ReportAbstractReportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var ReportAbstractReportModel string = "report.abstract_report"

type ReportAbstractReports []ReportAbstractReport

type ReportAbstractReportsNil []ReportAbstractReportNil

func (s *ReportAbstractReport) NilableType_() interface{} {
	return &ReportAbstractReportNil{}
}

func (ns *ReportAbstractReportNil) Type_() interface{} {
	s := &ReportAbstractReport{}
	return load(ns, s)
}

func (s *ReportAbstractReports) NilableType_() interface{} {
	return &ReportAbstractReportsNil{}
}

func (ns *ReportAbstractReportsNil) Type_() interface{} {
	s := &ReportAbstractReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAbstractReport))
	}
	return s
}
