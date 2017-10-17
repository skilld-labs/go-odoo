package types

import (
	"time"
)

type ResCurrency struct {
	Active               bool      `xmlrpc:"active"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	CurrencySubunitLabel string    `xmlrpc:"currency_subunit_label"`
	CurrencyUnitLabel    string    `xmlrpc:"currency_unit_label"`
	Date                 time.Time `xmlrpc:"date"`
	DecimalPlaces        int64     `xmlrpc:"decimal_places"`
	DisplayName          string    `xmlrpc:"display_name"`
	Id                   int64     `xmlrpc:"id"`
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	Name                 string    `xmlrpc:"name"`
	Position             string    `xmlrpc:"position"`
	Rate                 float64   `xmlrpc:"rate"`
	RateIds              []int64   `xmlrpc:"rate_ids"`
	Rounding             float64   `xmlrpc:"rounding"`
	Symbol               string    `xmlrpc:"symbol"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
}

type ResCurrencyNil struct {
	Active               bool        `xmlrpc:"active"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	CurrencySubunitLabel interface{} `xmlrpc:"currency_subunit_label"`
	CurrencyUnitLabel    interface{} `xmlrpc:"currency_unit_label"`
	Date                 interface{} `xmlrpc:"date"`
	DecimalPlaces        interface{} `xmlrpc:"decimal_places"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	Id                   interface{} `xmlrpc:"id"`
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	Name                 interface{} `xmlrpc:"name"`
	Position             interface{} `xmlrpc:"position"`
	Rate                 interface{} `xmlrpc:"rate"`
	RateIds              interface{} `xmlrpc:"rate_ids"`
	Rounding             interface{} `xmlrpc:"rounding"`
	Symbol               interface{} `xmlrpc:"symbol"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
}

var ResCurrencyModel string = "res.currency"

type ResCurrencys []ResCurrency

type ResCurrencysNil []ResCurrencyNil

func (s *ResCurrency) NilableType_() interface{} {
	return &ResCurrencyNil{}
}

func (ns *ResCurrencyNil) Type_() interface{} {
	s := &ResCurrency{}
	return load(ns, s)
}

func (s *ResCurrencys) NilableType_() interface{} {
	return &ResCurrencysNil{}
}

func (ns *ResCurrencysNil) Type_() interface{} {
	s := &ResCurrencys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCurrency))
	}
	return s
}
