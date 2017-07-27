package types

import (
	"time"
)

type MailChannelPartner struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	ChannelId     Many2One  `xmlrpc:"channel_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	FoldState     string    `xmlrpc:"fold_state"`
	Id            int64     `xmlrpc:"id"`
	IsMinimized   bool      `xmlrpc:"is_minimized"`
	IsPinned      bool      `xmlrpc:"is_pinned"`
	PartnerId     Many2One  `xmlrpc:"partner_id"`
	SeenMessageId Many2One  `xmlrpc:"seen_message_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type MailChannelPartnerNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	ChannelId     interface{} `xmlrpc:"channel_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	FoldState     interface{} `xmlrpc:"fold_state"`
	Id            interface{} `xmlrpc:"id"`
	IsMinimized   bool        `xmlrpc:"is_minimized"`
	IsPinned      bool        `xmlrpc:"is_pinned"`
	PartnerId     interface{} `xmlrpc:"partner_id"`
	SeenMessageId interface{} `xmlrpc:"seen_message_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var MailChannelPartnerModel string = "mail.channel.partner"

type MailChannelPartners []MailChannelPartner

type MailChannelPartnersNil []MailChannelPartnerNil

func (s *MailChannelPartner) NilableType_() interface{} {
	return &MailChannelPartnerNil{}
}

func (ns *MailChannelPartnerNil) Type_() interface{} {
	s := &MailChannelPartner{}
	return load(ns, s)
}

func (s *MailChannelPartners) NilableType_() interface{} {
	return &MailChannelPartnersNil{}
}

func (ns *MailChannelPartnersNil) Type_() interface{} {
	s := &MailChannelPartners{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailChannelPartner))
	}
	return s
}
