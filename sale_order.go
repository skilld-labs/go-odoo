package odoo

// SaleOrder represents sale.order model.
type SaleOrder struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken              *String    `xmlrpc:"access_token,omitempty"`
	ActivityDateDeadline     *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState            *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary          *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId           *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId           *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountTax                *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal              *Float     `xmlrpc:"amount_total,omitempty"`
	AmountUntaxed            *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AnalyticAccountId        *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	AutosalesConfigId        *Many2One  `xmlrpc:"autosales_config_id,omitempty"`
	CampaignId               *Many2One  `xmlrpc:"campaign_id,omitempty"`
	ClientOrderRef           *String    `xmlrpc:"client_order_ref,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	ConfirmationDate         *Time      `xmlrpc:"confirmation_date,omitempty"`
	ContributionsCreated     *Bool      `xmlrpc:"contributions_created,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateOrder                *Time      `xmlrpc:"date_order,omitempty"`
	DeliveryCount            *Int       `xmlrpc:"delivery_count,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	FiscalPositionId         *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	Incoterm                 *Many2One  `xmlrpc:"incoterm,omitempty"`
	InvoiceCount             *Int       `xmlrpc:"invoice_count,omitempty"`
	InvoiceIds               *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceStatus            *Selection `xmlrpc:"invoice_status,omitempty"`
	IsExpired                *Bool      `xmlrpc:"is_expired,omitempty"`
	MediumId                 *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	Note                     *String    `xmlrpc:"note,omitempty"`
	OpportunityId            *Many2One  `xmlrpc:"opportunity_id,omitempty"`
	OrderLine                *Relation  `xmlrpc:"order_line,omitempty"`
	Origin                   *String    `xmlrpc:"origin,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerInvoiceId         *Many2One  `xmlrpc:"partner_invoice_id,omitempty"`
	PartnerShippingId        *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentTermId            *Many2One  `xmlrpc:"payment_term_id,omitempty"`
	PickingIds               *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingPolicy            *Selection `xmlrpc:"picking_policy,omitempty"`
	PortalUrl                *String    `xmlrpc:"portal_url,omitempty"`
	PricelistId              *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProcurementGroupId       *Many2One  `xmlrpc:"procurement_group_id,omitempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omitempty"`
	ProjectIds               *Relation  `xmlrpc:"project_ids,omitempty"`
	ProjectProjectId         *Many2One  `xmlrpc:"project_project_id,omitempty"`
	SourceId                 *Many2One  `xmlrpc:"source_id,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omitempty"`
	TasksCount               *Int       `xmlrpc:"tasks_count,omitempty"`
	TasksCreated             *Bool      `xmlrpc:"tasks_created,omitempty"`
	TasksIds                 *Relation  `xmlrpc:"tasks_ids,omitempty"`
	TeamId                   *Many2One  `xmlrpc:"team_id,omitempty"`
	TimesheetCount           *Float     `xmlrpc:"timesheet_count,omitempty"`
	TimesheetIds             *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omitempty"`
	ValidityDate             *Time      `xmlrpc:"validity_date,omitempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
	XTitle                   *String    `xmlrpc:"x_title,omitempty"`
	XTitleReadOnly           *String    `xmlrpc:"x_title_read_only,omitempty"`
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

// CreateSaleOrders creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrders(sos []*SaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range sos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderModel, vv, nil)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by ids) will be updated by so values.
func (c *Client) UpdateSaleOrders(ids []int64, so *SaleOrder) error {
	return c.Update(SaleOrderModel, ids, so, nil)
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
	return &((*sos)[0]), nil
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
	return &((*sos)[0]), nil
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
	return c.Search(SaleOrderModel, criteria, options)
}

// FindSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
