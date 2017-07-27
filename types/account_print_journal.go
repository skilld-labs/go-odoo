package types

import (
	"time"
)

type AccountPrintJournal struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	AmountCurrency bool      `xmlrpc:"amount_currency"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DateFrom       time.Time `xmlrpc:"date_from"`
	DateTo         time.Time `xmlrpc:"date_to"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	JournalIds     []int64   `xmlrpc:"journal_ids"`
	SortSelection  string    `xmlrpc:"sort_selection"`
	TargetMove     string    `xmlrpc:"target_move"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type AccountPrintJournalNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	AmountCurrency bool        `xmlrpc:"amount_currency"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DateFrom       interface{} `xmlrpc:"date_from"`
	DateTo         interface{} `xmlrpc:"date_to"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	JournalIds     interface{} `xmlrpc:"journal_ids"`
	SortSelection  interface{} `xmlrpc:"sort_selection"`
	TargetMove     interface{} `xmlrpc:"target_move"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var AccountPrintJournalModel string = "account.print.journal"

type AccountPrintJournals []AccountPrintJournal

type AccountPrintJournalsNil []AccountPrintJournalNil

func (s *AccountPrintJournal) NilableType_() interface{} {
	return &AccountPrintJournalNil{}
}

func (ns *AccountPrintJournalNil) Type_() interface{} {
	s := &AccountPrintJournal{}
	return load(ns, s)
}

func (s *AccountPrintJournals) NilableType_() interface{} {
	return &AccountPrintJournalsNil{}
}

func (ns *AccountPrintJournalsNil) Type_() interface{} {
	s := &AccountPrintJournals{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountPrintJournal))
	}
	return s
}
