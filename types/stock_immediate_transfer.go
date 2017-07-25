package types

import (
	"time"
)

type StockImmediateTransfer struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	PickId      Many2One  `xmlrpc:"pick_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockImmediateTransferNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	PickId      interface{} `xmlrpc:"pick_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockImmediateTransferModel string = "stock.immediate.transfer"

type StockImmediateTransfers []StockImmediateTransfer

type StockImmediateTransfersNil []StockImmediateTransferNil

func (s *StockImmediateTransfer) NilableType_() interface{} {
	return &StockImmediateTransferNil{}
}

func (ns *StockImmediateTransferNil) Type_() interface{} {
	s := &StockImmediateTransfer{}
	return load(ns, s)
}

func (s *StockImmediateTransfers) NilableType_() interface{} {
	return &StockImmediateTransfersNil{}
}

func (ns *StockImmediateTransfersNil) Type_() interface{} {
	s := &StockImmediateTransfers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockImmediateTransfer))
	}
	return s
}
