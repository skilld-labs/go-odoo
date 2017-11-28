package types

import (
	"time"
)

type StockQuantityHistory struct {
	ComputeAtDate interface{} `xmlrpc:"compute_at_date"`
	CreateDate    time.Time   `xmlrpc:"create_date"`
	CreateUid     Many2One    `xmlrpc:"create_uid"`
	Date          time.Time   `xmlrpc:"date"`
	DisplayName   string      `xmlrpc:"display_name"`
	Id            int64       `xmlrpc:"id"`
	LastUpdate    time.Time   `xmlrpc:"__last_update"`
	WriteDate     time.Time   `xmlrpc:"write_date"`
	WriteUid      Many2One    `xmlrpc:"write_uid"`
}

type StockQuantityHistoryNil struct {
	ComputeAtDate interface{} `xmlrpc:"compute_at_date"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	Date          interface{} `xmlrpc:"date"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var StockQuantityHistoryModel string = "stock.quantity.history"

type StockQuantityHistorys []StockQuantityHistory

type StockQuantityHistorysNil []StockQuantityHistoryNil

func (s *StockQuantityHistory) NilableType_() interface{} {
	return &StockQuantityHistoryNil{}
}

func (ns *StockQuantityHistoryNil) Type_() interface{} {
	s := &StockQuantityHistory{}
	return load(ns, s)
}

func (s *StockQuantityHistorys) NilableType_() interface{} {
	return &StockQuantityHistorysNil{}
}

func (ns *StockQuantityHistorysNil) Type_() interface{} {
	s := &StockQuantityHistorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockQuantityHistory))
	}
	return s
}
