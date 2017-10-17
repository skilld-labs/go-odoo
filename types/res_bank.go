package types

import (
	"time"
)

type ResBank struct {
	Active      bool      `xmlrpc:"active"`
	Bic         string    `xmlrpc:"bic"`
	City        string    `xmlrpc:"city"`
	Country     Many2One  `xmlrpc:"country"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Email       string    `xmlrpc:"email"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	Phone       string    `xmlrpc:"phone"`
	State       Many2One  `xmlrpc:"state"`
	Street      string    `xmlrpc:"street"`
	Street2     string    `xmlrpc:"street2"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
	Zip         string    `xmlrpc:"zip"`
}

type ResBankNil struct {
	Active      bool        `xmlrpc:"active"`
	Bic         interface{} `xmlrpc:"bic"`
	City        interface{} `xmlrpc:"city"`
	Country     interface{} `xmlrpc:"country"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Email       interface{} `xmlrpc:"email"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	Phone       interface{} `xmlrpc:"phone"`
	State       interface{} `xmlrpc:"state"`
	Street      interface{} `xmlrpc:"street"`
	Street2     interface{} `xmlrpc:"street2"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
	Zip         interface{} `xmlrpc:"zip"`
}

var ResBankModel string = "res.bank"

type ResBanks []ResBank

type ResBanksNil []ResBankNil

func (s *ResBank) NilableType_() interface{} {
	return &ResBankNil{}
}

func (ns *ResBankNil) Type_() interface{} {
	s := &ResBank{}
	return load(ns, s)
}

func (s *ResBanks) NilableType_() interface{} {
	return &ResBanksNil{}
}

func (ns *ResBanksNil) Type_() interface{} {
	s := &ResBanks{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResBank))
	}
	return s
}
