package types

import (
	"time"
)

type ResCountry struct {
	AddressFormat   string      `xmlrpc:"address_format"`
	AddressViewId   Many2One    `xmlrpc:"address_view_id"`
	Code            string      `xmlrpc:"code"`
	CountryGroupIds []int64     `xmlrpc:"country_group_ids"`
	CreateDate      time.Time   `xmlrpc:"create_date"`
	CreateUid       Many2One    `xmlrpc:"create_uid"`
	CurrencyId      Many2One    `xmlrpc:"currency_id"`
	DisplayName     string      `xmlrpc:"display_name"`
	Id              int64       `xmlrpc:"id"`
	Image           string      `xmlrpc:"image"`
	LastUpdate      time.Time   `xmlrpc:"__last_update"`
	Name            string      `xmlrpc:"name"`
	NamePosition    interface{} `xmlrpc:"name_position"`
	PhoneCode       int64       `xmlrpc:"phone_code"`
	StateIds        []int64     `xmlrpc:"state_ids"`
	VatLabel        string      `xmlrpc:"vat_label"`
	WriteDate       time.Time   `xmlrpc:"write_date"`
	WriteUid        Many2One    `xmlrpc:"write_uid"`
}

type ResCountryNil struct {
	AddressFormat   interface{} `xmlrpc:"address_format"`
	AddressViewId   interface{} `xmlrpc:"address_view_id"`
	Code            interface{} `xmlrpc:"code"`
	CountryGroupIds interface{} `xmlrpc:"country_group_ids"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	CurrencyId      interface{} `xmlrpc:"currency_id"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	Image           interface{} `xmlrpc:"image"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	Name            interface{} `xmlrpc:"name"`
	NamePosition    interface{} `xmlrpc:"name_position"`
	PhoneCode       interface{} `xmlrpc:"phone_code"`
	StateIds        interface{} `xmlrpc:"state_ids"`
	VatLabel        interface{} `xmlrpc:"vat_label"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var ResCountryModel string = "res.country"

type ResCountrys []ResCountry

type ResCountrysNil []ResCountryNil

func (s *ResCountry) NilableType_() interface{} {
	return &ResCountryNil{}
}

func (ns *ResCountryNil) Type_() interface{} {
	s := &ResCountry{}
	return load(ns, s)
}

func (s *ResCountrys) NilableType_() interface{} {
	return &ResCountrysNil{}
}

func (ns *ResCountrysNil) Type_() interface{} {
	s := &ResCountrys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResCountry))
	}
	return s
}
