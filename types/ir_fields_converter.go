package types

import (
	"time"
)

type IrFieldsConverter struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrFieldsConverterNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrFieldsConverterModel string = "ir.fields.converter"

type IrFieldsConverters []IrFieldsConverter

type IrFieldsConvertersNil []IrFieldsConverterNil

func (s *IrFieldsConverter) NilableType_() interface{} {
	return &IrFieldsConverterNil{}
}

func (ns *IrFieldsConverterNil) Type_() interface{} {
	s := &IrFieldsConverter{}
	return load(ns, s)
}

func (s *IrFieldsConverters) NilableType_() interface{} {
	return &IrFieldsConvertersNil{}
}

func (ns *IrFieldsConvertersNil) Type_() interface{} {
	s := &IrFieldsConverters{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrFieldsConverter))
	}
	return s
}
