package types

import (
	"time"
)

type IrConfigParameter struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	GroupIds    []int64   `xmlrpc:"group_ids"`
	Id          int64     `xmlrpc:"id"`
	Key         string    `xmlrpc:"key"`
	Value       string    `xmlrpc:"value"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type IrConfigParameterNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	GroupIds    interface{} `xmlrpc:"group_ids"`
	Id          interface{} `xmlrpc:"id"`
	Key         interface{} `xmlrpc:"key"`
	Value       interface{} `xmlrpc:"value"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var IrConfigParameterModel string = "ir.config_parameter"

type IrConfigParameters []IrConfigParameter

type IrConfigParametersNil []IrConfigParameterNil

func (s *IrConfigParameter) NilableType_() interface{} {
	return &IrConfigParameterNil{}
}

func (ns *IrConfigParameterNil) Type_() interface{} {
	s := &IrConfigParameter{}
	return load(ns, s)
}

func (s *IrConfigParameters) NilableType_() interface{} {
	return &IrConfigParametersNil{}
}

func (ns *IrConfigParametersNil) Type_() interface{} {
	s := &IrConfigParameters{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrConfigParameter))
	}
	return s
}
