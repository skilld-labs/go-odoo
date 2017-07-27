package types

import (
	"time"
)

type StockPackOperationLot struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LotId       Many2One  `xmlrpc:"lot_id"`
	LotName     string    `xmlrpc:"lot_name"`
	OperationId Many2One  `xmlrpc:"operation_id"`
	PlusVisible bool      `xmlrpc:"plus_visible"`
	Qty         float64   `xmlrpc:"qty"`
	QtyTodo     float64   `xmlrpc:"qty_todo"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockPackOperationLotNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LotId       interface{} `xmlrpc:"lot_id"`
	LotName     interface{} `xmlrpc:"lot_name"`
	OperationId interface{} `xmlrpc:"operation_id"`
	PlusVisible bool        `xmlrpc:"plus_visible"`
	Qty         interface{} `xmlrpc:"qty"`
	QtyTodo     interface{} `xmlrpc:"qty_todo"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockPackOperationLotModel string = "stock.pack.operation.lot"

type StockPackOperationLots []StockPackOperationLot

type StockPackOperationLotsNil []StockPackOperationLotNil

func (s *StockPackOperationLot) NilableType_() interface{} {
	return &StockPackOperationLotNil{}
}

func (ns *StockPackOperationLotNil) Type_() interface{} {
	s := &StockPackOperationLot{}
	return load(ns, s)
}

func (s *StockPackOperationLots) NilableType_() interface{} {
	return &StockPackOperationLotsNil{}
}

func (ns *StockPackOperationLotsNil) Type_() interface{} {
	s := &StockPackOperationLots{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockPackOperationLot))
	}
	return s
}
