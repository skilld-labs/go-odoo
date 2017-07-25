package types

import (
	"time"
)

type IrActionsActWindowView struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ActWindowId Many2One  `xmlrpc:"act_window_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Multi       bool      `xmlrpc:"multi"`
	Sequence    int64     `xmlrpc:"sequence"`
	ViewId      Many2One  `xmlrpc:"view_id"`
	ViewMode    string    `xmlrpc:"view_mode"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrActionsActWindowViewNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ActWindowId interface{} `xmlrpc:"act_window_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Multi       bool        `xmlrpc:"multi"`
	Sequence    interface{} `xmlrpc:"sequence"`
	ViewId      interface{} `xmlrpc:"view_id"`
	ViewMode    interface{} `xmlrpc:"view_mode"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrActionsActWindowViewModel string = "ir.actions.act_window.view"

type IrActionsActWindowViews []IrActionsActWindowView

type IrActionsActWindowViewsNil []IrActionsActWindowViewNil

func (s *IrActionsActWindowView) NilableType_() interface{} {
	return &IrActionsActWindowViewNil{}
}

func (ns *IrActionsActWindowViewNil) Type_() interface{} {
	s := &IrActionsActWindowView{}
	return load(ns, s)
}

func (s *IrActionsActWindowViews) NilableType_() interface{} {
	return &IrActionsActWindowViewsNil{}
}

func (ns *IrActionsActWindowViewsNil) Type_() interface{} {
	s := &IrActionsActWindowViews{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrActionsActWindowView))
	}
	return s
}
