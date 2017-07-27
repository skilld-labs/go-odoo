package types

import (
	"time"
)

type IrModelConstraint struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DateInit    time.Time `xmlrpc:"date_init"`
	DateUpdate  time.Time `xmlrpc:"date_update"`
	Definition  string    `xmlrpc:"definition"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Model       Many2One  `xmlrpc:"model"`
	Module      Many2One  `xmlrpc:"module"`
	Name        string    `xmlrpc:"name"`
	Type        string    `xmlrpc:"type"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrModelConstraintNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DateInit    interface{} `xmlrpc:"date_init"`
	DateUpdate  interface{} `xmlrpc:"date_update"`
	Definition  interface{} `xmlrpc:"definition"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Model       interface{} `xmlrpc:"model"`
	Module      interface{} `xmlrpc:"module"`
	Name        interface{} `xmlrpc:"name"`
	Type        interface{} `xmlrpc:"type"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrModelConstraintModel string = "ir.model.constraint"

type IrModelConstraints []IrModelConstraint

type IrModelConstraintsNil []IrModelConstraintNil

func (s *IrModelConstraint) NilableType_() interface{} {
	return &IrModelConstraintNil{}
}

func (ns *IrModelConstraintNil) Type_() interface{} {
	s := &IrModelConstraint{}
	return load(ns, s)
}

func (s *IrModelConstraints) NilableType_() interface{} {
	return &IrModelConstraintsNil{}
}

func (ns *IrModelConstraintsNil) Type_() interface{} {
	s := &IrModelConstraints{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModelConstraint))
	}
	return s
}
