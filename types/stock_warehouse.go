package types

import (
	"time"
)

type StockWarehouse struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	Active              bool      `xmlrpc:"active"`
	BuyPullId           Many2One  `xmlrpc:"buy_pull_id"`
	BuyToResupply       bool      `xmlrpc:"buy_to_resupply"`
	Code                string    `xmlrpc:"code"`
	CompanyId           Many2One  `xmlrpc:"company_id"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	CreateUid           Many2One  `xmlrpc:"create_uid"`
	CrossdockRouteId    Many2One  `xmlrpc:"crossdock_route_id"`
	DefaultResupplyWhId Many2One  `xmlrpc:"default_resupply_wh_id"`
	DeliveryRouteId     Many2One  `xmlrpc:"delivery_route_id"`
	DeliverySteps       string    `xmlrpc:"delivery_steps"`
	DisplayName         string    `xmlrpc:"display_name"`
	Id                  int64     `xmlrpc:"id"`
	InTypeId            Many2One  `xmlrpc:"in_type_id"`
	IntTypeId           Many2One  `xmlrpc:"int_type_id"`
	LotStockId          Many2One  `xmlrpc:"lot_stock_id"`
	MtoPullId           Many2One  `xmlrpc:"mto_pull_id"`
	Name                string    `xmlrpc:"name"`
	OutTypeId           Many2One  `xmlrpc:"out_type_id"`
	PackTypeId          Many2One  `xmlrpc:"pack_type_id"`
	PartnerId           Many2One  `xmlrpc:"partner_id"`
	PickTypeId          Many2One  `xmlrpc:"pick_type_id"`
	ReceptionRouteId    Many2One  `xmlrpc:"reception_route_id"`
	ReceptionSteps      string    `xmlrpc:"reception_steps"`
	ResupplyRouteIds    []int64   `xmlrpc:"resupply_route_ids"`
	ResupplyWhIds       []int64   `xmlrpc:"resupply_wh_ids"`
	RouteIds            []int64   `xmlrpc:"route_ids"`
	ViewLocationId      Many2One  `xmlrpc:"view_location_id"`
	WhInputStockLocId   Many2One  `xmlrpc:"wh_input_stock_loc_id"`
	WhOutputStockLocId  Many2One  `xmlrpc:"wh_output_stock_loc_id"`
	WhPackStockLocId    Many2One  `xmlrpc:"wh_pack_stock_loc_id"`
	WhQcStockLocId      Many2One  `xmlrpc:"wh_qc_stock_loc_id"`
	WriteDate           time.Time `xmlrpc:"write_date"`
	WriteUid            Many2One  `xmlrpc:"write_uid"`
}

type StockWarehouseNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	Active              bool        `xmlrpc:"active"`
	BuyPullId           interface{} `xmlrpc:"buy_pull_id"`
	BuyToResupply       bool        `xmlrpc:"buy_to_resupply"`
	Code                interface{} `xmlrpc:"code"`
	CompanyId           interface{} `xmlrpc:"company_id"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	CreateUid           interface{} `xmlrpc:"create_uid"`
	CrossdockRouteId    interface{} `xmlrpc:"crossdock_route_id"`
	DefaultResupplyWhId interface{} `xmlrpc:"default_resupply_wh_id"`
	DeliveryRouteId     interface{} `xmlrpc:"delivery_route_id"`
	DeliverySteps       interface{} `xmlrpc:"delivery_steps"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	Id                  interface{} `xmlrpc:"id"`
	InTypeId            interface{} `xmlrpc:"in_type_id"`
	IntTypeId           interface{} `xmlrpc:"int_type_id"`
	LotStockId          interface{} `xmlrpc:"lot_stock_id"`
	MtoPullId           interface{} `xmlrpc:"mto_pull_id"`
	Name                interface{} `xmlrpc:"name"`
	OutTypeId           interface{} `xmlrpc:"out_type_id"`
	PackTypeId          interface{} `xmlrpc:"pack_type_id"`
	PartnerId           interface{} `xmlrpc:"partner_id"`
	PickTypeId          interface{} `xmlrpc:"pick_type_id"`
	ReceptionRouteId    interface{} `xmlrpc:"reception_route_id"`
	ReceptionSteps      interface{} `xmlrpc:"reception_steps"`
	ResupplyRouteIds    interface{} `xmlrpc:"resupply_route_ids"`
	ResupplyWhIds       interface{} `xmlrpc:"resupply_wh_ids"`
	RouteIds            interface{} `xmlrpc:"route_ids"`
	ViewLocationId      interface{} `xmlrpc:"view_location_id"`
	WhInputStockLocId   interface{} `xmlrpc:"wh_input_stock_loc_id"`
	WhOutputStockLocId  interface{} `xmlrpc:"wh_output_stock_loc_id"`
	WhPackStockLocId    interface{} `xmlrpc:"wh_pack_stock_loc_id"`
	WhQcStockLocId      interface{} `xmlrpc:"wh_qc_stock_loc_id"`
	WriteDate           interface{} `xmlrpc:"write_date"`
	WriteUid            interface{} `xmlrpc:"write_uid"`
}

var StockWarehouseModel string = "stock.warehouse"

type StockWarehouses []StockWarehouse

type StockWarehousesNil []StockWarehouseNil

func (s *StockWarehouse) NilableType_() interface{} {
	return &StockWarehouseNil{}
}

func (ns *StockWarehouseNil) Type_() interface{} {
	s := &StockWarehouse{}
	return load(ns, s)
}

func (s *StockWarehouses) NilableType_() interface{} {
	return &StockWarehousesNil{}
}

func (ns *StockWarehousesNil) Type_() interface{} {
	s := &StockWarehouses{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockWarehouse))
	}
	return s
}
