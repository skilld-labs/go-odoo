package types

import (
	"time"
)

type MailActivity struct {
	ActivityCategory          interface{} `xmlrpc:"activity_category"`
	ActivityTypeId            Many2One    `xmlrpc:"activity_type_id"`
	CalendarEventId           Many2One    `xmlrpc:"calendar_event_id"`
	CreateDate                time.Time   `xmlrpc:"create_date"`
	CreateUid                 Many2One    `xmlrpc:"create_uid"`
	DateDeadline              time.Time   `xmlrpc:"date_deadline"`
	DisplayName               string      `xmlrpc:"display_name"`
	Feedback                  string      `xmlrpc:"feedback"`
	HasRecommendedActivities  bool        `xmlrpc:"has_recommended_activities"`
	Icon                      string      `xmlrpc:"icon"`
	Id                        int64       `xmlrpc:"id"`
	LastUpdate                time.Time   `xmlrpc:"__last_update"`
	Note                      string      `xmlrpc:"note"`
	PreviousActivityTypeId    Many2One    `xmlrpc:"previous_activity_type_id"`
	RecommendedActivityTypeId Many2One    `xmlrpc:"recommended_activity_type_id"`
	ResId                     int64       `xmlrpc:"res_id"`
	ResModel                  string      `xmlrpc:"res_model"`
	ResModelId                Many2One    `xmlrpc:"res_model_id"`
	ResName                   string      `xmlrpc:"res_name"`
	State                     interface{} `xmlrpc:"state"`
	Summary                   string      `xmlrpc:"summary"`
	UserId                    Many2One    `xmlrpc:"user_id"`
	WriteDate                 time.Time   `xmlrpc:"write_date"`
	WriteUid                  Many2One    `xmlrpc:"write_uid"`
}

type MailActivityNil struct {
	ActivityCategory          interface{} `xmlrpc:"activity_category"`
	ActivityTypeId            interface{} `xmlrpc:"activity_type_id"`
	CalendarEventId           interface{} `xmlrpc:"calendar_event_id"`
	CreateDate                interface{} `xmlrpc:"create_date"`
	CreateUid                 interface{} `xmlrpc:"create_uid"`
	DateDeadline              interface{} `xmlrpc:"date_deadline"`
	DisplayName               interface{} `xmlrpc:"display_name"`
	Feedback                  interface{} `xmlrpc:"feedback"`
	HasRecommendedActivities  bool        `xmlrpc:"has_recommended_activities"`
	Icon                      interface{} `xmlrpc:"icon"`
	Id                        interface{} `xmlrpc:"id"`
	LastUpdate                interface{} `xmlrpc:"__last_update"`
	Note                      interface{} `xmlrpc:"note"`
	PreviousActivityTypeId    interface{} `xmlrpc:"previous_activity_type_id"`
	RecommendedActivityTypeId interface{} `xmlrpc:"recommended_activity_type_id"`
	ResId                     interface{} `xmlrpc:"res_id"`
	ResModel                  interface{} `xmlrpc:"res_model"`
	ResModelId                interface{} `xmlrpc:"res_model_id"`
	ResName                   interface{} `xmlrpc:"res_name"`
	State                     interface{} `xmlrpc:"state"`
	Summary                   interface{} `xmlrpc:"summary"`
	UserId                    interface{} `xmlrpc:"user_id"`
	WriteDate                 interface{} `xmlrpc:"write_date"`
	WriteUid                  interface{} `xmlrpc:"write_uid"`
}

var MailActivityModel string = "mail.activity"

type MailActivitys []MailActivity

type MailActivitysNil []MailActivityNil

func (s *MailActivity) NilableType_() interface{} {
	return &MailActivityNil{}
}

func (ns *MailActivityNil) Type_() interface{} {
	s := &MailActivity{}
	return load(ns, s)
}

func (s *MailActivitys) NilableType_() interface{} {
	return &MailActivitysNil{}
}

func (ns *MailActivitysNil) Type_() interface{} {
	s := &MailActivitys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailActivity))
	}
	return s
}
