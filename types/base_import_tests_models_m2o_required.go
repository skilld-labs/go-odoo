package types

import (
	"time"
)

type BaseImportTestsModelsM2oRequired struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       Many2One  `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsM2oRequiredNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsM2oRequiredModel string = "base_import.tests.models.m2o.required"

type BaseImportTestsModelsM2oRequireds []BaseImportTestsModelsM2oRequired

type BaseImportTestsModelsM2oRequiredsNil []BaseImportTestsModelsM2oRequiredNil

func (s *BaseImportTestsModelsM2oRequired) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oRequiredNil{}
}

func (ns *BaseImportTestsModelsM2oRequiredNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2oRequired{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsM2oRequireds) NilableType_() interface{} {
	return &BaseImportTestsModelsM2oRequiredsNil{}
}

func (ns *BaseImportTestsModelsM2oRequiredsNil) Type_() interface{} {
	s := &BaseImportTestsModelsM2oRequireds{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsM2oRequired))
	}
	return s
}
