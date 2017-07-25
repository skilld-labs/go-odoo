package types

import (
	"time"
)

type IrExports struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DisplayName  string    `xmlrpc:"display_name"`
	ExportFields []int64   `xmlrpc:"export_fields"`
	Id           int64     `xmlrpc:"id"`
	Name         string    `xmlrpc:"name"`
	Resource     string    `xmlrpc:"resource"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type IrExportsNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	ExportFields interface{} `xmlrpc:"export_fields"`
	Id           interface{} `xmlrpc:"id"`
	Name         interface{} `xmlrpc:"name"`
	Resource     interface{} `xmlrpc:"resource"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var IrExportsModel string = "ir.exports"

type IrExportss []IrExports

type IrExportssNil []IrExportsNil

func (s *IrExports) NilableType_() interface{} {
	return &IrExportsNil{}
}

func (ns *IrExportsNil) Type_() interface{} {
	s := &IrExports{}
	return load(ns, s)
}

func (s *IrExportss) NilableType_() interface{} {
	return &IrExportssNil{}
}

func (ns *IrExportssNil) Type_() interface{} {
	s := &IrExportss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrExports))
	}
	return s
}
