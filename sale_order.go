package odoo

import (
	"fmt"
)

// SaleOrder represents sale.order model.
type SaleOrder struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AmountByGroup               *String    `xmlrpc:"amount_by_group,omitempty"`
	AmountTax                   *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal                 *Float     `xmlrpc:"amount_total,omitempty"`
	AmountUndiscounted          *Float     `xmlrpc:"amount_undiscounted,omitempty"`
	AmountUntaxed               *Float     `xmlrpc:"amount_untaxed,omitempty"`
	AnalyticAccountId           *Many2One  `xmlrpc:"analytic_account_id,omitempty"`
	AuthorizedTransactionIds    *Relation  `xmlrpc:"authorized_transaction_ids,omitempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omitempty"`
	ClientOrderRef              *String    `xmlrpc:"client_order_ref,omitempty"`
	CommitmentDate              *Time      `xmlrpc:"commitment_date,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyRate                *Float     `xmlrpc:"currency_rate,omitempty"`
	DateOrder                   *Time      `xmlrpc:"date_order,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	ExpectedDate                *Time      `xmlrpc:"expected_date,omitempty"`
	FiscalPositionId            *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	InvoiceCount                *Int       `xmlrpc:"invoice_count,omitempty"`
	InvoiceIds                  *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceStatus               *Selection `xmlrpc:"invoice_status,omitempty"`
	IsExpired                   *Bool      `xmlrpc:"is_expired,omitempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	Note                        *String    `xmlrpc:"note,omitempty"`
	OrderLine                   *Relation  `xmlrpc:"order_line,omitempty"`
	Origin                      *String    `xmlrpc:"origin,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerInvoiceId            *Many2One  `xmlrpc:"partner_invoice_id,omitempty"`
	PartnerShippingId           *Many2One  `xmlrpc:"partner_shipping_id,omitempty"`
	PaymentTermId               *Many2One  `xmlrpc:"payment_term_id,omitempty"`
	PricelistId                 *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	Reference                   *String    `xmlrpc:"reference,omitempty"`
	RequirePayment              *Bool      `xmlrpc:"require_payment,omitempty"`
	RequireSignature            *Bool      `xmlrpc:"require_signature,omitempty"`
	SaleOrderOptionIds          *Relation  `xmlrpc:"sale_order_option_ids,omitempty"`
	SaleOrderTemplateId         *Many2One  `xmlrpc:"sale_order_template_id,omitempty"`
	ShowUpdatePricelist         *Bool      `xmlrpc:"show_update_pricelist,omitempty"`
	Signature                   *String    `xmlrpc:"signature,omitempty"`
	SignedBy                    *String    `xmlrpc:"signed_by,omitempty"`
	SignedOn                    *Time      `xmlrpc:"signed_on,omitempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omitempty"`
	TransactionIds              *Relation  `xmlrpc:"transaction_ids,omitempty"`
	TypeName                    *String    `xmlrpc:"type_name,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	ValidityDate                *Time      `xmlrpc:"validity_date,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(SaleOrderModel, so)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by IDs) will be updated by so values.
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
	return nil, fmt.Errorf("sale.order was not found")
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

// FindSaleOrderIds finds records IDs by querying it
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
	return -1, fmt.Errorf("sale.order was not found")
}
