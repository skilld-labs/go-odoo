package types

import (
	"time"
)

type SaleLayoutCategory struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	Pagebreak   bool      `xmlrpc:"pagebreak"`
	Sequence    int64     `xmlrpc:"sequence"`
	Subtotal    bool      `xmlrpc:"subtotal"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type SaleLayoutCategoryNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	Pagebreak   bool        `xmlrpc:"pagebreak"`
	Sequence    interface{} `xmlrpc:"sequence"`
	Subtotal    bool        `xmlrpc:"subtotal"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var SaleLayoutCategoryModel string = "sale.layout_category"

type SaleLayoutCategorys []SaleLayoutCategory

type SaleLayoutCategorysNil []SaleLayoutCategoryNil

func (s *SaleLayoutCategory) NilableType_() interface{} {
	return &SaleLayoutCategoryNil{}
}

func (ns *SaleLayoutCategoryNil) Type_() interface{} {
	s := &SaleLayoutCategory{}
	return load(ns, s)
}

func (s *SaleLayoutCategorys) NilableType_() interface{} {
	return &SaleLayoutCategorysNil{}
}

func (ns *SaleLayoutCategorysNil) Type_() interface{} {
	s := &SaleLayoutCategorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SaleLayoutCategory))
	}
	return s
}
