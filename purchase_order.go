package odoo

// PurchaseOrder represents purchase.order model.
type PurchaseOrder struct {
	AccessToken                  *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                    *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning                *String    `xmlrpc:"access_warning,omitempty"`
	ActivityDateDeadline         *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration  *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon        *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                  *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary              *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon             *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId               *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId               *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AlternativePoIds             *Relation  `xmlrpc:"alternative_po_ids,omitempty"`
	AmountTax                    *Float     `xmlrpc:"amount_tax,omitempty"`
	AmountTotal                  *Float     `xmlrpc:"amount_total,omitempty"`
	AmountUntaxed                *Float     `xmlrpc:"amount_untaxed,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                  *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyRate                 *Float     `xmlrpc:"currency_rate,omitempty"`
	DateApprove                  *Time      `xmlrpc:"date_approve,omitempty"`
	DateCalendarStart            *Time      `xmlrpc:"date_calendar_start,omitempty"`
	DateOrder                    *Time      `xmlrpc:"date_order,omitempty"`
	DatePlanned                  *Time      `xmlrpc:"date_planned,omitempty"`
	DatePlannedMps               *Time      `xmlrpc:"date_planned_mps,omitempty"`
	DefaultLocationDestIdUsage   *Selection `xmlrpc:"default_location_dest_id_usage,omitempty"`
	DestAddressId                *Many2One  `xmlrpc:"dest_address_id,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	EffectiveDate                *Time      `xmlrpc:"effective_date,omitempty"`
	FiscalPositionId             *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	Grid                         *String    `xmlrpc:"grid,omitempty"`
	GridProductTmplId            *Many2One  `xmlrpc:"grid_product_tmpl_id,omitempty"`
	GridUpdate                   *Bool      `xmlrpc:"grid_update,omitempty"`
	GroupId                      *Many2One  `xmlrpc:"group_id,omitempty"`
	HasAlternatives              *Bool      `xmlrpc:"has_alternatives,omitempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	IncomingPickingCount         *Int       `xmlrpc:"incoming_picking_count,omitempty"`
	IncotermId                   *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	IncotermLocation             *String    `xmlrpc:"incoterm_location,omitempty"`
	InvoiceCount                 *Int       `xmlrpc:"invoice_count,omitempty"`
	InvoiceIds                   *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceStatus                *Selection `xmlrpc:"invoice_status,omitempty"`
	IsQuantityCopy               *Selection `xmlrpc:"is_quantity_copy,omitempty"`
	IsShipped                    *Bool      `xmlrpc:"is_shipped,omitempty"`
	MailReceptionConfirmed       *Bool      `xmlrpc:"mail_reception_confirmed,omitempty"`
	MailReminderConfirmed        *Bool      `xmlrpc:"mail_reminder_confirmed,omitempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MrpProductionCount           *Int       `xmlrpc:"mrp_production_count,omitempty"`
	MyActivityDateDeadline       *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	Notes                        *String    `xmlrpc:"notes,omitempty"`
	OnTimeRate                   *Float     `xmlrpc:"on_time_rate,omitempty"`
	OnTimeRatePerc               *Float     `xmlrpc:"on_time_rate_perc,omitempty"`
	OrderLine                    *Relation  `xmlrpc:"order_line,omitempty"`
	Origin                       *String    `xmlrpc:"origin,omitempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerRef                   *String    `xmlrpc:"partner_ref,omitempty"`
	PaymentTermId                *Many2One  `xmlrpc:"payment_term_id,omitempty"`
	PickingIds                   *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingTypeId                *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	Priority                     *Selection `xmlrpc:"priority,omitempty"`
	ProductId                    *Many2One  `xmlrpc:"product_id,omitempty"`
	PurchaseGroupId              *Many2One  `xmlrpc:"purchase_group_id,omitempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReceiptReminderEmail         *Bool      `xmlrpc:"receipt_reminder_email,omitempty"`
	ReceiptStatus                *Selection `xmlrpc:"receipt_status,omitempty"`
	ReminderDateBeforeReceipt    *Int       `xmlrpc:"reminder_date_before_receipt,omitempty"`
	ReportGrids                  *Bool      `xmlrpc:"report_grids,omitempty"`
	RequisitionId                *Many2One  `xmlrpc:"requisition_id,omitempty"`
	SaleOrderCount               *Int       `xmlrpc:"sale_order_count,omitempty"`
	State                        *Selection `xmlrpc:"state,omitempty"`
	TaxCalculationRoundingMethod *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCountryId                 *Many2One  `xmlrpc:"tax_country_id,omitempty"`
	TaxTotals                    *String    `xmlrpc:"tax_totals,omitempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
	XStudioProductCategory       *Selection `xmlrpc:"x_studio_product_category,omitempty"`
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
