package types

import (
	"time"
)

type MailMassMailingCampaign struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	Bounced         int64     `xmlrpc:"bounced"`
	BouncedRatio    int64     `xmlrpc:"bounced_ratio"`
	CampaignId      Many2One  `xmlrpc:"campaign_id"`
	ClicksRatio     int64     `xmlrpc:"clicks_ratio"`
	Color           int64     `xmlrpc:"color"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	Delivered       int64     `xmlrpc:"delivered"`
	DisplayName     string    `xmlrpc:"display_name"`
	Failed          int64     `xmlrpc:"failed"`
	Id              int64     `xmlrpc:"id"`
	MassMailingIds  []int64   `xmlrpc:"mass_mailing_ids"`
	MediumId        Many2One  `xmlrpc:"medium_id"`
	Name            string    `xmlrpc:"name"`
	Opened          int64     `xmlrpc:"opened"`
	OpenedRatio     int64     `xmlrpc:"opened_ratio"`
	ReceivedRatio   int64     `xmlrpc:"received_ratio"`
	Replied         int64     `xmlrpc:"replied"`
	RepliedRatio    int64     `xmlrpc:"replied_ratio"`
	Scheduled       int64     `xmlrpc:"scheduled"`
	Sent            int64     `xmlrpc:"sent"`
	SourceId        Many2One  `xmlrpc:"source_id"`
	StageId         Many2One  `xmlrpc:"stage_id"`
	TagIds          []int64   `xmlrpc:"tag_ids"`
	Total           int64     `xmlrpc:"total"`
	TotalMailings   int64     `xmlrpc:"total_mailings"`
	UniqueAbTesting bool      `xmlrpc:"unique_ab_testing"`
	UserId          Many2One  `xmlrpc:"user_id"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingCampaignNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Bounced         interface{} `xmlrpc:"bounced"`
	BouncedRatio    interface{} `xmlrpc:"bounced_ratio"`
	CampaignId      interface{} `xmlrpc:"campaign_id"`
	ClicksRatio     interface{} `xmlrpc:"clicks_ratio"`
	Color           interface{} `xmlrpc:"color"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	Delivered       interface{} `xmlrpc:"delivered"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Failed          interface{} `xmlrpc:"failed"`
	Id              interface{} `xmlrpc:"id"`
	MassMailingIds  interface{} `xmlrpc:"mass_mailing_ids"`
	MediumId        interface{} `xmlrpc:"medium_id"`
	Name            interface{} `xmlrpc:"name"`
	Opened          interface{} `xmlrpc:"opened"`
	OpenedRatio     interface{} `xmlrpc:"opened_ratio"`
	ReceivedRatio   interface{} `xmlrpc:"received_ratio"`
	Replied         interface{} `xmlrpc:"replied"`
	RepliedRatio    interface{} `xmlrpc:"replied_ratio"`
	Scheduled       interface{} `xmlrpc:"scheduled"`
	Sent            interface{} `xmlrpc:"sent"`
	SourceId        interface{} `xmlrpc:"source_id"`
	StageId         interface{} `xmlrpc:"stage_id"`
	TagIds          interface{} `xmlrpc:"tag_ids"`
	Total           interface{} `xmlrpc:"total"`
	TotalMailings   interface{} `xmlrpc:"total_mailings"`
	UniqueAbTesting bool        `xmlrpc:"unique_ab_testing"`
	UserId          interface{} `xmlrpc:"user_id"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingCampaignModel string = "mail.mass_mailing.campaign"

type MailMassMailingCampaigns []MailMassMailingCampaign

type MailMassMailingCampaignsNil []MailMassMailingCampaignNil

func (s *MailMassMailingCampaign) NilableType_() interface{} {
	return &MailMassMailingCampaignNil{}
}

func (ns *MailMassMailingCampaignNil) Type_() interface{} {
	s := &MailMassMailingCampaign{}
	return load(ns, s)
}

func (s *MailMassMailingCampaigns) NilableType_() interface{} {
	return &MailMassMailingCampaignsNil{}
}

func (ns *MailMassMailingCampaignsNil) Type_() interface{} {
	s := &MailMassMailingCampaigns{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailingCampaign))
	}
	return s
}
