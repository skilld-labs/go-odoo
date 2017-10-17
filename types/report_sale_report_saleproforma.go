package types

import (
	"time"
)

type ReportSaleReportSaleproforma struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportSaleReportSaleproformaNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportSaleReportSaleproformaModel string = "report.sale.report_saleproforma"

type ReportSaleReportSaleproformas []ReportSaleReportSaleproforma

type ReportSaleReportSaleproformasNil []ReportSaleReportSaleproformaNil

func (s *ReportSaleReportSaleproforma) NilableType_() interface{} {
	return &ReportSaleReportSaleproformaNil{}
}

func (ns *ReportSaleReportSaleproformaNil) Type_() interface{} {
	s := &ReportSaleReportSaleproforma{}
	return load(ns, s)
}

func (s *ReportSaleReportSaleproformas) NilableType_() interface{} {
	return &ReportSaleReportSaleproformasNil{}
}

func (ns *ReportSaleReportSaleproformasNil) Type_() interface{} {
	s := &ReportSaleReportSaleproformas{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportSaleReportSaleproforma))
	}
	return s
}
