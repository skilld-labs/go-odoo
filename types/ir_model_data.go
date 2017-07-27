package types

import (
	"time"
)

type IrModelData struct {
	LastUpdate   time.Time `xmlrpc:"__last_update"`
	CompleteName string    `xmlrpc:"complete_name"`
	CreateDate   time.Time `xmlrpc:"create_date"`
	CreateUid    Many2One  `xmlrpc:"create_uid"`
	DateInit     time.Time `xmlrpc:"date_init"`
	DateUpdate   time.Time `xmlrpc:"date_update"`
	DisplayName  string    `xmlrpc:"display_name"`
	Id           int64     `xmlrpc:"id"`
	Model        string    `xmlrpc:"model"`
	Module       string    `xmlrpc:"module"`
	Name         string    `xmlrpc:"name"`
	Noupdate     bool      `xmlrpc:"noupdate"`
	Reference    string    `xmlrpc:"reference"`
	ResId        int64     `xmlrpc:"res_id"`
	WriteDate    time.Time `xmlrpc:"write_date"`
	WriteUid     Many2One  `xmlrpc:"write_uid"`
}

type IrModelDataNil struct {
	LastUpdate   interface{} `xmlrpc:"__last_update"`
	CompleteName interface{} `xmlrpc:"complete_name"`
	CreateDate   interface{} `xmlrpc:"create_date"`
	CreateUid    interface{} `xmlrpc:"create_uid"`
	DateInit     interface{} `xmlrpc:"date_init"`
	DateUpdate   interface{} `xmlrpc:"date_update"`
	DisplayName  interface{} `xmlrpc:"display_name"`
	Id           interface{} `xmlrpc:"id"`
	Model        interface{} `xmlrpc:"model"`
	Module       interface{} `xmlrpc:"module"`
	Name         interface{} `xmlrpc:"name"`
	Noupdate     bool        `xmlrpc:"noupdate"`
	Reference    interface{} `xmlrpc:"reference"`
	ResId        interface{} `xmlrpc:"res_id"`
	WriteDate    interface{} `xmlrpc:"write_date"`
	WriteUid     interface{} `xmlrpc:"write_uid"`
}

var IrModelDataModel string = "ir.model.data"

type IrModelDatas []IrModelData

type IrModelDatasNil []IrModelDataNil

func (s *IrModelData) NilableType_() interface{} {
	return &IrModelDataNil{}
}

func (ns *IrModelDataNil) Type_() interface{} {
	s := &IrModelData{}
	return load(ns, s)
}

func (s *IrModelDatas) NilableType_() interface{} {
	return &IrModelDatasNil{}
}

func (ns *IrModelDatasNil) Type_() interface{} {
	s := &IrModelDatas{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModelData))
	}
	return s
}
