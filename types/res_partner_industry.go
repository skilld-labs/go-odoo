package types

import (
	"time"
)

type ResPartnerIndustry struct {
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	FullName    string    `xmlrpc:"full_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResPartnerIndustryNil struct {
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	FullName    interface{} `xmlrpc:"full_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResPartnerIndustryModel string = "res.partner.industry"

type ResPartnerIndustrys []ResPartnerIndustry

type ResPartnerIndustrysNil []ResPartnerIndustryNil

func (s *ResPartnerIndustry) NilableType_() interface{} {
	return &ResPartnerIndustryNil{}
}

func (ns *ResPartnerIndustryNil) Type_() interface{} {
	s := &ResPartnerIndustry{}
	return load(ns, s)
}

func (s *ResPartnerIndustrys) NilableType_() interface{} {
	return &ResPartnerIndustrysNil{}
}

func (ns *ResPartnerIndustrysNil) Type_() interface{} {
	s := &ResPartnerIndustrys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResPartnerIndustry))
	}
	return s
}
