package types

import (
	"time"
)

type BaseImportTestsModelsCharRequired struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharRequiredNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsCharRequiredModel string = "base_import.tests.models.char.required"

type BaseImportTestsModelsCharRequireds []BaseImportTestsModelsCharRequired

type BaseImportTestsModelsCharRequiredsNil []BaseImportTestsModelsCharRequiredNil

func (s *BaseImportTestsModelsCharRequired) NilableType_() interface{} {
	return &BaseImportTestsModelsCharRequiredNil{}
}

func (ns *BaseImportTestsModelsCharRequiredNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharRequired{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsCharRequireds) NilableType_() interface{} {
	return &BaseImportTestsModelsCharRequiredsNil{}
}

func (ns *BaseImportTestsModelsCharRequiredsNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharRequireds{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsCharRequired))
	}
	return s
}
