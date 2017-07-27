package types

import (
	"time"
)

type AccountMove struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	Amount            float64   `xmlrpc:"amount"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	Date              time.Time `xmlrpc:"date"`
	DisplayName       string    `xmlrpc:"display_name"`
	DummyAccountId    Many2One  `xmlrpc:"dummy_account_id"`
	Id                int64     `xmlrpc:"id"`
	JournalId         Many2One  `xmlrpc:"journal_id"`
	LineIds           []int64   `xmlrpc:"line_ids"`
	MatchedPercentage float64   `xmlrpc:"matched_percentage"`
	Name              string    `xmlrpc:"name"`
	Narration         string    `xmlrpc:"narration"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	Ref               string    `xmlrpc:"ref"`
	State             string    `xmlrpc:"state"`
	StatementLineId   Many2One  `xmlrpc:"statement_line_id"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type AccountMoveNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	Amount            interface{} `xmlrpc:"amount"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	Date              interface{} `xmlrpc:"date"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	DummyAccountId    interface{} `xmlrpc:"dummy_account_id"`
	Id                interface{} `xmlrpc:"id"`
	JournalId         interface{} `xmlrpc:"journal_id"`
	LineIds           interface{} `xmlrpc:"line_ids"`
	MatchedPercentage interface{} `xmlrpc:"matched_percentage"`
	Name              interface{} `xmlrpc:"name"`
	Narration         interface{} `xmlrpc:"narration"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	Ref               interface{} `xmlrpc:"ref"`
	State             interface{} `xmlrpc:"state"`
	StatementLineId   interface{} `xmlrpc:"statement_line_id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var AccountMoveModel string = "account.move"

type AccountMoves []AccountMove

type AccountMovesNil []AccountMoveNil

func (s *AccountMove) NilableType_() interface{} {
	return &AccountMoveNil{}
}

func (ns *AccountMoveNil) Type_() interface{} {
	s := &AccountMove{}
	return load(ns, s)
}

func (s *AccountMoves) NilableType_() interface{} {
	return &AccountMovesNil{}
}

func (ns *AccountMovesNil) Type_() interface{} {
	s := &AccountMoves{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountMove))
	}
	return s
}
