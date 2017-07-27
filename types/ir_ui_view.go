package types

import (
	"time"
)

type IrUiView struct {
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	Active             bool      `xmlrpc:"active"`
	Arch               string    `xmlrpc:"arch"`
	ArchBase           string    `xmlrpc:"arch_base"`
	ArchDb             string    `xmlrpc:"arch_db"`
	ArchFs             string    `xmlrpc:"arch_fs"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	DisplayName        string    `xmlrpc:"display_name"`
	FieldParent        string    `xmlrpc:"field_parent"`
	GroupsId           []int64   `xmlrpc:"groups_id"`
	Id                 int64     `xmlrpc:"id"`
	InheritChildrenIds []int64   `xmlrpc:"inherit_children_ids"`
	InheritId          Many2One  `xmlrpc:"inherit_id"`
	Key                string    `xmlrpc:"key"`
	Mode               string    `xmlrpc:"mode"`
	Model              string    `xmlrpc:"model"`
	ModelDataId        Many2One  `xmlrpc:"model_data_id"`
	ModelIds           []int64   `xmlrpc:"model_ids"`
	Name               string    `xmlrpc:"name"`
	Priority           int64     `xmlrpc:"priority"`
	Type               string    `xmlrpc:"type"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
	XmlId              string    `xmlrpc:"xml_id"`
}

type IrUiViewNil struct {
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	Active             bool        `xmlrpc:"active"`
	Arch               interface{} `xmlrpc:"arch"`
	ArchBase           interface{} `xmlrpc:"arch_base"`
	ArchDb             interface{} `xmlrpc:"arch_db"`
	ArchFs             interface{} `xmlrpc:"arch_fs"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	FieldParent        interface{} `xmlrpc:"field_parent"`
	GroupsId           interface{} `xmlrpc:"groups_id"`
	Id                 interface{} `xmlrpc:"id"`
	InheritChildrenIds interface{} `xmlrpc:"inherit_children_ids"`
	InheritId          interface{} `xmlrpc:"inherit_id"`
	Key                interface{} `xmlrpc:"key"`
	Mode               interface{} `xmlrpc:"mode"`
	Model              interface{} `xmlrpc:"model"`
	ModelDataId        interface{} `xmlrpc:"model_data_id"`
	ModelIds           interface{} `xmlrpc:"model_ids"`
	Name               interface{} `xmlrpc:"name"`
	Priority           interface{} `xmlrpc:"priority"`
	Type               interface{} `xmlrpc:"type"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
	XmlId              interface{} `xmlrpc:"xml_id"`
}

var IrUiViewModel string = "ir.ui.view"

type IrUiViews []IrUiView

type IrUiViewsNil []IrUiViewNil

func (s *IrUiView) NilableType_() interface{} {
	return &IrUiViewNil{}
}

func (ns *IrUiViewNil) Type_() interface{} {
	s := &IrUiView{}
	return load(ns, s)
}

func (s *IrUiViews) NilableType_() interface{} {
	return &IrUiViewsNil{}
}

func (ns *IrUiViewsNil) Type_() interface{} {
	s := &IrUiViews{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrUiView))
	}
	return s
}
