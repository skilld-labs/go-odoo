package types

import (
	"time"
)

type BaseActionRuleLineTest struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LeadId      Many2One  `xmlrpc:"lead_id"`
	Name        string    `xmlrpc:"name"`
	UserId      Many2One  `xmlrpc:"user_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BaseActionRuleLineTestNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LeadId      interface{} `xmlrpc:"lead_id"`
	Name        interface{} `xmlrpc:"name"`
	UserId      interface{} `xmlrpc:"user_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BaseActionRuleLineTestModel string = "base.action.rule.line.test"

type BaseActionRuleLineTests []BaseActionRuleLineTest

type BaseActionRuleLineTestsNil []BaseActionRuleLineTestNil

func (s *BaseActionRuleLineTest) NilableType_() interface{} {
	return &BaseActionRuleLineTestNil{}
}

func (ns *BaseActionRuleLineTestNil) Type_() interface{} {
	s := &BaseActionRuleLineTest{}
	return load(ns, s)
}

func (s *BaseActionRuleLineTests) NilableType_() interface{} {
	return &BaseActionRuleLineTestsNil{}
}

func (ns *BaseActionRuleLineTestsNil) Type_() interface{} {
	s := &BaseActionRuleLineTests{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseActionRuleLineTest))
	}
	return s
}
