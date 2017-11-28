package types

import (
	"time"
)

type MailTest struct {
	AliasContact             interface{} `xmlrpc:"alias_contact"`
	AliasDefaults            string      `xmlrpc:"alias_defaults"`
	AliasDomain              string      `xmlrpc:"alias_domain"`
	AliasForceThreadId       int64       `xmlrpc:"alias_force_thread_id"`
	AliasId                  Many2One    `xmlrpc:"alias_id"`
	AliasModelId             Many2One    `xmlrpc:"alias_model_id"`
	AliasName                string      `xmlrpc:"alias_name"`
	AliasParentModelId       Many2One    `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId      int64       `xmlrpc:"alias_parent_thread_id"`
	AliasUserId              Many2One    `xmlrpc:"alias_user_id"`
	CreateDate               time.Time   `xmlrpc:"create_date"`
	CreateUid                Many2One    `xmlrpc:"create_uid"`
	Description              string      `xmlrpc:"description"`
	DisplayName              string      `xmlrpc:"display_name"`
	Id                       int64       `xmlrpc:"id"`
	LastUpdate               time.Time   `xmlrpc:"__last_update"`
	MessageChannelIds        []int64     `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64     `xmlrpc:"message_follower_ids"`
	MessageIds               []int64     `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time   `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64       `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64     `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64       `xmlrpc:"message_unread_counter"`
	Name                     string      `xmlrpc:"name"`
	WebsiteMessageIds        []int64     `xmlrpc:"website_message_ids"`
	WriteDate                time.Time   `xmlrpc:"write_date"`
	WriteUid                 Many2One    `xmlrpc:"write_uid"`
}

type MailTestNil struct {
	AliasContact             interface{} `xmlrpc:"alias_contact"`
	AliasDefaults            interface{} `xmlrpc:"alias_defaults"`
	AliasDomain              interface{} `xmlrpc:"alias_domain"`
	AliasForceThreadId       interface{} `xmlrpc:"alias_force_thread_id"`
	AliasId                  interface{} `xmlrpc:"alias_id"`
	AliasModelId             interface{} `xmlrpc:"alias_model_id"`
	AliasName                interface{} `xmlrpc:"alias_name"`
	AliasParentModelId       interface{} `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId      interface{} `xmlrpc:"alias_parent_thread_id"`
	AliasUserId              interface{} `xmlrpc:"alias_user_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Description              interface{} `xmlrpc:"description"`
	DisplayName              interface{} `xmlrpc:"display_name"`
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

var MailTestModel string = "mail.test"

type MailTests []MailTest

type MailTestsNil []MailTestNil

func (s *MailTest) NilableType_() interface{} {
	return &MailTestNil{}
}

func (ns *MailTestNil) Type_() interface{} {
	s := &MailTest{}
	return load(ns, s)
}

func (s *MailTests) NilableType_() interface{} {
	return &MailTestsNil{}
}

func (ns *MailTestsNil) Type_() interface{} {
	s := &MailTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailTest))
	}
	return s
}
