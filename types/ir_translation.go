package types

import (
	"time"
)

type IrTranslation struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Comments    string    `xmlrpc:"comments"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Lang        string    `xmlrpc:"lang"`
	Module      string    `xmlrpc:"module"`
	Name        string    `xmlrpc:"name"`
	ResId       int64     `xmlrpc:"res_id"`
	Source      string    `xmlrpc:"source"`
	Src         string    `xmlrpc:"src"`
	State       string    `xmlrpc:"state"`
	Type        string    `xmlrpc:"type"`
	Value       string    `xmlrpc:"value"`
}

type IrTranslationNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Comments    interface{} `xmlrpc:"comments"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Lang        interface{} `xmlrpc:"lang"`
	Module      interface{} `xmlrpc:"module"`
	Name        interface{} `xmlrpc:"name"`
	ResId       interface{} `xmlrpc:"res_id"`
	Source      interface{} `xmlrpc:"source"`
	Src         interface{} `xmlrpc:"src"`
	State       interface{} `xmlrpc:"state"`
	Type        interface{} `xmlrpc:"type"`
	Value       interface{} `xmlrpc:"value"`
}

var IrTranslationModel string = "ir.translation"

type IrTranslations []IrTranslation

type IrTranslationsNil []IrTranslationNil

func (s *IrTranslation) NilableType_() interface{} {
	return &IrTranslationNil{}
}

func (ns *IrTranslationNil) Type_() interface{} {
	s := &IrTranslation{}
	return load(ns, s)
}

func (s *IrTranslations) NilableType_() interface{} {
	return &IrTranslationsNil{}
}

func (ns *IrTranslationsNil) Type_() interface{} {
	s := &IrTranslations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrTranslation))
	}
	return s
}
