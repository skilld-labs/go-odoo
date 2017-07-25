package types

import (
	"time"
)

type BaseLanguageExport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	Data        string    `xmlrpc:"data"`
	DisplayName string    `xmlrpc:"display_name"`
	Format      string    `xmlrpc:"format"`
	Id          int64     `xmlrpc:"id"`
	Lang        string    `xmlrpc:"lang"`
	Modules     []int64   `xmlrpc:"modules"`
	Name        string    `xmlrpc:"name"`
	State       string    `xmlrpc:"state"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseLanguageExportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	Data        interface{} `xmlrpc:"data"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Format      interface{} `xmlrpc:"format"`
	Id          interface{} `xmlrpc:"id"`
	Lang        interface{} `xmlrpc:"lang"`
	Modules     interface{} `xmlrpc:"modules"`
	Name        interface{} `xmlrpc:"name"`
	State       interface{} `xmlrpc:"state"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseLanguageExportModel string = "base.language.export"

type BaseLanguageExports []BaseLanguageExport

type BaseLanguageExportsNil []BaseLanguageExportNil

func (s *BaseLanguageExport) NilableType_() interface{} {
	return &BaseLanguageExportNil{}
}

func (ns *BaseLanguageExportNil) Type_() interface{} {
	s := &BaseLanguageExport{}
	return load(ns, s)
}

func (s *BaseLanguageExports) NilableType_() interface{} {
	return &BaseLanguageExportsNil{}
}

func (ns *BaseLanguageExportsNil) Type_() interface{} {
	s := &BaseLanguageExports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseLanguageExport))
	}
	return s
}
