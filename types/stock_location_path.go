package types

import (
	"time"
)

type StockLocationPath struct {
	LastUpdate     time.Time `xmlrpc:"__last_update"`
	Active         bool      `xmlrpc:"active"`
	Auto           string    `xmlrpc:"auto"`
	CompanyId      Many2One  `xmlrpc:"company_id"`
	CreateDate     time.Time `xmlrpc:"create_date"`
	CreateUid      Many2One  `xmlrpc:"create_uid"`
	Delay          int64     `xmlrpc:"delay"`
	DisplayName    string    `xmlrpc:"display_name"`
	Id             int64     `xmlrpc:"id"`
	LocationDestId Many2One  `xmlrpc:"location_dest_id"`
	LocationFromId Many2One  `xmlrpc:"location_from_id"`
	Name           string    `xmlrpc:"name"`
	PickingTypeId  Many2One  `xmlrpc:"picking_type_id"`
	Propagate      bool      `xmlrpc:"propagate"`
	RouteId        Many2One  `xmlrpc:"route_id"`
	RouteSequence  int64     `xmlrpc:"route_sequence"`
	Sequence       int64     `xmlrpc:"sequence"`
	WarehouseId    Many2One  `xmlrpc:"warehouse_id"`
	WriteDate      time.Time `xmlrpc:"write_date"`
	WriteUid       Many2One  `xmlrpc:"write_uid"`
}

type StockLocationPathNil struct {
	LastUpdate     interface{} `xmlrpc:"__last_update"`
	Active         bool        `xmlrpc:"active"`
	Auto           interface{} `xmlrpc:"auto"`
	CompanyId      interface{} `xmlrpc:"company_id"`
	CreateDate     interface{} `xmlrpc:"create_date"`
	CreateUid      interface{} `xmlrpc:"create_uid"`
	Delay          interface{} `xmlrpc:"delay"`
	DisplayName    interface{} `xmlrpc:"display_name"`
	Id             interface{} `xmlrpc:"id"`
	LocationDestId interface{} `xmlrpc:"location_dest_id"`
	LocationFromId interface{} `xmlrpc:"location_from_id"`
	Name           interface{} `xmlrpc:"name"`
	PickingTypeId  interface{} `xmlrpc:"picking_type_id"`
	Propagate      bool        `xmlrpc:"propagate"`
	RouteId        interface{} `xmlrpc:"route_id"`
	RouteSequence  interface{} `xmlrpc:"route_sequence"`
	Sequence       interface{} `xmlrpc:"sequence"`
	WarehouseId    interface{} `xmlrpc:"warehouse_id"`
	WriteDate      interface{} `xmlrpc:"write_date"`
	WriteUid       interface{} `xmlrpc:"write_uid"`
}

var StockLocationPathModel string = "stock.location.path"

type StockLocationPaths []StockLocationPath

type StockLocationPathsNil []StockLocationPathNil

func (s *StockLocationPath) NilableType_() interface{} {
	return &StockLocationPathNil{}
}

func (ns *StockLocationPathNil) Type_() interface{} {
	s := &StockLocationPath{}
	return load(ns, s)
}

func (s *StockLocationPaths) NilableType_() interface{} {
	return &StockLocationPathsNil{}
}

func (ns *StockLocationPathsNil) Type_() interface{} {
	s := &StockLocationPaths{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockLocationPath))
	}
	return s
}
