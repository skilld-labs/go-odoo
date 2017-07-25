package types

import (
	"time"
)

type StockMove struct {
	LastUpdate             time.Time `xmlrpc:"__last_update"`
	Availability           float64   `xmlrpc:"availability"`
	BackorderId            Many2One  `xmlrpc:"backorder_id"`
	CompanyId              Many2One  `xmlrpc:"company_id"`
	CreateDate             time.Time `xmlrpc:"create_date"`
	CreateUid              Many2One  `xmlrpc:"create_uid"`
	Date                   time.Time `xmlrpc:"date"`
	DateExpected           time.Time `xmlrpc:"date_expected"`
	DisplayName            string    `xmlrpc:"display_name"`
	GroupId                Many2One  `xmlrpc:"group_id"`
	Id                     int64     `xmlrpc:"id"`
	InventoryId            Many2One  `xmlrpc:"inventory_id"`
	LinkedMoveOperationIds []int64   `xmlrpc:"linked_move_operation_ids"`
	LocationDestId         Many2One  `xmlrpc:"location_dest_id"`
	LocationId             Many2One  `xmlrpc:"location_id"`
	LotIds                 []int64   `xmlrpc:"lot_ids"`
	MoveDestId             Many2One  `xmlrpc:"move_dest_id"`
	MoveOrigIds            []int64   `xmlrpc:"move_orig_ids"`
	Name                   string    `xmlrpc:"name"`
	Note                   string    `xmlrpc:"note"`
	OrderedQty             float64   `xmlrpc:"ordered_qty"`
	Origin                 string    `xmlrpc:"origin"`
	OriginReturnedMoveId   Many2One  `xmlrpc:"origin_returned_move_id"`
	PartiallyAvailable     bool      `xmlrpc:"partially_available"`
	PartnerId              Many2One  `xmlrpc:"partner_id"`
	PickingId              Many2One  `xmlrpc:"picking_id"`
	PickingPartnerId       Many2One  `xmlrpc:"picking_partner_id"`
	PickingTypeId          Many2One  `xmlrpc:"picking_type_id"`
	PriceUnit              float64   `xmlrpc:"price_unit"`
	Priority               string    `xmlrpc:"priority"`
	ProcureMethod          string    `xmlrpc:"procure_method"`
	ProcurementId          Many2One  `xmlrpc:"procurement_id"`
	ProductId              Many2One  `xmlrpc:"product_id"`
	ProductPackaging       Many2One  `xmlrpc:"product_packaging"`
	ProductQty             float64   `xmlrpc:"product_qty"`
	ProductTmplId          Many2One  `xmlrpc:"product_tmpl_id"`
	ProductUom             Many2One  `xmlrpc:"product_uom"`
	ProductUomQty          float64   `xmlrpc:"product_uom_qty"`
	Propagate              bool      `xmlrpc:"propagate"`
	PurchaseLineId         Many2One  `xmlrpc:"purchase_line_id"`
	PushRuleId             Many2One  `xmlrpc:"push_rule_id"`
	QuantIds               []int64   `xmlrpc:"quant_ids"`
	RemainingQty           float64   `xmlrpc:"remaining_qty"`
	ReservedAvailability   float64   `xmlrpc:"reserved_availability"`
	ReservedQuantIds       []int64   `xmlrpc:"reserved_quant_ids"`
	RestrictLotId          Many2One  `xmlrpc:"restrict_lot_id"`
	RestrictPartnerId      Many2One  `xmlrpc:"restrict_partner_id"`
	ReturnedMoveIds        []int64   `xmlrpc:"returned_move_ids"`
	RouteIds               []int64   `xmlrpc:"route_ids"`
	RuleId                 Many2One  `xmlrpc:"rule_id"`
	Scrapped               bool      `xmlrpc:"scrapped"`
	Sequence               int64     `xmlrpc:"sequence"`
	SplitFrom              Many2One  `xmlrpc:"split_from"`
	State                  string    `xmlrpc:"state"`
	StringAvailabilityInfo string    `xmlrpc:"string_availability_info"`
	ToRefundSo             bool      `xmlrpc:"to_refund_so"`
	WarehouseId            Many2One  `xmlrpc:"warehouse_id"`
	WriteDate              time.Time `xmlrpc:"write_date"`
	WriteUid               Many2One  `xmlrpc:"write_uid"`
}

