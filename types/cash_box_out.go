package types

import (
	"time"
)

type CashBoxOut struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Amount      float64   `xmlrpc:"amount"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CashBoxOutNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Amount      interface{} `xmlrpc:"amount"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CashBoxOutModel string = "cash.box.out"

type CashBoxOuts []CashBoxOut

type CashBoxOutsNil []CashBoxOutNil

func (s *CashBoxOut) NilableType_() interface{} {
	return &CashBoxOutNil{}
}

func (ns *CashBoxOutNil) Type_() interface{} {
	s := &CashBoxOut{}
	return load(ns, s)
}

func (s *CashBoxOuts) NilableType_() interface{} {
	return &CashBoxOutsNil{}
}

func (ns *CashBoxOutsNil) Type_() interface{} {
	s := &CashBoxOuts{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CashBoxOut))
	}
	return s
}
