package types

import (
	"time"
)

type AccountMoveReversal struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Date        time.Time `xmlrpc:"date"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	JournalId   Many2One  `xmlrpc:"journal_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountMoveReversalNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Date        interface{} `xmlrpc:"date"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	JournalId   interface{} `xmlrpc:"journal_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountMoveReversalModel string = "account.move.reversal"

type AccountMoveReversals []AccountMoveReversal

type AccountMoveReversalsNil []AccountMoveReversalNil

func (s *AccountMoveReversal) NilableType_() interface{} {
	return &AccountMoveReversalNil{}
}

func (ns *AccountMoveReversalNil) Type_() interface{} {
	s := &AccountMoveReversal{}
	return load(ns, s)
}

func (s *AccountMoveReversals) NilableType_() interface{} {
	return &AccountMoveReversalsNil{}
}

func (ns *AccountMoveReversalsNil) Type_() interface{} {
	s := &AccountMoveReversals{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountMoveReversal))
	}
	return s
}
