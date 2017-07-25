package types

import (
	"time"
)

type IrFilters struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	ActionId    Many2One  `xmlrpc:"action_id"`
	Active      bool      `xmlrpc:"active"`
	Context     string    `xmlrpc:"context"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Domain      string    `xmlrpc:"domain"`
	Id          int64     `xmlrpc:"id"`
	IsDefault   bool      `xmlrpc:"is_default"`
	ModelId     string    `xmlrpc:"model_id"`
	Name        string    `xmlrpc:"name"`
	Sort        string    `xmlrpc:"sort"`
	UserId      Many2One  `xmlrpc:"user_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrFiltersNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	ActionId    interface{} `xmlrpc:"action_id"`
	Active      bool        `xmlrpc:"active"`
	Context     interface{} `xmlrpc:"context"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Domain      interface{} `xmlrpc:"domain"`
	Id          interface{} `xmlrpc:"id"`
	IsDefault   bool        `xmlrpc:"is_default"`
	ModelId     interface{} `xmlrpc:"model_id"`
	Name        interface{} `xmlrpc:"name"`
	Sort        interface{} `xmlrpc:"sort"`
	UserId      interface{} `xmlrpc:"user_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrFiltersModel string = "ir.filters"

type IrFilterss []IrFilters

type IrFilterssNil []IrFiltersNil

func (s *IrFilters) NilableType_() interface{} {
	return &IrFiltersNil{}
}

func (ns *IrFiltersNil) Type_() interface{} {
	s := &IrFilters{}
	return load(ns, s)
}

func (s *IrFilterss) NilableType_() interface{} {
	return &IrFilterssNil{}
}

func (ns *IrFilterssNil) Type_() interface{} {
	s := &IrFilterss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrFilters))
	}
	return s
}
