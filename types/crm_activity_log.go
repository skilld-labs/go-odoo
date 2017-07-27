package types

import (
	"time"
)

type CrmActivityLog struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DateAction            time.Time `xmlrpc:"date_action"`
	DateDeadline          time.Time `xmlrpc:"date_deadline"`
	DisplayName           string    `xmlrpc:"display_name"`
	Id                    int64     `xmlrpc:"id"`
	LastActivityId        Many2One  `xmlrpc:"last_activity_id"`
	LeadId                Many2One  `xmlrpc:"lead_id"`
	NextActivityId        Many2One  `xmlrpc:"next_activity_id"`
	Note                  string    `xmlrpc:"note"`
	PlannedRevenue        float64   `xmlrpc:"planned_revenue"`
	RecommendedActivityId Many2One  `xmlrpc:"recommended_activity_id"`
	TeamId                Many2One  `xmlrpc:"team_id"`
	TitleAction           string    `xmlrpc:"title_action"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type CrmActivityLogNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DateAction            interface{} `xmlrpc:"date_action"`
	DateDeadline          interface{} `xmlrpc:"date_deadline"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Id                    interface{} `xmlrpc:"id"`
	LastActivityId        interface{} `xmlrpc:"last_activity_id"`
	LeadId                interface{} `xmlrpc:"lead_id"`
	NextActivityId        interface{} `xmlrpc:"next_activity_id"`
	Note                  interface{} `xmlrpc:"note"`
	PlannedRevenue        interface{} `xmlrpc:"planned_revenue"`
	RecommendedActivityId interface{} `xmlrpc:"recommended_activity_id"`
	TeamId                interface{} `xmlrpc:"team_id"`
	TitleAction           interface{} `xmlrpc:"title_action"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var CrmActivityLogModel string = "crm.activity.log"

type CrmActivityLogs []CrmActivityLog

type CrmActivityLogsNil []CrmActivityLogNil

func (s *CrmActivityLog) NilableType_() interface{} {
	return &CrmActivityLogNil{}
}

func (ns *CrmActivityLogNil) Type_() interface{} {
	s := &CrmActivityLog{}
	return load(ns, s)
}

func (s *CrmActivityLogs) NilableType_() interface{} {
	return &CrmActivityLogsNil{}
}

func (ns *CrmActivityLogsNil) Type_() interface{} {
	s := &CrmActivityLogs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmActivityLog))
	}
	return s
}
