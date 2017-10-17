package types

import (
	"time"
)

type AccountGroup struct {
	CodePrefix  string    `xmlrpc:"code_prefix"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	ParentId    Many2One  `xmlrpc:"parent_id"`
	ParentLeft  int64     `xmlrpc:"parent_left"`
	ParentRight int64     `xmlrpc:"parent_right"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountGroupNil struct {
	CodePrefix  interface{} `xmlrpc:"code_prefix"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	ParentId    interface{} `xmlrpc:"parent_id"`
	ParentLeft  interface{} `xmlrpc:"parent_left"`
	ParentRight interface{} `xmlrpc:"parent_right"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountGroupModel string = "account.group"

type AccountGroups []AccountGroup

type AccountGroupsNil []AccountGroupNil

func (s *AccountGroup) NilableType_() interface{} {
	return &AccountGroupNil{}
}

func (ns *AccountGroupNil) Type_() interface{} {
	s := &AccountGroup{}
	return load(ns, s)
}

func (s *AccountGroups) NilableType_() interface{} {
	return &AccountGroupsNil{}
}

func (ns *AccountGroupsNil) Type_() interface{} {
	s := &AccountGroups{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountGroup))
	}
	return s
}
