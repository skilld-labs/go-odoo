package types

import (
	"time"
)

type BaseImportTestsModelsCharStillreadonly struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharStillreadonlyNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsCharStillreadonlyModel string = "base_import.tests.models.char.stillreadonly"

type BaseImportTestsModelsCharStillreadonlys []BaseImportTestsModelsCharStillreadonly

type BaseImportTestsModelsCharStillreadonlysNil []BaseImportTestsModelsCharStillreadonlyNil

func (s *BaseImportTestsModelsCharStillreadonly) NilableType_() interface{} {
	return &BaseImportTestsModelsCharStillreadonlyNil{}
}

func (ns *BaseImportTestsModelsCharStillreadonlyNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharStillreadonly{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsCharStillreadonlys) NilableType_() interface{} {
	return &BaseImportTestsModelsCharStillreadonlysNil{}
}

func (ns *BaseImportTestsModelsCharStillreadonlysNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharStillreadonlys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsCharStillreadonly))
	}
	return s
}
