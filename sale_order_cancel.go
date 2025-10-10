package odoo

// SaleOrderCancel represents sale.order.cancel model.
type SaleOrderCancel struct {
	AuthorId                   *Many2One `xmlrpc:"author_id,omitempty"`
	Body                       *String   `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue       *Bool     `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody                *Bool     `xmlrpc:"can_edit_body,omitempty"`
	CreateDate                 *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayDeliveryAlert       *Bool     `xmlrpc:"display_delivery_alert,omitempty"`
	DisplayInvoiceAlert        *Bool     `xmlrpc:"display_invoice_alert,omitempty"`
	DisplayName                *String   `xmlrpc:"display_name,omitempty"`
	DisplayPurchaseOrdersAlert *Bool     `xmlrpc:"display_purchase_orders_alert,omitempty"`
	Id                         *Int      `xmlrpc:"id,omitempty"`
	IsMailTemplateEditor       *Bool     `xmlrpc:"is_mail_template_editor,omitempty"`
	Lang                       *String   `xmlrpc:"lang,omitempty"`
	OrderId                    *Many2One `xmlrpc:"order_id,omitempty"`
	RecipientIds               *Relation `xmlrpc:"recipient_ids,omitempty"`
	RenderModel                *String   `xmlrpc:"render_model,omitempty"`
	Subject                    *String   `xmlrpc:"subject,omitempty"`
	TemplateId                 *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate                  *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderCancels represents array of sale.order.cancel model.
type SaleOrderCancels []SaleOrderCancel

// SaleOrderCancelModel is the odoo model name.
const SaleOrderCancelModel = "sale.order.cancel"

// Many2One convert SaleOrderCancel to *Many2One.
func (soc *SaleOrderCancel) Many2One() *Many2One {
	return NewMany2One(soc.Id.Get(), "")
}

// CreateSaleOrderCancel creates a new sale.order.cancel model and returns its id.
func (c *Client) CreateSaleOrderCancel(soc *SaleOrderCancel) (int64, error) {
	ids, err := c.CreateSaleOrderCancels([]*SaleOrderCancel{soc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderCancel creates a new sale.order.cancel model and returns its id.
func (c *Client) CreateSaleOrderCancels(socs []*SaleOrderCancel) ([]int64, error) {
	var vv []interface{}
	for _, v := range socs {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderCancelModel, vv, nil)
}

// UpdateSaleOrderCancel updates an existing sale.order.cancel record.
func (c *Client) UpdateSaleOrderCancel(soc *SaleOrderCancel) error {
	return c.UpdateSaleOrderCancels([]int64{soc.Id.Get()}, soc)
}

// UpdateSaleOrderCancels updates existing sale.order.cancel records.
// All records (represented by ids) will be updated by soc values.
func (c *Client) UpdateSaleOrderCancels(ids []int64, soc *SaleOrderCancel) error {
	return c.Update(SaleOrderCancelModel, ids, soc, nil)
}

// DeleteSaleOrderCancel deletes an existing sale.order.cancel record.
func (c *Client) DeleteSaleOrderCancel(id int64) error {
	return c.DeleteSaleOrderCancels([]int64{id})
}

// DeleteSaleOrderCancels deletes existing sale.order.cancel records.
func (c *Client) DeleteSaleOrderCancels(ids []int64) error {
	return c.Delete(SaleOrderCancelModel, ids)
}

// GetSaleOrderCancel gets sale.order.cancel existing record.
func (c *Client) GetSaleOrderCancel(id int64) (*SaleOrderCancel, error) {
	socs, err := c.GetSaleOrderCancels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*socs)[0]), nil
}

// GetSaleOrderCancels gets sale.order.cancel existing records.
func (c *Client) GetSaleOrderCancels(ids []int64) (*SaleOrderCancels, error) {
	socs := &SaleOrderCancels{}
	if err := c.Read(SaleOrderCancelModel, ids, nil, socs); err != nil {
		return nil, err
	}
	return socs, nil
}

// FindSaleOrderCancel finds sale.order.cancel record by querying it with criteria.
func (c *Client) FindSaleOrderCancel(criteria *Criteria) (*SaleOrderCancel, error) {
	socs := &SaleOrderCancels{}
	if err := c.SearchRead(SaleOrderCancelModel, criteria, NewOptions().Limit(1), socs); err != nil {
		return nil, err
	}
	return &((*socs)[0]), nil
}

// FindSaleOrderCancels finds sale.order.cancel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCancels(criteria *Criteria, options *Options) (*SaleOrderCancels, error) {
	socs := &SaleOrderCancels{}
	if err := c.SearchRead(SaleOrderCancelModel, criteria, options, socs); err != nil {
		return nil, err
	}
	return socs, nil
}

// FindSaleOrderCancelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCancelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderCancelModel, criteria, options)
}

// FindSaleOrderCancelId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderCancelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderCancelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
