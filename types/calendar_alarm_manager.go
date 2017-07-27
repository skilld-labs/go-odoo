package types

import (
	"time"
)

type CalendarAlarmManager struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type CalendarAlarmManagerNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var CalendarAlarmManagerModel string = "calendar.alarm_manager"

type CalendarAlarmManagers []CalendarAlarmManager

type CalendarAlarmManagersNil []CalendarAlarmManagerNil

func (s *CalendarAlarmManager) NilableType_() interface{} {
	return &CalendarAlarmManagerNil{}
}

func (ns *CalendarAlarmManagerNil) Type_() interface{} {
	s := &CalendarAlarmManager{}
	return load(ns, s)
}

func (s *CalendarAlarmManagers) NilableType_() interface{} {
	return &CalendarAlarmManagersNil{}
}

func (ns *CalendarAlarmManagersNil) Type_() interface{} {
	s := &CalendarAlarmManagers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CalendarAlarmManager))
	}
	return s
}
