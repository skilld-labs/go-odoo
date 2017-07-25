package types

import (
	"time"
)

type ResourceCalendarAttendance struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CalendarId  Many2One  `xmlrpc:"calendar_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	DateTo      time.Time `xmlrpc:"date_to"`
	Dayofweek   string    `xmlrpc:"dayofweek"`
	DisplayName string    `xmlrpc:"display_name"`
	HourFrom    float64   `xmlrpc:"hour_from"`
	HourTo      float64   `xmlrpc:"hour_to"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResourceCalendarAttendanceNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CalendarId  interface{} `xmlrpc:"calendar_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	DateTo      interface{} `xmlrpc:"date_to"`
	Dayofweek   interface{} `xmlrpc:"dayofweek"`
	DisplayName interface{} `xmlrpc:"display_name"`
	HourFrom    interface{} `xmlrpc:"hour_from"`
	HourTo      interface{} `xmlrpc:"hour_to"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResourceCalendarAttendanceModel string = "resource.calendar.attendance"

type ResourceCalendarAttendances []ResourceCalendarAttendance

type ResourceCalendarAttendancesNil []ResourceCalendarAttendanceNil

func (s *ResourceCalendarAttendance) NilableType_() interface{} {
	return &ResourceCalendarAttendanceNil{}
}

func (ns *ResourceCalendarAttendanceNil) Type_() interface{} {
	s := &ResourceCalendarAttendance{}
	return load(ns, s)
}

func (s *ResourceCalendarAttendances) NilableType_() interface{} {
	return &ResourceCalendarAttendancesNil{}
}

func (ns *ResourceCalendarAttendancesNil) Type_() interface{} {
	s := &ResourceCalendarAttendances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResourceCalendarAttendance))
	}
	return s
}
