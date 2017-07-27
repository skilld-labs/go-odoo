package types

import (
	"time"
)

type ReportProjectTaskUser struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	ClosingDays         float64   `xmlrpc:"closing_days"`
	CompanyId           Many2One  `xmlrpc:"company_id"`
	DateDeadline        time.Time `xmlrpc:"date_deadline"`
	DateEnd             time.Time `xmlrpc:"date_end"`
	DateLastStageUpdate time.Time `xmlrpc:"date_last_stage_update"`
	DateStart           time.Time `xmlrpc:"date_start"`
	DelayEndingsDays    float64   `xmlrpc:"delay_endings_days"`
	DisplayName         string    `xmlrpc:"display_name"`
	HoursDelay          float64   `xmlrpc:"hours_delay"`
	HoursEffective      float64   `xmlrpc:"hours_effective"`
	HoursPlanned        float64   `xmlrpc:"hours_planned"`
	Id                  int64     `xmlrpc:"id"`
	Name                string    `xmlrpc:"name"`
	Nbr                 int64     `xmlrpc:"nbr"`
	NoOfDays            int64     `xmlrpc:"no_of_days"`
	OpeningDays         float64   `xmlrpc:"opening_days"`
	PartnerId           Many2One  `xmlrpc:"partner_id"`
	Priority            string    `xmlrpc:"priority"`
	Progress            float64   `xmlrpc:"progress"`
	ProjectId           Many2One  `xmlrpc:"project_id"`
	RemainingHours      float64   `xmlrpc:"remaining_hours"`
	StageId             Many2One  `xmlrpc:"stage_id"`
	State               string    `xmlrpc:"state"`
	TotalHours          float64   `xmlrpc:"total_hours"`
	UserId              Many2One  `xmlrpc:"user_id"`
}

type ReportProjectTaskUserNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	ClosingDays         interface{} `xmlrpc:"closing_days"`
	CompanyId           interface{} `xmlrpc:"company_id"`
	DateDeadline        interface{} `xmlrpc:"date_deadline"`
	DateEnd             interface{} `xmlrpc:"date_end"`
	DateLastStageUpdate interface{} `xmlrpc:"date_last_stage_update"`
	DateStart           interface{} `xmlrpc:"date_start"`
	DelayEndingsDays    interface{} `xmlrpc:"delay_endings_days"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	HoursDelay          interface{} `xmlrpc:"hours_delay"`
	HoursEffective      interface{} `xmlrpc:"hours_effective"`
	HoursPlanned        interface{} `xmlrpc:"hours_planned"`
	Id                  interface{} `xmlrpc:"id"`
	Name                interface{} `xmlrpc:"name"`
	Nbr                 interface{} `xmlrpc:"nbr"`
	NoOfDays            interface{} `xmlrpc:"no_of_days"`
	OpeningDays         interface{} `xmlrpc:"opening_days"`
	PartnerId           interface{} `xmlrpc:"partner_id"`
	Priority            interface{} `xmlrpc:"priority"`
	Progress            interface{} `xmlrpc:"progress"`
	ProjectId           interface{} `xmlrpc:"project_id"`
	RemainingHours      interface{} `xmlrpc:"remaining_hours"`
	StageId             interface{} `xmlrpc:"stage_id"`
	State               interface{} `xmlrpc:"state"`
	TotalHours          interface{} `xmlrpc:"total_hours"`
	UserId              interface{} `xmlrpc:"user_id"`
}

var ReportProjectTaskUserModel string = "report.project.task.user"

type ReportProjectTaskUsers []ReportProjectTaskUser

type ReportProjectTaskUsersNil []ReportProjectTaskUserNil

func (s *ReportProjectTaskUser) NilableType_() interface{} {
	return &ReportProjectTaskUserNil{}
}

func (ns *ReportProjectTaskUserNil) Type_() interface{} {
	s := &ReportProjectTaskUser{}
	return load(ns, s)
}

func (s *ReportProjectTaskUsers) NilableType_() interface{} {
	return &ReportProjectTaskUsersNil{}
}

func (ns *ReportProjectTaskUsersNil) Type_() interface{} {
	s := &ReportProjectTaskUsers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ReportProjectTaskUser))
	}
	return s
}
