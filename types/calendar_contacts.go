package types

import (
	"time"
)

type CalendarContacts struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PartnerId   Many2One  `xmlrpc:"partner_id"`
	UserId      Many2One  `xmlrpc:"user_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CalendarContactsNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	UserId      interface{} `xmlrpc:"user_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CalendarContactsModel string = "calendar.contacts"

type CalendarContactss []CalendarContacts

type CalendarContactssNil []CalendarContactsNil

func (s *CalendarContacts) NilableType_() interface{} {
	return &CalendarContactsNil{}
}

func (ns *CalendarContactsNil) Type_() interface{} {
	s := &CalendarContacts{}
	return load(ns, s)
}

func (s *CalendarContactss) NilableType_() interface{} {
	return &CalendarContactssNil{}
}

func (ns *CalendarContactssNil) Type_() interface{} {
	s := &CalendarContactss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CalendarContacts))
	}
	return s
}
