package types

import (
	"time"
)

type BaseImportTestsModelsO2m struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       []int64   `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsO2mNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsO2mModel string = "base_import.tests.models.o2m"

type BaseImportTestsModelsO2ms []BaseImportTestsModelsO2m

type BaseImportTestsModelsO2msNil []BaseImportTestsModelsO2mNil

func (s *BaseImportTestsModelsO2m) NilableType_() interface{} {
	return &BaseImportTestsModelsO2mNil{}
}

func (ns *BaseImportTestsModelsO2mNil) Type_() interface{} {
	s := &BaseImportTestsModelsO2m{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsO2ms) NilableType_() interface{} {
	return &BaseImportTestsModelsO2msNil{}
}

func (ns *BaseImportTestsModelsO2msNil) Type_() interface{} {
	s := &BaseImportTestsModelsO2ms{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsO2m))
	}
	return s
}
