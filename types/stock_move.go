package types

import (
	"time"
)

type StockMove struct {
	AccountMoveIds           []int64     `xmlrpc:"account_move_ids"`
	Additional               bool        `xmlrpc:"additional"`
	Availability             float64     `xmlrpc:"availability"`
	BackorderId              Many2One    `xmlrpc:"backorder_id"`
	CompanyId                Many2One    `xmlrpc:"company_id"`
	CreateDate               time.Time   `xmlrpc:"create_date"`
	CreatedPurchaseLineId    Many2One    `xmlrpc:"created_purchase_line_id"`
	CreateUid                Many2One    `xmlrpc:"create_uid"`
	Date                     time.Time   `xmlrpc:"date"`
	DateExpected             time.Time   `xmlrpc:"date_expected"`
	DisplayName              string      `xmlrpc:"display_name"`
	GroupId                  Many2One    `xmlrpc:"group_id"`
	HasTracking              interface{} `xmlrpc:"has_tracking"`
	Id                       int64       `xmlrpc:"id"`
	InventoryId              Many2One    `xmlrpc:"inventory_id"`
	IsInitialDemandEditable  bool        `xmlrpc:"is_initial_demand_editable"`
	IsLocked                 bool        `xmlrpc:"is_locked"`
	IsQuantityDoneEditable   bool        `xmlrpc:"is_quantity_done_editable"`
	LastUpdate               time.Time   `xmlrpc:"__last_update"`
	LocationDestId           Many2One    `xmlrpc:"location_dest_id"`
	LocationId               Many2One    `xmlrpc:"location_id"`
	MoveDestIds              []int64     `xmlrpc:"move_dest_ids"`
	MoveLineIds              []int64     `xmlrpc:"move_line_ids"`
	MoveLineNosuggestIds     []int64     `xmlrpc:"move_line_nosuggest_ids"`
	MoveOrigIds              []int64     `xmlrpc:"move_orig_ids"`
	Name                     string      `xmlrpc:"name"`
	Note                     string      `xmlrpc:"note"`
	OrderedQty               float64     `xmlrpc:"ordered_qty"`
	Origin                   string      `xmlrpc:"origin"`
	OriginReturnedMoveId     Many2One    `xmlrpc:"origin_returned_move_id"`
	PartnerId                Many2One    `xmlrpc:"partner_id"`
	PickingCode              interface{} `xmlrpc:"picking_code"`
	PickingId                Many2One    `xmlrpc:"picking_id"`
	PickingPartnerId         Many2One    `xmlrpc:"picking_partner_id"`
	PickingTypeId            Many2One    `xmlrpc:"picking_type_id"`
	PriceUnit                float64     `xmlrpc:"price_unit"`
	Priority                 interface{} `xmlrpc:"priority"`
	ProcureMethod            interface{} `xmlrpc:"procure_method"`
	ProductId                Many2One    `xmlrpc:"product_id"`
	ProductPackaging         Many2One    `xmlrpc:"product_packaging"`
	ProductQty               float64     `xmlrpc:"product_qty"`
	ProductTmplId            Many2One    `xmlrpc:"product_tmpl_id"`
	ProductType              interface{} `xmlrpc:"product_type"`
	ProductUom               Many2One    `xmlrpc:"product_uom"`
	ProductUomQty            float64     `xmlrpc:"product_uom_qty"`
	Propagate                bool        `xmlrpc:"propagate"`
	PurchaseLineId           Many2One    `xmlrpc:"purchase_line_id"`
	PushRuleId               Many2One    `xmlrpc:"push_rule_id"`
	QuantityDone             float64     `xmlrpc:"quantity_done"`
	Reference                string      `xmlrpc:"reference"`
	RemainingQty             float64     `xmlrpc:"remaining_qty"`
	RemainingValue           float64     `xmlrpc:"remaining_value"`
	ReservedAvailability     float64     `xmlrpc:"reserved_availability"`
	RestrictPartnerId        Many2One    `xmlrpc:"restrict_partner_id"`
	ReturnedMoveIds          []int64     `xmlrpc:"returned_move_ids"`
	RouteIds                 []int64     `xmlrpc:"route_ids"`
	RuleId                   Many2One    `xmlrpc:"rule_id"`
	SaleLineId               Many2One    `xmlrpc:"sale_line_id"`
	ScrapIds                 []int64     `xmlrpc:"scrap_ids"`
	Scrapped                 bool        `xmlrpc:"scrapped"`
	Sequence                 int64       `xmlrpc:"sequence"`
	ShowDetailsVisible       bool        `xmlrpc:"show_details_visible"`
	ShowOperations           bool        `xmlrpc:"show_operations"`
	ShowReservedAvailability bool        `xmlrpc:"show_reserved_availability"`
	State                    interface{} `xmlrpc:"state"`
	StringAvailabilityInfo   string      `xmlrpc:"string_availability_info"`
	ToRefund                 bool        `xmlrpc:"to_refund"`
	Value                    float64     `xmlrpc:"value"`
	WarehouseId              Many2One    `xmlrpc:"warehouse_id"`
	WriteDate                time.Time   `xmlrpc:"write_date"`
	WriteUid                 Many2One    `xmlrpc:"write_uid"`
}

