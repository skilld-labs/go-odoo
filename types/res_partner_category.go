package types

import (
	"time"
)

type ResPartnerCategory struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	ChildIds    []int64   `xmlrpc:"child_ids"`
	Color       int64     `xmlrpc:"color"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	ParentId    Many2One  `xmlrpc:"parent_id"`
	ParentLeft  int64     `xmlrpc:"parent_left"`
	ParentRight int64     `xmlrpc:"parent_right"`
	PartnerIds  []int64   `xmlrpc:"partner_ids"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResPartnerCategoryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	ChildIds    interface{} `xmlrpc:"child_ids"`
	Color       interface{} `xmlrpc:"color"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	ParentId    interface{} `xmlrpc:"parent_id"`
	ParentLeft  interface{} `xmlrpc:"parent_left"`
	ParentRight interface{} `xmlrpc:"parent_right"`
	PartnerIds  interface{} `xmlrpc:"partner_ids"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResPartnerCategoryModel string = "res.partner.category"

type ResPartnerCategorys []ResPartnerCategory

type ResPartnerCategorysNil []ResPartnerCategoryNil

func (s *ResPartnerCategory) NilableType_() interface{} {
	return &ResPartnerCategoryNil{}
}

func (ns *ResPartnerCategoryNil) Type_() interface{} {
	s := &ResPartnerCategory{}
	return load(ns, s)
}

func (s *ResPartnerCategorys) NilableType_() interface{} {
	return &ResPartnerCategorysNil{}
}

func (ns *ResPartnerCategorysNil) Type_() interface{} {
	s := &ResPartnerCategorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResPartnerCategory))
	}
	return s
}
