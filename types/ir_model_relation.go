package types

import (
	"time"
)

type IrModelRelation struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateInit    time.Time `xmlrpc:"date_init"`
	DateUpdate  time.Time `xmlrpc:"date_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Model       Many2One  `xmlrpc:"model"`
	Module      Many2One  `xmlrpc:"module"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrModelRelationNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateInit    interface{} `xmlrpc:"date_init"`
	DateUpdate  interface{} `xmlrpc:"date_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Model       interface{} `xmlrpc:"model"`
	Module      interface{} `xmlrpc:"module"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrModelRelationModel string = "ir.model.relation"

type IrModelRelations []IrModelRelation

type IrModelRelationsNil []IrModelRelationNil

func (s *IrModelRelation) NilableType_() interface{} {
	return &IrModelRelationNil{}
}

func (ns *IrModelRelationNil) Type_() interface{} {
	s := &IrModelRelation{}
	return load(ns, s)
}

func (s *IrModelRelations) NilableType_() interface{} {
	return &IrModelRelationsNil{}
}

func (ns *IrModelRelationsNil) Type_() interface{} {
	s := &IrModelRelations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModelRelation))
	}
	return s
}
