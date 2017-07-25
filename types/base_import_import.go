package types

import (
	"time"
)

type BaseImportImport struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	File        string    `xmlrpc:"file"`
	FileName    string    `xmlrpc:"file_name"`
	FileType    string    `xmlrpc:"file_type"`
	Id          int64     `xmlrpc:"id"`
	ResModel    string    `xmlrpc:"res_model"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportImportNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	File        interface{} `xmlrpc:"file"`
	FileName    interface{} `xmlrpc:"file_name"`
	FileType    interface{} `xmlrpc:"file_type"`
	Id          interface{} `xmlrpc:"id"`
	ResModel    interface{} `xmlrpc:"res_model"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportImportModel string = "base_import.import"

type BaseImportImports []BaseImportImport

type BaseImportImportsNil []BaseImportImportNil

func (s *BaseImportImport) NilableType_() interface{} {
	return &BaseImportImportNil{}
}

func (ns *BaseImportImportNil) Type_() interface{} {
	s := &BaseImportImport{}
	return load(ns, s)
}

func (s *BaseImportImports) NilableType_() interface{} {
	return &BaseImportImportsNil{}
}

func (ns *BaseImportImportsNil) Type_() interface{} {
	s := &BaseImportImports{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportImport))
	}
	return s
}
