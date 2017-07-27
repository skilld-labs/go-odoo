package types

import (
	"time"
)

type MailMassMailingTest struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	EmailTo       string    `xmlrpc:"email_to"`
	Id            int64     `xmlrpc:"id"`
	MassMailingId Many2One  `xmlrpc:"mass_mailing_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingTestNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	EmailTo       interface{} `xmlrpc:"email_to"`
	Id            interface{} `xmlrpc:"id"`
	MassMailingId interface{} `xmlrpc:"mass_mailing_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingTestModel string = "mail.mass_mailing.test"

type MailMassMailingTests []MailMassMailingTest

type MailMassMailingTestsNil []MailMassMailingTestNil

func (s *MailMassMailingTest) NilableType_() interface{} {
	return &MailMassMailingTestNil{}
}

func (ns *MailMassMailingTestNil) Type_() interface{} {
	s := &MailMassMailingTest{}
	return load(ns, s)
}

func (s *MailMassMailingTests) NilableType_() interface{} {
	return &MailMassMailingTestsNil{}
}

func (ns *MailMassMailingTestsNil) Type_() interface{} {
	s := &MailMassMailingTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailingTest))
	}
	return s
}
