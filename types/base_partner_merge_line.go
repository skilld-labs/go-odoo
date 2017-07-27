package types

import (
	"time"
)

type BasePartnerMergeLine struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	AggrIds     string    `xmlrpc:"aggr_ids"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	MinId       int64     `xmlrpc:"min_id"`
	WizardId    Many2One  `xmlrpc:"wizard_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BasePartnerMergeLineNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	AggrIds     interface{} `xmlrpc:"aggr_ids"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	MinId       interface{} `xmlrpc:"min_id"`
	WizardId    interface{} `xmlrpc:"wizard_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BasePartnerMergeLineModel string = "base.partner.merge.line"

type BasePartnerMergeLines []BasePartnerMergeLine

type BasePartnerMergeLinesNil []BasePartnerMergeLineNil

func (s *BasePartnerMergeLine) NilableType_() interface{} {
	return &BasePartnerMergeLineNil{}
}

func (ns *BasePartnerMergeLineNil) Type_() interface{} {
	s := &BasePartnerMergeLine{}
	return load(ns, s)
}

func (s *BasePartnerMergeLines) NilableType_() interface{} {
	return &BasePartnerMergeLinesNil{}
}

func (ns *BasePartnerMergeLinesNil) Type_() interface{} {
	s := &BasePartnerMergeLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BasePartnerMergeLine))
	}
	return s
}
