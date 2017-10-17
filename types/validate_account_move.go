package types

import (
	"time"
)

type ValidateAccountMove struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ValidateAccountMoveNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ValidateAccountMoveModel string = "validate.account.move"

type ValidateAccountMoves []ValidateAccountMove

type ValidateAccountMovesNil []ValidateAccountMoveNil

func (s *ValidateAccountMove) NilableType_() interface{} {
	return &ValidateAccountMoveNil{}
}

func (ns *ValidateAccountMoveNil) Type_() interface{} {
	s := &ValidateAccountMove{}
	return load(ns, s)
}

func (s *ValidateAccountMoves) NilableType_() interface{} {
	return &ValidateAccountMovesNil{}
}

func (ns *ValidateAccountMovesNil) Type_() interface{} {
	s := &ValidateAccountMoves{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ValidateAccountMove))
	}
	return s
}
