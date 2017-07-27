package types

import (
	"time"
)

type CrmActivityReport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	AuthorId    Many2One  `xmlrpc:"author_id"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CountryId   Many2One  `xmlrpc:"country_id"`
	Date        time.Time `xmlrpc:"date"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LeadId      Many2One  `xmlrpc:"lead_id"`
	LeadType    string    `xmlrpc:"lead_type"`
	PartnerId   Many2One  `xmlrpc:"partner_id"`
	Probability float64   `xmlrpc:"probability"`
	StageId     Many2One  `xmlrpc:"stage_id"`
	Subject     string    `xmlrpc:"subject"`
	SubtypeId   Many2One  `xmlrpc:"subtype_id"`
	TeamId      Many2One  `xmlrpc:"team_id"`
	UserId      Many2One  `xmlrpc:"user_id"`
}

type CrmActivityReportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	AuthorId    interface{} `xmlrpc:"author_id"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	CountryId   interface{} `xmlrpc:"country_id"`
	Date        interface{} `xmlrpc:"date"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LeadId      interface{} `xmlrpc:"lead_id"`
	LeadType    interface{} `xmlrpc:"lead_type"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	Probability interface{} `xmlrpc:"probability"`
	StageId     interface{} `xmlrpc:"stage_id"`
	Subject     interface{} `xmlrpc:"subject"`
	SubtypeId   interface{} `xmlrpc:"subtype_id"`
	TeamId      interface{} `xmlrpc:"team_id"`
	UserId      interface{} `xmlrpc:"user_id"`
}

var CrmActivityReportModel string = "crm.activity.report"

type CrmActivityReports []CrmActivityReport

type CrmActivityReportsNil []CrmActivityReportNil

func (s *CrmActivityReport) NilableType_() interface{} {
	return &CrmActivityReportNil{}
}

func (ns *CrmActivityReportNil) Type_() interface{} {
	s := &CrmActivityReport{}
	return load(ns, s)
}

func (s *CrmActivityReports) NilableType_() interface{} {
	return &CrmActivityReportsNil{}
}

func (ns *CrmActivityReportsNil) Type_() interface{} {
	s := &CrmActivityReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmActivityReport))
	}
	return s
}
