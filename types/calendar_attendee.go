package types

import (
	"time"
)

type CalendarAttendee struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	AccessToken  string    `xmlrpc:"access_token"`
	Availability string    `xmlrpc:"availability"`
	CommonName   string    `xmlrpc:"common_name"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	Email        string    `xmlrpc:"email"`
	EventId      Many2One  `xmlrpc:"event_id"`
	Id           int64     `xmlrpc:"id"`
	PartnerId    Many2One  `xmlrpc:"partner_id"`
	State        string    `xmlrpc:"state"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type CalendarAttendeeNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	AccessToken  interface{} `xmlrpc:"access_token"`
	Availability interface{} `xmlrpc:"availability"`
	CommonName   interface{} `xmlrpc:"common_name"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Email        interface{} `xmlrpc:"email"`
	EventId      interface{} `xmlrpc:"event_id"`
	Id           interface{} `xmlrpc:"id"`
	PartnerId    interface{} `xmlrpc:"partner_id"`
	State        interface{} `xmlrpc:"state"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var CalendarAttendeeModel string = "calendar.attendee"

type CalendarAttendees []CalendarAttendee

type CalendarAttendeesNil []CalendarAttendeeNil

func (s *CalendarAttendee) NilableType_() interface{} {
	return &CalendarAttendeeNil{}
}

func (ns *CalendarAttendeeNil) Type_() interface{} {
	s := &CalendarAttendee{}
	return load(ns, s)
}

func (s *CalendarAttendees) NilableType_() interface{} {
	return &CalendarAttendeesNil{}
}

func (ns *CalendarAttendeesNil) Type_() interface{} {
	s := &CalendarAttendees{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CalendarAttendee))
	}
	return s
}
