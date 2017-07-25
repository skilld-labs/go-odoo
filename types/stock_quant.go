package types

import (
	"time"
)

type StockQuant struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	Cost                   float64   `xmlrpc:"cost"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	DisplayName            string    `xmlrpc:"display_name"`
	HistoryIds             []int64   `xmlrpc:"history_ids"`
	Id                     int64     `xmlrpc:"id"`
	InDate                 time.Time `xmlrpc:"in_date"`
	InventoryValue         float64   `xmlrpc:"inventory_value"`
	LocationId             Many2One  `xmlrpc:"location_id"`
	LotId                  Many2One  `xmlrpc:"lot_id"`
	Name                   string    `xmlrpc:"name"`
	NegativeDestLocationId Many2One  `xmlrpc:"negative_dest_location_id"`
	NegativeMoveId         Many2One  `xmlrpc:"negative_move_id"`
	OwnerId                Many2One  `xmlrpc:"owner_id"`
	PackageId              Many2One  `xmlrpc:"package_id"`
	PackagingTypeId        Many2One  `xmlrpc:"packaging_type_id"`
	ProductId              Many2One  `xmlrpc:"product_id"`
	ProductUomId           Many2One  `xmlrpc:"product_uom_id"`
	PropagatedFromId       Many2One  `xmlrpc:"propagated_from_id"`
	Qty                    float64   `xmlrpc:"qty"`
	ReservationId          Many2One  `xmlrpc:"reservation_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type StockQuantNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	Cost                   interface{} `xmlrpc:"cost"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	HistoryIds             interface{} `xmlrpc:"history_ids"`
	Id                     interface{} `xmlrpc:"id"`
	InDate                 interface{} `xmlrpc:"in_date"`
	InventoryValue         interface{} `xmlrpc:"inventory_value"`
	LocationId             interface{} `xmlrpc:"location_id"`
	LotId                  interface{} `xmlrpc:"lot_id"`
	Name                   interface{} `xmlrpc:"name"`
	NegativeDestLocationId interface{} `xmlrpc:"negative_dest_location_id"`
	NegativeMoveId         interface{} `xmlrpc:"negative_move_id"`
	OwnerId                interface{} `xmlrpc:"owner_id"`
	PackageId              interface{} `xmlrpc:"package_id"`
	PackagingTypeId        interface{} `xmlrpc:"packaging_type_id"`
	ProductId              interface{} `xmlrpc:"product_id"`
	ProductUomId           interface{} `xmlrpc:"product_uom_id"`
	PropagatedFromId       interface{} `xmlrpc:"propagated_from_id"`
	Qty                    interface{} `xmlrpc:"qty"`
	ReservationId          interface{} `xmlrpc:"reservation_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var StockQuantModel string = "stock.quant"

type StockQuants []StockQuant

type StockQuantsNil []StockQuantNil

func (s *StockQuant) NilableType_() interface{} {
	return &StockQuantNil{}
}

func (ns *StockQuantNil) Type_() interface{} {
	s := &StockQuant{}
	return load(ns, s)
}

func (s *StockQuants) NilableType_() interface{} {
	return &StockQuantsNil{}
}

func (ns *StockQuantsNil) Type_() interface{} {
	s := &StockQuants{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockQuant))
	}
	return s
}
