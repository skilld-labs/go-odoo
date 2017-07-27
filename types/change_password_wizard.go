package types

import (
	"time"
)

type ChangePasswordWizard struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	UserIds     []int64   `xmlrpc:"user_ids"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ChangePasswordWizardNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	UserIds     interface{} `xmlrpc:"user_ids"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ChangePasswordWizardModel string = "change.password.wizard"

type ChangePasswordWizards []ChangePasswordWizard

type ChangePasswordWizardsNil []ChangePasswordWizardNil

func (s *ChangePasswordWizard) NilableType_() interface{} {
	return &ChangePasswordWizardNil{}
}

func (ns *ChangePasswordWizardNil) Type_() interface{} {
	s := &ChangePasswordWizard{}
	return load(ns, s)
}

func (s *ChangePasswordWizards) NilableType_() interface{} {
	return &ChangePasswordWizardsNil{}
}

func (ns *ChangePasswordWizardsNil) Type_() interface{} {
	s := &ChangePasswordWizards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ChangePasswordWizard))
	}
	return s
}
