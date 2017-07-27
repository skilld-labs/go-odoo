package types

import (
	"time"
)

type StockPicking struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	BackorderId              Many2One  `xmlrpc:"backorder_id"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	Date                     time.Time `xmlrpc:"date"`
	DateDone                 time.Time `xmlrpc:"date_done"`
	DisplayName              string    `xmlrpc:"display_name"`
	GroupId                  Many2One  `xmlrpc:"group_id"`
	HasScrapMove             bool      `xmlrpc:"has_scrap_move"`
	Id                       int64     `xmlrpc:"id"`
	LaunchPackOperations     bool      `xmlrpc:"launch_pack_operations"`
	LocationDestId           Many2One  `xmlrpc:"location_dest_id"`
	LocationId               Many2One  `xmlrpc:"location_id"`
	MaxDate                  time.Time `xmlrpc:"max_date"`
	MessageChannelIds        []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       []int64   `xmlrpc:"message_follower_ids"`
	MessageIds               []int64   `xmlrpc:"message_ids"`
	MessageIsFollower        bool      `xmlrpc:"message_is_follower"`
	MessageLastPost          time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction        bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread            bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter     int64     `xmlrpc:"message_unread_counter"`
	MinDate                  time.Time `xmlrpc:"min_date"`
	MoveLines                []int64   `xmlrpc:"move_lines"`
	MoveType                 string    `xmlrpc:"move_type"`
	Name                     string    `xmlrpc:"name"`
	Note                     string    `xmlrpc:"note"`
	Origin                   string    `xmlrpc:"origin"`
	OwnerId                  Many2One  `xmlrpc:"owner_id"`
	PackOperationExist       bool      `xmlrpc:"pack_operation_exist"`
	PackOperationIds         []int64   `xmlrpc:"pack_operation_ids"`
	PackOperationPackIds     []int64   `xmlrpc:"pack_operation_pack_ids"`
	PackOperationProductIds  []int64   `xmlrpc:"pack_operation_product_ids"`
	PartnerId                Many2One  `xmlrpc:"partner_id"`
	PickingTypeCode          string    `xmlrpc:"picking_type_code"`
	PickingTypeEntirePacks   bool      `xmlrpc:"picking_type_entire_packs"`
	PickingTypeId            Many2One  `xmlrpc:"picking_type_id"`
	Printed                  bool      `xmlrpc:"printed"`
	Priority                 string    `xmlrpc:"priority"`
	ProductId                Many2One  `xmlrpc:"product_id"`
	PurchaseId               Many2One  `xmlrpc:"purchase_id"`
	QuantReservedExist       bool      `xmlrpc:"quant_reserved_exist"`
	RecomputePackOp          bool      `xmlrpc:"recompute_pack_op"`
	SaleId                   Many2One  `xmlrpc:"sale_id"`
	State                    string    `xmlrpc:"state"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type StockPickingNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	BackorderId              interface{} `xmlrpc:"backorder_id"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	Date                     interface{} `xmlrpc:"date"`
	DateDone                 interface{} `xmlrpc:"date_done"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	GroupId                  interface{} `xmlrpc:"group_id"`
	HasScrapMove             bool        `xmlrpc:"has_scrap_move"`
	Id                       interface{} `xmlrpc:"id"`
	LaunchPackOperations     bool        `xmlrpc:"launch_pack_operations"`
	LocationDestId           interface{} `xmlrpc:"location_dest_id"`
	LocationId               interface{} `xmlrpc:"location_id"`
	MaxDate                  interface{} `xmlrpc:"max_date"`
	MessageChannelIds        interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds       interface{} `xmlrpc:"message_follower_ids"`
	MessageIds               interface{} `xmlrpc:"message_ids"`
	MessageIsFollower        bool        `xmlrpc:"message_is_follower"`
	MessageLastPost          interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction        bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds        interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread            bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter     interface{} `xmlrpc:"message_unread_counter"`
	MinDate                  interface{} `xmlrpc:"min_date"`
	MoveLines                interface{} `xmlrpc:"move_lines"`
	MoveType                 interface{} `xmlrpc:"move_type"`
	Name                     interface{} `xmlrpc:"name"`
	Note                     interface{} `xmlrpc:"note"`
	Origin                   interface{} `xmlrpc:"origin"`
	OwnerId                  interface{} `xmlrpc:"owner_id"`
	PackOperationExist       bool        `xmlrpc:"pack_operation_exist"`
	PackOperationIds         interface{} `xmlrpc:"pack_operation_ids"`
	PackOperationPackIds     interface{} `xmlrpc:"pack_operation_pack_ids"`
	PackOperationProductIds  interface{} `xmlrpc:"pack_operation_product_ids"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PickingTypeCode          interface{} `xmlrpc:"picking_type_code"`
	PickingTypeEntirePacks   bool        `xmlrpc:"picking_type_entire_packs"`
	PickingTypeId            interface{} `xmlrpc:"picking_type_id"`
	Printed                  bool        `xmlrpc:"printed"`
	Priority                 interface{} `xmlrpc:"priority"`
	ProductId                interface{} `xmlrpc:"product_id"`
	PurchaseId               interface{} `xmlrpc:"purchase_id"`
	QuantReservedExist       bool        `xmlrpc:"quant_reserved_exist"`
	RecomputePackOp          bool        `xmlrpc:"recompute_pack_op"`
	SaleId                   interface{} `xmlrpc:"sale_id"`
	State                    interface{} `xmlrpc:"state"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var StockPickingModel string = "stock.picking"

type StockPickings []StockPicking

type StockPickingsNil []StockPickingNil

func (s *StockPicking) NilableType_() interface{} {
	return &StockPickingNil{}
}

func (ns *StockPickingNil) Type_() interface{} {
	s := &StockPicking{}
	return load(ns, s)
}

func (s *StockPickings) NilableType_() interface{} {
	return &StockPickingsNil{}
}

func (ns *StockPickingsNil) Type_() interface{} {
	s := &StockPickings{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockPicking))
	}
	return s
}
