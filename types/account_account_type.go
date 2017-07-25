package types

import (
	"time"
)

type AccountAccountType struct {
	LastUpdate            time.Time `xmlrpc:"__last_update"`
	CreateDate            time.Time `xmlrpc:"create_date"`
	CreateUid             Many2One  `xmlrpc:"create_uid"`
	DisplayName           string    `xmlrpc:"display_name"`
	Id                    int64     `xmlrpc:"id"`
	IncludeInitialBalance bool      `xmlrpc:"include_initial_balance"`
	Name                  string    `xmlrpc:"name"`
	Note                  string    `xmlrpc:"note"`
	Type                  string    `xmlrpc:"type"`
	WriteDate             time.Time `xmlrpc:"write_date"`
	WriteUid              Many2One  `xmlrpc:"write_uid"`
}

type AccountAccountTypeNil struct {
	LastUpdate            interface{} `xmlrpc:"__last_update"`
	CreateDate            interface{} `xmlrpc:"create_date"`
	CreateUid             interface{} `xmlrpc:"create_uid"`
	DisplayName           interface{} `xmlrpc:"display_name"`
	Id                    interface{} `xmlrpc:"id"`
	IncludeInitialBalance bool        `xmlrpc:"include_initial_balance"`
	Name                  interface{} `xmlrpc:"name"`
	Note                  interface{} `xmlrpc:"note"`
	Type                  interface{} `xmlrpc:"type"`
	WriteDate             interface{} `xmlrpc:"write_date"`
	WriteUid              interface{} `xmlrpc:"write_uid"`
}

var AccountAccountTypeModel string = "account.account.type"

type AccountAccountTypes []AccountAccountType

type AccountAccountTypesNil []AccountAccountTypeNil

func (s *AccountAccountType) NilableType_() interface{} {
	return &AccountAccountTypeNil{}
}

func (ns *AccountAccountTypeNil) Type_() interface{} {
	s := &AccountAccountType{}
	return load(ns, s)
}

func (s *AccountAccountTypes) NilableType_() interface{} {
	return &AccountAccountTypesNil{}
}

func (ns *AccountAccountTypesNil) Type_() interface{} {
	s := &AccountAccountTypes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAccountType))
	}
	return s
}
