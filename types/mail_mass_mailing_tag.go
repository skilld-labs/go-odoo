package types

import (
	"time"
)

type MailMassMailingTag struct {
	Color       int64     `xmlrpc:"color"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingTagNil struct {
	Color       interface{} `xmlrpc:"color"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingTagModel string = "mail.mass_mailing.tag"

type MailMassMailingTags []MailMassMailingTag

type MailMassMailingTagsNil []MailMassMailingTagNil

func (s *MailMassMailingTag) NilableType_() interface{} {
	return &MailMassMailingTagNil{}
}

func (ns *MailMassMailingTagNil) Type_() interface{} {
	s := &MailMassMailingTag{}
	return load(ns, s)
}

func (s *MailMassMailingTags) NilableType_() interface{} {
	return &MailMassMailingTagsNil{}
}

func (ns *MailMassMailingTagsNil) Type_() interface{} {
	s := &MailMassMailingTags{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailingTag))
	}
	return s
}
