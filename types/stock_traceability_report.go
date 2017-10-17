package types

import (
	"time"
)

type StockTraceabilityReport struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockTraceabilityReportNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockTraceabilityReportModel string = "stock.traceability.report"

type StockTraceabilityReports []StockTraceabilityReport

type StockTraceabilityReportsNil []StockTraceabilityReportNil

func (s *StockTraceabilityReport) NilableType_() interface{} {
	return &StockTraceabilityReportNil{}
}

func (ns *StockTraceabilityReportNil) Type_() interface{} {
	s := &StockTraceabilityReport{}
	return load(ns, s)
}

func (s *StockTraceabilityReports) NilableType_() interface{} {
	return &StockTraceabilityReportsNil{}
}

func (ns *StockTraceabilityReportsNil) Type_() interface{} {
	s := &StockTraceabilityReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockTraceabilityReport))
	}
	return s
}
