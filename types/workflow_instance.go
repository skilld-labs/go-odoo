package types

import (
	"time"
)

type WorkflowInstance struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	ResId         int64     `xmlrpc:"res_id"`
	ResType       string    `xmlrpc:"res_type"`
	State         string    `xmlrpc:"state"`
	TransitionIds []int64   `xmlrpc:"transition_ids"`
	Uid           int64     `xmlrpc:"uid"`
	WkfId         Many2One  `xmlrpc:"wkf_id"`
}

type WorkflowInstanceNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	ResId         interface{} `xmlrpc:"res_id"`
	ResType       interface{} `xmlrpc:"res_type"`
	State         interface{} `xmlrpc:"state"`
	TransitionIds interface{} `xmlrpc:"transition_ids"`
	Uid           interface{} `xmlrpc:"uid"`
	WkfId         interface{} `xmlrpc:"wkf_id"`
}

var WorkflowInstanceModel string = "workflow.instance"

type WorkflowInstances []WorkflowInstance

type WorkflowInstancesNil []WorkflowInstanceNil

func (s *WorkflowInstance) NilableType_() interface{} {
	return &WorkflowInstanceNil{}
}

func (ns *WorkflowInstanceNil) Type_() interface{} {
	s := &WorkflowInstance{}
	return load(ns, s)
}

func (s *WorkflowInstances) NilableType_() interface{} {
	return &WorkflowInstancesNil{}
}

func (ns *WorkflowInstancesNil) Type_() interface{} {
	s := &WorkflowInstances{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WorkflowInstance))
	}
	return s
}
