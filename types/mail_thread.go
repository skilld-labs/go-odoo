package types

import (
	"time"
)

type MailThread struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	DisplayName              string    `xmlrpc:"display_name"`
	Id                       int64     `xmlrpc:"id"`
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
}

type MailThreadNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Id                       interface{} `xmlrpc:"id"`
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
}

var MailThreadModel string = "mail.thread"

type MailThreads []MailThread

type MailThreadsNil []MailThreadNil

func (s *MailThread) NilableType_() interface{} {
	return &MailThreadNil{}
}

func (ns *MailThreadNil) Type_() interface{} {
	s := &MailThread{}
	return load(ns, s)
}

func (s *MailThreads) NilableType_() interface{} {
	return &MailThreadsNil{}
}

func (ns *MailThreadsNil) Type_() interface{} {
	s := &MailThreads{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailThread))
	}
	return s
}
