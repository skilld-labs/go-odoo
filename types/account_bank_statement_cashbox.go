package types

import (
	"time"
)

type AccountBankStatementCashbox struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CashboxLinesIds []int64   `xmlrpc:"cashbox_lines_ids"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type AccountBankStatementCashboxNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CashboxLinesIds interface{} `xmlrpc:"cashbox_lines_ids"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var AccountBankStatementCashboxModel string = "account.bank.statement.cashbox"

type AccountBankStatementCashboxs []AccountBankStatementCashbox

type AccountBankStatementCashboxsNil []AccountBankStatementCashboxNil

func (s *AccountBankStatementCashbox) NilableType_() interface{} {
	return &AccountBankStatementCashboxNil{}
}

func (ns *AccountBankStatementCashboxNil) Type_() interface{} {
	s := &AccountBankStatementCashbox{}
	return load(ns, s)
}

func (s *AccountBankStatementCashboxs) NilableType_() interface{} {
	return &AccountBankStatementCashboxsNil{}
}

func (ns *AccountBankStatementCashboxsNil) Type_() interface{} {
	s := &AccountBankStatementCashboxs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBankStatementCashbox))
	}
	return s
}
