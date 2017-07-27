package types

import (
	"time"
)

type MailWizardInvite struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ChannelIds  []int64   `xmlrpc:"channel_ids"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Message     string    `xmlrpc:"message"`
	PartnerIds  []int64   `xmlrpc:"partner_ids"`
	ResId       int64     `xmlrpc:"res_id"`
	ResModel    string    `xmlrpc:"res_model"`
	SendMail    bool      `xmlrpc:"send_mail"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type MailWizardInviteNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ChannelIds  interface{} `xmlrpc:"channel_ids"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Message     interface{} `xmlrpc:"message"`
	PartnerIds  interface{} `xmlrpc:"partner_ids"`
	ResId       interface{} `xmlrpc:"res_id"`
	ResModel    interface{} `xmlrpc:"res_model"`
	SendMail    bool        `xmlrpc:"send_mail"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var MailWizardInviteModel string = "mail.wizard.invite"

type MailWizardInvites []MailWizardInvite

type MailWizardInvitesNil []MailWizardInviteNil

func (s *MailWizardInvite) NilableType_() interface{} {
	return &MailWizardInviteNil{}
}

func (ns *MailWizardInviteNil) Type_() interface{} {
	s := &MailWizardInvite{}
	return load(ns, s)
}

func (s *MailWizardInvites) NilableType_() interface{} {
	return &MailWizardInvitesNil{}
}

func (ns *MailWizardInvitesNil) Type_() interface{} {
	s := &MailWizardInvites{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailWizardInvite))
	}
	return s
}
