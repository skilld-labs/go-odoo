package types

import (
	"time"
)

type IrValues struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	ActionId      Many2One  `xmlrpc:"action_id"`
	CompanyId     Many2One  `xmlrpc:"company_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	Key           string    `xmlrpc:"key"`
	Key2          string    `xmlrpc:"key2"`
	Model         string    `xmlrpc:"model"`
	ModelId       Many2One  `xmlrpc:"model_id"`
	Name          string    `xmlrpc:"name"`
	ResId         int64     `xmlrpc:"res_id"`
	UserId        Many2One  `xmlrpc:"user_id"`
	Value         string    `xmlrpc:"value"`
	ValueUnpickle string    `xmlrpc:"value_unpickle"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type IrValuesNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	ActionId      interface{} `xmlrpc:"action_id"`
	CompanyId     interface{} `xmlrpc:"company_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	Key           interface{} `xmlrpc:"key"`
	Key2          interface{} `xmlrpc:"key2"`
	Model         interface{} `xmlrpc:"model"`
	ModelId       interface{} `xmlrpc:"model_id"`
	Name          interface{} `xmlrpc:"name"`
	ResId         interface{} `xmlrpc:"res_id"`
	UserId        interface{} `xmlrpc:"user_id"`
	Value         interface{} `xmlrpc:"value"`
	ValueUnpickle interface{} `xmlrpc:"value_unpickle"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var IrValuesModel string = "ir.values"

type IrValuess []IrValues

type IrValuessNil []IrValuesNil

func (s *IrValues) NilableType_() interface{} {
	return &IrValuesNil{}
}

func (ns *IrValuesNil) Type_() interface{} {
	s := &IrValues{}
	return load(ns, s)
}

func (s *IrValuess) NilableType_() interface{} {
	return &IrValuessNil{}
}

func (ns *IrValuessNil) Type_() interface{} {
	s := &IrValuess{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrValues))
	}
	return s
}
