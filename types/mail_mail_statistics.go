package types

import (
	"time"
)

type MailMailStatistics struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	Bounced               time.Time `xmlrpc:"bounced"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	Exception             time.Time `xmlrpc:"exception"`
	Id                    int64     `xmlrpc:"id"`
	LinksClickIds         []int64   `xmlrpc:"links_click_ids"`
	MailMailId            Many2One  `xmlrpc:"mail_mail_id"`
	MailMailIdInt         int64     `xmlrpc:"mail_mail_id_int"`
	MassMailingCampaignId Many2One  `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         Many2One  `xmlrpc:"mass_mailing_id"`
	MessageId             string    `xmlrpc:"message_id"`
	Model                 string    `xmlrpc:"model"`
	Opened                time.Time `xmlrpc:"opened"`
	Recipient             string    `xmlrpc:"recipient"`
	Replied               time.Time `xmlrpc:"replied"`
	ResId                 int64     `xmlrpc:"res_id"`
	Scheduled             time.Time `xmlrpc:"scheduled"`
	Sent                  time.Time `xmlrpc:"sent"`
	State                 string    `xmlrpc:"state"`
	StateUpdate           time.Time `xmlrpc:"state_update"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type MailMailStatisticsNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	Bounced               interface{} `xmlrpc:"bounced"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Exception             interface{} `xmlrpc:"exception"`
	Id                    interface{} `xmlrpc:"id"`
	LinksClickIds         interface{} `xmlrpc:"links_click_ids"`
	MailMailId            interface{} `xmlrpc:"mail_mail_id"`
	MailMailIdInt         interface{} `xmlrpc:"mail_mail_id_int"`
	MassMailingCampaignId interface{} `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         interface{} `xmlrpc:"mass_mailing_id"`
	MessageId             interface{} `xmlrpc:"message_id"`
	Model                 interface{} `xmlrpc:"model"`
	Opened                interface{} `xmlrpc:"opened"`
	Recipient             interface{} `xmlrpc:"recipient"`
	Replied               interface{} `xmlrpc:"replied"`
	ResId                 interface{} `xmlrpc:"res_id"`
	Scheduled             interface{} `xmlrpc:"scheduled"`
	Sent                  interface{} `xmlrpc:"sent"`
	State                 interface{} `xmlrpc:"state"`
	StateUpdate           interface{} `xmlrpc:"state_update"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var MailMailStatisticsModel string = "mail.mail.statistics"

type MailMailStatisticss []MailMailStatistics

type MailMailStatisticssNil []MailMailStatisticsNil

func (s *MailMailStatistics) NilableType_() interface{} {
	return &MailMailStatisticsNil{}
}

func (ns *MailMailStatisticsNil) Type_() interface{} {
	s := &MailMailStatistics{}
	return load(ns, s)
}

func (s *MailMailStatisticss) NilableType_() interface{} {
	return &MailMailStatisticssNil{}
}

func (ns *MailMailStatisticssNil) Type_() interface{} {
	s := &MailMailStatisticss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMailStatistics))
	}
	return s
}
