package types

import (
	"time"
)

type BaseLanguageImport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Code        string    `xmlrpc:"code"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Data        string    `xmlrpc:"data"`
	DisplayName string    `xmlrpc:"display_name"`
	Filename    string    `xmlrpc:"filename"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Overwrite   bool      `xmlrpc:"overwrite"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseLanguageImportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Code        interface{} `xmlrpc:"code"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Data        interface{} `xmlrpc:"data"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Filename    interface{} `xmlrpc:"filename"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Overwrite   bool        `xmlrpc:"overwrite"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseLanguageImportModel string = "base.language.import"

type BaseLanguageImports []BaseLanguageImport

type BaseLanguageImportsNil []BaseLanguageImportNil

func (s *BaseLanguageImport) NilableType_() interface{} {
	return &BaseLanguageImportNil{}
}

func (ns *BaseLanguageImportNil) Type_() interface{} {
	s := &BaseLanguageImport{}
	return load(ns, s)
}

func (s *BaseLanguageImports) NilableType_() interface{} {
	return &BaseLanguageImportsNil{}
}

func (ns *BaseLanguageImportsNil) Type_() interface{} {
	s := &BaseLanguageImports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseLanguageImport))
	}
	return s
}
