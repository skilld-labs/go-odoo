package types

import (
	"time"
)

type AccountUnreconcile struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountUnreconcileNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountUnreconcileModel string = "account.unreconcile"

type AccountUnreconciles []AccountUnreconcile

type AccountUnreconcilesNil []AccountUnreconcileNil

func (s *AccountUnreconcile) NilableType_() interface{} {
	return &AccountUnreconcileNil{}
}

func (ns *AccountUnreconcileNil) Type_() interface{} {
	s := &AccountUnreconcile{}
	return load(ns, s)
}

func (s *AccountUnreconciles) NilableType_() interface{} {
	return &AccountUnreconcilesNil{}
}

func (ns *AccountUnreconcilesNil) Type_() interface{} {
	s := &AccountUnreconciles{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountUnreconcile))
	}
	return s
}
