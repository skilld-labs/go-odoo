package types

import (
	"time"
)

type StockWarehouseOrderpoint struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Active         bool      `xmlrpc:"active"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	DisplayName    string    `xmlrpc:"display_name"`
	GroupId        Many2One  `xmlrpc:"group_id"`
	Id             int64     `xmlrpc:"id"`
	LeadDays       int64     `xmlrpc:"lead_days"`
	LeadType       string    `xmlrpc:"lead_type"`
	LocationId     Many2One  `xmlrpc:"location_id"`
	Name           string    `xmlrpc:"name"`
	ProcurementIds []int64   `xmlrpc:"procurement_ids"`
	ProductId      Many2One  `xmlrpc:"product_id"`
	ProductMaxQty  float64   `xmlrpc:"product_max_qty"`
	ProductMinQty  float64   `xmlrpc:"product_min_qty"`
	ProductUom     Many2One  `xmlrpc:"product_uom"`
	QtyMultiple    float64   `xmlrpc:"qty_multiple"`
	WarehouseId    Many2One  `xmlrpc:"warehouse_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type StockWarehouseOrderpointNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Active         bool        `xmlrpc:"active"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	GroupId        interface{} `xmlrpc:"group_id"`
	Id             interface{} `xmlrpc:"id"`
	LeadDays       interface{} `xmlrpc:"lead_days"`
	LeadType       interface{} `xmlrpc:"lead_type"`
	LocationId     interface{} `xmlrpc:"location_id"`
	Name           interface{} `xmlrpc:"name"`
	ProcurementIds interface{} `xmlrpc:"procurement_ids"`
	ProductId      interface{} `xmlrpc:"product_id"`
	ProductMaxQty  interface{} `xmlrpc:"product_max_qty"`
	ProductMinQty  interface{} `xmlrpc:"product_min_qty"`
	ProductUom     interface{} `xmlrpc:"product_uom"`
	QtyMultiple    interface{} `xmlrpc:"qty_multiple"`
	WarehouseId    interface{} `xmlrpc:"warehouse_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var StockWarehouseOrderpointModel string = "stock.warehouse.orderpoint"

type StockWarehouseOrderpoints []StockWarehouseOrderpoint

type StockWarehouseOrderpointsNil []StockWarehouseOrderpointNil

func (s *StockWarehouseOrderpoint) NilableType_() interface{} {
	return &StockWarehouseOrderpointNil{}
}

func (ns *StockWarehouseOrderpointNil) Type_() interface{} {
	s := &StockWarehouseOrderpoint{}
	return load(ns, s)
}

func (s *StockWarehouseOrderpoints) NilableType_() interface{} {
	return &StockWarehouseOrderpointsNil{}
}

func (ns *StockWarehouseOrderpointsNil) Type_() interface{} {
	s := &StockWarehouseOrderpoints{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockWarehouseOrderpoint))
	}
	return s
}
