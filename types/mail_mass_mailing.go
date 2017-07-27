package types

import (
	"time"
)

type MailMassMailing struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	Active                bool      `xmlrpc:"active"`
	AttachmentIds         []int64   `xmlrpc:"attachment_ids"`
	BodyHtml              string    `xmlrpc:"body_html"`
	Bounced               int64     `xmlrpc:"bounced"`
	BouncedRatio          int64     `xmlrpc:"bounced_ratio"`
	CampaignId            Many2One  `xmlrpc:"campaign_id"`
	ClicksRatio           int64     `xmlrpc:"clicks_ratio"`
	Color                 int64     `xmlrpc:"color"`
	ContactAbPc           int64     `xmlrpc:"contact_ab_pc"`
	ContactListIds        []int64   `xmlrpc:"contact_list_ids"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	Delivered             int64     `xmlrpc:"delivered"`
	DisplayName           string    `xmlrpc:"display_name"`
	EmailFrom             string    `xmlrpc:"email_from"`
	Failed                int64     `xmlrpc:"failed"`
	Id                    int64     `xmlrpc:"id"`
	KeepArchives          bool      `xmlrpc:"keep_archives"`
	MailingDomain         string    `xmlrpc:"mailing_domain"`
	MailingModel          string    `xmlrpc:"mailing_model"`
	MassMailingCampaignId Many2One  `xmlrpc:"mass_mailing_campaign_id"`
	MediumId              Many2One  `xmlrpc:"medium_id"`
	Name                  string    `xmlrpc:"name"`
	NextDeparture         time.Time `xmlrpc:"next_departure"`
	Opened                int64     `xmlrpc:"opened"`
	OpenedRatio           int64     `xmlrpc:"opened_ratio"`
	ReceivedRatio         int64     `xmlrpc:"received_ratio"`
	Replied               int64     `xmlrpc:"replied"`
	RepliedRatio          int64     `xmlrpc:"replied_ratio"`
	ReplyTo               string    `xmlrpc:"reply_to"`
	ReplyToMode           string    `xmlrpc:"reply_to_mode"`
	ScheduleDate          time.Time `xmlrpc:"schedule_date"`
	Scheduled             int64     `xmlrpc:"scheduled"`
	Sent                  int64     `xmlrpc:"sent"`
	SentDate              time.Time `xmlrpc:"sent_date"`
	SourceId              Many2One  `xmlrpc:"source_id"`
	State                 string    `xmlrpc:"state"`
	StatisticsIds         []int64   `xmlrpc:"statistics_ids"`
	Total                 int64     `xmlrpc:"total"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type MailMassMailingNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	Active                bool        `xmlrpc:"active"`
	AttachmentIds         interface{} `xmlrpc:"attachment_ids"`
	BodyHtml              interface{} `xmlrpc:"body_html"`
	Bounced               interface{} `xmlrpc:"bounced"`
	BouncedRatio          interface{} `xmlrpc:"bounced_ratio"`
	CampaignId            interface{} `xmlrpc:"campaign_id"`
	ClicksRatio           interface{} `xmlrpc:"clicks_ratio"`
	Color                 interface{} `xmlrpc:"color"`
	ContactAbPc           interface{} `xmlrpc:"contact_ab_pc"`
	ContactListIds        interface{} `xmlrpc:"contact_list_ids"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	Delivered             interface{} `xmlrpc:"delivered"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	EmailFrom             interface{} `xmlrpc:"email_from"`
	Failed                interface{} `xmlrpc:"failed"`
	Id                    interface{} `xmlrpc:"id"`
	KeepArchives          bool        `xmlrpc:"keep_archives"`
	MailingDomain         interface{} `xmlrpc:"mailing_domain"`
	MailingModel          interface{} `xmlrpc:"mailing_model"`
	MassMailingCampaignId interface{} `xmlrpc:"mass_mailing_campaign_id"`
	MediumId              interface{} `xmlrpc:"medium_id"`
	Name                  interface{} `xmlrpc:"name"`
	NextDeparture         interface{} `xmlrpc:"next_departure"`
	Opened                interface{} `xmlrpc:"opened"`
	OpenedRatio           interface{} `xmlrpc:"opened_ratio"`
	ReceivedRatio         interface{} `xmlrpc:"received_ratio"`
	Replied               interface{} `xmlrpc:"replied"`
	RepliedRatio          interface{} `xmlrpc:"replied_ratio"`
	ReplyTo               interface{} `xmlrpc:"reply_to"`
	ReplyToMode           interface{} `xmlrpc:"reply_to_mode"`
	ScheduleDate          interface{} `xmlrpc:"schedule_date"`
	Scheduled             interface{} `xmlrpc:"scheduled"`
	Sent                  interface{} `xmlrpc:"sent"`
	SentDate              interface{} `xmlrpc:"sent_date"`
	SourceId              interface{} `xmlrpc:"source_id"`
	State                 interface{} `xmlrpc:"state"`
	StatisticsIds         interface{} `xmlrpc:"statistics_ids"`
	Total                 interface{} `xmlrpc:"total"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var MailMassMailingModel string = "mail.mass_mailing"

type MailMassMailings []MailMassMailing

type MailMassMailingsNil []MailMassMailingNil

func (s *MailMassMailing) NilableType_() interface{} {
	return &MailMassMailingNil{}
}

func (ns *MailMassMailingNil) Type_() interface{} {
	s := &MailMassMailing{}
	return load(ns, s)
}

func (s *MailMassMailings) NilableType_() interface{} {
	return &MailMassMailingsNil{}
}

func (ns *MailMassMailingsNil) Type_() interface{} {
	s := &MailMassMailings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MailMassMailing))
	}
	return s
}
