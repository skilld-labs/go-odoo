package types

import (
	"time"
)

type StockFixedPutawayStrat struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CategoryId      Many2One  `xmlrpc:"category_id"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DisplayName     string    `xmlrpc:"display_name"`
	FixedLocationId Many2One  `xmlrpc:"fixed_location_id"`
	Id              int64     `xmlrpc:"id"`
	PutawayId       Many2One  `xmlrpc:"putaway_id"`
	Sequence        int64     `xmlrpc:"sequence"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type StockFixedPutawayStratNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CategoryId      interface{} `xmlrpc:"category_id"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	FixedLocationId interface{} `xmlrpc:"fixed_location_id"`
	Id              interface{} `xmlrpc:"id"`
	PutawayId       interface{} `xmlrpc:"putaway_id"`
	Sequence        interface{} `xmlrpc:"sequence"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var StockFixedPutawayStratModel string = "stock.fixed.putaway.strat"

type StockFixedPutawayStrats []StockFixedPutawayStrat

type StockFixedPutawayStratsNil []StockFixedPutawayStratNil

func (s *StockFixedPutawayStrat) NilableType_() interface{} {
	return &StockFixedPutawayStratNil{}
}

func (ns *StockFixedPutawayStratNil) Type_() interface{} {
	s := &StockFixedPutawayStrat{}
	return load(ns, s)
}

func (s *StockFixedPutawayStrats) NilableType_() interface{} {
	return &StockFixedPutawayStratsNil{}
}

func (ns *StockFixedPutawayStratsNil) Type_() interface{} {
	s := &StockFixedPutawayStrats{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockFixedPutawayStrat))
	}
	return s
}
