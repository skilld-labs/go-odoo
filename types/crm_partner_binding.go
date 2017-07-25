package types

import (
	"time"
)

type CrmPartnerBinding struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Action      string    `xmlrpc:"action"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PartnerId   Many2One  `xmlrpc:"partner_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CrmPartnerBindingNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Action      interface{} `xmlrpc:"action"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PartnerId   interface{} `xmlrpc:"partner_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CrmPartnerBindingModel string = "crm.partner.binding"

type CrmPartnerBindings []CrmPartnerBinding

type CrmPartnerBindingsNil []CrmPartnerBindingNil

func (s *CrmPartnerBinding) NilableType_() interface{} {
	return &CrmPartnerBindingNil{}
}

func (ns *CrmPartnerBindingNil) Type_() interface{} {
	s := &CrmPartnerBinding{}
	return load(ns, s)
}

func (s *CrmPartnerBindings) NilableType_() interface{} {
	return &CrmPartnerBindingsNil{}
}

func (ns *CrmPartnerBindingsNil) Type_() interface{} {
	s := &CrmPartnerBindings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmPartnerBinding))
	}
	return s
}
