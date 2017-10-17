package types

import (
	"time"
)

type StockSchedulerCompute struct {
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type StockSchedulerComputeNil struct {
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var StockSchedulerComputeModel string = "stock.scheduler.compute"

type StockSchedulerComputes []StockSchedulerCompute

type StockSchedulerComputesNil []StockSchedulerComputeNil

func (s *StockSchedulerCompute) NilableType_() interface{} {
	return &StockSchedulerComputeNil{}
}

func (ns *StockSchedulerComputeNil) Type_() interface{} {
	s := &StockSchedulerCompute{}
	return load(ns, s)
}

func (s *StockSchedulerComputes) NilableType_() interface{} {
	return &StockSchedulerComputesNil{}
}

func (ns *StockSchedulerComputesNil) Type_() interface{} {
	s := &StockSchedulerComputes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockSchedulerCompute))
	}
	return s
}
