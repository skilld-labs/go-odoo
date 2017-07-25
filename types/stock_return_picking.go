package types

import (
	"time"
)

type StockReturnPicking struct {
	LastUpdate         time.Time `xmlrpc:"__last_update"`
	CreateDate         time.Time `xmlrpc:"create_date"`
	CreateUid          Many2One  `xmlrpc:"create_uid"`
	DisplayName        string    `xmlrpc:"display_name"`
	Id                 int64     `xmlrpc:"id"`
	LocationId         Many2One  `xmlrpc:"location_id"`
	MoveDestExists     bool      `xmlrpc:"move_dest_exists"`
	OriginalLocationId Many2One  `xmlrpc:"original_location_id"`
	ParentLocationId   Many2One  `xmlrpc:"parent_location_id"`
	ProductReturnMoves []int64   `xmlrpc:"product_return_moves"`
	WriteDate          time.Time `xmlrpc:"write_date"`
	WriteUid           Many2One  `xmlrpc:"write_uid"`
}

type StockReturnPickingNil struct {
	LastUpdate         interface{} `xmlrpc:"__last_update"`
	CreateDate         interface{} `xmlrpc:"create_date"`
	CreateUid          interface{} `xmlrpc:"create_uid"`
	DisplayName        interface{} `xmlrpc:"display_name"`
	Id                 interface{} `xmlrpc:"id"`
	LocationId         interface{} `xmlrpc:"location_id"`
	MoveDestExists     bool        `xmlrpc:"move_dest_exists"`
	OriginalLocationId interface{} `xmlrpc:"original_location_id"`
	ParentLocationId   interface{} `xmlrpc:"parent_location_id"`
	ProductReturnMoves interface{} `xmlrpc:"product_return_moves"`
	WriteDate          interface{} `xmlrpc:"write_date"`
	WriteUid           interface{} `xmlrpc:"write_uid"`
}

var StockReturnPickingModel string = "stock.return.picking"

type StockReturnPickings []StockReturnPicking

type StockReturnPickingsNil []StockReturnPickingNil

func (s *StockReturnPicking) NilableType_() interface{} {
	return &StockReturnPickingNil{}
}

func (ns *StockReturnPickingNil) Type_() interface{} {
	s := &StockReturnPicking{}
	return load(ns, s)
}

func (s *StockReturnPickings) NilableType_() interface{} {
	return &StockReturnPickingsNil{}
}

func (ns *StockReturnPickingsNil) Type_() interface{} {
	s := &StockReturnPickings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockReturnPicking))
	}
	return s
}
