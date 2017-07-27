package types

import (
	"time"
)

type CrmLeadTag struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Color       int64     `xmlrpc:"color"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CrmLeadTagNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Color       interface{} `xmlrpc:"color"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CrmLeadTagModel string = "crm.lead.tag"

type CrmLeadTags []CrmLeadTag

type CrmLeadTagsNil []CrmLeadTagNil

func (s *CrmLeadTag) NilableType_() interface{} {
	return &CrmLeadTagNil{}
}

func (ns *CrmLeadTagNil) Type_() interface{} {
	s := &CrmLeadTag{}
	return load(ns, s)
}

func (s *CrmLeadTags) NilableType_() interface{} {
	return &CrmLeadTagsNil{}
}

func (ns *CrmLeadTagsNil) Type_() interface{} {
	s := &CrmLeadTags{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CrmLeadTag))
	}
	return s
}
