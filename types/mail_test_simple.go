package types

import (
	"time"
)

type MailTestSimple struct {
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	Description              string    `xmlrpc:"description"`
	DisplayName              string    `xmlrpc:"display_name"`
	EmailFrom                string    `xmlrpc:"email_from"`
	Id                       int64     `xmlrpc:"id"`
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	MessageChannelIds        []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64   `xmlrpc:"message_follower_ids"`
	MessageIds               []int64   `xmlrpc:"message_ids"`
	MessageIsFollower        bool      `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction        bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread            bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64     `xmlrpc:"message_unread_counter"`
	Name                     string    `xmlrpc:"name"`
	WebsiteMessageIds        []int64   `xmlrpc:"website_message_ids"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type MailTestSimpleNil struct {
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Description              interface{} `xmlrpc:"description"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	EmailFrom                interface{} `xmlrpc:"email_from"`
	Id                       interface{} `xmlrpc:"id"`
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	Name                     interface{} `xmlrpc:"name"`
	WebsiteMessageIds        interface{} `xmlrpc:"website_message_ids"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var MailTestSimpleModel string = "mail.test.simple"

type MailTestSimples []MailTestSimple

type MailTestSimplesNil []MailTestSimpleNil

func (s *MailTestSimple) NilableType_() interface{} {
	return &MailTestSimpleNil{}
}

func (ns *MailTestSimpleNil) Type_() interface{} {
	s := &MailTestSimple{}
	return load(ns, s)
}

func (s *MailTestSimples) NilableType_() interface{} {
	return &MailTestSimplesNil{}
}

func (ns *MailTestSimplesNil) Type_() interface{} {
	s := &MailTestSimples{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailTestSimple))
	}
	return s
}
