package types

import (
	"time"
)

type UtmMixin struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CampaignId  Many2One  `xmlrpc:"campaign_id"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	MediumId    Many2One  `xmlrpc:"medium_id"`
	SourceId    Many2One  `xmlrpc:"source_id"`
}

type UtmMixinNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CampaignId  interface{} `xmlrpc:"campaign_id"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	MediumId    interface{} `xmlrpc:"medium_id"`
	SourceId    interface{} `xmlrpc:"source_id"`
}

var UtmMixinModel string = "utm.mixin"

type UtmMixins []UtmMixin

type UtmMixinsNil []UtmMixinNil

func (s *UtmMixin) NilableType_() interface{} {
	return &UtmMixinNil{}
}

func (ns *UtmMixinNil) Type_() interface{} {
	s := &UtmMixin{}
	return load(ns, s)
}

func (s *UtmMixins) NilableType_() interface{} {
	return &UtmMixinsNil{}
}

func (ns *UtmMixinsNil) Type_() interface{} {
	s := &UtmMixins{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*UtmMixin))
	}
	return s
}
