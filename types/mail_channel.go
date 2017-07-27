package types

import (
	"time"
)

type MailChannel struct {
	LastUpdate                time.Time `xmlrpc:"__last_update"`
	AliasContact              string    `xmlrpc:"alias_contact"`
	AliasDefaults             string    `xmlrpc:"alias_defaults"`
	AliasDomain               string    `xmlrpc:"alias_domain"`
	AliasForceThreadId        int64     `xmlrpc:"alias_force_thread_id"`
	AliasId                   Many2One  `xmlrpc:"alias_id"`
	AliasModelId              Many2One  `xmlrpc:"alias_model_id"`
	AliasName                 string    `xmlrpc:"alias_name"`
	AliasParentModelId        Many2One  `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId       int64     `xmlrpc:"alias_parent_thread_id"`
	AliasUserId               Many2One  `xmlrpc:"alias_user_id"`
	AnonymousName             string    `xmlrpc:"anonymous_name"`
	ChannelLastSeenPartnerIds []int64   `xmlrpc:"channel_last_seen_partner_ids"`
	ChannelMessageIds         []int64   `xmlrpc:"channel_message_ids"`
	ChannelPartnerIds         []int64   `xmlrpc:"channel_partner_ids"`
	ChannelType               string    `xmlrpc:"channel_type"`
	CreateDate                time.Time `xmlrpc:"create_date"`
	CreateUid                 Many2One  `xmlrpc:"create_uid"`
	Description               string    `xmlrpc:"description"`
	DisplayName               string    `xmlrpc:"display_name"`
	EmailSend                 bool      `xmlrpc:"email_send"`
	GroupIds                  []int64   `xmlrpc:"group_ids"`
	GroupPublicId             Many2One  `xmlrpc:"group_public_id"`
	Id                        int64     `xmlrpc:"id"`
	Image                     string    `xmlrpc:"image"`
	ImageMedium               string    `xmlrpc:"image_medium"`
	ImageSmall                string    `xmlrpc:"image_small"`
	IsMember                  bool      `xmlrpc:"is_member"`
	IsSubscribed              bool      `xmlrpc:"is_subscribed"`
	LivechatChannelId         Many2One  `xmlrpc:"livechat_channel_id"`
	MessageChannelIds         []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds        []int64   `xmlrpc:"message_follower_ids"`
	MessageIds                []int64   `xmlrpc:"message_ids"`
	MessageIsFollower         bool      `xmlrpc:"message_is_follower"`
	MessageLastPost           time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction         bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter  int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds         []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread             bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter      int64     `xmlrpc:"message_unread_counter"`
	Name                      string    `xmlrpc:"name"`
	Public                    string    `xmlrpc:"public"`
	RatingCount               int64     `xmlrpc:"rating_count"`
	RatingIds                 []int64   `xmlrpc:"rating_ids"`
	RatingLastValue           float64   `xmlrpc:"rating_last_value"`
	Uuid                      string    `xmlrpc:"uuid"`
	WriteDate                 time.Time `xmlrpc:"write_date"`
	WriteUid                  Many2One  `xmlrpc:"write_uid"`
}

type MailChannelNil struct {
	LastUpdate                interface{} `xmlrpc:"__last_update"`
	AliasContact              interface{} `xmlrpc:"alias_contact"`
	AliasDefaults             interface{} `xmlrpc:"alias_defaults"`
	AliasDomain               interface{} `xmlrpc:"alias_domain"`
	AliasForceThreadId        interface{} `xmlrpc:"alias_force_thread_id"`
	AliasId                   interface{} `xmlrpc:"alias_id"`
	AliasModelId              interface{} `xmlrpc:"alias_model_id"`
	AliasName                 interface{} `xmlrpc:"alias_name"`
	AliasParentModelId        interface{} `xmlrpc:"alias_parent_model_id"`
	AliasParentThreadId       interface{} `xmlrpc:"alias_parent_thread_id"`
	AliasUserId               interface{} `xmlrpc:"alias_user_id"`
	AnonymousName             interface{} `xmlrpc:"anonymous_name"`
	ChannelLastSeenPartnerIds interface{} `xmlrpc:"channel_last_seen_partner_ids"`
	ChannelMessageIds         interface{} `xmlrpc:"channel_message_ids"`
	ChannelPartnerIds         interface{} `xmlrpc:"channel_partner_ids"`
	ChannelType               interface{} `xmlrpc:"channel_type"`
	CreateDate                interface{} `xmlrpc:"create_date"`
	CreateUid                 interface{} `xmlrpc:"create_uid"`
	Description               interface{} `xmlrpc:"description"`
	DisplayName               interface{} `xmlrpc:"display_name"`
	EmailSend                 bool        `xmlrpc:"email_send"`
	GroupIds                  interface{} `xmlrpc:"group_ids"`
	GroupPublicId             interface{} `xmlrpc:"group_public_id"`
	Id                        interface{} `xmlrpc:"id"`
	Image                     interface{} `xmlrpc:"image"`
	ImageMedium               interface{} `xmlrpc:"image_medium"`
	ImageSmall                interface{} `xmlrpc:"image_small"`
	IsMember                  bool        `xmlrpc:"is_member"`
	IsSubscribed              bool        `xmlrpc:"is_subscribed"`
	LivechatChannelId         interface{} `xmlrpc:"livechat_channel_id"`
	MessageChannelIds         interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds        interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                interface{} `xmlrpc:"message_ids"`
	MessageIsFollower         bool        `xmlrpc:"message_is_follower"`
	MessageLastPost           interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction         bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter  interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds         interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread             bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter      interface{} `xmlrpc:"message_unread_counter"`
	Name                      interface{} `xmlrpc:"name"`
	Public                    interface{} `xmlrpc:"public"`
	RatingCount               interface{} `xmlrpc:"rating_count"`
	RatingIds                 interface{} `xmlrpc:"rating_ids"`
	RatingLastValue           interface{} `xmlrpc:"rating_last_value"`
	Uuid                      interface{} `xmlrpc:"uuid"`
	WriteDate                 interface{} `xmlrpc:"write_date"`
	WriteUid                  interface{} `xmlrpc:"write_uid"`
}

var MailChannelModel string = "mail.channel"

type MailChannels []MailChannel

type MailChannelsNil []MailChannelNil

func (s *MailChannel) NilableType_() interface{} {
	return &MailChannelNil{}
}

func (ns *MailChannelNil) Type_() interface{} {
	s := &MailChannel{}
	return load(ns, s)
}

func (s *MailChannels) NilableType_() interface{} {
	return &MailChannelsNil{}
}

func (ns *MailChannelsNil) Type_() interface{} {
	s := &MailChannels{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailChannel))
	}
	return s
}
