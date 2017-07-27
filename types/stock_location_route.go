package types

import (
	"time"
)

type StockLocationRoute struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	Active                 bool      `xmlrpc:"active"`
	CategIds               []int64   `xmlrpc:"categ_ids"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	DisplayName            string    `xmlrpc:"display_name"`
	Id                     int64     `xmlrpc:"id"`
	Name                   string    `xmlrpc:"name"`
	ProductCategSelectable bool      `xmlrpc:"product_categ_selectable"`
	ProductIds             []int64   `xmlrpc:"product_ids"`
	ProductSelectable      bool      `xmlrpc:"product_selectable"`
	PullIds                []int64   `xmlrpc:"pull_ids"`
	PushIds                []int64   `xmlrpc:"push_ids"`
	SaleSelectable         bool      `xmlrpc:"sale_selectable"`
	Sequence               int64     `xmlrpc:"sequence"`
	SuppliedWhId           Many2One  `xmlrpc:"supplied_wh_id"`
	SupplierWhId           Many2One  `xmlrpc:"supplier_wh_id"`
	WarehouseIds           []int64   `xmlrpc:"warehouse_ids"`
	WarehouseSelectable    bool      `xmlrpc:"warehouse_selectable"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type StockLocationRouteNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Active                 bool        `xmlrpc:"active"`
	CategIds               interface{} `xmlrpc:"categ_ids"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	Id                     interface{} `xmlrpc:"id"`
	Name                   interface{} `xmlrpc:"name"`
	ProductCategSelectable bool        `xmlrpc:"product_categ_selectable"`
	ProductIds             interface{} `xmlrpc:"product_ids"`
	ProductSelectable      bool        `xmlrpc:"product_selectable"`
	PullIds                interface{} `xmlrpc:"pull_ids"`
	PushIds                interface{} `xmlrpc:"push_ids"`
	SaleSelectable         bool        `xmlrpc:"sale_selectable"`
	Sequence               interface{} `xmlrpc:"sequence"`
	SuppliedWhId           interface{} `xmlrpc:"supplied_wh_id"`
	SupplierWhId           interface{} `xmlrpc:"supplier_wh_id"`
	WarehouseIds           interface{} `xmlrpc:"warehouse_ids"`
	WarehouseSelectable    bool        `xmlrpc:"warehouse_selectable"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var StockLocationRouteModel string = "stock.location.route"

type StockLocationRoutes []StockLocationRoute

type StockLocationRoutesNil []StockLocationRouteNil

func (s *StockLocationRoute) NilableType_() interface{} {
	return &StockLocationRouteNil{}
}

func (ns *StockLocationRouteNil) Type_() interface{} {
	s := &StockLocationRoute{}
	return load(ns, s)
}

func (s *StockLocationRoutes) NilableType_() interface{} {
	return &StockLocationRoutesNil{}
}

func (ns *StockLocationRoutesNil) Type_() interface{} {
	s := &StockLocationRoutes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockLocationRoute))
	}
	return s
}
