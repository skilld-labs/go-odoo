package types

import (
	"time"
)

type CrmMergeOpportunity struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	OpportunityIds []int64   `xmlrpc:"opportunity_ids"`
	TeamId         Many2One  `xmlrpc:"team_id"`
	UserId         Many2One  `xmlrpc:"user_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type CrmMergeOpportunityNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	OpportunityIds interface{} `xmlrpc:"opportunity_ids"`
	TeamId         interface{} `xmlrpc:"team_id"`
	UserId         interface{} `xmlrpc:"user_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var CrmMergeOpportunityModel string = "crm.merge.opportunity"

type CrmMergeOpportunitys []CrmMergeOpportunity

type CrmMergeOpportunitysNil []CrmMergeOpportunityNil

func (s *CrmMergeOpportunity) NilableType_() interface{} {
	return &CrmMergeOpportunityNil{}
}

func (ns *CrmMergeOpportunityNil) Type_() interface{} {
	s := &CrmMergeOpportunity{}
	return load(ns, s)
}

func (s *CrmMergeOpportunitys) NilableType_() interface{} {
	return &CrmMergeOpportunitysNil{}
}

func (ns *CrmMergeOpportunitysNil) Type_() interface{} {
	s := &CrmMergeOpportunitys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmMergeOpportunity))
	}
	return s
}
