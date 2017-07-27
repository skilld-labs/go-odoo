package types

import (
	"time"
)

type IrUiViewCustom struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Arch        string    `xmlrpc:"arch"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	RefId       Many2One  `xmlrpc:"ref_id"`
	UserId      Many2One  `xmlrpc:"user_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrUiViewCustomNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Arch        interface{} `xmlrpc:"arch"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	RefId       interface{} `xmlrpc:"ref_id"`
	UserId      interface{} `xmlrpc:"user_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrUiViewCustomModel string = "ir.ui.view.custom"

type IrUiViewCustoms []IrUiViewCustom

type IrUiViewCustomsNil []IrUiViewCustomNil

func (s *IrUiViewCustom) NilableType_() interface{} {
	return &IrUiViewCustomNil{}
}

func (ns *IrUiViewCustomNil) Type_() interface{} {
	s := &IrUiViewCustom{}
	return load(ns, s)
}

func (s *IrUiViewCustoms) NilableType_() interface{} {
	return &IrUiViewCustomsNil{}
}

func (ns *IrUiViewCustomsNil) Type_() interface{} {
	s := &IrUiViewCustoms{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrUiViewCustom))
	}
	return s
}
