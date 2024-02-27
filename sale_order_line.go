package odoo

// SaleOrderLine represents sale.order.line model.
type SaleOrderLine struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	AmtInvoiced            *Float     `xmlrpc:"amt_invoiced,omitempty"`
	AmtToInvoice           *Float     `xmlrpc:"amt_to_invoice,omitempty"`
	AnalyticTagIds         *Relation  `xmlrpc:"analytic_tag_ids,omitempty"`
	AutosalesBaseOrderLine *Many2One  `xmlrpc:"autosales_base_order_line,omitempty"`
	AutosalesLine          *Bool      `xmlrpc:"autosales_line,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omitempty"`
	CustomerLead           *Float     `xmlrpc:"customer_lead,omitempty"`
	Discount               *Float     `xmlrpc:"discount,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	InvoiceLines           *Relation  `xmlrpc:"invoice_lines,omitempty"`
	InvoiceStatus          *Selection `xmlrpc:"invoice_status,omitempty"`
	IsDownpayment          *Bool      `xmlrpc:"is_downpayment,omitempty"`
	IsService              *Bool      `xmlrpc:"is_service,omitempty"`
	LayoutCategoryId       *Many2One  `xmlrpc:"layout_category_id,omitempty"`
	LayoutCategorySequence *Int       `xmlrpc:"layout_category_sequence,omitempty"`
	MoveIds                *Relation  `xmlrpc:"move_ids,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	OrderId                *Many2One  `xmlrpc:"order_id,omitempty"`
	OrderPartnerId         *Many2One  `xmlrpc:"order_partner_id,omitempty"`
	PriceReduce            *Float     `xmlrpc:"price_reduce,omitempty"`
	PriceReduceTaxexcl     *Float     `xmlrpc:"price_reduce_taxexcl,omitempty"`
	PriceReduceTaxinc      *Float     `xmlrpc:"price_reduce_taxinc,omitempty"`
	PriceSubtotal          *Float     `xmlrpc:"price_subtotal,omitempty"`
	PriceTax               *Float     `xmlrpc:"price_tax,omitempty"`
	PriceTotal             *Float     `xmlrpc:"price_total,omitempty"`
	PriceUnit              *Float     `xmlrpc:"price_unit,omitempty"`
	ProductId              *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductImage           *String    `xmlrpc:"product_image,omitempty"`
	ProductPackaging       *Many2One  `xmlrpc:"product_packaging,omitempty"`
	ProductUom             *Many2One  `xmlrpc:"product_uom,omitempty"`
	ProductUomQty          *Float     `xmlrpc:"product_uom_qty,omitempty"`
	ProductUpdatable       *Bool      `xmlrpc:"product_updatable,omitempty"`
	QtyDelivered           *Float     `xmlrpc:"qty_delivered,omitempty"`
	QtyDeliveredUpdateable *Bool      `xmlrpc:"qty_delivered_updateable,omitempty"`
	QtyInvoiced            *Float     `xmlrpc:"qty_invoiced,omitempty"`
	QtyToInvoice           *Float     `xmlrpc:"qty_to_invoice,omitempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omitempty"`
	SalesmanId             *Many2One  `xmlrpc:"salesman_id,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	State                  *Selection `xmlrpc:"state,omitempty"`
	TaskId                 *Many2One  `xmlrpc:"task_id,omitempty"`
	TaxId                  *Relation  `xmlrpc:"tax_id,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderLines represents array of sale.order.line model.
type SaleOrderLines []SaleOrderLine

// SaleOrderLineModel is the odoo model name.
const SaleOrderLineModel = "sale.order.line"

// Many2One convert SaleOrderLine to *Many2One.
func (sol *SaleOrderLine) Many2One() *Many2One {
	return NewMany2One(sol.Id.Get(), "")
}

// CreateSaleOrderLine creates a new sale.order.line model and returns its id.
func (c *Client) CreateSaleOrderLine(sol *SaleOrderLine) (int64, error) {
	ids, err := c.CreateSaleOrderLines([]*SaleOrderLine{sol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderLine creates a new sale.order.line model and returns its id.
func (c *Client) CreateSaleOrderLines(sols []*SaleOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range sols {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderLineModel, vv, nil)
}

// UpdateSaleOrderLine updates an existing sale.order.line record.
func (c *Client) UpdateSaleOrderLine(sol *SaleOrderLine) error {
	return c.UpdateSaleOrderLines([]int64{sol.Id.Get()}, sol)
}

// UpdateSaleOrderLines updates existing sale.order.line records.
// All records (represented by ids) will be updated by sol values.
func (c *Client) UpdateSaleOrderLines(ids []int64, sol *SaleOrderLine) error {
	return c.Update(SaleOrderLineModel, ids, sol, nil)
}

// DeleteSaleOrderLine deletes an existing sale.order.line record.
func (c *Client) DeleteSaleOrderLine(id int64) error {
	return c.DeleteSaleOrderLines([]int64{id})
}

// DeleteSaleOrderLines deletes existing sale.order.line records.
func (c *Client) DeleteSaleOrderLines(ids []int64) error {
	return c.Delete(SaleOrderLineModel, ids)
}

// GetSaleOrderLine gets sale.order.line existing record.
func (c *Client) GetSaleOrderLine(id int64) (*SaleOrderLine, error) {
	sols, err := c.GetSaleOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sols)[0]), nil
}

// GetSaleOrderLines gets sale.order.line existing records.
func (c *Client) GetSaleOrderLines(ids []int64) (*SaleOrderLines, error) {
	sols := &SaleOrderLines{}
	if err := c.Read(SaleOrderLineModel, ids, nil, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLine finds sale.order.line record by querying it with criteria.
func (c *Client) FindSaleOrderLine(criteria *Criteria) (*SaleOrderLine, error) {
	sols := &SaleOrderLines{}
	if err := c.SearchRead(SaleOrderLineModel, criteria, NewOptions().Limit(1), sols); err != nil {
		return nil, err
	}
	return &((*sols)[0]), nil
}

// FindSaleOrderLines finds sale.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLines(criteria *Criteria, options *Options) (*SaleOrderLines, error) {
	sols := &SaleOrderLines{}
	if err := c.SearchRead(SaleOrderLineModel, criteria, options, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderLineModel, criteria, options)
}

// FindSaleOrderLineId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
