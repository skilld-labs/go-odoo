package types

import (
	"time"
)

type ProductAttributeLine struct {
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	AttributeId   Many2One  `xmlrpc:"attribute_id"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	ProductTmplId Many2One  `xmlrpc:"product_tmpl_id"`
	ValueIds      []int64   `xmlrpc:"value_ids"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type ProductAttributeLineNil struct {
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	AttributeId   interface{} `xmlrpc:"attribute_id"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	ProductTmplId interface{} `xmlrpc:"product_tmpl_id"`
	ValueIds      interface{} `xmlrpc:"value_ids"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var ProductAttributeLineModel string = "product.attribute.line"

type ProductAttributeLines []ProductAttributeLine

type ProductAttributeLinesNil []ProductAttributeLineNil

func (s *ProductAttributeLine) NilableType_() interface{} {
	return &ProductAttributeLineNil{}
}

func (ns *ProductAttributeLineNil) Type_() interface{} {
	s := &ProductAttributeLine{}
	return load(ns, s)
}

func (s *ProductAttributeLines) NilableType_() interface{} {
	return &ProductAttributeLinesNil{}
}

func (ns *ProductAttributeLinesNil) Type_() interface{} {
	s := &ProductAttributeLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductAttributeLine))
	}
	return s
}