type StockMoveNil struct {
	AccountMoveIds           interface{} `xmlrpc:"account_move_ids"`
	Additional               bool        `xmlrpc:"additional"`
	Availability             interface{} `xmlrpc:"availability"`
	BackorderId              interface{} `xmlrpc:"backorder_id"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreatedPurchaseLineId    interface{} `xmlrpc:"created_purchase_line_id"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Date                     interface{} `xmlrpc:"date"`
	DateExpected             interface{} `xmlrpc:"date_expected"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	GroupId                  interface{} `xmlrpc:"group_id"`
	HasTracking              interface{} `xmlrpc:"has_tracking"`
	Id                       interface{} `xmlrpc:"id"`
	InventoryId              interface{} `xmlrpc:"inventory_id"`
	IsInitialDemandEditable  bool        `xmlrpc:"is_initial_demand_editable"`
	IsLocked                 bool        `xmlrpc:"is_locked"`
	IsQuantityDoneEditable   bool        `xmlrpc:"is_quantity_done_editable"`
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	LocationDestId           interface{} `xmlrpc:"location_dest_id"`
	LocationId               interface{} `xmlrpc:"location_id"`
	MoveDestIds              interface{} `xmlrpc:"move_dest_ids"`
	MoveLineIds              interface{} `xmlrpc:"move_line_ids"`
	MoveLineNosuggestIds     interface{} `xmlrpc:"move_line_nosuggest_ids"`
	MoveOrigIds              interface{} `xmlrpc:"move_orig_ids"`
	Name                     interface{} `xmlrpc:"name"`
	Note                     interface{} `xmlrpc:"note"`
	OrderedQty               interface{} `xmlrpc:"ordered_qty"`
	Origin                   interface{} `xmlrpc:"origin"`
	OriginReturnedMoveId     interface{} `xmlrpc:"origin_returned_move_id"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PickingCode              interface{} `xmlrpc:"picking_code"`
	PickingId                interface{} `xmlrpc:"picking_id"`
	PickingPartnerId         interface{} `xmlrpc:"picking_partner_id"`
	PickingTypeId            interface{} `xmlrpc:"picking_type_id"`
	PriceUnit                interface{} `xmlrpc:"price_unit"`
	Priority                 interface{} `xmlrpc:"priority"`
	ProcureMethod            interface{} `xmlrpc:"procure_method"`
	ProductId                interface{} `xmlrpc:"product_id"`
	ProductPackaging         interface{} `xmlrpc:"product_packaging"`
	ProductQty               interface{} `xmlrpc:"product_qty"`
	ProductTmplId            interface{} `xmlrpc:"product_tmpl_id"`
	ProductType              interface{} `xmlrpc:"product_type"`
	ProductUom               interface{} `xmlrpc:"product_uom"`
	ProductUomQty            interface{} `xmlrpc:"product_uom_qty"`
	Propagate                bool        `xmlrpc:"propagate"`
	PurchaseLineId           interface{} `xmlrpc:"purchase_line_id"`
	PushRuleId               interface{} `xmlrpc:"push_rule_id"`
	QuantityDone             interface{} `xmlrpc:"quantity_done"`
	Reference                interface{} `xmlrpc:"reference"`
	RemainingQty             interface{} `xmlrpc:"remaining_qty"`
	RemainingValue           interface{} `xmlrpc:"remaining_value"`
	ReservedAvailability     interface{} `xmlrpc:"reserved_availability"`
	RestrictPartnerId        interface{} `xmlrpc:"restrict_partner_id"`
	ReturnedMoveIds          interface{} `xmlrpc:"returned_move_ids"`
	RouteIds                 interface{} `xmlrpc:"route_ids"`
	RuleId                   interface{} `xmlrpc:"rule_id"`
	SaleLineId               interface{} `xmlrpc:"sale_line_id"`
	ScrapIds                 interface{} `xmlrpc:"scrap_ids"`
	Scrapped                 bool        `xmlrpc:"scrapped"`
	Sequence                 interface{} `xmlrpc:"sequence"`
	ShowDetailsVisible       bool        `xmlrpc:"show_details_visible"`
	ShowOperations           bool        `xmlrpc:"show_operations"`
	ShowReservedAvailability bool        `xmlrpc:"show_reserved_availability"`
	State                    interface{} `xmlrpc:"state"`
	StringAvailabilityInfo   interface{} `xmlrpc:"string_availability_info"`
	ToRefund                 bool        `xmlrpc:"to_refund"`
	Value                    interface{} `xmlrpc:"value"`
	WarehouseId              interface{} `xmlrpc:"warehouse_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var StockMoveModel string = "stock.move"

type StockMoves []StockMove

type StockMovesNil []StockMoveNil

func (s *StockMove) NilableType_() interface{} {
	return &StockMoveNil{}
}

func (ns *StockMoveNil) Type_() interface{} {
	s := &StockMove{}
	return load(ns, s)
}

func (s *StockMoves) NilableType_() interface{} {
	return &StockMovesNil{}
}

func (ns *StockMovesNil) Type_() interface{} {
	s := &StockMoves{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockMove))
	}
	return s
}
