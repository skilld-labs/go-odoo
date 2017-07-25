package types

import (
	"time"
)

type UtmCampaign struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type UtmCampaignNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var UtmCampaignModel string = "utm.campaign"

type UtmCampaigns []UtmCampaign

type UtmCampaignsNil []UtmCampaignNil

func (s *UtmCampaign) NilableType_() interface{} {
	return &UtmCampaignNil{}
}

func (ns *UtmCampaignNil) Type_() interface{} {
	s := &UtmCampaign{}
	return load(ns, s)
}

func (s *UtmCampaigns) NilableType_() interface{} {
	return &UtmCampaignsNil{}
}

func (ns *UtmCampaignsNil) Type_() interface{} {
	s := &UtmCampaigns{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*UtmCampaign))
	}
	return s
}
