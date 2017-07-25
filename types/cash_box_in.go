package types

import (
	"time"
)

type CashBoxIn struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Amount      float64   `xmlrpc:"amount"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Ref         string    `xmlrpc:"ref"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type CashBoxInNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Amount      interface{} `xmlrpc:"amount"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Ref         interface{} `xmlrpc:"ref"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var CashBoxInModel string = "cash.box.in"

type CashBoxIns []CashBoxIn

type CashBoxInsNil []CashBoxInNil

func (s *CashBoxIn) NilableType_() interface{} {
	return &CashBoxInNil{}
}

func (ns *CashBoxInNil) Type_() interface{} {
	s := &CashBoxIn{}
	return load(ns, s)
}

func (s *CashBoxIns) NilableType_() interface{} {
	return &CashBoxInsNil{}
}

func (ns *CashBoxInsNil) Type_() interface{} {
	s := &CashBoxIns{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*CashBoxIn))
	}
	return s
}
