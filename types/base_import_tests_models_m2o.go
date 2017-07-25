package types

import (
	"time"
)

type BaseImportTestsModelsM2o struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       Many2One  `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsM2oNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsM2oModel string = "base_import.tests.models.m2o"

type BaseImportTestsModelsM2os []BaseImportTestsModelsM2o

type BaseImportTestsModelsM2osNil []BaseImportTestsModelsM2oNil

func (s *BaseImportTestsModelsM2o) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oNil{}
}

func (ns *BaseImportTestsModelsM2oNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2o{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsM2os) NilableType_() interface{} {
	return &BaseImportTestsModelsM2osNil{}
}

func (ns *BaseImportTestsModelsM2osNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2os{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsM2o))
	}
	return s
}
