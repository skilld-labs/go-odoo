package types

import (
	"time"
)

type BaseImportTestsModelsM2oRequiredRelated struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       int64     `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsM2oRequiredRelatedNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsM2oRequiredRelatedModel string = "base_import.tests.models.m2o.required.related"

type BaseImportTestsModelsM2oRequiredRelateds []BaseImportTestsModelsM2oRequiredRelated

type BaseImportTestsModelsM2oRequiredRelatedsNil []BaseImportTestsModelsM2oRequiredRelatedNil

func (s *BaseImportTestsModelsM2oRequiredRelated) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oRequiredRelatedNil{}
}

func (ns *BaseImportTestsModelsM2oRequiredRelatedNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2oRequiredRelated{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsM2oRequiredRelateds) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oRequiredRelatedsNil{}
}

func (ns *BaseImportTestsModelsM2oRequiredRelatedsNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2oRequiredRelateds{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsM2oRequiredRelated))
	}
	return s
}
