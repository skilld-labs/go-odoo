package types

import (
	"time"
)

type CrmLead2opportunityPartnerMass struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	Action           string    `xmlrpc:"action"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	Deduplicate      bool      `xmlrpc:"deduplicate"`
	DisplayName      string    `xmlrpc:"display_name"`
	ForceAssignation bool      `xmlrpc:"force_assignation"`
	Id               int64     `xmlrpc:"id"`
	Name             string    `xmlrpc:"name"`
	OpportunityIds   []int64   `xmlrpc:"opportunity_ids"`
	PartnerId        Many2One  `xmlrpc:"partner_id"`
	TeamId           Many2One  `xmlrpc:"team_id"`
	UserId           Many2One  `xmlrpc:"user_id"`
	UserIds          []int64   `xmlrpc:"user_ids"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type CrmLead2opportunityPartnerMassNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	Action           interface{} `xmlrpc:"action"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	Deduplicate      bool        `xmlrpc:"deduplicate"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	ForceAssignation bool        `xmlrpc:"force_assignation"`
	Id               interface{} `xmlrpc:"id"`
	Name             interface{} `xmlrpc:"name"`
	OpportunityIds   interface{} `xmlrpc:"opportunity_ids"`
	PartnerId        interface{} `xmlrpc:"partner_id"`
	TeamId           interface{} `xmlrpc:"team_id"`
	UserId           interface{} `xmlrpc:"user_id"`
	UserIds          interface{} `xmlrpc:"user_ids"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var CrmLead2opportunityPartnerMassModel string = "crm.lead2opportunity.partner.mass"

type CrmLead2opportunityPartnerMasss []CrmLead2opportunityPartnerMass

type CrmLead2opportunityPartnerMasssNil []CrmLead2opportunityPartnerMassNil

func (s *CrmLead2opportunityPartnerMass) NilableType_() interface{} {
	return &CrmLead2opportunityPartnerMassNil{}
}

func (ns *CrmLead2opportunityPartnerMassNil) Type_() interface{} {
	s := &CrmLead2opportunityPartnerMass{}
	return load(ns, s)
}

func (s *CrmLead2opportunityPartnerMasss) NilableType_() interface{} {
	return &CrmLead2opportunityPartnerMasssNil{}
}

func (ns *CrmLead2opportunityPartnerMasssNil) Type_() interface{} {
	s := &CrmLead2opportunityPartnerMasss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmLead2opportunityPartnerMass))
	}
	return s
}
