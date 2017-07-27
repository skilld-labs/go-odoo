package types

import (
	"time"
)

type IrActionsActWindow struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	AutoSearch   bool      `xmlrpc:"auto_search"`
	Context      string    `xmlrpc:"context"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	Domain       string    `xmlrpc:"domain"`
	Filter       bool      `xmlrpc:"filter"`
	GroupsId     []int64   `xmlrpc:"groups_id"`
	Help         string    `xmlrpc:"help"`
	Id           int64     `xmlrpc:"id"`
	Limit        int64     `xmlrpc:"limit"`
	Multi        bool      `xmlrpc:"multi"`
	Name         string    `xmlrpc:"name"`
	ResId        int64     `xmlrpc:"res_id"`
	ResModel     string    `xmlrpc:"res_model"`
	SearchView   string    `xmlrpc:"search_view"`
	SearchViewId Many2One  `xmlrpc:"search_view_id"`
	SrcModel     string    `xmlrpc:"src_model"`
	Target       string    `xmlrpc:"target"`
	Type         string    `xmlrpc:"type"`
	Usage        string    `xmlrpc:"usage"`
	ViewId       Many2One  `xmlrpc:"view_id"`
	ViewIds      []int64   `xmlrpc:"view_ids"`
	ViewMode     string    `xmlrpc:"view_mode"`
	ViewType     string    `xmlrpc:"view_type"`
	Views        string    `xmlrpc:"views"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
	XmlId        string    `xmlrpc:"xml_id"`
}

type IrActionsActWindowNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	AutoSearch   bool        `xmlrpc:"auto_search"`
	Context      interface{} `xmlrpc:"context"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Domain       interface{} `xmlrpc:"domain"`
	Filter       bool        `xmlrpc:"filter"`
	GroupsId     interface{} `xmlrpc:"groups_id"`
	Help         interface{} `xmlrpc:"help"`
	Id           interface{} `xmlrpc:"id"`
	Limit        interface{} `xmlrpc:"limit"`
	Multi        bool        `xmlrpc:"multi"`
	Name         interface{} `xmlrpc:"name"`
	ResId        interface{} `xmlrpc:"res_id"`
	ResModel     interface{} `xmlrpc:"res_model"`
	SearchView   interface{} `xmlrpc:"search_view"`
	SearchViewId interface{} `xmlrpc:"search_view_id"`
	SrcModel     interface{} `xmlrpc:"src_model"`
	Target       interface{} `xmlrpc:"target"`
	Type         interface{} `xmlrpc:"type"`
	Usage        interface{} `xmlrpc:"usage"`
	ViewId       interface{} `xmlrpc:"view_id"`
	ViewIds      interface{} `xmlrpc:"view_ids"`
	ViewMode     interface{} `xmlrpc:"view_mode"`
	ViewType     interface{} `xmlrpc:"view_type"`
	Views        interface{} `xmlrpc:"views"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
	XmlId        interface{} `xmlrpc:"xml_id"`
}

var IrActionsActWindowModel string = "ir.actions.act_window"

type IrActionsActWindows []IrActionsActWindow

type IrActionsActWindowsNil []IrActionsActWindowNil

func (s *IrActionsActWindow) NilableType_() interface{} {
	return &IrActionsActWindowNil{}
}

func (ns *IrActionsActWindowNil) Type_() interface{} {
	s := &IrActionsActWindow{}
	return load(ns, s)
}

func (s *IrActionsActWindows) NilableType_() interface{} {
	return &IrActionsActWindowsNil{}
}

func (ns *IrActionsActWindowsNil) Type_() interface{} {
	s := &IrActionsActWindows{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsActWindow))
	}
	return s
}
