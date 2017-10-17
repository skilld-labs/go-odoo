package types

import (
	"time"
)

type StockQuantPackage struct {
	CompanyId                 Many2One  `xmlrpc:"company_id"`
	CreateDate                time.Time `xmlrpc:"create_date"`
	CreateUid                 Many2One  `xmlrpc:"create_uid"`
	CurrentPickingId          bool      `xmlrpc:"current_picking_id"`
	CurrentPickingMoveLineIds []int64   `xmlrpc:"current_picking_move_line_ids"`
	DisplayName               string    `xmlrpc:"display_name"`
	Id                        int64     `xmlrpc:"id"`
	LastUpdate                time.Time `xmlrpc:"__last_update"`
	LocationId                Many2One  `xmlrpc:"location_id"`
	MoveLineIds               []int64   `xmlrpc:"move_line_ids"`
	Name                      string    `xmlrpc:"name"`
	OwnerId                   Many2One  `xmlrpc:"owner_id"`
	PackagingId               Many2One  `xmlrpc:"packaging_id"`
	QuantIds                  []int64   `xmlrpc:"quant_ids"`
	WriteDate                 time.Time `xmlrpc:"write_date"`
	WriteUid                  Many2One  `xmlrpc:"write_uid"`
}

type StockQuantPackageNil struct {
	CompanyId                 interface{} `xmlrpc:"company_id"`
	CreateDate                interface{} `xmlrpc:"create_date"`
	CreateUid                 interface{} `xmlrpc:"create_uid"`
	CurrentPickingId          bool        `xmlrpc:"current_picking_id"`
	CurrentPickingMoveLineIds interface{} `xmlrpc:"current_picking_move_line_ids"`
	DisplayName               interface{} `xmlrpc:"display_name"`
	Id                        interface{} `xmlrpc:"id"`
	LastUpdate                interface{} `xmlrpc:"__last_update"`
	LocationId                interface{} `xmlrpc:"location_id"`
	MoveLineIds               interface{} `xmlrpc:"move_line_ids"`
	Name                      interface{} `xmlrpc:"name"`
	OwnerId                   interface{} `xmlrpc:"owner_id"`
	PackagingId               interface{} `xmlrpc:"packaging_id"`
	QuantIds                  interface{} `xmlrpc:"quant_ids"`
	WriteDate                 interface{} `xmlrpc:"write_date"`
	WriteUid                  interface{} `xmlrpc:"write_uid"`
}

var StockQuantPackageModel string = "stock.quant.package"

type StockQuantPackages []StockQuantPackage

type StockQuantPackagesNil []StockQuantPackageNil

func (s *StockQuantPackage) NilableType_() interface{} {
	return &StockQuantPackageNil{}
}

func (ns *StockQuantPackageNil) Type_() interface{} {
	s := &StockQuantPackage{}
	return load(ns, s)
}

func (s *StockQuantPackages) NilableType_() interface{} {
	return &StockQuantPackagesNil{}
}

func (ns *StockQuantPackagesNil) Type_() interface{} {
	s := &StockQuantPackages{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockQuantPackage))
	}
	return s
}
