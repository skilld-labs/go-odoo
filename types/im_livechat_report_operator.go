package types

import (
	"time"
)

type ImLivechatReportOperator struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	ChannelId         Many2One  `xmlrpc:"channel_id"`
	DisplayName       string    `xmlrpc:"display_name"`
	Duration          float64   `xmlrpc:"duration"`
	Id                int64     `xmlrpc:"id"`
	LivechatChannelId Many2One  `xmlrpc:"livechat_channel_id"`
	NbrChannel        int64     `xmlrpc:"nbr_channel"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	StartDate         time.Time `xmlrpc:"start_date"`
	TimeToAnswer      float64   `xmlrpc:"time_to_answer"`
}

type ImLivechatReportOperatorNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	ChannelId         interface{} `xmlrpc:"channel_id"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Duration          interface{} `xmlrpc:"duration"`
	Id                interface{} `xmlrpc:"id"`
	LivechatChannelId interface{} `xmlrpc:"livechat_channel_id"`
	NbrChannel        interface{} `xmlrpc:"nbr_channel"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	StartDate         interface{} `xmlrpc:"start_date"`
	TimeToAnswer      interface{} `xmlrpc:"time_to_answer"`
}

var ImLivechatReportOperatorModel string = "im_livechat.report.operator"

type ImLivechatReportOperators []ImLivechatReportOperator

type ImLivechatReportOperatorsNil []ImLivechatReportOperatorNil

func (s *ImLivechatReportOperator) NilableType_() interface{} {
	return &ImLivechatReportOperatorNil{}
}

func (ns *ImLivechatReportOperatorNil) Type_() interface{} {
	s := &ImLivechatReportOperator{}
	return load(ns, s)
}

func (s *ImLivechatReportOperators) NilableType_() interface{} {
	return &ImLivechatReportOperatorsNil{}
}

func (ns *ImLivechatReportOperatorsNil) Type_() interface{} {
	s := &ImLivechatReportOperators{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ImLivechatReportOperator))
	}
	return s
}
