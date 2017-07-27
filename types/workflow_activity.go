package types

import (
	"time"
)

type WorkflowActivity struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Action         string    `xmlrpc:"action"`
	ActionId       Many2One  `xmlrpc:"action_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	FlowStart      bool      `xmlrpc:"flow_start"`
	FlowStop       bool      `xmlrpc:"flow_stop"`
	Id             int64     `xmlrpc:"id"`
	InTransitions  []int64   `xmlrpc:"in_transitions"`
	JoinMode       string    `xmlrpc:"join_mode"`
	Kind           string    `xmlrpc:"kind"`
	Name           string    `xmlrpc:"name"`
	OutTransitions []int64   `xmlrpc:"out_transitions"`
	SignalSend     string    `xmlrpc:"signal_send"`
	SplitMode      string    `xmlrpc:"split_mode"`
	SubflowId      Many2One  `xmlrpc:"subflow_id"`
	WkfId          Many2One  `xmlrpc:"wkf_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type WorkflowActivityNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Action         interface{} `xmlrpc:"action"`
	ActionId       interface{} `xmlrpc:"action_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	FlowStart      bool        `xmlrpc:"flow_start"`
	FlowStop       bool        `xmlrpc:"flow_stop"`
	Id             interface{} `xmlrpc:"id"`
	InTransitions  interface{} `xmlrpc:"in_transitions"`
	JoinMode       interface{} `xmlrpc:"join_mode"`
	Kind           interface{} `xmlrpc:"kind"`
	Name           interface{} `xmlrpc:"name"`
	OutTransitions interface{} `xmlrpc:"out_transitions"`
	SignalSend     interface{} `xmlrpc:"signal_send"`
	SplitMode      interface{} `xmlrpc:"split_mode"`
	SubflowId      interface{} `xmlrpc:"subflow_id"`
	WkfId          interface{} `xmlrpc:"wkf_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var WorkflowActivityModel string = "workflow.activity"

type WorkflowActivitys []WorkflowActivity

type WorkflowActivitysNil []WorkflowActivityNil

func (s *WorkflowActivity) NilableType_() interface{} {
	return &WorkflowActivityNil{}
}

func (ns *WorkflowActivityNil) Type_() interface{} {
	s := &WorkflowActivity{}
	return load(ns, s)
}

func (s *WorkflowActivitys) NilableType_() interface{} {
	return &WorkflowActivitysNil{}
}

func (ns *WorkflowActivitysNil) Type_() interface{} {
	s := &WorkflowActivitys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WorkflowActivity))
	}
	return s
}
