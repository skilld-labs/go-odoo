package types

import (
	"time"
)

type IrModel struct {
	LastUpdate        time.Time `xmlrpc:"__last_update"`
	AccessIds         []int64   `xmlrpc:"access_ids"`
	CreateDate        time.Time `xmlrpc:"create_date"`
	CreateUid         Many2One  `xmlrpc:"create_uid"`
	DisplayName       string    `xmlrpc:"display_name"`
	FieldId           []int64   `xmlrpc:"field_id"`
	Id                int64     `xmlrpc:"id"`
	Info              string    `xmlrpc:"info"`
	InheritedModelIds []int64   `xmlrpc:"inherited_model_ids"`
	Model             string    `xmlrpc:"model"`
	Modules           string    `xmlrpc:"modules"`
	Name              string    `xmlrpc:"name"`
	State             string    `xmlrpc:"state"`
	Transient         bool      `xmlrpc:"transient"`
	ViewIds           []int64   `xmlrpc:"view_ids"`
	WriteDate         time.Time `xmlrpc:"write_date"`
	WriteUid          Many2One  `xmlrpc:"write_uid"`
}

type IrModelNil struct {
	LastUpdate        interface{} `xmlrpc:"__last_update"`
	AccessIds         interface{} `xmlrpc:"access_ids"`
	CreateDate        interface{} `xmlrpc:"create_date"`
	CreateUid         interface{} `xmlrpc:"create_uid"`
	DisplayName       interface{} `xmlrpc:"display_name"`
	FieldId           interface{} `xmlrpc:"field_id"`
	Id                interface{} `xmlrpc:"id"`
	Info              interface{} `xmlrpc:"info"`
	InheritedModelIds interface{} `xmlrpc:"inherited_model_ids"`
	Model             interface{} `xmlrpc:"model"`
	Modules           interface{} `xmlrpc:"modules"`
	Name              interface{} `xmlrpc:"name"`
	State             interface{} `xmlrpc:"state"`
	Transient         bool        `xmlrpc:"transient"`
	ViewIds           interface{} `xmlrpc:"view_ids"`
	WriteDate         interface{} `xmlrpc:"write_date"`
	WriteUid          interface{} `xmlrpc:"write_uid"`
}

var IrModelModel string = "ir.model"

type IrModels []IrModel

type IrModelsNil []IrModelNil

func (s *IrModel) NilableType_() interface{} {
	return &IrModelNil{}
}

func (ns *IrModelNil) Type_() interface{} {
	s := &IrModel{}
	return load(ns, s)
}

func (s *IrModels) NilableType_() interface{} {
	return &IrModelsNil{}
}

func (ns *IrModelsNil) Type_() interface{} {
	s := &IrModels{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrModel))
	}
	return s
}
