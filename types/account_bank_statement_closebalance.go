package types

import (
	"time"
)

type AccountBankStatementClosebalance struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountBankStatementClosebalanceNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountBankStatementClosebalanceModel string = "account.bank.statement.closebalance"

type AccountBankStatementClosebalances []AccountBankStatementClosebalance

type AccountBankStatementClosebalancesNil []AccountBankStatementClosebalanceNil

func (s *AccountBankStatementClosebalance) NilableType_() interface{} {
	return &AccountBankStatementClosebalanceNil{}
}

func (ns *AccountBankStatementClosebalanceNil) Type_() interface{} {
	s := &AccountBankStatementClosebalance{}
	return load(ns, s)
}

func (s *AccountBankStatementClosebalances) NilableType_() interface{} {
	return &AccountBankStatementClosebalancesNil{}
}

func (ns *AccountBankStatementClosebalancesNil) Type_() interface{} {
	s := &AccountBankStatementClosebalances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBankStatementClosebalance))
	}
	return s
}
