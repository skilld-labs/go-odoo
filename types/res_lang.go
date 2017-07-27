package types

import (
	"time"
)

type ResLang struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	Active       bool      `xmlrpc:"active"`
	Code         string    `xmlrpc:"code"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DateFormat   string    `xmlrpc:"date_format"`
	DecimalPoint string    `xmlrpc:"decimal_point"`
	Direction    string    `xmlrpc:"direction"`
	DisplayName  string    `xmlrpc:"display_name"`
	Grouping     string    `xmlrpc:"grouping"`
	Id           int64     `xmlrpc:"id"`
	IsoCode      string    `xmlrpc:"iso_code"`
	Name         string    `xmlrpc:"name"`
	ThousandsSep string    `xmlrpc:"thousands_sep"`
	TimeFormat   string    `xmlrpc:"time_format"`
	Translatable bool      `xmlrpc:"translatable"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type ResLangNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	Active       bool        `xmlrpc:"active"`
	Code         interface{} `xmlrpc:"code"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DateFormat   interface{} `xmlrpc:"date_format"`
	DecimalPoint interface{} `xmlrpc:"decimal_point"`
	Direction    interface{} `xmlrpc:"direction"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Grouping     interface{} `xmlrpc:"grouping"`
	Id           interface{} `xmlrpc:"id"`
	IsoCode      interface{} `xmlrpc:"iso_code"`
	Name         interface{} `xmlrpc:"name"`
	ThousandsSep interface{} `xmlrpc:"thousands_sep"`
	TimeFormat   interface{} `xmlrpc:"time_format"`
	Translatable bool        `xmlrpc:"translatable"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var ResLangModel string = "res.lang"

type ResLangs []ResLang

type ResLangsNil []ResLangNil

func (s *ResLang) NilableType_() interface{} {
	return &ResLangNil{}
}

func (ns *ResLangNil) Type_() interface{} {
	s := &ResLang{}
	return load(ns, s)
}

func (s *ResLangs) NilableType_() interface{} {
	return &ResLangsNil{}
}

func (ns *ResLangsNil) Type_() interface{} {
	s := &ResLangs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResLang))
	}
	return s
}
