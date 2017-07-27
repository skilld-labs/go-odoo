package types

import (
	"time"
)

type IrQwebFieldSelection struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
}

type IrQwebFieldSelectionNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
}

var IrQwebFieldSelectionModel string = "ir.qweb.field.selection"

type IrQwebFieldSelections []IrQwebFieldSelection

type IrQwebFieldSelectionsNil []IrQwebFieldSelectionNil

func (s *IrQwebFieldSelection) NilableType_() interface{} {
	return &IrQwebFieldSelectionNil{}
}

func (ns *IrQwebFieldSelectionNil) Type_() interface{} {
	s := &IrQwebFieldSelection{}
	return load(ns, s)
}

func (s *IrQwebFieldSelections) NilableType_() interface{} {
	return &IrQwebFieldSelectionsNil{}
}

func (ns *IrQwebFieldSelectionsNil) Type_() interface{} {
	s := &IrQwebFieldSelections{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldSelection))
	}
	return s
}
