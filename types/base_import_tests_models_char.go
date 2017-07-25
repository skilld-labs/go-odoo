package types

import (
	"time"
)

type BaseImportTestsModelsChar struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
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
