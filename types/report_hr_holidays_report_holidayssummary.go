package types

import (
	"time"
)

type ReportHrHolidaysReportHolidayssummary struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportHrHolidaysReportHolidayssummaryNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
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
