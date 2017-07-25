package types

import (
	"time"
)

type CrmLead2opportunityPartner struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Action         string    `xmlrpc:"action"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	Name           string    `xmlrpc:"name"`
	OpportunityIds []int64   `xmlrpc:"opportunity_ids"`
	PartnerId      Many2One  `xmlrpc:"partner_id"`
	TeamId         Many2One  `xmlrpc:"team_id"`
	UserId         Many2One  `xmlrpc:"user_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type CrmLead2opportunityPartnerNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Action         interface{} `xmlrpc:"action"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	Name           interface{} `xmlrpc:"name"`
	OpportunityIds interface{} `xmlrpc:"opportunity_ids"`
	PartnerId      interface{} `xmlrpc:"partner_id"`
	TeamId         interface{} `xmlrpc:"team_id"`
	UserId         interface{} `xmlrpc:"user_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var CrmLead2opportunityPartnerModel string = "crm.lead2opportunity.partner"

type CrmLead2opportunityPartners []CrmLead2opportunityPartner

type CrmLead2opportunityPartnersNil []CrmLead2opportunityPartnerNil

func (s *CrmLead2opportunityPartner) NilableType_() interface{} {
	return &CrmLead2opportunityPartnerNil{}
}

func (ns *CrmLead2opportunityPartnerNil) Type_() interface{} {
	s := &CrmLead2opportunityPartner{}
	return load(ns, s)
}

func (s *CrmLead2opportunityPartners) NilableType_() interface{} {
	return &CrmLead2opportunityPartnersNil{}
}

func (ns *CrmLead2opportunityPartnersNil) Type_() interface{} {
	s := &CrmLead2opportunityPartners{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmLead2opportunityPartner))
	}
	return s
}
