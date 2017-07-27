package types

import (
	"time"
)

type ProductRemoval struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Method      string    `xmlrpc:"method"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductRemovalNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Method      interface{} `xmlrpc:"method"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductRemovalModel string = "product.removal"

type ProductRemovals []ProductRemoval

type ProductRemovalsNil []ProductRemovalNil

func (s *ProductRemoval) NilableType_() interface{} {
	return &ProductRemovalNil{}
}

func (ns *ProductRemovalNil) Type_() interface{} {
	s := &ProductRemoval{}
	return load(ns, s)
}

func (s *ProductRemovals) NilableType_() interface{} {
	return &ProductRemovalsNil{}
}

func (ns *ProductRemovalsNil) Type_() interface{} {
	s := &ProductRemovals{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductRemoval))
	}
	return s
}
