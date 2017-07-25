package types

import (
	"time"
)

type MailMassMailingList struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	ContactNbr  int64     `xmlrpc:"contact_nbr"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingListNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	ContactNbr  interface{} `xmlrpc:"contact_nbr"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingListModel string = "mail.mass_mailing.list"

type MailMassMailingLists []MailMassMailingList

type MailMassMailingListsNil []MailMassMailingListNil

func (s *MailMassMailingList) NilableType_() interface{} {
	return &MailMassMailingListNil{}
}

func (ns *MailMassMailingListNil) Type_() interface{} {
	s := &MailMassMailingList{}
	return load(ns, s)
}

func (s *MailMassMailingLists) NilableType_() interface{} {
	return &MailMassMailingListsNil{}
}

func (ns *MailMassMailingListsNil) Type_() interface{} {
	s := &MailMassMailingLists{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailingList))
	}
	return s
}
