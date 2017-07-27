package types

import (
	"time"
)

type AccountBankStatementImport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DataFile    string    `xmlrpc:"data_file"`
	DisplayName string    `xmlrpc:"display_name"`
	Filename    string    `xmlrpc:"filename"`
	Id          int64     `xmlrpc:"id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type AccountBankStatementImportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DataFile    interface{} `xmlrpc:"data_file"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Filename    interface{} `xmlrpc:"filename"`
	Id          interface{} `xmlrpc:"id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var AccountBankStatementImportModel string = "account.bank.statement.import"

type AccountBankStatementImports []AccountBankStatementImport

type AccountBankStatementImportsNil []AccountBankStatementImportNil

func (s *AccountBankStatementImport) NilableType_() interface{} {
	return &AccountBankStatementImportNil{}
}

func (ns *AccountBankStatementImportNil) Type_() interface{} {
	s := &AccountBankStatementImport{}
	return load(ns, s)
}

func (s *AccountBankStatementImports) NilableType_() interface{} {
	return &AccountBankStatementImportsNil{}
}

func (ns *AccountBankStatementImportsNil) Type_() interface{} {
	s := &AccountBankStatementImports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBankStatementImport))
	}
	return s
}
