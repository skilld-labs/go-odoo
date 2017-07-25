package types

import (
	"time"
)

type IrRule struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Domain      string    `xmlrpc:"domain"`
	DomainForce string    `xmlrpc:"domain_force"`
	Global      bool      `xmlrpc:"global"`
	Groups      []int64   `xmlrpc:"groups"`
	Id          int64     `xmlrpc:"id"`
	ModelId     Many2One  `xmlrpc:"model_id"`
	Name        string    `xmlrpc:"name"`
	PermCreate  bool      `xmlrpc:"perm_create"`
	PermRead    bool      `xmlrpc:"perm_read"`
	PermUnlink  bool      `xmlrpc:"perm_unlink"`
	PermWrite   bool      `xmlrpc:"perm_write"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrRuleNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Domain      interface{} `xmlrpc:"domain"`
	DomainForce interface{} `xmlrpc:"domain_force"`
	Global      bool        `xmlrpc:"global"`
	Groups      interface{} `xmlrpc:"groups"`
	Id          interface{} `xmlrpc:"id"`
	ModelId     interface{} `xmlrpc:"model_id"`
	Name        interface{} `xmlrpc:"name"`
	PermCreate  bool        `xmlrpc:"perm_create"`
	PermRead    bool        `xmlrpc:"perm_read"`
	PermUnlink  bool        `xmlrpc:"perm_unlink"`
	PermWrite   bool        `xmlrpc:"perm_write"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrRuleModel string = "ir.rule"

type IrRules []IrRule

type IrRulesNil []IrRuleNil

func (s *IrRule) NilableType_() interface{} {
	return &IrRuleNil{}
}

func (ns *IrRuleNil) Type_() interface{} {
	s := &IrRule{}
	return load(ns, s)
}

func (s *IrRules) NilableType_() interface{} {
	return &IrRulesNil{}
}

func (ns *IrRulesNil) Type_() interface{} {
	s := &IrRules{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrRule))
	}
	return s
}
