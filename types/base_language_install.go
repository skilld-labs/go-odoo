package types

import (
	"time"
)

type BaseLanguageInstall struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Lang        string    `xmlrpc:"lang"`
	Overwrite   bool      `xmlrpc:"overwrite"`
	State       string    `xmlrpc:"state"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseLanguageInstallNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Lang        interface{} `xmlrpc:"lang"`
	Overwrite   bool        `xmlrpc:"overwrite"`
	State       interface{} `xmlrpc:"state"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseLanguageInstallModel string = "base.language.install"

type BaseLanguageInstalls []BaseLanguageInstall

type BaseLanguageInstallsNil []BaseLanguageInstallNil

func (s *BaseLanguageInstall) NilableType_() interface{} {
	return &BaseLanguageInstallNil{}
}

func (ns *BaseLanguageInstallNil) Type_() interface{} {
	s := &BaseLanguageInstall{}
	return load(ns, s)
}

func (s *BaseLanguageInstalls) NilableType_() interface{} {
	return &BaseLanguageInstallsNil{}
}

func (ns *BaseLanguageInstallsNil) Type_() interface{} {
	s := &BaseLanguageInstalls{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseLanguageInstall))
	}
	return s
}
