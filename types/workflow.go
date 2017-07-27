package types

import (
	"time"
)

type Workflow struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Activities  []int64   `xmlrpc:"activities"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	OnCreate    bool      `xmlrpc:"on_create"`
	Osv         string    `xmlrpc:"osv"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type WorkflowNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Activities  interface{} `xmlrpc:"activities"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	OnCreate    bool        `xmlrpc:"on_create"`
	Osv         interface{} `xmlrpc:"osv"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var WorkflowModel string = "workflow"

type Workflows []Workflow

type WorkflowsNil []WorkflowNil

func (s *Workflow) NilableType_() interface{} {
	return &WorkflowNil{}
}

func (ns *WorkflowNil) Type_() interface{} {
	s := &Workflow{}
	return load(ns, s)
}

func (s *Workflows) NilableType_() interface{} {
	return &WorkflowsNil{}
}

func (ns *WorkflowsNil) Type_() interface{} {
	s := &Workflows{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*Workflow))
	}
	return s
}
