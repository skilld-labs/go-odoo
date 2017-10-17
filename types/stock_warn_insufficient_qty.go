package types

import (
	"time"
)

type StockWarnInsufficientQty struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	LocationId  Many2One  `xmlrpc:"location_id"`
	ProductId   Many2One  `xmlrpc:"product_id"`
	QuantIds    []int64   `xmlrpc:"quant_ids"`
}

type StockWarnInsufficientQtyNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	LocationId  interface{} `xmlrpc:"location_id"`
	ProductId   interface{} `xmlrpc:"product_id"`
	QuantIds    interface{} `xmlrpc:"quant_ids"`
}

var StockWarnInsufficientQtyModel string = "stock.warn.insufficient.qty"

type StockWarnInsufficientQtys []StockWarnInsufficientQty

type StockWarnInsufficientQtysNil []StockWarnInsufficientQtyNil

func (s *StockWarnInsufficientQty) NilableType_() interface{} {
	return &StockWarnInsufficientQtyNil{}
}

func (ns *StockWarnInsufficientQtyNil) Type_() interface{} {
	s := &StockWarnInsufficientQty{}
	return load(ns, s)
}

func (s *StockWarnInsufficientQtys) NilableType_() interface{} {
	return &StockWarnInsufficientQtysNil{}
}

func (ns *StockWarnInsufficientQtysNil) Type_() interface{} {
	s := &StockWarnInsufficientQtys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockWarnInsufficientQty))
	}
	return s
}
