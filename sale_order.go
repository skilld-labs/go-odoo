package odoo

import (
	"fmt"
)

// SaleOrder represents sale.order model.
type SaleOrder struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken              *String    `xmlrpc:"access_token,omptempty"`
	ActivityDateDeadline     *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState            *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary          *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId           *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId           *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AmountTax                *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTotal              *Float     `xmlrpc:"amount_total,omptempty"`
	AmountUntaxed            *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AnalyticAccountId        *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	AutosalesConfigId        *Many2One  `xmlrpc:"autosales_config_id,omptempty"`
	CampaignId               *Many2One  `xmlrpc:"campaign_id,omptempty"`
	ClientOrderRef           *String    `xmlrpc:"client_order_ref,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfirmationDate         *Time      `xmlrpc:"confirmation_date,omptempty"`
	ContributionsCreated     *Bool      `xmlrpc:"contributions_created,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateOrder                *Time      `xmlrpc:"date_order,omptempty"`
	DeliveryCount            *Int       `xmlrpc:"delivery_count,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionId         *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Incoterm                 *Many2One  `xmlrpc:"incoterm,omptempty"`
	InvoiceCount             *Int       `xmlrpc:"invoice_count,omptempty"`
	InvoiceIds               *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceStatus            *Selection `xmlrpc:"invoice_status,omptempty"`
	IsExpired                *Bool      `xmlrpc:"is_expired,omptempty"`
	MediumId                 *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	Note                     *String    `xmlrpc:"note,omptempty"`
	OpportunityId            *Many2One  `xmlrpc:"opportunity_id,omptempty"`
	OrderLine                *Relation  `xmlrpc:"order_line,omptempty"`
	Origin                   *String    `xmlrpc:"origin,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerInvoiceId         *Many2One  `xmlrpc:"partner_invoice_id,omptempty"`
	PartnerShippingId        *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PaymentTermId            *Many2One  `xmlrpc:"payment_term_id,omptempty"`
	PickingIds               *Relation  `xmlrpc:"picking_ids,omptempty"`
	PickingPolicy            *Selection `xmlrpc:"picking_policy,omptempty"`
	PortalUrl                *String    `xmlrpc:"portal_url,omptempty"`
	PricelistId              *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProcurementGroupId       *Many2One  `xmlrpc:"procurement_group_id,omptempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omptempty"`
	ProjectIds               *Relation  `xmlrpc:"project_ids,omptempty"`
	ProjectProjectId         *Many2One  `xmlrpc:"project_project_id,omptempty"`
	SourceId                 *Many2One  `xmlrpc:"source_id,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omptempty"`
	TasksCount               *Int       `xmlrpc:"tasks_count,omptempty"`
	TasksIds                 *Relation  `xmlrpc:"tasks_ids,omptempty"`
	TeamId                   *Many2One  `xmlrpc:"team_id,omptempty"`
	TimesheetCount           *Float     `xmlrpc:"timesheet_count,omptempty"`
	TimesheetIds             *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	ValidityDate             *Time      `xmlrpc:"validity_date,omptempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
	XTitle                   *String    `xmlrpc:"x_title,omptempty"`
	XTitleReadOnly           *String    `xmlrpc:"x_title_read_only,omptempty"`
}

// SaleOrders represents array of sale.order model.
type SaleOrders []SaleOrder

// SaleOrderModel is the odoo model name.
const SaleOrderModel = "sale.order"

// Many2One convert SaleOrder to *Many2One.
func (so *SaleOrder) Many2One() *Many2One {
	return NewMany2One(so.Id.Get(), "")
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrder(so *SaleOrder) (int64, error) {
	ids, err := c.CreateSaleOrders([]*SaleOrder{so})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrders(sos []*SaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range sos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderModel, vv)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by ids) will be updated by so values.
func (c *Client) UpdateSaleOrders(ids []int64, so *SaleOrder) error {
	return c.Update(SaleOrderModel, ids, so)
}

// DeleteSaleOrder deletes an existing sale.order record.
func (c *Client) DeleteSaleOrder(id int64) error {
	return c.DeleteSaleOrders([]int64{id})
}

// DeleteSaleOrders deletes existing sale.order records.
func (c *Client) DeleteSaleOrders(ids []int64) error {
	return c.Delete(SaleOrderModel, ids)
}

// GetSaleOrder gets sale.order existing record.
func (c *Client) GetSaleOrder(id int64) (*SaleOrder, error) {
	sos, err := c.GetSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if sos != nil && len(*sos) > 0 {
		return &((*sos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order not found", id)
}

// GetSaleOrders gets sale.order existing records.
func (c *Client) GetSaleOrders(ids []int64) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.Read(SaleOrderModel, ids, nil, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrder finds sale.order record by querying it with criteria.
func (c *Client) FindSaleOrder(criteria *Criteria) (*SaleOrder, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, NewOptions().Limit(1), sos); err != nil {
		return nil, err
	}
	if sos != nil && len(*sos) > 0 {
		return &((*sos)[0]), nil
	}
	return nil, fmt.Errorf("sale.order was not found with criteria %v", criteria)
}

// FindSaleOrders finds sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrders(criteria *Criteria, options *Options) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, options, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order was not found with criteria %v and options %v", criteria, options)
}
