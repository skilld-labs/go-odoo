package types

import (
	"time"
)

type AccountBankStatementLine struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	AccountId         Many2One  `xmlrpc:"account_id"`
	Amount            float64   `xmlrpc:"amount"`
	AmountCurrency    float64   `xmlrpc:"amount_currency"`
	BankAccountId     Many2One  `xmlrpc:"bank_account_id"`
	CompanyId         Many2One  `xmlrpc:"company_id"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	Date              time.Time `xmlrpc:"date"`
	DisplayName       string    `xmlrpc:"display_name"`
	Id                int64     `xmlrpc:"id"`
	JournalCurrencyId Many2One  `xmlrpc:"journal_currency_id"`
	JournalEntryIds   []int64   `xmlrpc:"journal_entry_ids"`
	JournalId         Many2One  `xmlrpc:"journal_id"`
	MoveName          string    `xmlrpc:"move_name"`
	Name              string    `xmlrpc:"name"`
	Note              string    `xmlrpc:"note"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	PartnerName       string    `xmlrpc:"partner_name"`
	Ref               string    `xmlrpc:"ref"`
	Sequence          int64     `xmlrpc:"sequence"`
	State             string    `xmlrpc:"state"`
	StatementId       Many2One  `xmlrpc:"statement_id"`
	UniqueImportId    string    `xmlrpc:"unique_import_id"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type AccountBankStatementLineNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	AccountId         interface{} `xmlrpc:"account_id"`
	Amount            interface{} `xmlrpc:"amount"`
	AmountCurrency    interface{} `xmlrpc:"amount_currency"`
	BankAccountId     interface{} `xmlrpc:"bank_account_id"`
	CompanyId         interface{} `xmlrpc:"company_id"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	Date              interface{} `xmlrpc:"date"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Id                interface{} `xmlrpc:"id"`
	JournalCurrencyId interface{} `xmlrpc:"journal_currency_id"`
	JournalEntryIds   interface{} `xmlrpc:"journal_entry_ids"`
	JournalId         interface{} `xmlrpc:"journal_id"`
	MoveName          interface{} `xmlrpc:"move_name"`
	Name              interface{} `xmlrpc:"name"`
	Note              interface{} `xmlrpc:"note"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	PartnerName       interface{} `xmlrpc:"partner_name"`
	Ref               interface{} `xmlrpc:"ref"`
	Sequence          interface{} `xmlrpc:"sequence"`
	State             interface{} `xmlrpc:"state"`
	StatementId       interface{} `xmlrpc:"statement_id"`
	UniqueImportId    interface{} `xmlrpc:"unique_import_id"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var AccountBankStatementLineModel string = "account.bank.statement.line"

type AccountBankStatementLines []AccountBankStatementLine

type AccountBankStatementLinesNil []AccountBankStatementLineNil

func (s *AccountBankStatementLine) NilableType_() interface{} {
	return &AccountBankStatementLineNil{}
}

func (ns *AccountBankStatementLineNil) Type_() interface{} {
	s := &AccountBankStatementLine{}
	return load(ns, s)
}

func (s *AccountBankStatementLines) NilableType_() interface{} {
	return &AccountBankStatementLinesNil{}
}

func (ns *AccountBankStatementLinesNil) Type_() interface{} {
	s := &AccountBankStatementLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBankStatementLine))
	}
	return s
}
