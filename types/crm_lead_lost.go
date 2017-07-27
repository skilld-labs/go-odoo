package types

import (
	"time"
)

type CrmLeadLost struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	Id           int64     `xmlrpc:"id"`
	LostReasonId Many2One  `xmlrpc:"lost_reason_id"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type CrmLeadLostNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Id           interface{} `xmlrpc:"id"`
	LostReasonId interface{} `xmlrpc:"lost_reason_id"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var CrmLeadLostModel string = "crm.lead.lost"

type CrmLeadLosts []CrmLeadLost

type CrmLeadLostsNil []CrmLeadLostNil

func (s *CrmLeadLost) NilableType_() interface{} {
	return &CrmLeadLostNil{}
}

func (ns *CrmLeadLostNil) Type_() interface{} {
	s := &CrmLeadLost{}
	return load(ns, s)
}

func (s *CrmLeadLosts) NilableType_() interface{} {
	return &CrmLeadLostsNil{}
}

func (ns *CrmLeadLostsNil) Type_() interface{} {
	s := &CrmLeadLosts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmLeadLost))
	}
	return s
}
