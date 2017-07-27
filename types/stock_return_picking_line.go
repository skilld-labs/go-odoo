package types

import (
	"time"
)

type StockReturnPickingLine struct {
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	MoveId      Many2One  `xmlrpc:"move_id"`
	ProductId   Many2One  `xmlrpc:"product_id"`
	Quantity    float64   `xmlrpc:"quantity"`
	ToRefundSo  bool      `xmlrpc:"to_refund_so"`
	WizardId    Many2One  `xmlrpc:"wizard_id"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockReturnPickingLineNil struct {
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	MoveId      interface{} `xmlrpc:"move_id"`
	ProductId   interface{} `xmlrpc:"product_id"`
	Quantity    interface{} `xmlrpc:"quantity"`
	ToRefundSo  bool        `xmlrpc:"to_refund_so"`
	WizardId    interface{} `xmlrpc:"wizard_id"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockReturnPickingLineModel string = "stock.return.picking.line"

type StockReturnPickingLines []StockReturnPickingLine

type StockReturnPickingLinesNil []StockReturnPickingLineNil

func (s *StockReturnPickingLine) NilableType_() interface{} {
	return &StockReturnPickingLineNil{}
}

func (ns *StockReturnPickingLineNil) Type_() interface{} {
	s := &StockReturnPickingLine{}
	return load(ns, s)
}

func (s *StockReturnPickingLines) NilableType_() interface{} {
	return &StockReturnPickingLinesNil{}
}

func (ns *StockReturnPickingLinesNil) Type_() interface{} {
	s := &StockReturnPickingLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockReturnPickingLine))
	}
	return s
}