type StockMoveNil struct {
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	Availability           interface{} `xmlrpc:"availability"`
	BackorderId            interface{} `xmlrpc:"backorder_id"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	Date                   interface{} `xmlrpc:"date"`
	DateExpected           interface{} `xmlrpc:"date_expected"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	GroupId                interface{} `xmlrpc:"group_id"`
	Id                     interface{} `xmlrpc:"id"`
	InventoryId            interface{} `xmlrpc:"inventory_id"`
	LinkedMoveOperationIds interface{} `xmlrpc:"linked_move_operation_ids"`
	LocationDestId         interface{} `xmlrpc:"location_dest_id"`
	LocationId             interface{} `xmlrpc:"location_id"`
	LotIds                 interface{} `xmlrpc:"lot_ids"`
	MoveDestId             interface{} `xmlrpc:"move_dest_id"`
	MoveOrigIds            interface{} `xmlrpc:"move_orig_ids"`
	Name                   interface{} `xmlrpc:"name"`
	Note                   interface{} `xmlrpc:"note"`
	OrderedQty             interface{} `xmlrpc:"ordered_qty"`
	Origin                 interface{} `xmlrpc:"origin"`
	OriginReturnedMoveId   interface{} `xmlrpc:"origin_returned_move_id"`
	PartiallyAvailable     bool        `xmlrpc:"partially_available"`
	PartnerId              interface{} `xmlrpc:"partner_id"`
	PickingId              interface{} `xmlrpc:"picking_id"`
	PickingPartnerId       interface{} `xmlrpc:"picking_partner_id"`
	PickingTypeId          interface{} `xmlrpc:"picking_type_id"`
	PriceUnit              interface{} `xmlrpc:"price_unit"`
	Priority               interface{} `xmlrpc:"priority"`
	ProcureMethod          interface{} `xmlrpc:"procure_method"`
	ProcurementId          interface{} `xmlrpc:"procurement_id"`
	ProductId              interface{} `xmlrpc:"product_id"`
	ProductPackaging       interface{} `xmlrpc:"product_packaging"`
	ProductQty             interface{} `xmlrpc:"product_qty"`
	ProductTmplId          interface{} `xmlrpc:"product_tmpl_id"`
	ProductUom             interface{} `xmlrpc:"product_uom"`
	ProductUomQty          interface{} `xmlrpc:"product_uom_qty"`
	Propagate              bool        `xmlrpc:"propagate"`
	PurchaseLineId         interface{} `xmlrpc:"purchase_line_id"`
	PushRuleId             interface{} `xmlrpc:"push_rule_id"`
	QuantIds               interface{} `xmlrpc:"quant_ids"`
	RemainingQty           interface{} `xmlrpc:"remaining_qty"`
	ReservedAvailability   interface{} `xmlrpc:"reserved_availability"`
	ReservedQuantIds       interface{} `xmlrpc:"reserved_quant_ids"`
	RestrictLotId          interface{} `xmlrpc:"restrict_lot_id"`
	RestrictPartnerId      interface{} `xmlrpc:"restrict_partner_id"`
	ReturnedMoveIds        interface{} `xmlrpc:"returned_move_ids"`
	RouteIds               interface{} `xmlrpc:"route_ids"`
	RuleId                 interface{} `xmlrpc:"rule_id"`
	Scrapped               bool        `xmlrpc:"scrapped"`
	Sequence               interface{} `xmlrpc:"sequence"`
	SplitFrom              interface{} `xmlrpc:"split_from"`
	State                  interface{} `xmlrpc:"state"`
	StringAvailabilityInfo interface{} `xmlrpc:"string_availability_info"`
	ToRefundSo             bool        `xmlrpc:"to_refund_so"`
	WarehouseId            interface{} `xmlrpc:"warehouse_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
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
