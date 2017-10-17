package types

import (
	"time"
)

type ReportAccountReportJournal struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type ReportAccountReportJournalNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var ReportAccountReportJournalModel string = "report.account.report_journal"

type ReportAccountReportJournals []ReportAccountReportJournal

type ReportAccountReportJournalsNil []ReportAccountReportJournalNil

func (s *ReportAccountReportJournal) NilableType_() interface{} {
	return &ReportAccountReportJournalNil{}
}

func (ns *ReportAccountReportJournalNil) Type_() interface{} {
	s := &ReportAccountReportJournal{}
	return load(ns, s)
}

func (s *ReportAccountReportJournals) NilableType_() interface{} {
	return &ReportAccountReportJournalsNil{}
}

func (ns *ReportAccountReportJournalsNil) Type_() interface{} {
	s := &ReportAccountReportJournals{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportAccountReportJournal))
	}
	return s
}
