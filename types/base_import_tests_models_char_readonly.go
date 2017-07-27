package types

import (
	"time"
)

type BaseImportTestsModelsCharReadonly struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharReadonlyNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsCharReadonlyModel string = "base_import.tests.models.char.readonly"

type BaseImportTestsModelsCharReadonlys []BaseImportTestsModelsCharReadonly

type BaseImportTestsModelsCharReadonlysNil []BaseImportTestsModelsCharReadonlyNil

func (s *BaseImportTestsModelsCharReadonly) NilableType_() interface{} {
	return &BaseImportTestsModelsCharReadonlyNil{}
}

func (ns *BaseImportTestsModelsCharReadonlyNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharReadonly{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsCharReadonlys) NilableType_() interface{} {
	return &BaseImportTestsModelsCharReadonlysNil{}
}

func (ns *BaseImportTestsModelsCharReadonlysNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharReadonlys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsCharReadonly))
	}
	return s
}
