package types

import (
	"time"
)

type BaseImportTestsModelsM2oRelated struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       int64     `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsM2oRelatedNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsM2oRelatedModel string = "base_import.tests.models.m2o.related"

type BaseImportTestsModelsM2oRelateds []BaseImportTestsModelsM2oRelated

type BaseImportTestsModelsM2oRelatedsNil []BaseImportTestsModelsM2oRelatedNil

func (s *BaseImportTestsModelsM2oRelated) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oRelatedNil{}
}

func (ns *BaseImportTestsModelsM2oRelatedNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2oRelated{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsM2oRelateds) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oRelatedsNil{}
}

func (ns *BaseImportTestsModelsM2oRelatedsNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2oRelateds{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsM2oRelated))
	}
	return s
}
