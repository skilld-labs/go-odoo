package types

import (
	"time"
)

type ImLivechatChannel struct {
	LastUpdate                   time.Time `xmlrpc:"__last_update"`
	AreYouInside                 bool      `xmlrpc:"are_you_inside"`
	ButtonText                   string    `xmlrpc:"button_text"`
	ChannelIds                   []int64   `xmlrpc:"channel_ids"`
	CreateDate                   time.Time `xmlrpc:"create_date"`
	CreateUid                    Many2One  `xmlrpc:"create_uid"`
	DefaultMessage               string    `xmlrpc:"default_message"`
	DisplayName                  string    `xmlrpc:"display_name"`
	Id                           int64     `xmlrpc:"id"`
	Image                        string    `xmlrpc:"image"`
	ImageMedium                  string    `xmlrpc:"image_medium"`
	ImageSmall                   string    `xmlrpc:"image_small"`
	InputPlaceholder             string    `xmlrpc:"input_placeholder"`
	Name                         string    `xmlrpc:"name"`
	NbrChannel                   int64     `xmlrpc:"nbr_channel"`
	RatingPercentageSatisfaction int64     `xmlrpc:"rating_percentage_satisfaction"`
	RuleIds                      []int64   `xmlrpc:"rule_ids"`
	ScriptExternal               string    `xmlrpc:"script_external"`
	UserIds                      []int64   `xmlrpc:"user_ids"`
	WebPage                      string    `xmlrpc:"web_page"`
	WriteDate                    time.Time `xmlrpc:"write_date"`
	WriteUid                     Many2One  `xmlrpc:"write_uid"`
}

type ImLivechatChannelNil struct {
	LastUpdate                   interface{} `xmlrpc:"__last_update"`
	AreYouInside                 bool        `xmlrpc:"are_you_inside"`
	ButtonText                   interface{} `xmlrpc:"button_text"`
	ChannelIds                   interface{} `xmlrpc:"channel_ids"`
	CreateDate                   interface{} `xmlrpc:"create_date"`
	CreateUid                    interface{} `xmlrpc:"create_uid"`
	DefaultMessage               interface{} `xmlrpc:"default_message"`
	DisplayName                  interface{} `xmlrpc:"display_name"`
	Id                           interface{} `xmlrpc:"id"`
	Image                        interface{} `xmlrpc:"image"`
	ImageMedium                  interface{} `xmlrpc:"image_medium"`
	ImageSmall                   interface{} `xmlrpc:"image_small"`
	InputPlaceholder             interface{} `xmlrpc:"input_placeholder"`
	Name                         interface{} `xmlrpc:"name"`
	NbrChannel                   interface{} `xmlrpc:"nbr_channel"`
	RatingPercentageSatisfaction interface{} `xmlrpc:"rating_percentage_satisfaction"`
	RuleIds                      interface{} `xmlrpc:"rule_ids"`
	ScriptExternal               interface{} `xmlrpc:"script_external"`
	UserIds                      interface{} `xmlrpc:"user_ids"`
	WebPage                      interface{} `xmlrpc:"web_page"`
	WriteDate                    interface{} `xmlrpc:"write_date"`
	WriteUid                     interface{} `xmlrpc:"write_uid"`
}

var ImLivechatChannelModel string = "im_livechat.channel"

type ImLivechatChannels []ImLivechatChannel

type ImLivechatChannelsNil []ImLivechatChannelNil

func (s *ImLivechatChannel) NilableType_() interface{} {
	return &ImLivechatChannelNil{}
}

func (ns *ImLivechatChannelNil) Type_() interface{} {
	s := &ImLivechatChannel{}
	return load(ns, s)
}

func (s *ImLivechatChannels) NilableType_() interface{} {
	return &ImLivechatChannelsNil{}
}

func (ns *ImLivechatChannelsNil) Type_() interface{} {
	s := &ImLivechatChannels{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ImLivechatChannel))
	}
	return s
}
