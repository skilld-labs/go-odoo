package types

import (
	"time"
)

type AccountAnalyticTag struct {
	Active      bool      `xmlrpc:"active"`
	Color       int64     `xmlrpc:"color"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountAnalyticTagNil struct {
	Active      bool        `xmlrpc:"active"`
	Color       interface{} `xmlrpc:"color"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountAnalyticTagModel string = "account.analytic.tag"

type AccountAnalyticTags []AccountAnalyticTag

type AccountAnalyticTagsNil []AccountAnalyticTagNil

func (s *AccountAnalyticTag) NilableType_() interface{} {
	return &AccountAnalyticTagNil{}
}

func (ns *AccountAnalyticTagNil) Type_() interface{} {
	s := &AccountAnalyticTag{}
	return load(ns, s)
}

func (s *AccountAnalyticTags) NilableType_() interface{} {
	return &AccountAnalyticTagsNil{}
}

func (ns *AccountAnalyticTagsNil) Type_() interface{} {
	s := &AccountAnalyticTags{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAnalyticTag))
	}
	return s
}
