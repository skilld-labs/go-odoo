package types

import (
	"time"
)

type AccountCashRounding struct {
	AccountId      Many2One  `xmlrpc:"account_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Name           string    `xmlrpc:"name"`
	Rounding       float64   `xmlrpc:"rounding"`
	RoundingMethod string    `xmlrpc:"rounding_method"`
	Strategy       string    `xmlrpc:"strategy"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type AccountCashRoundingNil struct {
	AccountId      interface{} `xmlrpc:"account_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Name           interface{} `xmlrpc:"name"`
	Rounding       interface{} `xmlrpc:"rounding"`
	RoundingMethod interface{} `xmlrpc:"rounding_method"`
	Strategy       interface{} `xmlrpc:"strategy"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var AccountCashRoundingModel string = "account.cash.rounding"

type AccountCashRoundings []AccountCashRounding

type AccountCashRoundingsNil []AccountCashRoundingNil

func (s *AccountCashRounding) NilableType_() interface{} {
	return &AccountCashRoundingNil{}
}

func (ns *AccountCashRoundingNil) Type_() interface{} {
	s := &AccountCashRounding{}
	return load(ns, s)
}

func (s *AccountCashRoundings) NilableType_() interface{} {
	return &AccountCashRoundingsNil{}
}

func (ns *AccountCashRoundingsNil) Type_() interface{} {
	s := &AccountCashRoundings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountCashRounding))
	}
	return s
}
