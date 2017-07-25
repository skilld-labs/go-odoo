package types

import (
	"time"
)

type StockMoveOperationLink struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	MoveId          Many2One  `xmlrpc:"move_id"`
	OperationId     Many2One  `xmlrpc:"operation_id"`
	Qty             float64   `xmlrpc:"qty"`
	ReservedQuantId Many2One  `xmlrpc:"reserved_quant_id"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type StockMoveOperationLinkNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	MoveId          interface{} `xmlrpc:"move_id"`
	OperationId     interface{} `xmlrpc:"operation_id"`
	Qty             interface{} `xmlrpc:"qty"`
	ReservedQuantId interface{} `xmlrpc:"reserved_quant_id"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var StockMoveOperationLinkModel string = "stock.move.operation.link"

type StockMoveOperationLinks []StockMoveOperationLink

type StockMoveOperationLinksNil []StockMoveOperationLinkNil

func (s *StockMoveOperationLink) NilableType_() interface{} {
	return &StockMoveOperationLinkNil{}
}

func (ns *StockMoveOperationLinkNil) Type_() interface{} {
	s := &StockMoveOperationLink{}
	return load(ns, s)
}

func (s *StockMoveOperationLinks) NilableType_() interface{} {
	return &StockMoveOperationLinksNil{}
}

func (ns *StockMoveOperationLinksNil) Type_() interface{} {
	s := &StockMoveOperationLinks{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockMoveOperationLink))
	}
	return s
}
