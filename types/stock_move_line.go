package types

import (
	"time"
)

type StockMoveLine struct {
	ConsumeLineIds          []int64   `xmlrpc:"consume_line_ids"`
	CreateDate              time.Time `xmlrpc:"create_date"`
	CreateUid               Many2One  `xmlrpc:"create_uid"`
	Date                    time.Time `xmlrpc:"date"`
	DisplayName             string    `xmlrpc:"display_name"`
	FromLoc                 string    `xmlrpc:"from_loc"`
	Id                      int64     `xmlrpc:"id"`
	IsInitialDemandEditable bool      `xmlrpc:"is_initial_demand_editable"`
	IsLocked                bool      `xmlrpc:"is_locked"`
	LastUpdate              time.Time `xmlrpc:"__last_update"`
	LocationDestId          Many2One  `xmlrpc:"location_dest_id"`
	LocationId              Many2One  `xmlrpc:"location_id"`
	LotId                   Many2One  `xmlrpc:"lot_id"`
	LotName                 string    `xmlrpc:"lot_name"`
	LotsVisible             bool      `xmlrpc:"lots_visible"`
	MoveId                  Many2One  `xmlrpc:"move_id"`
	OrderedQty              float64   `xmlrpc:"ordered_qty"`
	OwnerId                 Many2One  `xmlrpc:"owner_id"`
	PackageId               Many2One  `xmlrpc:"package_id"`
	PickingId               Many2One  `xmlrpc:"picking_id"`
	ProduceLineIds          []int64   `xmlrpc:"produce_line_ids"`
	ProductId               Many2One  `xmlrpc:"product_id"`
	ProductQty              float64   `xmlrpc:"product_qty"`
	ProductUomId            Many2One  `xmlrpc:"product_uom_id"`
	ProductUomQty           float64   `xmlrpc:"product_uom_qty"`
	QtyDone                 float64   `xmlrpc:"qty_done"`
	Reference               string    `xmlrpc:"reference"`
	ResultPackageId         Many2One  `xmlrpc:"result_package_id"`
	State                   string    `xmlrpc:"state"`
	ToLoc                   string    `xmlrpc:"to_loc"`
	WriteDate               time.Time `xmlrpc:"write_date"`
	WriteUid                Many2One  `xmlrpc:"write_uid"`
}

type StockMoveLineNil struct {
	ConsumeLineIds          interface{} `xmlrpc:"consume_line_ids"`
	CreateDate              interface{} `xmlrpc:"create_date"`
	CreateUid               interface{} `xmlrpc:"create_uid"`
	Date                    interface{} `xmlrpc:"date"`
	DisplayName             interface{} `xmlrpc:"display_name"`
	FromLoc                 interface{} `xmlrpc:"from_loc"`
	Id                      interface{} `xmlrpc:"id"`
	IsInitialDemandEditable bool        `xmlrpc:"is_initial_demand_editable"`
	IsLocked                bool        `xmlrpc:"is_locked"`
	LastUpdate              interface{} `xmlrpc:"__last_update"`
	LocationDestId          interface{} `xmlrpc:"location_dest_id"`
	LocationId              interface{} `xmlrpc:"location_id"`
	LotId                   interface{} `xmlrpc:"lot_id"`
	LotName                 interface{} `xmlrpc:"lot_name"`
	LotsVisible             bool        `xmlrpc:"lots_visible"`
	MoveId                  interface{} `xmlrpc:"move_id"`
	OrderedQty              interface{} `xmlrpc:"ordered_qty"`
	OwnerId                 interface{} `xmlrpc:"owner_id"`
	PackageId               interface{} `xmlrpc:"package_id"`
	PickingId               interface{} `xmlrpc:"picking_id"`
	ProduceLineIds          interface{} `xmlrpc:"produce_line_ids"`
	ProductId               interface{} `xmlrpc:"product_id"`
	ProductQty              interface{} `xmlrpc:"product_qty"`
	ProductUomId            interface{} `xmlrpc:"product_uom_id"`
	ProductUomQty           interface{} `xmlrpc:"product_uom_qty"`
	QtyDone                 interface{} `xmlrpc:"qty_done"`
	Reference               interface{} `xmlrpc:"reference"`
	ResultPackageId         interface{} `xmlrpc:"result_package_id"`
	State                   interface{} `xmlrpc:"state"`
	ToLoc                   interface{} `xmlrpc:"to_loc"`
	WriteDate               interface{} `xmlrpc:"write_date"`
	WriteUid                interface{} `xmlrpc:"write_uid"`
}

var StockMoveLineModel string = "stock.move.line"

type StockMoveLines []StockMoveLine

type StockMoveLinesNil []StockMoveLineNil

func (s *StockMoveLine) NilableType_() interface{} {
	return &StockMoveLineNil{}
}

func (ns *StockMoveLineNil) Type_() interface{} {
	s := &StockMoveLine{}
	return load(ns, s)
}

func (s *StockMoveLines) NilableType_() interface{} {
	return &StockMoveLinesNil{}
}

func (ns *StockMoveLinesNil) Type_() interface{} {
	s := &StockMoveLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockMoveLine))
	}
	return s
}
