package types

import (
	"time"
)

type ReportHrHolidaysReportHolidayssummary struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type ReportHrHolidaysReportHolidayssummaryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var ReportHrHolidaysReportHolidayssummaryModel string = "report.hr_holidays.report_holidayssummary"

type ReportHrHolidaysReportHolidayssummarys []ReportHrHolidaysReportHolidayssummary

type ReportHrHolidaysReportHolidayssummarysNil []ReportHrHolidaysReportHolidayssummaryNil

func (s *ReportHrHolidaysReportHolidayssummary) NilableType_() interface{} {
	return &ReportHrHolidaysReportHolidayssummaryNil{}
}

func (ns *ReportHrHolidaysReportHolidayssummaryNil) Type_() interface{} {
	s := &ReportHrHolidaysReportHolidayssummary{}
	return load(ns, s)
}

func (s *ReportHrHolidaysReportHolidayssummarys) NilableType_() interface{} {
	return &ReportHrHolidaysReportHolidayssummarysNil{}
}

func (ns *ReportHrHolidaysReportHolidayssummarysNil) Type_() interface{} {
	s := &ReportHrHolidaysReportHolidayssummarys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportHrHolidaysReportHolidayssummary))
	}
	return s
}
