package types

import (
	"time"
)

type AccountCommonAccountReport struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DateFrom       time.Time `xmlrpc:"date_from"`
	DateTo         time.Time `xmlrpc:"date_to"`
	DisplayAccount string    `xmlrpc:"display_account"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	JournalIds     []int64   `xmlrpc:"journal_ids"`
	TargetMove     string    `xmlrpc:"target_move"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type AccountCommonAccountReportNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DateFrom       interface{} `xmlrpc:"date_from"`
	DateTo         interface{} `xmlrpc:"date_to"`
	DisplayAccount interface{} `xmlrpc:"display_account"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	JournalIds     interface{} `xmlrpc:"journal_ids"`
	TargetMove     interface{} `xmlrpc:"target_move"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var AccountCommonAccountReportModel string = "account.common.account.report"

type AccountCommonAccountReports []AccountCommonAccountReport

type AccountCommonAccountReportsNil []AccountCommonAccountReportNil

func (s *AccountCommonAccountReport) NilableType_() interface{} {
	return &AccountCommonAccountReportNil{}
}

func (ns *AccountCommonAccountReportNil) Type_() interface{} {
	s := &AccountCommonAccountReport{}
	return load(ns, s)
}

func (s *AccountCommonAccountReports) NilableType_() interface{} {
	return &AccountCommonAccountReportsNil{}
}

func (ns *AccountCommonAccountReportsNil) Type_() interface{} {
	s := &AccountCommonAccountReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountCommonAccountReport))
	}
	return s
}
