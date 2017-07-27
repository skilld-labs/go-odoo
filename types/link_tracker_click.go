package types

import (
	"time"
)

type LinkTrackerClick struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	ClickDate             time.Time `xmlrpc:"click_date"`
	CountryId             Many2One  `xmlrpc:"country_id"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	Id                    int64     `xmlrpc:"id"`
	Ip                    string    `xmlrpc:"ip"`
	LinkId                Many2One  `xmlrpc:"link_id"`
	MailStatId            Many2One  `xmlrpc:"mail_stat_id"`
	MassMailingCampaignId Many2One  `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         Many2One  `xmlrpc:"mass_mailing_id"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type LinkTrackerClickNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	ClickDate             interface{} `xmlrpc:"click_date"`
	CountryId             interface{} `xmlrpc:"country_id"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Id                    interface{} `xmlrpc:"id"`
	Ip                    interface{} `xmlrpc:"ip"`
	LinkId                interface{} `xmlrpc:"link_id"`
	MailStatId            interface{} `xmlrpc:"mail_stat_id"`
	MassMailingCampaignId interface{} `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         interface{} `xmlrpc:"mass_mailing_id"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var LinkTrackerClickModel string = "link.tracker.click"

type LinkTrackerClicks []LinkTrackerClick

type LinkTrackerClicksNil []LinkTrackerClickNil

func (s *LinkTrackerClick) NilableType_() interface{} {
	return &LinkTrackerClickNil{}
}

func (ns *LinkTrackerClickNil) Type_() interface{} {
	s := &LinkTrackerClick{}
	return load(ns, s)
}

func (s *LinkTrackerClicks) NilableType_() interface{} {
	return &LinkTrackerClicksNil{}
}

func (ns *LinkTrackerClicksNil) Type_() interface{} {
	s := &LinkTrackerClicks{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*LinkTrackerClick))
	}
	return s
}
