package types

import (
	"time"
)

type ReportBaseReportIrmodulereference struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportBaseReportIrmodulereferenceNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportBaseReportIrmodulereferenceModel string = "report.base.report_irmodulereference"

type ReportBaseReportIrmodulereferences []ReportBaseReportIrmodulereference

type ReportBaseReportIrmodulereferencesNil []ReportBaseReportIrmodulereferenceNil

func (s *ReportBaseReportIrmodulereference) NilableType_() interface{} {
	return &ReportBaseReportIrmodulereferenceNil{}
}

func (ns *ReportBaseReportIrmodulereferenceNil) Type_() interface{} {
	s := &ReportBaseReportIrmodulereference{}
	return load(ns, s)
}

func (s *ReportBaseReportIrmodulereferences) NilableType_() interface{} {
	return &ReportBaseReportIrmodulereferencesNil{}
}

func (ns *ReportBaseReportIrmodulereferencesNil) Type_() interface{} {
	s := &ReportBaseReportIrmodulereferences{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportBaseReportIrmodulereference))
	}
	return s
}
