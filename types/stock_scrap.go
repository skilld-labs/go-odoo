package types

import (
	"time"
)

type StockScrap struct {
	LastUpdate      time.Time `xmlrpc:"__last_update"`
	CreateDate      time.Time `xmlrpc:"create_date"`
	CreateUid       Many2One  `xmlrpc:"create_uid"`
	DateExpected    time.Time `xmlrpc:"date_expected"`
	DisplayName     string    `xmlrpc:"display_name"`
	Id              int64     `xmlrpc:"id"`
	LocationId      Many2One  `xmlrpc:"location_id"`
	LotId           Many2One  `xmlrpc:"lot_id"`
	MoveId          Many2One  `xmlrpc:"move_id"`
	Name            string    `xmlrpc:"name"`
	Origin          string    `xmlrpc:"origin"`
	OwnerId         Many2One  `xmlrpc:"owner_id"`
	PackageId       Many2One  `xmlrpc:"package_id"`
	PickingId       Many2One  `xmlrpc:"picking_id"`
	ProductId       Many2One  `xmlrpc:"product_id"`
	ProductUomId    Many2One  `xmlrpc:"product_uom_id"`
	ScrapLocationId Many2One  `xmlrpc:"scrap_location_id"`
	ScrapQty        float64   `xmlrpc:"scrap_qty"`
	State           string    `xmlrpc:"state"`
	Tracking        string    `xmlrpc:"tracking"`
	WriteDate       time.Time `xmlrpc:"write_date"`
	WriteUid        Many2One  `xmlrpc:"write_uid"`
}

type StockScrapNil struct {
	LastUpdate      interface{} `xmlrpc:"__last_update"`
	CreateDate      interface{} `xmlrpc:"create_date"`
	CreateUid       interface{} `xmlrpc:"create_uid"`
	DateExpected    interface{} `xmlrpc:"date_expected"`
	DisplayName     interface{} `xmlrpc:"display_name"`
	Id              interface{} `xmlrpc:"id"`
	LocationId      interface{} `xmlrpc:"location_id"`
	LotId           interface{} `xmlrpc:"lot_id"`
	MoveId          interface{} `xmlrpc:"move_id"`
	Name            interface{} `xmlrpc:"name"`
	Origin          interface{} `xmlrpc:"origin"`
	OwnerId         interface{} `xmlrpc:"owner_id"`
	PackageId       interface{} `xmlrpc:"package_id"`
	PickingId       interface{} `xmlrpc:"picking_id"`
	ProductId       interface{} `xmlrpc:"product_id"`
	ProductUomId    interface{} `xmlrpc:"product_uom_id"`
	ScrapLocationId interface{} `xmlrpc:"scrap_location_id"`
	ScrapQty        interface{} `xmlrpc:"scrap_qty"`
	State           interface{} `xmlrpc:"state"`
	Tracking        interface{} `xmlrpc:"tracking"`
	WriteDate       interface{} `xmlrpc:"write_date"`
	WriteUid        interface{} `xmlrpc:"write_uid"`
}

var StockScrapModel string = "stock.scrap"

type StockScraps []StockScrap

type StockScrapsNil []StockScrapNil

func (s *StockScrap) NilableType_() interface{} {
	return &StockScrapNil{}
}

func (ns *StockScrapNil) Type_() interface{} {
	s := &StockScrap{}
	return load(ns, s)
}

func (s *StockScraps) NilableType_() interface{} {
	return &StockScrapsNil{}
}

func (ns *StockScrapsNil) Type_() interface{} {
	s := &StockScraps{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockScrap))
	}
	return s
}
