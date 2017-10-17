package types

import (
	"time"
)

type MailMassMailingStage struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	Sequence    int64     `xmlrpc:"sequence"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingStageNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	Sequence    interface{} `xmlrpc:"sequence"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingStageModel string = "mail.mass_mailing.stage"

type MailMassMailingStages []MailMassMailingStage

type MailMassMailingStagesNil []MailMassMailingStageNil

func (s *MailMassMailingStage) NilableType_() interface{} {
	return &MailMassMailingStageNil{}
}

func (ns *MailMassMailingStageNil) Type_() interface{} {
	s := &MailMassMailingStage{}
	return load(ns, s)
}

func (s *MailMassMailingStages) NilableType_() interface{} {
	return &MailMassMailingStagesNil{}
}

func (ns *MailMassMailingStagesNil) Type_() interface{} {
	s := &MailMassMailingStages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailingStage))
	}
	return s
}
