package types

import (
	"time"
)

type BaseActionRuleLeadTest struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Active         bool      `xmlrpc:"active"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	Customer       bool      `xmlrpc:"customer"`
	DateActionLast time.Time `xmlrpc:"date_action_last"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	LineIds        []int64   `xmlrpc:"line_ids"`
	Name           string    `xmlrpc:"name"`
	PartnerId      Many2One  `xmlrpc:"partner_id"`
	State          string    `xmlrpc:"state"`
	UserId         Many2One  `xmlrpc:"user_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type BaseActionRuleLeadTestNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Active         bool        `xmlrpc:"active"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	Customer       bool        `xmlrpc:"customer"`
	DateActionLast interface{} `xmlrpc:"date_action_last"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	LineIds        interface{} `xmlrpc:"line_ids"`
	Name           interface{} `xmlrpc:"name"`
	PartnerId      interface{} `xmlrpc:"partner_id"`
	State          interface{} `xmlrpc:"state"`
	UserId         interface{} `xmlrpc:"user_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var BaseActionRuleLeadTestModel string = "base.action.rule.lead.test"

type BaseActionRuleLeadTests []BaseActionRuleLeadTest

type BaseActionRuleLeadTestsNil []BaseActionRuleLeadTestNil

func (s *BaseActionRuleLeadTest) NilableType_() interface{} {
	return &BaseActionRuleLeadTestNil{}
}

func (ns *BaseActionRuleLeadTestNil) Type_() interface{} {
	s := &BaseActionRuleLeadTest{}
	return load(ns, s)
}

func (s *BaseActionRuleLeadTests) NilableType_() interface{} {
	return &BaseActionRuleLeadTestsNil{}
}

func (ns *BaseActionRuleLeadTestsNil) Type_() interface{} {
	s := &BaseActionRuleLeadTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseActionRuleLeadTest))
	}
	return s
}
