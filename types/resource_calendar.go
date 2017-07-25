package types

import (
	"time"
)

type ResourceCalendar struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	AttendanceIds []int64   `xmlrpc:"attendance_ids"`
	CompanyId     Many2One  `xmlrpc:"company_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	LeaveIds      []int64   `xmlrpc:"leave_ids"`
	Manager       Many2One  `xmlrpc:"manager"`
	Name          string    `xmlrpc:"name"`
	UomId         Many2One  `xmlrpc:"uom_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type ResourceCalendarNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	AttendanceIds interface{} `xmlrpc:"attendance_ids"`
	CompanyId     interface{} `xmlrpc:"company_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	LeaveIds      interface{} `xmlrpc:"leave_ids"`
	Manager       interface{} `xmlrpc:"manager"`
	Name          interface{} `xmlrpc:"name"`
	UomId         interface{} `xmlrpc:"uom_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
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
