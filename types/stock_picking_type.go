package types

import (
	"time"
)

type StockPickingType struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	Active                 bool      `xmlrpc:"active"`
	BarcodeNomenclatureId  Many2One  `xmlrpc:"barcode_nomenclature_id"`
	Code                   string    `xmlrpc:"code"`
	Color                  int64     `xmlrpc:"color"`
	CountPicking           int64     `xmlrpc:"count_picking"`
	CountPickingBackorders int64     `xmlrpc:"count_picking_backorders"`
	CountPickingDraft      int64     `xmlrpc:"count_picking_draft"`
	CountPickingLate       int64     `xmlrpc:"count_picking_late"`
	CountPickingReady      int64     `xmlrpc:"count_picking_ready"`
	CountPickingWaiting    int64     `xmlrpc:"count_picking_waiting"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	DefaultLocationDestId  Many2One  `xmlrpc:"default_location_dest_id"`
	DefaultLocationSrcId   Many2One  `xmlrpc:"default_location_src_id"`
	DisplayName            string    `xmlrpc:"display_name"`
	Id                     int64     `xmlrpc:"id"`
	LastDonePicking        string    `xmlrpc:"last_done_picking"`
	Name                   string    `xmlrpc:"name"`
	RatePickingBackorders  int64     `xmlrpc:"rate_picking_backorders"`
	RatePickingLate        int64     `xmlrpc:"rate_picking_late"`
	ReturnPickingTypeId    Many2One  `xmlrpc:"return_picking_type_id"`
	Sequence               int64     `xmlrpc:"sequence"`
	SequenceId             Many2One  `xmlrpc:"sequence_id"`
	ShowEntirePacks        bool      `xmlrpc:"show_entire_packs"`
	UseCreateLots          bool      `xmlrpc:"use_create_lots"`
	UseExistingLots        bool      `xmlrpc:"use_existing_lots"`
	WarehouseId            Many2One  `xmlrpc:"warehouse_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type StockPickingTypeNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Active                 bool        `xmlrpc:"active"`
	BarcodeNomenclatureId  interface{} `xmlrpc:"barcode_nomenclature_id"`
	Code                   interface{} `xmlrpc:"code"`
	Color                  interface{} `xmlrpc:"color"`
	CountPicking           interface{} `xmlrpc:"count_picking"`
	CountPickingBackorders interface{} `xmlrpc:"count_picking_backorders"`
	CountPickingDraft      interface{} `xmlrpc:"count_picking_draft"`
	CountPickingLate       interface{} `xmlrpc:"count_picking_late"`
	CountPickingReady      interface{} `xmlrpc:"count_picking_ready"`
	CountPickingWaiting    interface{} `xmlrpc:"count_picking_waiting"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	DefaultLocationDestId  interface{} `xmlrpc:"default_location_dest_id"`
	DefaultLocationSrcId   interface{} `xmlrpc:"default_location_src_id"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	Id                     interface{} `xmlrpc:"id"`
	LastDonePicking        interface{} `xmlrpc:"last_done_picking"`
	Name                   interface{} `xmlrpc:"name"`
	RatePickingBackorders  interface{} `xmlrpc:"rate_picking_backorders"`
	RatePickingLate        interface{} `xmlrpc:"rate_picking_late"`
	ReturnPickingTypeId    interface{} `xmlrpc:"return_picking_type_id"`
	Sequence               interface{} `xmlrpc:"sequence"`
	SequenceId             interface{} `xmlrpc:"sequence_id"`
	ShowEntirePacks        bool        `xmlrpc:"show_entire_packs"`
	UseCreateLots          bool        `xmlrpc:"use_create_lots"`
	UseExistingLots        bool        `xmlrpc:"use_existing_lots"`
	WarehouseId            interface{} `xmlrpc:"warehouse_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var StockPickingTypeModel string = "stock.picking.type"

type StockPickingTypes []StockPickingType

type StockPickingTypesNil []StockPickingTypeNil

func (s *StockPickingType) NilableType_() interface{} {
	return &StockPickingTypeNil{}
}

func (ns *StockPickingTypeNil) Type_() interface{} {
	s := &StockPickingType{}
	return load(ns, s)
}

func (s *StockPickingTypes) NilableType_() interface{} {
	return &StockPickingTypesNil{}
}

func (ns *StockPickingTypesNil) Type_() interface{} {
	s := &StockPickingTypes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockPickingType))
	}
	return s
}
