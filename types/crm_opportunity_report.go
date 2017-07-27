package types

import (
	"time"
)

type CrmOpportunityReport struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	Active              bool      `xmlrpc:"active"`
	CampaignId          Many2One  `xmlrpc:"campaign_id"`
	City                string    `xmlrpc:"city"`
	CompanyId           Many2One  `xmlrpc:"company_id"`
	CountryId           Many2One  `xmlrpc:"country_id"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	DateClosed          time.Time `xmlrpc:"date_closed"`
	DateConversion      time.Time `xmlrpc:"date_conversion"`
	DateDeadline        time.Time `xmlrpc:"date_deadline"`
	DateLastStageUpdate time.Time `xmlrpc:"date_last_stage_update"`
	DelayClose          float64   `xmlrpc:"delay_close"`
	DelayExpected       float64   `xmlrpc:"delay_expected"`
	DelayOpen           float64   `xmlrpc:"delay_open"`
	DisplayName         string    `xmlrpc:"display_name"`
	ExpectedRevenue     float64   `xmlrpc:"expected_revenue"`
	Id                  int64     `xmlrpc:"id"`
	LostReason          Many2One  `xmlrpc:"lost_reason"`
	MediumId            Many2One  `xmlrpc:"medium_id"`
	NbrActivities       int64     `xmlrpc:"nbr_activities"`
	OpeningDate         time.Time `xmlrpc:"opening_date"`
	PartnerId           Many2One  `xmlrpc:"partner_id"`
	Priority            string    `xmlrpc:"priority"`
	Probability         float64   `xmlrpc:"probability"`
	SourceId            Many2One  `xmlrpc:"source_id"`
	StageId             Many2One  `xmlrpc:"stage_id"`
	StageName           string    `xmlrpc:"stage_name"`
	TeamId              Many2One  `xmlrpc:"team_id"`
	TotalRevenue        float64   `xmlrpc:"total_revenue"`
	Type                string    `xmlrpc:"type"`
	UserId              Many2One  `xmlrpc:"user_id"`
}

type CrmOpportunityReportNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	Active              bool        `xmlrpc:"active"`
	CampaignId          interface{} `xmlrpc:"campaign_id"`
	City                interface{} `xmlrpc:"city"`
	CompanyId           interface{} `xmlrpc:"company_id"`
	CountryId           interface{} `xmlrpc:"country_id"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	DateClosed          interface{} `xmlrpc:"date_closed"`
	DateConversion      interface{} `xmlrpc:"date_conversion"`
	DateDeadline        interface{} `xmlrpc:"date_deadline"`
	DateLastStageUpdate interface{} `xmlrpc:"date_last_stage_update"`
	DelayClose          interface{} `xmlrpc:"delay_close"`
	DelayExpected       interface{} `xmlrpc:"delay_expected"`
	DelayOpen           interface{} `xmlrpc:"delay_open"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	ExpectedRevenue     interface{} `xmlrpc:"expected_revenue"`
	Id                  interface{} `xmlrpc:"id"`
	LostReason          interface{} `xmlrpc:"lost_reason"`
	MediumId            interface{} `xmlrpc:"medium_id"`
	NbrActivities       interface{} `xmlrpc:"nbr_activities"`
	OpeningDate         interface{} `xmlrpc:"opening_date"`
	PartnerId           interface{} `xmlrpc:"partner_id"`
	Priority            interface{} `xmlrpc:"priority"`
	Probability         interface{} `xmlrpc:"probability"`
	SourceId            interface{} `xmlrpc:"source_id"`
	StageId             interface{} `xmlrpc:"stage_id"`
	StageName           interface{} `xmlrpc:"stage_name"`
	TeamId              interface{} `xmlrpc:"team_id"`
	TotalRevenue        interface{} `xmlrpc:"total_revenue"`
	Type                interface{} `xmlrpc:"type"`
	UserId              interface{} `xmlrpc:"user_id"`
}

var CrmOpportunityReportModel string = "crm.opportunity.report"

type CrmOpportunityReports []CrmOpportunityReport

type CrmOpportunityReportsNil []CrmOpportunityReportNil

func (s *CrmOpportunityReport) NilableType_() interface{} {
	return &CrmOpportunityReportNil{}
}

func (ns *CrmOpportunityReportNil) Type_() interface{} {
	s := &CrmOpportunityReport{}
	return load(ns, s)
}

func (s *CrmOpportunityReports) NilableType_() interface{} {
	return &CrmOpportunityReportsNil{}
}

func (ns *CrmOpportunityReportsNil) Type_() interface{} {
	s := &CrmOpportunityReports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmOpportunityReport))
	}
	return s
}
