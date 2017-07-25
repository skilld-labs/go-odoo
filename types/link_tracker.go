package types

import (
	"time"
)

type LinkTracker struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	CampaignId            Many2One  `xmlrpc:"campaign_id"`
	Code                  string    `xmlrpc:"code"`
	Count                 int64     `xmlrpc:"count"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	Favicon               string    `xmlrpc:"favicon"`
	IconSrc               string    `xmlrpc:"icon_src"`
	Id                    int64     `xmlrpc:"id"`
	LinkClickIds          []int64   `xmlrpc:"link_click_ids"`
	LinkCodeIds           []int64   `xmlrpc:"link_code_ids"`
	MassMailingCampaignId Many2One  `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         Many2One  `xmlrpc:"mass_mailing_id"`
	MediumId              Many2One  `xmlrpc:"medium_id"`
	RedirectedUrl         string    `xmlrpc:"redirected_url"`
	ShortUrl              string    `xmlrpc:"short_url"`
	ShortUrlHost          string    `xmlrpc:"short_url_host"`
	SourceId              Many2One  `xmlrpc:"source_id"`
	Title                 string    `xmlrpc:"title"`
	Url                   string    `xmlrpc:"url"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type LinkTrackerNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	CampaignId            interface{} `xmlrpc:"campaign_id"`
	Code                  interface{} `xmlrpc:"code"`
	Count                 interface{} `xmlrpc:"count"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Favicon               interface{} `xmlrpc:"favicon"`
	IconSrc               interface{} `xmlrpc:"icon_src"`
	Id                    interface{} `xmlrpc:"id"`
	LinkClickIds          interface{} `xmlrpc:"link_click_ids"`
	LinkCodeIds           interface{} `xmlrpc:"link_code_ids"`
	MassMailingCampaignId interface{} `xmlrpc:"mass_mailing_campaign_id"`
	MassMailingId         interface{} `xmlrpc:"mass_mailing_id"`
	MediumId              interface{} `xmlrpc:"medium_id"`
	RedirectedUrl         interface{} `xmlrpc:"redirected_url"`
	ShortUrl              interface{} `xmlrpc:"short_url"`
	ShortUrlHost          interface{} `xmlrpc:"short_url_host"`
	SourceId              interface{} `xmlrpc:"source_id"`
	Title                 interface{} `xmlrpc:"title"`
	Url                   interface{} `xmlrpc:"url"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var LinkTrackerModel string = "link.tracker"

type LinkTrackers []LinkTracker

type LinkTrackersNil []LinkTrackerNil

func (s *LinkTracker) NilableType_() interface{} {
	return &LinkTrackerNil{}
}

func (ns *LinkTrackerNil) Type_() interface{} {
	s := &LinkTracker{}
	return load(ns, s)
}

func (s *LinkTrackers) NilableType_() interface{} {
	return &LinkTrackersNil{}
}

func (ns *LinkTrackersNil) Type_() interface{} {
	s := &LinkTrackers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*LinkTracker))
	}
	return s
}
