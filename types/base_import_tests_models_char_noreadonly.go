package types

import (
	"time"
)

type BaseImportTestsModelsCharNoreadonly struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharNoreadonlyNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsCharNoreadonlyModel string = "base_import.tests.models.char.noreadonly"

type BaseImportTestsModelsCharNoreadonlys []BaseImportTestsModelsCharNoreadonly

type BaseImportTestsModelsCharNoreadonlysNil []BaseImportTestsModelsCharNoreadonlyNil

func (s *BaseImportTestsModelsCharNoreadonly) NilableType_() interface{} {
	return &BaseImportTestsModelsCharNoreadonlyNil{}
}

func (ns *BaseImportTestsModelsCharNoreadonlyNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharNoreadonly{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsCharNoreadonlys) NilableType_() interface{} {
	return &BaseImportTestsModelsCharNoreadonlysNil{}
}

func (ns *BaseImportTestsModelsCharNoreadonlysNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharNoreadonlys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsCharNoreadonly))
	}
	return s
}
