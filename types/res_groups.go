package types

import (
	"time"
)

type ResGroups struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CategoryId      Many2One  `xmlrpc:"category_id"`
	Color           int64     `xmlrpc:"color"`
	Comment         string    `xmlrpc:"comment"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	FullName        string    `xmlrpc:"full_name"`
	Id              int64     `xmlrpc:"id"`
	ImpliedIds      []int64   `xmlrpc:"implied_ids"`
	IsPortal        bool      `xmlrpc:"is_portal"`
	MenuAccess      []int64   `xmlrpc:"menu_access"`
	ModelAccess     []int64   `xmlrpc:"model_access"`
	Name            string    `xmlrpc:"name"`
	RuleGroups      []int64   `xmlrpc:"rule_groups"`
	Share           bool      `xmlrpc:"share"`
	TransImpliedIds []int64   `xmlrpc:"trans_implied_ids"`
	Users           []int64   `xmlrpc:"users"`
	ViewAccess      []int64   `xmlrpc:"view_access"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type ResGroupsNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CategoryId      interface{} `xmlrpc:"category_id"`
	Color           interface{} `xmlrpc:"color"`
	Comment         interface{} `xmlrpc:"comment"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	FullName        interface{} `xmlrpc:"full_name"`
	Id              interface{} `xmlrpc:"id"`
	ImpliedIds      interface{} `xmlrpc:"implied_ids"`
	IsPortal        bool        `xmlrpc:"is_portal"`
	MenuAccess      interface{} `xmlrpc:"menu_access"`
	ModelAccess     interface{} `xmlrpc:"model_access"`
	Name            interface{} `xmlrpc:"name"`
	RuleGroups      interface{} `xmlrpc:"rule_groups"`
	Share           bool        `xmlrpc:"share"`
	TransImpliedIds interface{} `xmlrpc:"trans_implied_ids"`
	Users           interface{} `xmlrpc:"users"`
	ViewAccess      interface{} `xmlrpc:"view_access"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var ResGroupsModel string = "res.groups"

type ResGroupss []ResGroups

type ResGroupssNil []ResGroupsNil

func (s *ResGroups) NilableType_() interface{} {
	return &ResGroupsNil{}
}

func (ns *ResGroupsNil) Type_() interface{} {
	s := &ResGroups{}
	return load(ns, s)
}

func (s *ResGroupss) NilableType_() interface{} {
	return &ResGroupssNil{}
}

func (ns *ResGroupssNil) Type_() interface{} {
	s := &ResGroupss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResGroups))
	}
	return s
}
