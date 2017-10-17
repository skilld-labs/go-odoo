package types

import (
	"time"
)

type StockOverprocessedTransfer struct {
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DisplayName              string    `xmlrpc:"display_name"`
	Id                       int64     `xmlrpc:"id"`
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	OverprocessedProductName string    `xmlrpc:"overprocessed_product_name"`
	PickingId                Many2One  `xmlrpc:"picking_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type StockOverprocessedTransferNil struct {
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	Id                       interface{} `xmlrpc:"id"`
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	OverprocessedProductName interface{} `xmlrpc:"overprocessed_product_name"`
	PickingId                interface{} `xmlrpc:"picking_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var StockOverprocessedTransferModel string = "stock.overprocessed.transfer"

type StockOverprocessedTransfers []StockOverprocessedTransfer

type StockOverprocessedTransfersNil []StockOverprocessedTransferNil

func (s *StockOverprocessedTransfer) NilableType_() interface{} {
	return &StockOverprocessedTransferNil{}
}

func (ns *StockOverprocessedTransferNil) Type_() interface{} {
	s := &StockOverprocessedTransfer{}
	return load(ns, s)
}

func (s *StockOverprocessedTransfers) NilableType_() interface{} {
	return &StockOverprocessedTransfersNil{}
}

func (ns *StockOverprocessedTransfersNil) Type_() interface{} {
	s := &StockOverprocessedTransfers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockOverprocessedTransfer))
	}
	return s
}
