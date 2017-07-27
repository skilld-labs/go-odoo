package types

import (
	"time"
)

type BaseImportTestsModelsO2mChild struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	ParentId    Many2One  `xmlrpc:"parent_id"`
	Value       int64     `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsO2mChildNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	ParentId    interface{} `xmlrpc:"parent_id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsO2mChildModel string = "base_import.tests.models.o2m.child"

type BaseImportTestsModelsO2mChilds []BaseImportTestsModelsO2mChild

type BaseImportTestsModelsO2mChildsNil []BaseImportTestsModelsO2mChildNil

func (s *BaseImportTestsModelsO2mChild) NilableType_() interface{} {
	return &BaseImportTestsModelsO2mChildNil{}
}

func (ns *BaseImportTestsModelsO2mChildNil) Type_() interface{} {
	s := &BaseImportTestsModelsO2mChild{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsO2mChilds) NilableType_() interface{} {
	return &BaseImportTestsModelsO2mChildsNil{}
}

func (ns *BaseImportTestsModelsO2mChildsNil) Type_() interface{} {
	s := &BaseImportTestsModelsO2mChilds{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsO2mChild))
	}
	return s
}
