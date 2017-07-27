package types

import (
	"time"
)

type IrExportsLine struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	ExportId    Many2One  `xmlrpc:"export_id"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrExportsLineNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	ExportId    interface{} `xmlrpc:"export_id"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrExportsLineModel string = "ir.exports.line"

type IrExportsLines []IrExportsLine

type IrExportsLinesNil []IrExportsLineNil

func (s *IrExportsLine) NilableType_() interface{} {
	return &IrExportsLineNil{}
}

func (ns *IrExportsLineNil) Type_() interface{} {
	s := &IrExportsLine{}
	return load(ns, s)
}

func (s *IrExportsLines) NilableType_() interface{} {
	return &IrExportsLinesNil{}
}

func (ns *IrExportsLinesNil) Type_() interface{} {
	s := &IrExportsLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrExportsLine))
	}
	return s
}
