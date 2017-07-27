package types

import (
	"time"
)

type ResCurrencyRate struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CompanyId   Many2One  `xmlrpc:"company_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	CurrencyId  Many2One  `xmlrpc:"currency_id"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        time.Time `xmlrpc:"name"`
	Rate        float64   `xmlrpc:"rate"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ResCurrencyRateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CompanyId   interface{} `xmlrpc:"company_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	CurrencyId  interface{} `xmlrpc:"currency_id"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Rate        interface{} `xmlrpc:"rate"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ResCurrencyRateModel string = "res.currency.rate"

type ResCurrencyRates []ResCurrencyRate

type ResCurrencyRatesNil []ResCurrencyRateNil

func (s *ResCurrencyRate) NilableType_() interface{} {
	return &ResCurrencyRateNil{}
}

func (ns *ResCurrencyRateNil) Type_() interface{} {
	s := &ResCurrencyRate{}
	return load(ns, s)
}

func (s *ResCurrencyRates) NilableType_() interface{} {
	return &ResCurrencyRatesNil{}
}

func (ns *ResCurrencyRatesNil) Type_() interface{} {
	s := &ResCurrencyRates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCurrencyRate))
	}
	return s
}
