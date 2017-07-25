package types

import (
	"time"
)

type CalendarAlarm struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Duration        int64     `xmlrpc:"duration"`
	DurationMinutes int64     `xmlrpc:"duration_minutes"`
	Id              int64     `xmlrpc:"id"`
	Interval        string    `xmlrpc:"interval"`
	Name            string    `xmlrpc:"name"`
	Type            string    `xmlrpc:"type"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type CalendarAlarmNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Duration        interface{} `xmlrpc:"duration"`
	DurationMinutes interface{} `xmlrpc:"duration_minutes"`
	Id              interface{} `xmlrpc:"id"`
	Interval        interface{} `xmlrpc:"interval"`
	Name            interface{} `xmlrpc:"name"`
	Type            interface{} `xmlrpc:"type"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var CalendarAlarmModel string = "calendar.alarm"

type CalendarAlarms []CalendarAlarm

type CalendarAlarmsNil []CalendarAlarmNil

func (s *CalendarAlarm) NilableType_() interface{} {
	return &CalendarAlarmNil{}
}

func (ns *CalendarAlarmNil) Type_() interface{} {
	s := &CalendarAlarm{}
	return load(ns, s)
}

func (s *CalendarAlarms) NilableType_() interface{} {
	return &CalendarAlarmsNil{}
}

func (ns *CalendarAlarmsNil) Type_() interface{} {
	s := &CalendarAlarms{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CalendarAlarm))
	}
	return s
}
