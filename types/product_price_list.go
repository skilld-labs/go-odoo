package types

import (
	"time"
)

type ProductPriceList struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PriceList   Many2One  `xmlrpc:"price_list"`
	Qty1        int64     `xmlrpc:"qty1"`
	Qty2        int64     `xmlrpc:"qty2"`
	Qty3        int64     `xmlrpc:"qty3"`
	Qty4        int64     `xmlrpc:"qty4"`
	Qty5        int64     `xmlrpc:"qty5"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductPriceListNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PriceList   interface{} `xmlrpc:"price_list"`
	Qty1        interface{} `xmlrpc:"qty1"`
	Qty2        interface{} `xmlrpc:"qty2"`
	Qty3        interface{} `xmlrpc:"qty3"`
	Qty4        interface{} `xmlrpc:"qty4"`
	Qty5        interface{} `xmlrpc:"qty5"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductPriceListModel string = "product.price_list"

type ProductPriceLists []ProductPriceList

type ProductPriceListsNil []ProductPriceListNil

func (s *ProductPriceList) NilableType_() interface{} {
	return &ProductPriceListNil{}
}

func (ns *ProductPriceListNil) Type_() interface{} {
	s := &ProductPriceList{}
	return load(ns, s)
}

func (s *ProductPriceLists) NilableType_() interface{} {
	return &ProductPriceListsNil{}
}

func (ns *ProductPriceListsNil) Type_() interface{} {
	s := &ProductPriceLists{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductPriceList))
	}
	return s
}
