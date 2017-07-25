package types

import (
	"time"
)

type WorkflowTriggers struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	InstanceId  Many2One  `xmlrpc:"instance_id"`
	Model       string    `xmlrpc:"model"`
	ResId       int64     `xmlrpc:"res_id"`
	WorkitemId  Many2One  `xmlrpc:"workitem_id"`
}

type WorkflowTriggersNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	InstanceId  interface{} `xmlrpc:"instance_id"`
	Model       interface{} `xmlrpc:"model"`
	ResId       interface{} `xmlrpc:"res_id"`
	WorkitemId  interface{} `xmlrpc:"workitem_id"`
}

var WorkflowTriggersModel string = "workflow.triggers"

type WorkflowTriggerss []WorkflowTriggers

type WorkflowTriggerssNil []WorkflowTriggersNil

func (s *WorkflowTriggers) NilableType_() interface{} {
	return &WorkflowTriggersNil{}
}

func (ns *WorkflowTriggersNil) Type_() interface{} {
	s := &WorkflowTriggers{}
	return load(ns, s)
}

func (s *WorkflowTriggerss) NilableType_() interface{} {
	return &WorkflowTriggerssNil{}
}

func (ns *WorkflowTriggerssNil) Type_() interface{} {
	s := &WorkflowTriggerss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WorkflowTriggers))
	}
	return s
}
