package types

import (
	"time"
)

type PurchaseOrder struct {
	LastUpdate                 time.Time `xmlrpc:"__last_update"`
	AmountTax                  float64   `xmlrpc:"amount_tax"`
	AmountTotal                float64   `xmlrpc:"amount_total"`
	AmountUntaxed              float64   `xmlrpc:"amount_untaxed"`
	CompanyId                  Many2One  `xmlrpc:"company_id"`
	CreateDate                 time.Time `xmlrpc:"create_date"`
	CreateUid                  Many2One  `xmlrpc:"create_uid"`
	CurrencyId                 Many2One  `xmlrpc:"currency_id"`
	DateApprove                time.Time `xmlrpc:"date_approve"`
	DateOrder                  time.Time `xmlrpc:"date_order"`
	DatePlanned                time.Time `xmlrpc:"date_planned"`
	DefaultLocationDestIdUsage string    `xmlrpc:"default_location_dest_id_usage"`
	DestAddressId              Many2One  `xmlrpc:"dest_address_id"`
	DisplayName                string    `xmlrpc:"display_name"`
	FiscalPositionId           Many2One  `xmlrpc:"fiscal_position_id"`
	GroupId                    Many2One  `xmlrpc:"group_id"`
	Id                         int64     `xmlrpc:"id"`
	IncotermId                 Many2One  `xmlrpc:"incoterm_id"`
	InvoiceCount               int64     `xmlrpc:"invoice_count"`
	InvoiceIds                 []int64   `xmlrpc:"invoice_ids"`
	InvoiceStatus              string    `xmlrpc:"invoice_status"`
	IsShipped                  bool      `xmlrpc:"is_shipped"`
	MessageChannelIds          []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds         []int64   `xmlrpc:"message_follower_ids"`
	MessageIds                 []int64   `xmlrpc:"message_ids"`
	MessageIsFollower          bool      `xmlrpc:"message_is_follower"`
	MessageLastPost            time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction          bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter   int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds          []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread              bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter       int64     `xmlrpc:"message_unread_counter"`
	Name                       string    `xmlrpc:"name"`
	Notes                      string    `xmlrpc:"notes"`
	OrderLine                  []int64   `xmlrpc:"order_line"`
	Origin                     string    `xmlrpc:"origin"`
	PartnerId                  Many2One  `xmlrpc:"partner_id"`
	PartnerRef                 string    `xmlrpc:"partner_ref"`
	PaymentTermId              Many2One  `xmlrpc:"payment_term_id"`
	PickingCount               int64     `xmlrpc:"picking_count"`
	PickingIds                 []int64   `xmlrpc:"picking_ids"`
	PickingTypeId              Many2One  `xmlrpc:"picking_type_id"`
	ProductId                  Many2One  `xmlrpc:"product_id"`
	State                      string    `xmlrpc:"state"`
	WriteDate                  time.Time `xmlrpc:"write_date"`
	WriteUid                   Many2One  `xmlrpc:"write_uid"`
}

type PurchaseOrderNil struct {
	LastUpdate                 interface{} `xmlrpc:"__last_update"`
	AmountTax                  interface{} `xmlrpc:"amount_tax"`
	AmountTotal                interface{} `xmlrpc:"amount_total"`
	AmountUntaxed              interface{} `xmlrpc:"amount_untaxed"`
	CompanyId                  interface{} `xmlrpc:"company_id"`
	CreateDate                 interface{} `xmlrpc:"create_date"`
	CreateUid                  interface{} `xmlrpc:"create_uid"`
	CurrencyId                 interface{} `xmlrpc:"currency_id"`
	DateApprove                interface{} `xmlrpc:"date_approve"`
	DateOrder                  interface{} `xmlrpc:"date_order"`
	DatePlanned                interface{} `xmlrpc:"date_planned"`
	DefaultLocationDestIdUsage interface{} `xmlrpc:"default_location_dest_id_usage"`
	DestAddressId              interface{} `xmlrpc:"dest_address_id"`
	DisplayName                interface{} `xmlrpc:"display_name"`
	FiscalPositionId           interface{} `xmlrpc:"fiscal_position_id"`
	GroupId                    interface{} `xmlrpc:"group_id"`
	Id                         interface{} `xmlrpc:"id"`
	IncotermId                 interface{} `xmlrpc:"incoterm_id"`
	InvoiceCount               interface{} `xmlrpc:"invoice_count"`
	InvoiceIds                 interface{} `xmlrpc:"invoice_ids"`
	InvoiceStatus              interface{} `xmlrpc:"invoice_status"`
	IsShipped                  bool        `xmlrpc:"is_shipped"`
	MessageChannelIds          interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds         interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                 interface{} `xmlrpc:"message_ids"`
	MessageIsFollower          bool        `xmlrpc:"message_is_follower"`
	MessageLastPost            interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction          bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter   interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds          interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread              bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter       interface{} `xmlrpc:"message_unread_counter"`
	Name                       interface{} `xmlrpc:"name"`
	Notes                      interface{} `xmlrpc:"notes"`
	OrderLine                  interface{} `xmlrpc:"order_line"`
	Origin                     interface{} `xmlrpc:"origin"`
	PartnerId                  interface{} `xmlrpc:"partner_id"`
	PartnerRef                 interface{} `xmlrpc:"partner_ref"`
	PaymentTermId              interface{} `xmlrpc:"payment_term_id"`
	PickingCount               interface{} `xmlrpc:"picking_count"`
	PickingIds                 interface{} `xmlrpc:"picking_ids"`
	PickingTypeId              interface{} `xmlrpc:"picking_type_id"`
	ProductId                  interface{} `xmlrpc:"product_id"`
	State                      interface{} `xmlrpc:"state"`
	WriteDate                  interface{} `xmlrpc:"write_date"`
	WriteUid                   interface{} `xmlrpc:"write_uid"`
}

var PurchaseOrderModel string = "purchase.order"

type PurchaseOrders []PurchaseOrder

type PurchaseOrdersNil []PurchaseOrderNil

func (s *PurchaseOrder) NilableType_() interface{} {
	return &PurchaseOrderNil{}
}

func (ns *PurchaseOrderNil) Type_() interface{} {
	s := &PurchaseOrder{}
	return load(ns, s)
}

func (s *PurchaseOrders) NilableType_() interface{} {
	return &PurchaseOrdersNil{}
}

func (ns *PurchaseOrdersNil) Type_() interface{} {
	s := &PurchaseOrders{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PurchaseOrder))
	}
	return s
}
