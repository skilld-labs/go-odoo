package types

import (
	"time"
)

type WizardIrModelMenuCreate struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	MenuId      Many2One  `xmlrpc:"menu_id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type WizardIrModelMenuCreateNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	MenuId      interface{} `xmlrpc:"menu_id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var WizardIrModelMenuCreateModel string = "wizard.ir.model.menu.create"

type WizardIrModelMenuCreates []WizardIrModelMenuCreate

type WizardIrModelMenuCreatesNil []WizardIrModelMenuCreateNil

func (s *WizardIrModelMenuCreate) NilableType_() interface{} {
	return &WizardIrModelMenuCreateNil{}
}

func (ns *WizardIrModelMenuCreateNil) Type_() interface{} {
	s := &WizardIrModelMenuCreate{}
	return load(ns, s)
}

func (s *WizardIrModelMenuCreates) NilableType_() interface{} {
	return &WizardIrModelMenuCreatesNil{}
}

func (ns *WizardIrModelMenuCreatesNil) Type_() interface{} {
	s := &WizardIrModelMenuCreates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WizardIrModelMenuCreate))
	}
	return s
}
