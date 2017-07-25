package types

import (
	"time"
)

type BaseActionRule struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	ActFollowers      []int64   `xmlrpc:"act_followers"`
	ActUserId         Many2One  `xmlrpc:"act_user_id"`
	Active            bool      `xmlrpc:"active"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	DisplayName       string    `xmlrpc:"display_name"`
	FilterDomain      string    `xmlrpc:"filter_domain"`
	FilterId          Many2One  `xmlrpc:"filter_id"`
	FilterPreDomain   string    `xmlrpc:"filter_pre_domain"`
	FilterPreId       Many2One  `xmlrpc:"filter_pre_id"`
	Id                int64     `xmlrpc:"id"`
	Kind              string    `xmlrpc:"kind"`
	LastRun           time.Time `xmlrpc:"last_run"`
	Model             string    `xmlrpc:"model"`
	ModelId           Many2One  `xmlrpc:"model_id"`
	Name              string    `xmlrpc:"name"`
	OnChangeFields    string    `xmlrpc:"on_change_fields"`
	Sequence          int64     `xmlrpc:"sequence"`
	ServerActionIds   []int64   `xmlrpc:"server_action_ids"`
	TrgDateCalendarId Many2One  `xmlrpc:"trg_date_calendar_id"`
	TrgDateId         Many2One  `xmlrpc:"trg_date_id"`
	TrgDateRange      int64     `xmlrpc:"trg_date_range"`
	TrgDateRangeType  string    `xmlrpc:"trg_date_range_type"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type BaseActionRuleNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	ActFollowers      interface{} `xmlrpc:"act_followers"`
	ActUserId         interface{} `xmlrpc:"act_user_id"`
	Active            bool        `xmlrpc:"active"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	FilterDomain      interface{} `xmlrpc:"filter_domain"`
	FilterId          interface{} `xmlrpc:"filter_id"`
	FilterPreDomain   interface{} `xmlrpc:"filter_pre_domain"`
	FilterPreId       interface{} `xmlrpc:"filter_pre_id"`
	Id                interface{} `xmlrpc:"id"`
	Kind              interface{} `xmlrpc:"kind"`
	LastRun           interface{} `xmlrpc:"last_run"`
	Model             interface{} `xmlrpc:"model"`
	ModelId           interface{} `xmlrpc:"model_id"`
	Name              interface{} `xmlrpc:"name"`
	OnChangeFields    interface{} `xmlrpc:"on_change_fields"`
	Sequence          interface{} `xmlrpc:"sequence"`
	ServerActionIds   interface{} `xmlrpc:"server_action_ids"`
	TrgDateCalendarId interface{} `xmlrpc:"trg_date_calendar_id"`
	TrgDateId         interface{} `xmlrpc:"trg_date_id"`
	TrgDateRange      interface{} `xmlrpc:"trg_date_range"`
	TrgDateRangeType  interface{} `xmlrpc:"trg_date_range_type"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var BaseActionRuleModel string = "base.action.rule"

type BaseActionRules []BaseActionRule

type BaseActionRulesNil []BaseActionRuleNil

func (s *BaseActionRule) NilableType_() interface{} {
	return &BaseActionRuleNil{}
}

func (ns *BaseActionRuleNil) Type_() interface{} {
	s := &BaseActionRule{}
	return load(ns, s)
}

func (s *BaseActionRules) NilableType_() interface{} {
	return &BaseActionRulesNil{}
}

func (ns *BaseActionRulesNil) Type_() interface{} {
	s := &BaseActionRules{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseActionRule))
	}
	return s
}
