package types

import (
	"time"
)

type WorkflowTransition struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	ActFrom       Many2One  `xmlrpc:"act_from"`
	ActTo         Many2One  `xmlrpc:"act_to"`
	Condition     string    `xmlrpc:"condition"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	GroupId       Many2One  `xmlrpc:"group_id"`
	Id            int64     `xmlrpc:"id"`
	Sequence      int64     `xmlrpc:"sequence"`
	Signal        string    `xmlrpc:"signal"`
	TriggerExprId string    `xmlrpc:"trigger_expr_id"`
	TriggerModel  string    `xmlrpc:"trigger_model"`
	WkfId         Many2One  `xmlrpc:"wkf_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type WorkflowTransitionNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	ActFrom       interface{} `xmlrpc:"act_from"`
	ActTo         interface{} `xmlrpc:"act_to"`
	Condition     interface{} `xmlrpc:"condition"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	GroupId       interface{} `xmlrpc:"group_id"`
	Id            interface{} `xmlrpc:"id"`
	Sequence      interface{} `xmlrpc:"sequence"`
	Signal        interface{} `xmlrpc:"signal"`
	TriggerExprId interface{} `xmlrpc:"trigger_expr_id"`
	TriggerModel  interface{} `xmlrpc:"trigger_model"`
	WkfId         interface{} `xmlrpc:"wkf_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var WorkflowTransitionModel string = "workflow.transition"

type WorkflowTransitions []WorkflowTransition

type WorkflowTransitionsNil []WorkflowTransitionNil

func (s *WorkflowTransition) NilableType_() interface{} {
	return &WorkflowTransitionNil{}
}

func (ns *WorkflowTransitionNil) Type_() interface{} {
	s := &WorkflowTransition{}
	return load(ns, s)
}

func (s *WorkflowTransitions) NilableType_() interface{} {
	return &WorkflowTransitionsNil{}
}

func (ns *WorkflowTransitionsNil) Type_() interface{} {
	s := &WorkflowTransitions{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WorkflowTransition))
	}
	return s
}
