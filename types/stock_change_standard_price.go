package types

import (
	"time"
)

type StockChangeStandardPrice struct {
	LastUpdate           time.Time `xmlrpc:"__last_update"`
	CounterpartAccountId Many2One  `xmlrpc:"counterpart_account_id"`
	CreateDate           time.Time `xmlrpc:"create_date"`
	CreateUid            Many2One  `xmlrpc:"create_uid"`
	DisplayName          string    `xmlrpc:"display_name"`
	Id                   int64     `xmlrpc:"id"`
	NewPrice             float64   `xmlrpc:"new_price"`
	WriteDate            time.Time `xmlrpc:"write_date"`
	WriteUid             Many2One  `xmlrpc:"write_uid"`
}

type StockChangeStandardPriceNil struct {
	LastUpdate           interface{} `xmlrpc:"__last_update"`
	CounterpartAccountId interface{} `xmlrpc:"counterpart_account_id"`
	CreateDate           interface{} `xmlrpc:"create_date"`
	CreateUid            interface{} `xmlrpc:"create_uid"`
	DisplayName          interface{} `xmlrpc:"display_name"`
	Id                   interface{} `xmlrpc:"id"`
	NewPrice             interface{} `xmlrpc:"new_price"`
	WriteDate            interface{} `xmlrpc:"write_date"`
	WriteUid             interface{} `xmlrpc:"write_uid"`
}

var StockChangeStandardPriceModel string = "stock.change.standard.price"

type StockChangeStandardPrices []StockChangeStandardPrice

type StockChangeStandardPricesNil []StockChangeStandardPriceNil

func (s *StockChangeStandardPrice) NilableType_() interface{} {
	return &StockChangeStandardPriceNil{}
}

func (ns *StockChangeStandardPriceNil) Type_() interface{} {
	s := &StockChangeStandardPrice{}
	return load(ns, s)
}

func (s *StockChangeStandardPrices) NilableType_() interface{} {
	return &StockChangeStandardPricesNil{}
}

func (ns *StockChangeStandardPricesNil) Type_() interface{} {
	s := &StockChangeStandardPrices{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockChangeStandardPrice))
	}
	return s
}
