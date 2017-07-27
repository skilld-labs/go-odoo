package types

import (
	"time"
)

type ReportAccountReportTrialbalance struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type ReportAccountReportTrialbalanceNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var ReportAccountReportTrialbalanceModel string = "report.account.report_trialbalance"

type ReportAccountReportTrialbalances []ReportAccountReportTrialbalance

type ReportAccountReportTrialbalancesNil []ReportAccountReportTrialbalanceNil

func (s *ReportAccountReportTrialbalance) NilableType_() interface{} {
	return &ReportAccountReportTrialbalanceNil{}
}

func (ns *ReportAccountReportTrialbalanceNil) Type_() interface{} {
	s := &ReportAccountReportTrialbalance{}
	return load(ns, s)
}

func (s *ReportAccountReportTrialbalances) NilableType_() interface{} {
	return &ReportAccountReportTrialbalancesNil{}
}

func (ns *ReportAccountReportTrialbalancesNil) Type_() interface{} {
	s := &ReportAccountReportTrialbalances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportTrialbalance))
	}
	return s
}
