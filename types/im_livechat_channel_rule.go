package types

import (
	"time"
)

type ImLivechatChannelRule struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Action         string    `xmlrpc:"action"`
	AutoPopupTimer int64     `xmlrpc:"auto_popup_timer"`
	ChannelId      Many2One  `xmlrpc:"channel_id"`
	CountryIds     []int64   `xmlrpc:"country_ids"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	RegexUrl       string    `xmlrpc:"regex_url"`
	Sequence       int64     `xmlrpc:"sequence"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type ImLivechatChannelRuleNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Action         interface{} `xmlrpc:"action"`
	AutoPopupTimer interface{} `xmlrpc:"auto_popup_timer"`
	ChannelId      interface{} `xmlrpc:"channel_id"`
	CountryIds     interface{} `xmlrpc:"country_ids"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	RegexUrl       interface{} `xmlrpc:"regex_url"`
	Sequence       interface{} `xmlrpc:"sequence"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var ImLivechatChannelRuleModel string = "im_livechat.channel.rule"

type ImLivechatChannelRules []ImLivechatChannelRule

type ImLivechatChannelRulesNil []ImLivechatChannelRuleNil

func (s *ImLivechatChannelRule) NilableType_() interface{} {
	return &ImLivechatChannelRuleNil{}
}

func (ns *ImLivechatChannelRuleNil) Type_() interface{} {
	s := &ImLivechatChannelRule{}
	return load(ns, s)
}

func (s *ImLivechatChannelRules) NilableType_() interface{} {
	return &ImLivechatChannelRulesNil{}
}

func (ns *ImLivechatChannelRulesNil) Type_() interface{} {
	s := &ImLivechatChannelRules{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ImLivechatChannelRule))
	}
	return s
}
