package types

import (
	"time"
)

type ResourceCalendarLeaves struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CalendarId  Many2One  `xmlrpc:"calendar_id"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	DateTo      time.Time `xmlrpc:"date_to"`
	DisplayName string    `xmlrpc:"display_name"`
	HolidayId   Many2One  `xmlrpc:"holiday_id"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	ResourceId  Many2One  `xmlrpc:"resource_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResourceCalendarLeavesNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CalendarId  interface{} `xmlrpc:"calendar_id"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	DateTo      interface{} `xmlrpc:"date_to"`
	DisplayName interface{} `xmlrpc:"display_name"`
	HolidayId   interface{} `xmlrpc:"holiday_id"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	ResourceId  interface{} `xmlrpc:"resource_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResourceCalendarLeavesModel string = "resource.calendar.leaves"

type ResourceCalendarLeavess []ResourceCalendarLeaves

type ResourceCalendarLeavessNil []ResourceCalendarLeavesNil

func (s *ResourceCalendarLeaves) NilableType_() interface{} {
	return &ResourceCalendarLeavesNil{}
}

func (ns *ResourceCalendarLeavesNil) Type_() interface{} {
	s := &ResourceCalendarLeaves{}
	return load(ns, s)
}

func (s *ResourceCalendarLeavess) NilableType_() interface{} {
	return &ResourceCalendarLeavessNil{}
}

func (ns *ResourceCalendarLeavessNil) Type_() interface{} {
	s := &ResourceCalendarLeavess{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResourceCalendarLeaves))
	}
	return s
}
