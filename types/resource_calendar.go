package types

import (
	"time"
)

type ResourceCalendar struct {
	AttendanceIds  []int64   `xmlrpc:"attendance_ids"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	GlobalLeaveIds []int64   `xmlrpc:"global_leave_ids"`
	Id             int64     `xmlrpc:"id"`
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	LeaveIds       []int64   `xmlrpc:"leave_ids"`
	Name           string    `xmlrpc:"name"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type ResourceCalendarNil struct {
	AttendanceIds  interface{} `xmlrpc:"attendance_ids"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	GlobalLeaveIds interface{} `xmlrpc:"global_leave_ids"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	LeaveIds       interface{} `xmlrpc:"leave_ids"`
	Name           interface{} `xmlrpc:"name"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var ResourceCalendarModel string = "resource.calendar"

type ResourceCalendars []ResourceCalendar

type ResourceCalendarsNil []ResourceCalendarNil

func (s *ResourceCalendar) NilableType_() interface{} {
	return &ResourceCalendarNil{}
}

func (ns *ResourceCalendarNil) Type_() interface{} {
	s := &ResourceCalendar{}
	return load(ns, s)
}

func (s *ResourceCalendars) NilableType_() interface{} {
	return &ResourceCalendarsNil{}
}

func (ns *ResourceCalendarsNil) Type_() interface{} {
	s := &ResourceCalendars{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResourceCalendar))
	}
	return s
}
