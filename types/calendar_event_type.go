package types

import (
	"time"
)

type CalendarEventType struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CalendarEventTypeNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CalendarEventTypeModel string = "calendar.event.type"

type CalendarEventTypes []CalendarEventType

type CalendarEventTypesNil []CalendarEventTypeNil

func (s *CalendarEventType) NilableType_() interface{} {
	return &CalendarEventTypeNil{}
}

func (ns *CalendarEventTypeNil) Type_() interface{} {
	s := &CalendarEventType{}
	return load(ns, s)
}

func (s *CalendarEventTypes) NilableType_() interface{} {
	return &CalendarEventTypesNil{}
}

func (ns *CalendarEventTypesNil) Type_() interface{} {
	s := &CalendarEventTypes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CalendarEventType))
	}
	return s
}
