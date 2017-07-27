package types

import (
	"time"
)

type ResPartnerBank struct {
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	AccNumber          string    `xmlrpc:"acc_number"`
	AccType            string    `xmlrpc:"acc_type"`
	BankBic            string    `xmlrpc:"bank_bic"`
	BankId             Many2One  `xmlrpc:"bank_id"`
	BankName           string    `xmlrpc:"bank_name"`
	CompanyId          Many2One  `xmlrpc:"company_id"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	CurrencyId         Many2One  `xmlrpc:"currency_id"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	JournalId          []int64   `xmlrpc:"journal_id"`
	PartnerId          Many2One  `xmlrpc:"partner_id"`
	SanitizedAccNumber string    `xmlrpc:"sanitized_acc_number"`
	Sequence           int64     `xmlrpc:"sequence"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type ResPartnerBankNil struct {
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	AccNumber          interface{} `xmlrpc:"acc_number"`
	AccType            interface{} `xmlrpc:"acc_type"`
	BankBic            interface{} `xmlrpc:"bank_bic"`
	BankId             interface{} `xmlrpc:"bank_id"`
	BankName           interface{} `xmlrpc:"bank_name"`
	CompanyId          interface{} `xmlrpc:"company_id"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	CurrencyId         interface{} `xmlrpc:"currency_id"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	JournalId          interface{} `xmlrpc:"journal_id"`
	PartnerId          interface{} `xmlrpc:"partner_id"`
	SanitizedAccNumber interface{} `xmlrpc:"sanitized_acc_number"`
	Sequence           interface{} `xmlrpc:"sequence"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var ResPartnerBankModel string = "res.partner.bank"

type ResPartnerBanks []ResPartnerBank

type ResPartnerBanksNil []ResPartnerBankNil

func (s *ResPartnerBank) NilableType_() interface{} {
	return &ResPartnerBankNil{}
}

func (ns *ResPartnerBankNil) Type_() interface{} {
	s := &ResPartnerBank{}
	return load(ns, s)
}

func (s *ResPartnerBanks) NilableType_() interface{} {
	return &ResPartnerBanksNil{}
}

func (ns *ResPartnerBanksNil) Type_() interface{} {
	s := &ResPartnerBanks{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResPartnerBank))
	}
	return s
}
