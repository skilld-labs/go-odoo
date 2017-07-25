package types

import (
	"time"
)

type AccountCommonPartnerReport struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CompanyId       Many2One  `xmlrpc:"company_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DateFrom        time.Time `xmlrpc:"date_from"`
	DateTo          time.Time `xmlrpc:"date_to"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	JournalIds      []int64   `xmlrpc:"journal_ids"`
	ResultSelection string    `xmlrpc:"result_selection"`
	TargetMove      string    `xmlrpc:"target_move"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountCommonPartnerReportNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CompanyId       interface{} `xmlrpc:"company_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DateFrom        interface{} `xmlrpc:"date_from"`
	DateTo          interface{} `xmlrpc:"date_to"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	JournalIds      interface{} `xmlrpc:"journal_ids"`
	ResultSelection interface{} `xmlrpc:"result_selection"`
	TargetMove      interface{} `xmlrpc:"target_move"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountCommonPartnerReportModel string = "account.common.partner.report"

type AccountCommonPartnerReports []AccountCommonPartnerReport

type AccountCommonPartnerReportsNil []AccountCommonPartnerReportNil

func (s *AccountCommonPartnerReport) NilableType_() interface{} {
	return &AccountCommonPartnerReportNil{}
}

func (ns *AccountCommonPartnerReportNil) Type_() interface{} {
	s := &AccountCommonPartnerReport{}
	return load(ns, s)
}

func (s *AccountCommonPartnerReports) NilableType_() interface{} {
	return &AccountCommonPartnerReportsNil{}
}

func (ns *AccountCommonPartnerReportsNil) Type_() interface{} {
	s := &AccountCommonPartnerReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountCommonPartnerReport))
	}
	return s
}
