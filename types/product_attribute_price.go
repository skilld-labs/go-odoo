package types

import (
	"time"
)

type ProductAttributePrice struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	PriceExtra    float64   `xmlrpc:"price_extra"`
	ProductTmplId Many2One  `xmlrpc:"product_tmpl_id"`
	ValueId       Many2One  `xmlrpc:"value_id"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type ProductAttributePriceNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	PriceExtra    interface{} `xmlrpc:"price_extra"`
	ProductTmplId interface{} `xmlrpc:"product_tmpl_id"`
	ValueId       interface{} `xmlrpc:"value_id"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var ProductAttributePriceModel string = "product.attribute.price"

type ProductAttributePrices []ProductAttributePrice

type ProductAttributePricesNil []ProductAttributePriceNil

func (s *ProductAttributePrice) NilableType_() interface{} {
	return &ProductAttributePriceNil{}
}

func (ns *ProductAttributePriceNil) Type_() interface{} {
	s := &ProductAttributePrice{}
	return load(ns, s)
}

func (s *ProductAttributePrices) NilableType_() interface{} {
	return &ProductAttributePricesNil{}
}

func (ns *ProductAttributePricesNil) Type_() interface{} {
	s := &ProductAttributePrices{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductAttributePrice))
	}
	return s
}
