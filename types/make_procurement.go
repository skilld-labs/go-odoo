package types

import (
	"time"
)

type MakeProcurement struct {
	LastUpdate          time.Time `xmlrpc:"__last_update"`
	CreateDate          time.Time `xmlrpc:"create_date"`
	CreateUid           Many2One  `xmlrpc:"create_uid"`
	DatePlanned         time.Time `xmlrpc:"date_planned"`
	DisplayName         string    `xmlrpc:"display_name"`
	Id                  int64     `xmlrpc:"id"`
	ProductId           Many2One  `xmlrpc:"product_id"`
	ProductTmplId       Many2One  `xmlrpc:"product_tmpl_id"`
	ProductVariantCount int64     `xmlrpc:"product_variant_count"`
	Qty                 float64   `xmlrpc:"qty"`
	ResModel            string    `xmlrpc:"res_model"`
	RouteIds            []int64   `xmlrpc:"route_ids"`
	UomId               Many2One  `xmlrpc:"uom_id"`
	WarehouseId         Many2One  `xmlrpc:"warehouse_id"`
	WriteDate           time.Time `xmlrpc:"write_date"`
	WriteUid            Many2One  `xmlrpc:"write_uid"`
}

type MakeProcurementNil struct {
	LastUpdate          interface{} `xmlrpc:"__last_update"`
	CreateDate          interface{} `xmlrpc:"create_date"`
	CreateUid           interface{} `xmlrpc:"create_uid"`
	DatePlanned         interface{} `xmlrpc:"date_planned"`
	DisplayName         interface{} `xmlrpc:"display_name"`
	Id                  interface{} `xmlrpc:"id"`
	ProductId           interface{} `xmlrpc:"product_id"`
	ProductTmplId       interface{} `xmlrpc:"product_tmpl_id"`
	ProductVariantCount interface{} `xmlrpc:"product_variant_count"`
	Qty                 interface{} `xmlrpc:"qty"`
	ResModel            interface{} `xmlrpc:"res_model"`
	RouteIds            interface{} `xmlrpc:"route_ids"`
	UomId               interface{} `xmlrpc:"uom_id"`
	WarehouseId         interface{} `xmlrpc:"warehouse_id"`
	WriteDate           interface{} `xmlrpc:"write_date"`
	WriteUid            interface{} `xmlrpc:"write_uid"`
}

var MakeProcurementModel string = "make.procurement"

type MakeProcurements []MakeProcurement

type MakeProcurementsNil []MakeProcurementNil

func (s *MakeProcurement) NilableType_() interface{} {
	return &MakeProcurementNil{}
}

func (ns *MakeProcurementNil) Type_() interface{} {
	s := &MakeProcurement{}
	return load(ns, s)
}

func (s *MakeProcurements) NilableType_() interface{} {
	return &MakeProcurementsNil{}
}

func (ns *MakeProcurementsNil) Type_() interface{} {
	s := &MakeProcurements{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*MakeProcurement))
	}
	return s
}
