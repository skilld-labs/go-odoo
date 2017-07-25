package types

import (
	"time"
)

type AccountAccount struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	Code                   string    `xmlrpc:"code"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	CurrencyId             Many2One  `xmlrpc:"currency_id"`
	Deprecated             bool      `xmlrpc:"deprecated"`
	DisplayName            string    `xmlrpc:"display_name"`
	Id                     int64     `xmlrpc:"id"`
	InternalType           string    `xmlrpc:"internal_type"`
	LastTimeEntriesChecked time.Time `xmlrpc:"last_time_entries_checked"`
	Name                   string    `xmlrpc:"name"`
	Note                   string    `xmlrpc:"note"`
	Reconcile              bool      `xmlrpc:"reconcile"`
	TagIds                 []int64   `xmlrpc:"tag_ids"`
	TaxIds                 []int64   `xmlrpc:"tax_ids"`
	UserTypeId             Many2One  `xmlrpc:"user_type_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type AccountAccountNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Code                   interface{} `xmlrpc:"code"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	CurrencyId             interface{} `xmlrpc:"currency_id"`
	Deprecated             bool        `xmlrpc:"deprecated"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	Id                     interface{} `xmlrpc:"id"`
	InternalType           interface{} `xmlrpc:"internal_type"`
	LastTimeEntriesChecked interface{} `xmlrpc:"last_time_entries_checked"`
	Name                   interface{} `xmlrpc:"name"`
	Note                   interface{} `xmlrpc:"note"`
	Reconcile              bool        `xmlrpc:"reconcile"`
	TagIds                 interface{} `xmlrpc:"tag_ids"`
	TaxIds                 interface{} `xmlrpc:"tax_ids"`
	UserTypeId             interface{} `xmlrpc:"user_type_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var AccountAccountModel string = "account.account"

type AccountAccounts []AccountAccount

type AccountAccountsNil []AccountAccountNil

func (s *AccountAccount) NilableType_() interface{} {
	return &AccountAccountNil{}
}

func (ns *AccountAccountNil) Type_() interface{} {
	s := &AccountAccount{}
	return load(ns, s)
}

func (s *AccountAccounts) NilableType_() interface{} {
	return &AccountAccountsNil{}
}

func (ns *AccountAccountsNil) Type_() interface{} {
	s := &AccountAccounts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAccount))
	}
	return s
}
