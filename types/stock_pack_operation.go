package types

import (
	"time"
)

type StockPackOperation struct {
	LastUpdate                   time.Time `xmlrpc:"__last_update"`
	CreateDate                   time.Time `xmlrpc:"create_date"`
	CreateUid                    Many2One  `xmlrpc:"create_uid"`
	Date                         time.Time `xmlrpc:"date"`
	DisplayName                  string    `xmlrpc:"display_name"`
	FreshRecord                  bool      `xmlrpc:"fresh_record"`
	FromLoc                      string    `xmlrpc:"from_loc"`
	Id                           int64     `xmlrpc:"id"`
	IsDone                       bool      `xmlrpc:"is_done"`
	LinkedMoveOperationIds       []int64   `xmlrpc:"linked_move_operation_ids"`
	LocationDestId               Many2One  `xmlrpc:"location_dest_id"`
	LocationId                   Many2One  `xmlrpc:"location_id"`
	LotsVisible                  bool      `xmlrpc:"lots_visible"`
	OrderedQty                   float64   `xmlrpc:"ordered_qty"`
	OwnerId                      Many2One  `xmlrpc:"owner_id"`
	PackLotIds                   []int64   `xmlrpc:"pack_lot_ids"`
	PackageId                    Many2One  `xmlrpc:"package_id"`
	PickingDestinationLocationId Many2One  `xmlrpc:"picking_destination_location_id"`
	PickingId                    Many2One  `xmlrpc:"picking_id"`
	PickingSourceLocationId      Many2One  `xmlrpc:"picking_source_location_id"`
	ProductId                    Many2One  `xmlrpc:"product_id"`
	ProductQty                   float64   `xmlrpc:"product_qty"`
	ProductUomId                 Many2One  `xmlrpc:"product_uom_id"`
	QtyDone                      float64   `xmlrpc:"qty_done"`
	RemainingQty                 float64   `xmlrpc:"remaining_qty"`
	ResultPackageId              Many2One  `xmlrpc:"result_package_id"`
	State                        string    `xmlrpc:"state"`
	ToLoc                        string    `xmlrpc:"to_loc"`
	WriteDate                    time.Time `xmlrpc:"write_date"`
	WriteUid                     Many2One  `xmlrpc:"write_uid"`
}

type StockPackOperationNil struct {
	LastUpdate                   interface{} `xmlrpc:"__last_update"`
	CreateDate                   interface{} `xmlrpc:"create_date"`
	CreateUid                    interface{} `xmlrpc:"create_uid"`
	Date                         interface{} `xmlrpc:"date"`
	DisplayName                  interface{} `xmlrpc:"display_name"`
	FreshRecord                  bool        `xmlrpc:"fresh_record"`
	FromLoc                      interface{} `xmlrpc:"from_loc"`
	Id                           interface{} `xmlrpc:"id"`
	IsDone                       bool        `xmlrpc:"is_done"`
	LinkedMoveOperationIds       interface{} `xmlrpc:"linked_move_operation_ids"`
	LocationDestId               interface{} `xmlrpc:"location_dest_id"`
	LocationId                   interface{} `xmlrpc:"location_id"`
	LotsVisible                  bool        `xmlrpc:"lots_visible"`
	OrderedQty                   interface{} `xmlrpc:"ordered_qty"`
	OwnerId                      interface{} `xmlrpc:"owner_id"`
	PackLotIds                   interface{} `xmlrpc:"pack_lot_ids"`
	PackageId                    interface{} `xmlrpc:"package_id"`
	PickingDestinationLocationId interface{} `xmlrpc:"picking_destination_location_id"`
	PickingId                    interface{} `xmlrpc:"picking_id"`
	PickingSourceLocationId      interface{} `xmlrpc:"picking_source_location_id"`
	ProductId                    interface{} `xmlrpc:"product_id"`
	ProductQty                   interface{} `xmlrpc:"product_qty"`
	ProductUomId                 interface{} `xmlrpc:"product_uom_id"`
	QtyDone                      interface{} `xmlrpc:"qty_done"`
	RemainingQty                 interface{} `xmlrpc:"remaining_qty"`
	ResultPackageId              interface{} `xmlrpc:"result_package_id"`
	State                        interface{} `xmlrpc:"state"`
	ToLoc                        interface{} `xmlrpc:"to_loc"`
	WriteDate                    interface{} `xmlrpc:"write_date"`
	WriteUid                     interface{} `xmlrpc:"write_uid"`
}

var StockPackOperationModel string = "stock.pack.operation"

type StockPackOperations []StockPackOperation

type StockPackOperationsNil []StockPackOperationNil

func (s *StockPackOperation) NilableType_() interface{} {
	return &StockPackOperationNil{}
}

func (ns *StockPackOperationNil) Type_() interface{} {
	s := &StockPackOperation{}
	return load(ns, s)
}

func (s *StockPackOperations) NilableType_() interface{} {
	return &StockPackOperationsNil{}
}

func (ns *StockPackOperationsNil) Type_() interface{} {
	s := &StockPackOperations{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockPackOperation))
	}
	return s
}
