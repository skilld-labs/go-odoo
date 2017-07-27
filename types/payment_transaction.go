package types

import (
	"time"
)

type PaymentTransaction struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	AcquirerId        Many2One  `xmlrpc:"acquirer_id"`
	AcquirerReference string    `xmlrpc:"acquirer_reference"`
	Amount            float64   `xmlrpc:"amount"`
	CallbackEval      string    `xmlrpc:"callback_eval"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	CurrencyId        Many2One  `xmlrpc:"currency_id"`
	DateValidate      time.Time `xmlrpc:"date_validate"`
	DisplayName       string    `xmlrpc:"display_name"`
	Fees              float64   `xmlrpc:"fees"`
	Html3ds           string    `xmlrpc:"html_3ds"`
	Id                int64     `xmlrpc:"id"`
	PartnerAddress    string    `xmlrpc:"partner_address"`
	PartnerCity       string    `xmlrpc:"partner_city"`
	PartnerCountryId  Many2One  `xmlrpc:"partner_country_id"`
	PartnerEmail      string    `xmlrpc:"partner_email"`
	PartnerId         Many2One  `xmlrpc:"partner_id"`
	PartnerLang       string    `xmlrpc:"partner_lang"`
	PartnerName       string    `xmlrpc:"partner_name"`
	PartnerPhone      string    `xmlrpc:"partner_phone"`
	PartnerZip        string    `xmlrpc:"partner_zip"`
	PaymentTokenId    Many2One  `xmlrpc:"payment_token_id"`
	Reference         string    `xmlrpc:"reference"`
	State             string    `xmlrpc:"state"`
	StateMessage      string    `xmlrpc:"state_message"`
	Type              string    `xmlrpc:"type"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type PaymentTransactionNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	AcquirerId        interface{} `xmlrpc:"acquirer_id"`
	AcquirerReference interface{} `xmlrpc:"acquirer_reference"`
	Amount            interface{} `xmlrpc:"amount"`
	CallbackEval      interface{} `xmlrpc:"callback_eval"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	CurrencyId        interface{} `xmlrpc:"currency_id"`
	DateValidate      interface{} `xmlrpc:"date_validate"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	Fees              interface{} `xmlrpc:"fees"`
	Html3ds           interface{} `xmlrpc:"html_3ds"`
	Id                interface{} `xmlrpc:"id"`
	PartnerAddress    interface{} `xmlrpc:"partner_address"`
	PartnerCity       interface{} `xmlrpc:"partner_city"`
	PartnerCountryId  interface{} `xmlrpc:"partner_country_id"`
	PartnerEmail      interface{} `xmlrpc:"partner_email"`
	PartnerId         interface{} `xmlrpc:"partner_id"`
	PartnerLang       interface{} `xmlrpc:"partner_lang"`
	PartnerName       interface{} `xmlrpc:"partner_name"`
	PartnerPhone      interface{} `xmlrpc:"partner_phone"`
	PartnerZip        interface{} `xmlrpc:"partner_zip"`
	PaymentTokenId    interface{} `xmlrpc:"payment_token_id"`
	Reference         interface{} `xmlrpc:"reference"`
	State             interface{} `xmlrpc:"state"`
	StateMessage      interface{} `xmlrpc:"state_message"`
	Type              interface{} `xmlrpc:"type"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var PaymentTransactionModel string = "payment.transaction"

type PaymentTransactions []PaymentTransaction

type PaymentTransactionsNil []PaymentTransactionNil

func (s *PaymentTransaction) NilableType_() interface{} {
	return &PaymentTransactionNil{}
}

func (ns *PaymentTransactionNil) Type_() interface{} {
	s := &PaymentTransaction{}
	return load(ns, s)
}

func (s *PaymentTransactions) NilableType_() interface{} {
	return &PaymentTransactionsNil{}
}

func (ns *PaymentTransactionsNil) Type_() interface{} {
	s := &PaymentTransactions{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PaymentTransaction))
	}
	return s
}
