package types

import (
	"time"
)

type ProjectTaskMergeWizard struct {
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateNewTask   bool      `xmlrpc:"create_new_task"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	TargetProjectId Many2One  `xmlrpc:"target_project_id"`
	TargetTaskId    Many2One  `xmlrpc:"target_task_id"`
	TargetTaskName  string    `xmlrpc:"target_task_name"`
	TaskIds         []int64   `xmlrpc:"task_ids"`
	UserId          Many2One  `xmlrpc:"user_id"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type ProjectTaskMergeWizardNil struct {
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateNewTask   bool        `xmlrpc:"create_new_task"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	TargetProjectId interface{} `xmlrpc:"target_project_id"`
	TargetTaskId    interface{} `xmlrpc:"target_task_id"`
	TargetTaskName  interface{} `xmlrpc:"target_task_name"`
	TaskIds         interface{} `xmlrpc:"task_ids"`
	UserId          interface{} `xmlrpc:"user_id"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var ProjectTaskMergeWizardModel string = "project.task.merge.wizard"

type ProjectTaskMergeWizards []ProjectTaskMergeWizard

type ProjectTaskMergeWizardsNil []ProjectTaskMergeWizardNil

func (s *ProjectTaskMergeWizard) NilableType_() interface{} {
	return &ProjectTaskMergeWizardNil{}
}

func (ns *ProjectTaskMergeWizardNil) Type_() interface{} {
	s := &ProjectTaskMergeWizard{}
	return load(ns, s)
}

func (s *ProjectTaskMergeWizards) NilableType_() interface{} {
	return &ProjectTaskMergeWizardsNil{}
}

func (ns *ProjectTaskMergeWizardsNil) Type_() interface{} {
	s := &ProjectTaskMergeWizards{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProjectTaskMergeWizard))
	}
	return s
}
