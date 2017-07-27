package types

import (
	"time"
)

type ProjectTaskType struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	Description    string    `xmlrpc:"description"`
	DisplayName    string    `xmlrpc:"display_name"`
	Fold           bool      `xmlrpc:"fold"`
	Id             int64     `xmlrpc:"id"`
	LegendBlocked  string    `xmlrpc:"legend_blocked"`
	LegendDone     string    `xmlrpc:"legend_done"`
	LegendNormal   string    `xmlrpc:"legend_normal"`
	LegendPriority string    `xmlrpc:"legend_priority"`
	MailTemplateId Many2One  `xmlrpc:"mail_template_id"`
	Name           string    `xmlrpc:"name"`
	ProjectIds     []int64   `xmlrpc:"project_ids"`
	Sequence       int64     `xmlrpc:"sequence"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type ProjectTaskTypeNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	Description    interface{} `xmlrpc:"description"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Fold           bool        `xmlrpc:"fold"`
	Id             interface{} `xmlrpc:"id"`
	LegendBlocked  interface{} `xmlrpc:"legend_blocked"`
	LegendDone     interface{} `xmlrpc:"legend_done"`
	LegendNormal   interface{} `xmlrpc:"legend_normal"`
	LegendPriority interface{} `xmlrpc:"legend_priority"`
	MailTemplateId interface{} `xmlrpc:"mail_template_id"`
	Name           interface{} `xmlrpc:"name"`
	ProjectIds     interface{} `xmlrpc:"project_ids"`
	Sequence       interface{} `xmlrpc:"sequence"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var ProjectTaskTypeModel string = "project.task.type"

type ProjectTaskTypes []ProjectTaskType

type ProjectTaskTypesNil []ProjectTaskTypeNil

func (s *ProjectTaskType) NilableType_() interface{} {
	return &ProjectTaskTypeNil{}
}

func (ns *ProjectTaskTypeNil) Type_() interface{} {
	s := &ProjectTaskType{}
	return load(ns, s)
}

func (s *ProjectTaskTypes) NilableType_() interface{} {
	return &ProjectTaskTypesNil{}
}

func (ns *ProjectTaskTypesNil) Type_() interface{} {
	s := &ProjectTaskTypes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProjectTaskType))
	}
	return s
}
