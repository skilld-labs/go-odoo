package types

import (
	"time"
)

type StockWarnInsufficientQtyScrap struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	LocationId  Many2One  `xmlrpc:"location_id"`
	ProductId   Many2One  `xmlrpc:"product_id"`
	QuantIds    []int64   `xmlrpc:"quant_ids"`
	ScrapId     Many2One  `xmlrpc:"scrap_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockWarnInsufficientQtyScrapNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	LocationId  interface{} `xmlrpc:"location_id"`
	ProductId   interface{} `xmlrpc:"product_id"`
	QuantIds    interface{} `xmlrpc:"quant_ids"`
	ScrapId     interface{} `xmlrpc:"scrap_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockWarnInsufficientQtyScrapModel string = "stock.warn.insufficient.qty.scrap"

type StockWarnInsufficientQtyScraps []StockWarnInsufficientQtyScrap

type StockWarnInsufficientQtyScrapsNil []StockWarnInsufficientQtyScrapNil

func (s *StockWarnInsufficientQtyScrap) NilableType_() interface{} {
	return &StockWarnInsufficientQtyScrapNil{}
}

func (ns *StockWarnInsufficientQtyScrapNil) Type_() interface{} {
	s := &StockWarnInsufficientQtyScrap{}
	return load(ns, s)
}

func (s *StockWarnInsufficientQtyScraps) NilableType_() interface{} {
	return &StockWarnInsufficientQtyScrapsNil{}
}

func (ns *StockWarnInsufficientQtyScrapsNil) Type_() interface{} {
	s := &StockWarnInsufficientQtyScraps{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockWarnInsufficientQtyScrap))
	}
	return s
}
