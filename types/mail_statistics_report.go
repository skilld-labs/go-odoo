package types

import (
	"time"
)

type MailStatisticsReport struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	Bounced       int64     `xmlrpc:"bounced"`
	Campaign      string    `xmlrpc:"campaign"`
	Delivered     int64     `xmlrpc:"delivered"`
	DisplayName   string    `xmlrpc:"display_name"`
	EmailFrom     string    `xmlrpc:"email_from"`
	Id            int64     `xmlrpc:"id"`
	Name          string    `xmlrpc:"name"`
	Opened        int64     `xmlrpc:"opened"`
	Replied       int64     `xmlrpc:"replied"`
	ScheduledDate time.Time `xmlrpc:"scheduled_date"`
	Sent          int64     `xmlrpc:"sent"`
	State         string    `xmlrpc:"state"`
}

type MailStatisticsReportNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	Bounced       interface{} `xmlrpc:"bounced"`
	Campaign      interface{} `xmlrpc:"campaign"`
	Delivered     interface{} `xmlrpc:"delivered"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	EmailFrom     interface{} `xmlrpc:"email_from"`
	Id            interface{} `xmlrpc:"id"`
	Name          interface{} `xmlrpc:"name"`
	Opened        interface{} `xmlrpc:"opened"`
	Replied       interface{} `xmlrpc:"replied"`
	ScheduledDate interface{} `xmlrpc:"scheduled_date"`
	Sent          interface{} `xmlrpc:"sent"`
	State         interface{} `xmlrpc:"state"`
}

var MailStatisticsReportModel string = "mail.statistics.report"

type MailStatisticsReports []MailStatisticsReport

type MailStatisticsReportsNil []MailStatisticsReportNil

func (s *MailStatisticsReport) NilableType_() interface{} {
	return &MailStatisticsReportNil{}
}

func (ns *MailStatisticsReportNil) Type_() interface{} {
	s := &MailStatisticsReport{}
	return load(ns, s)
}

func (s *MailStatisticsReports) NilableType_() interface{} {
	return &MailStatisticsReportsNil{}
}

func (ns *MailStatisticsReportsNil) Type_() interface{} {
	s := &MailStatisticsReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailStatisticsReport))
	}
	return s
}
