package types

import (
	"time"
)

type ProductUomCateg struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductUomCategNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductUomCategModel string = "product.uom.categ"

type ProductUomCategs []ProductUomCateg

type ProductUomCategsNil []ProductUomCategNil

func (s *ProductUomCateg) NilableType_() interface{} {
	return &ProductUomCategNil{}
}

func (ns *ProductUomCategNil) Type_() interface{} {
	s := &ProductUomCateg{}
	return load(ns, s)
}

func (s *ProductUomCategs) NilableType_() interface{} {
	return &ProductUomCategsNil{}
}

func (ns *ProductUomCategsNil) Type_() interface{} {
	s := &ProductUomCategs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductUomCateg))
	}
	return s
}
