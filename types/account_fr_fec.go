package types

import (
	"time"
)

type AccountFrFec struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateFrom    time.Time `xmlrpc:"date_from"`
	DateTo      time.Time `xmlrpc:"date_to"`
	DisplayName string    `xmlrpc:"display_name"`
	ExportType  string    `xmlrpc:"export_type"`
	FecData     string    `xmlrpc:"fec_data"`
	Filename    string    `xmlrpc:"filename"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountFrFecNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateFrom    interface{} `xmlrpc:"date_from"`
	DateTo      interface{} `xmlrpc:"date_to"`
	DisplayName interface{} `xmlrpc:"display_name"`
	ExportType  interface{} `xmlrpc:"export_type"`
	FecData     interface{} `xmlrpc:"fec_data"`
	Filename    interface{} `xmlrpc:"filename"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountFrFecModel string = "account.fr.fec"

type AccountFrFecs []AccountFrFec

type AccountFrFecsNil []AccountFrFecNil

func (s *AccountFrFec) NilableType_() interface{} {
	return &AccountFrFecNil{}
}

func (ns *AccountFrFecNil) Type_() interface{} {
	s := &AccountFrFec{}
	return load(ns, s)
}

func (s *AccountFrFecs) NilableType_() interface{} {
	return &AccountFrFecsNil{}
}

func (ns *AccountFrFecsNil) Type_() interface{} {
	s := &AccountFrFecs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountFrFec))
	}
	return s
}
