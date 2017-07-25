package types

import (
	"time"
)

type ProductUom struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Active      bool      `xmlrpc:"active"`
	CategoryId  Many2One  `xmlrpc:"category_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Factor      float64   `xmlrpc:"factor"`
	FactorInv   float64   `xmlrpc:"factor_inv"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Rounding    float64   `xmlrpc:"rounding"`
	UomType     string    `xmlrpc:"uom_type"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductUomNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Active      bool        `xmlrpc:"active"`
	CategoryId  interface{} `xmlrpc:"category_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Factor      interface{} `xmlrpc:"factor"`
	FactorInv   interface{} `xmlrpc:"factor_inv"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Rounding    interface{} `xmlrpc:"rounding"`
	UomType     interface{} `xmlrpc:"uom_type"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductUomModel string = "product.uom"

type ProductUoms []ProductUom

type ProductUomsNil []ProductUomNil

func (s *ProductUom) NilableType_() interface{} {
	return &ProductUomNil{}
}

func (ns *ProductUomNil) Type_() interface{} {
	s := &ProductUom{}
	return load(ns, s)
}

func (s *ProductUoms) NilableType_() interface{} {
	return &ProductUomsNil{}
}

func (ns *ProductUomsNil) Type_() interface{} {
	s := &ProductUoms{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductUom))
	}
	return s
}
