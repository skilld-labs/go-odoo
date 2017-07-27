package types

import (
	"time"
)

type AccountAccountTag struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	Applicability string    `xmlrpc:"applicability"`
	Color         int64     `xmlrpc:"color"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	Name          string    `xmlrpc:"name"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type AccountAccountTagNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	Applicability interface{} `xmlrpc:"applicability"`
	Color         interface{} `xmlrpc:"color"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	Name          interface{} `xmlrpc:"name"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var AccountAccountTagModel string = "account.account.tag"

type AccountAccountTags []AccountAccountTag

type AccountAccountTagsNil []AccountAccountTagNil

func (s *AccountAccountTag) NilableType_() interface{} {
	return &AccountAccountTagNil{}
}

func (ns *AccountAccountTagNil) Type_() interface{} {
	s := &AccountAccountTag{}
	return load(ns, s)
}

func (s *AccountAccountTags) NilableType_() interface{} {
	return &AccountAccountTagsNil{}
}

func (ns *AccountAccountTagsNil) Type_() interface{} {
	s := &AccountAccountTags{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAccountTag))
	}
	return s
}
