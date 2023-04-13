package odoo

import (
	"fmt"
)

// PurchaseOrder represents purchase.order model.
type PurchaseOrder struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadline       *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds                *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState              *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary            *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId             *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId             *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AmountTax                  *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTotal                *Float     `xmlrpc:"amount_total,omptempty"`
	AmountUntaxed              *Float     `xmlrpc:"amount_untaxed,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                 *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateApprove                *Time      `xmlrpc:"date_approve,omptempty"`
	DateOrder                  *Time      `xmlrpc:"date_order,omptempty"`
	DatePlanned                *Time      `xmlrpc:"date_planned,omptempty"`
	DefaultLocationDestIdUsage *Selection `xmlrpc:"default_location_dest_id_usage,omptempty"`
	DestAddressId              *Many2One  `xmlrpc:"dest_address_id,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionId           *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	GroupId                    *Many2One  `xmlrpc:"group_id,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	IncotermId                 *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InvoiceCount               *Int       `xmlrpc:"invoice_count,omptempty"`
	InvoiceIds                 *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceStatus              *Selection `xmlrpc:"invoice_status,omptempty"`
	IsShipped                  *Bool      `xmlrpc:"is_shipped,omptempty"`
	MessageChannelIds          *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds         *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                 *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower          *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost            *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction          *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter   *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread              *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter       *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	Notes                      *String    `xmlrpc:"notes,omptempty"`
	OrderLine                  *Relation  `xmlrpc:"order_line,omptempty"`
	Origin                     *String    `xmlrpc:"origin,omptempty"`
	PartnerId                  *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerRef                 *String    `xmlrpc:"partner_ref,omptempty"`
	PaymentTermId              *Many2One  `xmlrpc:"payment_term_id,omptempty"`
	PickingCount               *Int       `xmlrpc:"picking_count,omptempty"`
	PickingIds                 *Relation  `xmlrpc:"picking_ids,omptempty"`
	PickingTypeId              *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	ProductId                  *Many2One  `xmlrpc:"product_id,omptempty"`
	State                      *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds          *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteUrl                 *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
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

// CreatePurchaseOrder creates a new purchase.order model and returns its id.
func (c *Client) CreatePurchaseOrders(pos []*PurchaseOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range pos {
		vv = append(vv, v)
	}
	return c.Create(PurchaseOrderModel, vv)
}

// UpdatePurchaseOrder updates an existing purchase.order record.
func (c *Client) UpdatePurchaseOrder(po *PurchaseOrder) error {
	return c.UpdatePurchaseOrders([]int64{po.Id.Get()}, po)
}

// UpdatePurchaseOrders updates existing purchase.order records.
// All records (represented by ids) will be updated by po values.
func (c *Client) UpdatePurchaseOrders(ids []int64, po *PurchaseOrder) error {
	return c.Update(PurchaseOrderModel, ids, po)
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
	if pos != nil && len(*pos) > 0 {
		return &((*pos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of purchase.order not found", id)
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
	if pos != nil && len(*pos) > 0 {
		return &((*pos)[0]), nil
	}
	return nil, fmt.Errorf("purchase.order was not found with criteria %v", criteria)
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
	ids, err := c.Search(PurchaseOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPurchaseOrderId finds record id by querying it with criteria.
func (c *Client) FindPurchaseOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("purchase.order was not found with criteria %v and options %v", criteria, options)
}
