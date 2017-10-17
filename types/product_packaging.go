package types

import (
	"time"
)

type ProductPackaging struct {
	Barcode     string    `xmlrpc:"barcode"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	ProductId   Many2One  `xmlrpc:"product_id"`
	Qty         float64   `xmlrpc:"qty"`
	Sequence    int64     `xmlrpc:"sequence"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type ProductPackagingNil struct {
	Barcode     interface{} `xmlrpc:"barcode"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	ProductId   interface{} `xmlrpc:"product_id"`
	Qty         interface{} `xmlrpc:"qty"`
	Sequence    interface{} `xmlrpc:"sequence"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var ProductPackagingModel string = "product.packaging"

type ProductPackagings []ProductPackaging

type ProductPackagingsNil []ProductPackagingNil

func (s *ProductPackaging) NilableType_() interface{} {
	return &ProductPackagingNil{}
}

func (ns *ProductPackagingNil) Type_() interface{} {
	s := &ProductPackaging{}
	return load(ns, s)
}

func (s *ProductPackagings) NilableType_() interface{} {
	return &ProductPackagingsNil{}
}

func (ns *ProductPackagingsNil) Type_() interface{} {
	s := &ProductPackagings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductPackaging))
	}
	return s
}
