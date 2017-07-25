package types

import (
	"time"
)

type BarcodeNomenclature struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	Name        string    `xmlrpc:"name"`
	RuleIds     []int64   `xmlrpc:"rule_ids"`
	UpcEanConv  string    `xmlrpc:"upc_ean_conv"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type BarcodeNomenclatureNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	Name        interface{} `xmlrpc:"name"`
	RuleIds     interface{} `xmlrpc:"rule_ids"`
	UpcEanConv  interface{} `xmlrpc:"upc_ean_conv"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var BarcodeNomenclatureModel string = "barcode.nomenclature"

type BarcodeNomenclatures []BarcodeNomenclature

type BarcodeNomenclaturesNil []BarcodeNomenclatureNil

func (s *BarcodeNomenclature) NilableType_() interface{} {
	return &BarcodeNomenclatureNil{}
}

func (ns *BarcodeNomenclatureNil) Type_() interface{} {
	s := &BarcodeNomenclature{}
	return load(ns, s)
}

func (s *BarcodeNomenclatures) NilableType_() interface{} {
	return &BarcodeNomenclaturesNil{}
}

func (ns *BarcodeNomenclaturesNil) Type_() interface{} {
	s := &BarcodeNomenclatures{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BarcodeNomenclature))
	}
	return s
}
