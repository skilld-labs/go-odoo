package types

import (
	"time"
)

type StockQuant struct {
	CompanyId        Many2One  `xmlrpc:"company_id"`
	CreateDate       time.Time `xmlrpc:"create_date"`
	CreateUid        Many2One  `xmlrpc:"create_uid"`
	DisplayName      string    `xmlrpc:"display_name"`
	Id               int64     `xmlrpc:"id"`
	InDate           time.Time `xmlrpc:"in_date"`
	LastUpdate       time.Time `xmlrpc:"__last_update"`
	LocationId       Many2One  `xmlrpc:"location_id"`
	LotId            Many2One  `xmlrpc:"lot_id"`
	OwnerId          Many2One  `xmlrpc:"owner_id"`
	PackageId        Many2One  `xmlrpc:"package_id"`
	ProductId        Many2One  `xmlrpc:"product_id"`
	ProductTmplId    Many2One  `xmlrpc:"product_tmpl_id"`
	ProductUomId     Many2One  `xmlrpc:"product_uom_id"`
	Quantity         float64   `xmlrpc:"quantity"`
	ReservedQuantity float64   `xmlrpc:"reserved_quantity"`
	WriteDate        time.Time `xmlrpc:"write_date"`
	WriteUid         Many2One  `xmlrpc:"write_uid"`
}

type StockQuantNil struct {
	CompanyId        interface{} `xmlrpc:"company_id"`
	CreateDate       interface{} `xmlrpc:"create_date"`
	CreateUid        interface{} `xmlrpc:"create_uid"`
	DisplayName      interface{} `xmlrpc:"display_name"`
	Id               interface{} `xmlrpc:"id"`
	InDate           interface{} `xmlrpc:"in_date"`
	LastUpdate       interface{} `xmlrpc:"__last_update"`
	LocationId       interface{} `xmlrpc:"location_id"`
	LotId            interface{} `xmlrpc:"lot_id"`
	OwnerId          interface{} `xmlrpc:"owner_id"`
	PackageId        interface{} `xmlrpc:"package_id"`
	ProductId        interface{} `xmlrpc:"product_id"`
	ProductTmplId    interface{} `xmlrpc:"product_tmpl_id"`
	ProductUomId     interface{} `xmlrpc:"product_uom_id"`
	Quantity         interface{} `xmlrpc:"quantity"`
	ReservedQuantity interface{} `xmlrpc:"reserved_quantity"`
	WriteDate        interface{} `xmlrpc:"write_date"`
	WriteUid         interface{} `xmlrpc:"write_uid"`
}

var StockQuantModel string = "stock.quant"

type StockQuants []StockQuant

type StockQuantsNil []StockQuantNil

func (s *StockQuant) NilableType_() interface{} {
	return &StockQuantNil{}
}

func (ns *StockQuantNil) Type_() interface{} {
	s := &StockQuant{}
	return load(ns, s)
}

func (s *StockQuants) NilableType_() interface{} {
	return &StockQuantsNil{}
}

func (ns *StockQuantsNil) Type_() interface{} {
	s := &StockQuants{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockQuant))
	}
	return s
}
