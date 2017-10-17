package types

import (
	"time"
)

type MailMassMailingContact struct {
	CompanyName              string    `xmlrpc:"company_name"`
	CountryId                Many2One  `xmlrpc:"country_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DisplayName              string    `xmlrpc:"display_name"`
	Email                    string    `xmlrpc:"email"`
	Id                       int64     `xmlrpc:"id"`
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	ListIds                  []int64   `xmlrpc:"list_ids"`
	MessageBounce            int64     `xmlrpc:"message_bounce"`
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
	OptOut                   bool      `xmlrpc:"opt_out"`
	TagIds                   []int64   `xmlrpc:"tag_ids"`
	TitleId                  Many2One  `xmlrpc:"title_id"`
	UnsubscriptionDate       time.Time `xmlrpc:"unsubscription_date"`
	WebsiteMessageIds        []int64   `xmlrpc:"website_message_ids"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingContactNil struct {
	CompanyName              interface{} `xmlrpc:"company_name"`
	CountryId                interface{} `xmlrpc:"country_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Email                    interface{} `xmlrpc:"email"`
	Id                       interface{} `xmlrpc:"id"`
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	ListIds                  interface{} `xmlrpc:"list_ids"`
	MessageBounce            interface{} `xmlrpc:"message_bounce"`
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
	OptOut                   bool        `xmlrpc:"opt_out"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	TitleId                  interface{} `xmlrpc:"title_id"`
	UnsubscriptionDate       interface{} `xmlrpc:"unsubscription_date"`
	WebsiteMessageIds        interface{} `xmlrpc:"website_message_ids"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingContactModel string = "mail.mass_mailing.contact"

type MailMassMailingContacts []MailMassMailingContact

type MailMassMailingContactsNil []MailMassMailingContactNil

func (s *MailMassMailingContact) NilableType_() interface{} {
	return &MailMassMailingContactNil{}
}

func (ns *MailMassMailingContactNil) Type_() interface{} {
	s := &MailMassMailingContact{}
	return load(ns, s)
}

func (s *MailMassMailingContacts) NilableType_() interface{} {
	return &MailMassMailingContactsNil{}
}

func (ns *MailMassMailingContactsNil) Type_() interface{} {
	s := &MailMassMailingContacts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailingContact))
	}
	return s
}
