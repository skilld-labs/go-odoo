package types

import (
	"time"
)

type ReportAccountReportPartnerledger struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportAccountReportPartnerledgerNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportAccountReportPartnerledgerModel string = "report.account.report_partnerledger"

type ReportAccountReportPartnerledgers []ReportAccountReportPartnerledger

type ReportAccountReportPartnerledgersNil []ReportAccountReportPartnerledgerNil

func (s *ReportAccountReportPartnerledger) NilableType_() interface{} {
	return &ReportAccountReportPartnerledgerNil{}
}

func (ns *ReportAccountReportPartnerledgerNil) Type_() interface{} {
	s := &ReportAccountReportPartnerledger{}
	return load(ns, s)
}

func (s *ReportAccountReportPartnerledgers) NilableType_() interface{} {
	return &ReportAccountReportPartnerledgersNil{}
}

func (ns *ReportAccountReportPartnerledgersNil) Type_() interface{} {
	s := &ReportAccountReportPartnerledgers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportPartnerledger))
	}
	return s
}
