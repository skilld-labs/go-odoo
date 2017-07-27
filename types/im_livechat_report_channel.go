package types

import (
	"time"
)

type ImLivechatReportChannel struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	ChannelId         Many2One  `xmlrpc:"channel_id"`
	ChannelName       string    `xmlrpc:"channel_name"`
	DisplayName       string    `xmlrpc:"display_name"`
	Duration          float64   `xmlrpc:"duration"`
	Id                int64     `xmlrpc:"id"`
	LivechatChannelId Many2One  `xmlrpc:"livechat_channel_id"`
	NbrMessage        int64     `xmlrpc:"nbr_message"`
	NbrSpeaker        int64     `xmlrpc:"nbr_speaker"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	StartDate         time.Time `xmlrpc:"start_date"`
	StartDateHour     string    `xmlrpc:"start_date_hour"`
	TechnicalName     string    `xmlrpc:"technical_name"`
	Uuid              string    `xmlrpc:"uuid"`
}

type ImLivechatReportChannelNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	ChannelId         interface{} `xmlrpc:"channel_id"`
	ChannelName       interface{} `xmlrpc:"channel_name"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Duration          interface{} `xmlrpc:"duration"`
	Id                interface{} `xmlrpc:"id"`
	LivechatChannelId interface{} `xmlrpc:"livechat_channel_id"`
	NbrMessage        interface{} `xmlrpc:"nbr_message"`
	NbrSpeaker        interface{} `xmlrpc:"nbr_speaker"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	StartDate         interface{} `xmlrpc:"start_date"`
	StartDateHour     interface{} `xmlrpc:"start_date_hour"`
	TechnicalName     interface{} `xmlrpc:"technical_name"`
	Uuid              interface{} `xmlrpc:"uuid"`
}

var ImLivechatReportChannelModel string = "im_livechat.report.channel"

type ImLivechatReportChannels []ImLivechatReportChannel

type ImLivechatReportChannelsNil []ImLivechatReportChannelNil

func (s *ImLivechatReportChannel) NilableType_() interface{} {
	return &ImLivechatReportChannelNil{}
}

func (ns *ImLivechatReportChannelNil) Type_() interface{} {
	s := &ImLivechatReportChannel{}
	return load(ns, s)
}

func (s *ImLivechatReportChannels) NilableType_() interface{} {
	return &ImLivechatReportChannelsNil{}
}

func (ns *ImLivechatReportChannelsNil) Type_() interface{} {
	s := &ImLivechatReportChannels{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ImLivechatReportChannel))
	}
	return s
}
