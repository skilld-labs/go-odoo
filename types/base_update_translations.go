package types

import (
	"time"
)

type BaseUpdateTranslations struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Lang        string    `xmlrpc:"lang"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseUpdateTranslationsNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Lang        interface{} `xmlrpc:"lang"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseUpdateTranslationsModel string = "base.update.translations"

type BaseUpdateTranslationss []BaseUpdateTranslations

type BaseUpdateTranslationssNil []BaseUpdateTranslationsNil

func (s *BaseUpdateTranslations) NilableType_() interface{} {
	return &BaseUpdateTranslationsNil{}
}

func (ns *BaseUpdateTranslationsNil) Type_() interface{} {
	s := &BaseUpdateTranslations{}
	return load(ns, s)
}

func (s *BaseUpdateTranslationss) NilableType_() interface{} {
	return &BaseUpdateTranslationssNil{}
}

func (ns *BaseUpdateTranslationssNil) Type_() interface{} {
	s := &BaseUpdateTranslationss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseUpdateTranslations))
	}
	return s
}
