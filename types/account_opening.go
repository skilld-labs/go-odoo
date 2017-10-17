package types

import (
	"time"
)

type AccountOpening struct {
	CompanyId          Many2One  `xmlrpc:"company_id"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	CurrencyId         Many2One  `xmlrpc:"currency_id"`
	Date               time.Time `xmlrpc:"date"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	JournalId          Many2One  `xmlrpc:"journal_id"`
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	OpeningMoveId      Many2One  `xmlrpc:"opening_move_id"`
	OpeningMoveLineIds []int64   `xmlrpc:"opening_move_line_ids"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type AccountOpeningNil struct {
	CompanyId          interface{} `xmlrpc:"company_id"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	CurrencyId         interface{} `xmlrpc:"currency_id"`
	Date               interface{} `xmlrpc:"date"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	JournalId          interface{} `xmlrpc:"journal_id"`
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	OpeningMoveId      interface{} `xmlrpc:"opening_move_id"`
	OpeningMoveLineIds interface{} `xmlrpc:"opening_move_line_ids"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var AccountOpeningModel string = "account.opening"

type AccountOpenings []AccountOpening

type AccountOpeningsNil []AccountOpeningNil

func (s *AccountOpening) NilableType_() interface{} {
	return &AccountOpeningNil{}
}

func (ns *AccountOpeningNil) Type_() interface{} {
	s := &AccountOpening{}
	return load(ns, s)
}

func (s *AccountOpenings) NilableType_() interface{} {
	return &AccountOpeningsNil{}
}

func (ns *AccountOpeningsNil) Type_() interface{} {
	s := &AccountOpenings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountOpening))
	}
	return s
}
