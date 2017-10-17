package types

import (
	"time"
)

type IapAccount struct {
	AccountToken string    `xmlrpc:"account_token"`
	CompanyId    Many2One  `xmlrpc:"company_id"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	Id           int64     `xmlrpc:"id"`
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	ServiceName  string    `xmlrpc:"service_name"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type IapAccountNil struct {
	AccountToken interface{} `xmlrpc:"account_token"`
	CompanyId    interface{} `xmlrpc:"company_id"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Id           interface{} `xmlrpc:"id"`
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	ServiceName  interface{} `xmlrpc:"service_name"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var IapAccountModel string = "iap.account"

type IapAccounts []IapAccount

type IapAccountsNil []IapAccountNil

func (s *IapAccount) NilableType_() interface{} {
	return &IapAccountNil{}
}

func (ns *IapAccountNil) Type_() interface{} {
	s := &IapAccount{}
	return load(ns, s)
}

func (s *IapAccounts) NilableType_() interface{} {
	return &IapAccountsNil{}
}

func (ns *IapAccountsNil) Type_() interface{} {
	s := &IapAccounts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IapAccount))
	}
	return s
}
