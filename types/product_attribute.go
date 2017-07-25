package types

import (
	"time"
)

type ProductAttribute struct {
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	AttributeLineIds []int64   `xmlrpc:"attribute_line_ids"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	CreateVariant    bool      `xmlrpc:"create_variant"`
	DisplayName      string    `xmlrpc:"display_name"`
	Id               int64     `xmlrpc:"id"`
	Name             string    `xmlrpc:"name"`
	Sequence         int64     `xmlrpc:"sequence"`
	ValueIds         []int64   `xmlrpc:"value_ids"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type ProductAttributeNil struct {
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	AttributeLineIds interface{} `xmlrpc:"attribute_line_ids"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	CreateVariant    bool        `xmlrpc:"create_variant"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Id               interface{} `xmlrpc:"id"`
	Name             interface{} `xmlrpc:"name"`
	Sequence         interface{} `xmlrpc:"sequence"`
	ValueIds         interface{} `xmlrpc:"value_ids"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var ProductAttributeModel string = "product.attribute"

type ProductAttributes []ProductAttribute

type ProductAttributesNil []ProductAttributeNil

func (s *ProductAttribute) NilableType_() interface{} {
	return &ProductAttributeNil{}
}

func (ns *ProductAttributeNil) Type_() interface{} {
	s := &ProductAttribute{}
	return load(ns, s)
}

func (s *ProductAttributes) NilableType_() interface{} {
	return &ProductAttributesNil{}
}

func (ns *ProductAttributesNil) Type_() interface{} {
	s := &ProductAttributes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductAttribute))
	}
	return s
}
