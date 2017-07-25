package types

import (
	"time"
)

type BaseImportTestsModelsCharStates struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseImportTestsModelsCharStatesNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseImportTestsModelsCharStatesModel string = "base_import.tests.models.char.states"

type BaseImportTestsModelsCharStatess []BaseImportTestsModelsCharStates

type BaseImportTestsModelsCharStatessNil []BaseImportTestsModelsCharStatesNil

func (s *BaseImportTestsModelsCharStates) NilableType_() interface{} {
	return &BaseImportTestsModelsCharStatesNil{}
}

func (ns *BaseImportTestsModelsCharStatesNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharStates{}
	return load(ns, s)
}

func (s *BaseImportTestsModelsCharStatess) NilableType_() interface{} {
	return &BaseImportTestsModelsCharStatessNil{}
}

func (ns *BaseImportTestsModelsCharStatessNil) Type_() interface{} {
	s := &BaseImportTestsModelsCharStatess{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseImportTestsModelsCharStates))
	}
	return s
}
