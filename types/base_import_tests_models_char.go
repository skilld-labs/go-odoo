package types

import (
	"time"
)

type BaseImportTestsModelsChar struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsCharModel string = "base_import.tests.models.char"

type BaseImportTestsModelsChars []BaseImportTestsModelsChar

type BaseImportTestsModelsCharsNil []BaseImportTestsModelsCharNil

func (s *BaseImportTestsModelsChar) NilableType_() interface{} {
	return &BaseImportTestsModelsCharNil{}
}

func (ns *BaseImportTestsModelsCharNil) Type_() interface{} {
	s := &BaseImportTestsModelsChar{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsChars) NilableType_() interface{} {
	return &BaseImportTestsModelsCharsNil{}
}

func (ns *BaseImportTestsModelsCharsNil) Type_() interface{} {
	s := &BaseImportTestsModelsChars{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsChar))
	}
	return s
}
