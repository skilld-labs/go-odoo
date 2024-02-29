package odoo

// PurchaseOrder represents purchase.order model.
type PurchaseOrder struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline       *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds                *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState              *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary            *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId             *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId             *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountTax                  *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal                *Float     `xmlrpc:"amount_total,omitempty"`
	AmountUntaxed              *Float     `xmlrpc:"amount_untaxed,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                 *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateApprove                *Time      `xmlrpc:"date_approve,omitempty"`
	DateOrder                  *Time      `xmlrpc:"date_order,omitempty"`
	DatePlanned                *Time      `xmlrpc:"date_planned,omitempty"`
	DefaultLocationDestIdUsage *Selection `xmlrpc:"default_location_dest_id_usage,omitempty"`
	DestAddressId              *Many2One  `xmlrpc:"dest_address_id,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	FiscalPositionId           *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	GroupId                    *Many2One  `xmlrpc:"group_id,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	IncotermId                 *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InvoiceCount               *Int       `xmlrpc:"invoice_count,omitempty"`
	InvoiceIds                 *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceStatus              *Selection `xmlrpc:"invoice_status,omitempty"`
	IsShipped                  *Bool      `xmlrpc:"is_shipped,omitempty"`
	MessageChannelIds          *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds         *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds                 *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower          *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost            *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction          *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter   *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread              *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter       *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	Notes                      *String    `xmlrpc:"notes,omitempty"`
	OrderLine                  *Relation  `xmlrpc:"order_line,omitempty"`
	Origin                     *String    `xmlrpc:"origin,omitempty"`
	PartnerId                  *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerRef                 *String    `xmlrpc:"partner_ref,omitempty"`
	PaymentTermId              *Many2One  `xmlrpc:"payment_term_id,omitempty"`
	PickingCount               *Int       `xmlrpc:"picking_count,omitempty"`
	PickingIds                 *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingTypeId              *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	ProductId                  *Many2One  `xmlrpc:"product_id,omitempty"`
	State                      *Selection `xmlrpc:"state,omitempty"`
	WebsiteMessageIds          *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteUrl                 *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PurchaseOrders represents array of purchase.order model.
type PurchaseOrders []PurchaseOrder

// PurchaseOrderModel is the odoo model name.
const PurchaseOrderModel = "purchase.order"

// Many2One convert PurchaseOrder to *Many2One.
func (po *PurchaseOrder) Many2One() *Many2One {
	return NewMany2One(po.Id.Get(), "")
}

// CreatePurchaseOrder creates a new purchase.order model and returns its id.
func (c *Client) CreatePurchaseOrder(po *PurchaseOrder) (int64, error) {
	ids, err := c.CreatePurchaseOrders([]*PurchaseOrder{po})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseOrders creates a new purchase.order model and returns its id.
func (c *Client) CreatePurchaseOrders(pos []*PurchaseOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range pos {
		vv = append(vv, v)
	}
	return c.Create(PurchaseOrderModel, vv, nil)
}

// UpdatePurchaseOrder updates an existing purchase.order record.
func (c *Client) UpdatePurchaseOrder(po *PurchaseOrder) error {
	return c.UpdatePurchaseOrders([]int64{po.Id.Get()}, po)
}

// UpdatePurchaseOrders updates existing purchase.order records.
// All records (represented by ids) will be updated by po values.
func (c *Client) UpdatePurchaseOrders(ids []int64, po *PurchaseOrder) error {
	return c.Update(PurchaseOrderModel, ids, po, nil)
}

// DeletePurchaseOrder deletes an existing purchase.order record.
func (c *Client) DeletePurchaseOrder(id int64) error {
	return c.DeletePurchaseOrders([]int64{id})
}

// DeletePurchaseOrders deletes existing purchase.order records.
func (c *Client) DeletePurchaseOrders(ids []int64) error {
	return c.Delete(PurchaseOrderModel, ids)
}

// GetPurchaseOrder gets purchase.order existing record.
func (c *Client) GetPurchaseOrder(id int64) (*PurchaseOrder, error) {
	pos, err := c.GetPurchaseOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// GetPurchaseOrders gets purchase.order existing records.
func (c *Client) GetPurchaseOrders(ids []int64) (*PurchaseOrders, error) {
	pos := &PurchaseOrders{}
	if err := c.Read(PurchaseOrderModel, ids, nil, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPurchaseOrder finds purchase.order record by querying it with criteria.
func (c *Client) FindPurchaseOrder(criteria *Criteria) (*PurchaseOrder, error) {
	pos := &PurchaseOrders{}
	if err := c.SearchRead(PurchaseOrderModel, criteria, NewOptions().Limit(1), pos); err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// FindPurchaseOrders finds purchase.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseOrders(criteria *Criteria, options *Options) (*PurchaseOrders, error) {
	pos := &PurchaseOrders{}
	if err := c.SearchRead(PurchaseOrderModel, criteria, options, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPurchaseOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PurchaseOrderModel, criteria, options)
}

// FindPurchaseOrderId finds record id by querying it with criteria.
func (c *Client) FindPurchaseOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
