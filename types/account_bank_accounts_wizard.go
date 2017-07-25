package types

import (
	"time"
)

type AccountBankAccountsWizard struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	AccName       string    `xmlrpc:"acc_name"`
	AccountType   string    `xmlrpc:"account_type"`
	BankAccountId Many2One  `xmlrpc:"bank_account_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	CurrencyId    Many2One  `xmlrpc:"currency_id"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type AccountBankAccountsWizardNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	AccName       interface{} `xmlrpc:"acc_name"`
	AccountType   interface{} `xmlrpc:"account_type"`
	BankAccountId interface{} `xmlrpc:"bank_account_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	CurrencyId    interface{} `xmlrpc:"currency_id"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var AccountBankAccountsWizardModel string = "account.bank.accounts.wizard"

type AccountBankAccountsWizards []AccountBankAccountsWizard

type AccountBankAccountsWizardsNil []AccountBankAccountsWizardNil

func (s *AccountBankAccountsWizard) NilableType_() interface{} {
	return &AccountBankAccountsWizardNil{}
}

func (ns *AccountBankAccountsWizardNil) Type_() interface{} {
	s := &AccountBankAccountsWizard{}
	return load(ns, s)
}

func (s *AccountBankAccountsWizards) NilableType_() interface{} {
	return &AccountBankAccountsWizardsNil{}
}

func (ns *AccountBankAccountsWizardsNil) Type_() interface{} {
	s := &AccountBankAccountsWizards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountBankAccountsWizard))
	}
	return s
}
