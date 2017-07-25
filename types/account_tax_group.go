package types

import (
	"time"
)

type AccountTaxGroup struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Sequence    int64     `xmlrpc:"sequence"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountTaxGroupNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Sequence    interface{} `xmlrpc:"sequence"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountTaxGroupModel string = "account.tax.group"

type AccountTaxGroups []AccountTaxGroup

type AccountTaxGroupsNil []AccountTaxGroupNil

func (s *AccountTaxGroup) NilableType_() interface{} {
	return &AccountTaxGroupNil{}
}

func (ns *AccountTaxGroupNil) Type_() interface{} {
	s := &AccountTaxGroup{}
	return load(ns, s)
}

func (s *AccountTaxGroups) NilableType_() interface{} {
	return &AccountTaxGroupsNil{}
}

func (ns *AccountTaxGroupsNil) Type_() interface{} {
	s := &AccountTaxGroups{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountTaxGroup))
	}
	return s
}
