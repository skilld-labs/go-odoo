package types

import (
	"time"
)

type AccountCashboxLine struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CashboxId   Many2One  `xmlrpc:"cashbox_id"`
	CoinValue   float64   `xmlrpc:"coin_value"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Number      int64     `xmlrpc:"number"`
	Subtotal    float64   `xmlrpc:"subtotal"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountCashboxLineNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CashboxId   interface{} `xmlrpc:"cashbox_id"`
	CoinValue   interface{} `xmlrpc:"coin_value"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Number      interface{} `xmlrpc:"number"`
	Subtotal    interface{} `xmlrpc:"subtotal"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountCashboxLineModel string = "account.cashbox.line"

type AccountCashboxLines []AccountCashboxLine

type AccountCashboxLinesNil []AccountCashboxLineNil

func (s *AccountCashboxLine) NilableType_() interface{} {
	return &AccountCashboxLineNil{}
}

func (ns *AccountCashboxLineNil) Type_() interface{} {
	s := &AccountCashboxLine{}
	return load(ns, s)
}

func (s *AccountCashboxLines) NilableType_() interface{} {
	return &AccountCashboxLinesNil{}
}

func (ns *AccountCashboxLinesNil) Type_() interface{} {
	s := &AccountCashboxLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountCashboxLine))
	}
	return s
}
