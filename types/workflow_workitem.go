package types

import (
	"time"
)

type WorkflowWorkitem struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ActId       Many2One  `xmlrpc:"act_id"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	InstId      Many2One  `xmlrpc:"inst_id"`
	State       string    `xmlrpc:"state"`
	SubflowId   Many2One  `xmlrpc:"subflow_id"`
	WkfId       Many2One  `xmlrpc:"wkf_id"`
}

type WorkflowWorkitemNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ActId       interface{} `xmlrpc:"act_id"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	InstId      interface{} `xmlrpc:"inst_id"`
	State       interface{} `xmlrpc:"state"`
	SubflowId   interface{} `xmlrpc:"subflow_id"`
	WkfId       interface{} `xmlrpc:"wkf_id"`
}

var WorkflowWorkitemModel string = "workflow.workitem"

type WorkflowWorkitems []WorkflowWorkitem

type WorkflowWorkitemsNil []WorkflowWorkitemNil

func (s *WorkflowWorkitem) NilableType_() interface{} {
	return &WorkflowWorkitemNil{}
}

func (ns *WorkflowWorkitemNil) Type_() interface{} {
	s := &WorkflowWorkitem{}
	return load(ns, s)
}

func (s *WorkflowWorkitems) NilableType_() interface{} {
	return &WorkflowWorkitemsNil{}
}

func (ns *WorkflowWorkitemsNil) Type_() interface{} {
	s := &WorkflowWorkitems{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*WorkflowWorkitem))
	}
	return s
}
