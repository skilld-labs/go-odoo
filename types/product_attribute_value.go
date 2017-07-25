package types

import (
	"time"
)

type ProductAttributeValue struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	AttributeId Many2One  `xmlrpc:"attribute_id"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	PriceExtra  float64   `xmlrpc:"price_extra"`
	PriceIds    []int64   `xmlrpc:"price_ids"`
	ProductIds  []int64   `xmlrpc:"product_ids"`
	Sequence    int64     `xmlrpc:"sequence"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductAttributeValueNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	AttributeId interface{} `xmlrpc:"attribute_id"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	PriceExtra  interface{} `xmlrpc:"price_extra"`
	PriceIds    interface{} `xmlrpc:"price_ids"`
	ProductIds  interface{} `xmlrpc:"product_ids"`
	Sequence    interface{} `xmlrpc:"sequence"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductAttributeValueModel string = "product.attribute.value"

type ProductAttributeValues []ProductAttributeValue

type ProductAttributeValuesNil []ProductAttributeValueNil

func (s *ProductAttributeValue) NilableType_() interface{} {
	return &ProductAttributeValueNil{}
}

func (ns *ProductAttributeValueNil) Type_() interface{} {
	s := &ProductAttributeValue{}
	return load(ns, s)
}

func (s *ProductAttributeValues) NilableType_() interface{} {
	return &ProductAttributeValuesNil{}
}

func (ns *ProductAttributeValuesNil) Type_() interface{} {
	s := &ProductAttributeValues{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductAttributeValue))
	}
	return s
}
