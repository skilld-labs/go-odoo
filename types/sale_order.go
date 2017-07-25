package types

import (
	"time"
)

type SaleOrder struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	AmountTax                float64   `xmlrpc:"amount_tax"`
	AmountTotal              float64   `xmlrpc:"amount_total"`
	AmountUntaxed            float64   `xmlrpc:"amount_untaxed"`
	AutosalesConfigId        Many2One  `xmlrpc:"autosales_config_id"`
	CampaignId               Many2One  `xmlrpc:"campaign_id"`
	ClientOrderRef           string    `xmlrpc:"client_order_ref"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	ConfirmationDate         time.Time `xmlrpc:"confirmation_date"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	CurrencyId               Many2One  `xmlrpc:"currency_id"`
	DateOrder                time.Time `xmlrpc:"date_order"`
	DeliveryCount            int64     `xmlrpc:"delivery_count"`
	DisplayName              string    `xmlrpc:"display_name"`
	FiscalPositionId         Many2One  `xmlrpc:"fiscal_position_id"`
	Id                       int64     `xmlrpc:"id"`
	Incoterm                 Many2One  `xmlrpc:"incoterm"`
	InvoiceCount             int64     `xmlrpc:"invoice_count"`
	InvoiceIds               []int64   `xmlrpc:"invoice_ids"`
	InvoiceStatus            string    `xmlrpc:"invoice_status"`
	MediumId                 Many2One  `xmlrpc:"medium_id"`
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
	Name                     string    `xmlrpc:"name"`
	Note                     string    `xmlrpc:"note"`
	OpportunityId            Many2One  `xmlrpc:"opportunity_id"`
	OrderLine                []int64   `xmlrpc:"order_line"`
	Origin                   string    `xmlrpc:"origin"`
	PartnerId                Many2One  `xmlrpc:"partner_id"`
	PartnerInvoiceId         Many2One  `xmlrpc:"partner_invoice_id"`
	PartnerShippingId        Many2One  `xmlrpc:"partner_shipping_id"`
	PaymentTermId            Many2One  `xmlrpc:"payment_term_id"`
	PickingIds               []int64   `xmlrpc:"picking_ids"`
	PickingPolicy            string    `xmlrpc:"picking_policy"`
	PricelistId              Many2One  `xmlrpc:"pricelist_id"`
	ProcurementGroupId       Many2One  `xmlrpc:"procurement_group_id"`
	ProductId                Many2One  `xmlrpc:"product_id"`
	ProjectId                Many2One  `xmlrpc:"project_id"`
	ProjectProjectId         Many2One  `xmlrpc:"project_project_id"`
	RelatedProjectId         Many2One  `xmlrpc:"related_project_id"`
	SourceId                 Many2One  `xmlrpc:"source_id"`
	State                    string    `xmlrpc:"state"`
	TagIds                   []int64   `xmlrpc:"tag_ids"`
	TasksCount               int64     `xmlrpc:"tasks_count"`
	TasksIds                 []int64   `xmlrpc:"tasks_ids"`
	TeamId                   Many2One  `xmlrpc:"team_id"`
	TimesheetCount           float64   `xmlrpc:"timesheet_count"`
	TimesheetIds             []int64   `xmlrpc:"timesheet_ids"`
	UserId                   Many2One  `xmlrpc:"user_id"`
	ValidityDate             time.Time `xmlrpc:"validity_date"`
	WarehouseId              Many2One  `xmlrpc:"warehouse_id"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type SaleOrderNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	AmountTax                interface{} `xmlrpc:"amount_tax"`
	AmountTotal              interface{} `xmlrpc:"amount_total"`
	AmountUntaxed            interface{} `xmlrpc:"amount_untaxed"`
	AutosalesConfigId        interface{} `xmlrpc:"autosales_config_id"`
	CampaignId               interface{} `xmlrpc:"campaign_id"`
	ClientOrderRef           interface{} `xmlrpc:"client_order_ref"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	ConfirmationDate         interface{} `xmlrpc:"confirmation_date"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	CurrencyId               interface{} `xmlrpc:"currency_id"`
	DateOrder                interface{} `xmlrpc:"date_order"`
	DeliveryCount            interface{} `xmlrpc:"delivery_count"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	FiscalPositionId         interface{} `xmlrpc:"fiscal_position_id"`
	Id                       interface{} `xmlrpc:"id"`
	Incoterm                 interface{} `xmlrpc:"incoterm"`
	InvoiceCount             interface{} `xmlrpc:"invoice_count"`
	InvoiceIds               interface{} `xmlrpc:"invoice_ids"`
	InvoiceStatus            interface{} `xmlrpc:"invoice_status"`
	MediumId                 interface{} `xmlrpc:"medium_id"`
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
	Name                     interface{} `xmlrpc:"name"`
	Note                     interface{} `xmlrpc:"note"`
	OpportunityId            interface{} `xmlrpc:"opportunity_id"`
	OrderLine                interface{} `xmlrpc:"order_line"`
	Origin                   interface{} `xmlrpc:"origin"`
	PartnerId                interface{} `xmlrpc:"partner_id"`
	PartnerInvoiceId         interface{} `xmlrpc:"partner_invoice_id"`
	PartnerShippingId        interface{} `xmlrpc:"partner_shipping_id"`
	PaymentTermId            interface{} `xmlrpc:"payment_term_id"`
	PickingIds               interface{} `xmlrpc:"picking_ids"`
	PickingPolicy            interface{} `xmlrpc:"picking_policy"`
	PricelistId              interface{} `xmlrpc:"pricelist_id"`
	ProcurementGroupId       interface{} `xmlrpc:"procurement_group_id"`
	ProductId                interface{} `xmlrpc:"product_id"`
	ProjectId                interface{} `xmlrpc:"project_id"`
	ProjectProjectId         interface{} `xmlrpc:"project_project_id"`
	RelatedProjectId         interface{} `xmlrpc:"related_project_id"`
	SourceId                 interface{} `xmlrpc:"source_id"`
	State                    interface{} `xmlrpc:"state"`
	TagIds                   interface{} `xmlrpc:"tag_ids"`
	TasksCount               interface{} `xmlrpc:"tasks_count"`
	TasksIds                 interface{} `xmlrpc:"tasks_ids"`
	TeamId                   interface{} `xmlrpc:"team_id"`
	TimesheetCount           interface{} `xmlrpc:"timesheet_count"`
	TimesheetIds             interface{} `xmlrpc:"timesheet_ids"`
	UserId                   interface{} `xmlrpc:"user_id"`
	ValidityDate             interface{} `xmlrpc:"validity_date"`
	WarehouseId              interface{} `xmlrpc:"warehouse_id"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var SaleOrderModel string = "sale.order"

type SaleOrders []SaleOrder

type SaleOrdersNil []SaleOrderNil

func (s *SaleOrder) NilableType_() interface{} {
	return &SaleOrderNil{}
}

func (ns *SaleOrderNil) Type_() interface{} {
	s := &SaleOrder{}
	return load(ns, s)
}

func (s *SaleOrders) NilableType_() interface{} {
	return &SaleOrdersNil{}
}

func (ns *SaleOrdersNil) Type_() interface{} {
	s := &SaleOrders{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SaleOrder))
	}
	return s
}
