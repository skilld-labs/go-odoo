package types

import (
	"time"
)

type BaseImportTestsModelsPreview struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Othervalue  int64     `xmlrpc:"othervalue"`
	Somevalue   int64     `xmlrpc:"somevalue"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsPreviewNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Othervalue  interface{} `xmlrpc:"othervalue"`
	Somevalue   interface{} `xmlrpc:"somevalue"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsPreviewModel string = "base_import.tests.models.preview"

type BaseImportTestsModelsPreviews []BaseImportTestsModelsPreview

type BaseImportTestsModelsPreviewsNil []BaseImportTestsModelsPreviewNil

func (s *BaseImportTestsModelsPreview) NilableType_() interface{} {
	return &BaseImportTestsModelsPreviewNil{}
}

func (ns *BaseImportTestsModelsPreviewNil) Type_() interface{} {
	s := &BaseImportTestsModelsPreview{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsPreviews) NilableType_() interface{} {
	return &BaseImportTestsModelsPreviewsNil{}
}

func (ns *BaseImportTestsModelsPreviewsNil) Type_() interface{} {
	s := &BaseImportTestsModelsPreviews{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsPreview))
	}
	return s
}
