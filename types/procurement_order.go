package types

import (
	"time"
)

type ProcurementOrder struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DatePlanned              time.Time `xmlrpc:"date_planned"`
	DisplayName              string    `xmlrpc:"display_name"`
	GroupId                  Many2One  `xmlrpc:"group_id"`
	Id                       int64     `xmlrpc:"id"`
	LocationId               Many2One  `xmlrpc:"location_id"`
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
	MoveDestId               Many2One  `xmlrpc:"move_dest_id"`
	MoveIds                  []int64   `xmlrpc:"move_ids"`
	Name                     string    `xmlrpc:"name"`
	OrderpointId             Many2One  `xmlrpc:"orderpoint_id"`
	Origin                   string    `xmlrpc:"origin"`
	PartnerDestId            Many2One  `xmlrpc:"partner_dest_id"`
	Priority                 string    `xmlrpc:"priority"`
	ProductId                Many2One  `xmlrpc:"product_id"`
	ProductQty               float64   `xmlrpc:"product_qty"`
	ProductUom               Many2One  `xmlrpc:"product_uom"`
	PurchaseId               Many2One  `xmlrpc:"purchase_id"`
	PurchaseLineId           Many2One  `xmlrpc:"purchase_line_id"`
	RouteIds                 []int64   `xmlrpc:"route_ids"`
	RuleId                   Many2One  `xmlrpc:"rule_id"`
	SaleLineId               Many2One  `xmlrpc:"sale_line_id"`
	State                    string    `xmlrpc:"state"`
	TaskId                   Many2One  `xmlrpc:"task_id"`
	WarehouseId              Many2One  `xmlrpc:"warehouse_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type ProcurementOrderNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DatePlanned              interface{} `xmlrpc:"date_planned"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	GroupId                  interface{} `xmlrpc:"group_id"`
	Id                       interface{} `xmlrpc:"id"`
	LocationId               interface{} `xmlrpc:"location_id"`
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
	MoveDestId               interface{} `xmlrpc:"move_dest_id"`
	MoveIds                  interface{} `xmlrpc:"move_ids"`
	Name                     interface{} `xmlrpc:"name"`
	OrderpointId             interface{} `xmlrpc:"orderpoint_id"`
	Origin                   interface{} `xmlrpc:"origin"`
	PartnerDestId            interface{} `xmlrpc:"partner_dest_id"`
	Priority                 interface{} `xmlrpc:"priority"`
	ProductId                interface{} `xmlrpc:"product_id"`
	ProductQty               interface{} `xmlrpc:"product_qty"`
	ProductUom               interface{} `xmlrpc:"product_uom"`
	PurchaseId               interface{} `xmlrpc:"purchase_id"`
	PurchaseLineId           interface{} `xmlrpc:"purchase_line_id"`
	RouteIds                 interface{} `xmlrpc:"route_ids"`
	RuleId                   interface{} `xmlrpc:"rule_id"`
	SaleLineId               interface{} `xmlrpc:"sale_line_id"`
	State                    interface{} `xmlrpc:"state"`
	TaskId                   interface{} `xmlrpc:"task_id"`
	WarehouseId              interface{} `xmlrpc:"warehouse_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var ProcurementOrderModel string = "procurement.order"

type ProcurementOrders []ProcurementOrder

type ProcurementOrdersNil []ProcurementOrderNil

func (s *ProcurementOrder) NilableType_() interface{} {
	return &ProcurementOrderNil{}
}

func (ns *ProcurementOrderNil) Type_() interface{} {
	s := &ProcurementOrder{}
	return load(ns, s)
}

func (s *ProcurementOrders) NilableType_() interface{} {
	return &ProcurementOrdersNil{}
}

func (ns *ProcurementOrdersNil) Type_() interface{} {
	s := &ProcurementOrders{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProcurementOrder))
	}
	return s
}
