package types

import (
	"time"
)

type AccountAccount struct {
	Code                   string      `xmlrpc:"code"`
	CompanyId              Many2One    `xmlrpc:"company_id"`
	CreateDate             time.Time   `xmlrpc:"create_date"`
	CreateUid              Many2One    `xmlrpc:"create_uid"`
	CurrencyId             Many2One    `xmlrpc:"currency_id"`
	Deprecated             bool        `xmlrpc:"deprecated"`
	DisplayName            string      `xmlrpc:"display_name"`
	GroupId                Many2One    `xmlrpc:"group_id"`
	Id                     int64       `xmlrpc:"id"`
	InternalType           interface{} `xmlrpc:"internal_type"`
	LastTimeEntriesChecked time.Time   `xmlrpc:"last_time_entries_checked"`
	LastUpdate             time.Time   `xmlrpc:"__last_update"`
	Name                   string      `xmlrpc:"name"`
	Note                   string      `xmlrpc:"note"`
	OpeningCredit          float64     `xmlrpc:"opening_credit"`
	OpeningDebit           float64     `xmlrpc:"opening_debit"`
	Reconcile              bool        `xmlrpc:"reconcile"`
	TagIds                 []int64     `xmlrpc:"tag_ids"`
	TaxIds                 []int64     `xmlrpc:"tax_ids"`
	UserTypeId             Many2One    `xmlrpc:"user_type_id"`
	WriteDate              time.Time   `xmlrpc:"write_date"`
	WriteUid               Many2One    `xmlrpc:"write_uid"`
}

type AccountAccountNil struct {
	Code                   interface{} `xmlrpc:"code"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	CurrencyId             interface{} `xmlrpc:"currency_id"`
	Deprecated             bool        `xmlrpc:"deprecated"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	GroupId                interface{} `xmlrpc:"group_id"`
	Id                     interface{} `xmlrpc:"id"`
	InternalType           interface{} `xmlrpc:"internal_type"`
	LastTimeEntriesChecked interface{} `xmlrpc:"last_time_entries_checked"`
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Name                   interface{} `xmlrpc:"name"`
	Note                   interface{} `xmlrpc:"note"`
	OpeningCredit          interface{} `xmlrpc:"opening_credit"`
	OpeningDebit           interface{} `xmlrpc:"opening_debit"`
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
