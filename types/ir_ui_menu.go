package types

import (
	"time"
)

type IrUiMenu struct {
	Action       string    `xmlrpc:"action"`
	Active       bool      `xmlrpc:"active"`
	ChildId      []int64   `xmlrpc:"child_id"`
	CompleteName string    `xmlrpc:"complete_name"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	GroupsId     []int64   `xmlrpc:"groups_id"`
	Id           int64     `xmlrpc:"id"`
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	Name         string    `xmlrpc:"name"`
	ParentId     Many2One  `xmlrpc:"parent_id"`
	ParentLeft   int64     `xmlrpc:"parent_left"`
	ParentRight  int64     `xmlrpc:"parent_right"`
	Sequence     int64     `xmlrpc:"sequence"`
	WebIcon      string    `xmlrpc:"web_icon"`
	WebIconData  string    `xmlrpc:"web_icon_data"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type IrUiMenuNil struct {
	Action       interface{} `xmlrpc:"action"`
	Active       bool        `xmlrpc:"active"`
	ChildId      interface{} `xmlrpc:"child_id"`
	CompleteName interface{} `xmlrpc:"complete_name"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	GroupsId     interface{} `xmlrpc:"groups_id"`
	Id           interface{} `xmlrpc:"id"`
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	Name         interface{} `xmlrpc:"name"`
	ParentId     interface{} `xmlrpc:"parent_id"`
	ParentLeft   interface{} `xmlrpc:"parent_left"`
	ParentRight  interface{} `xmlrpc:"parent_right"`
	Sequence     interface{} `xmlrpc:"sequence"`
	WebIcon      interface{} `xmlrpc:"web_icon"`
	WebIconData  interface{} `xmlrpc:"web_icon_data"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var IrUiMenuModel string = "ir.ui.menu"

type IrUiMenus []IrUiMenu

type IrUiMenusNil []IrUiMenuNil

func (s *IrUiMenu) NilableType_() interface{} {
	return &IrUiMenuNil{}
}

func (ns *IrUiMenuNil) Type_() interface{} {
	s := &IrUiMenu{}
	return load(ns, s)
}

func (s *IrUiMenus) NilableType_() interface{} {
	return &IrUiMenusNil{}
}

func (ns *IrUiMenusNil) Type_() interface{} {
	s := &IrUiMenus{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrUiMenu))
	}
	return s
}
