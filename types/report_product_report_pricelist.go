package types

import (
	"time"
)

type ReportProductReportPricelist struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportProductReportPricelistNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportProductReportPricelistModel string = "report.product.report_pricelist"

type ReportProductReportPricelists []ReportProductReportPricelist

type ReportProductReportPricelistsNil []ReportProductReportPricelistNil

func (s *ReportProductReportPricelist) NilableType_() interface{} {
	return &ReportProductReportPricelistNil{}
}

func (ns *ReportProductReportPricelistNil) Type_() interface{} {
	s := &ReportProductReportPricelist{}
	return load(ns, s)
}

func (s *ReportProductReportPricelists) NilableType_() interface{} {
	return &ReportProductReportPricelistsNil{}
}

func (ns *ReportProductReportPricelistsNil) Type_() interface{} {
	s := &ReportProductReportPricelists{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportProductReportPricelist))
	}
	return s
}
