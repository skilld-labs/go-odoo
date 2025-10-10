package odoo

// PurchaseOrderLine represents purchase.order.line model.
type PurchaseOrderLine struct {
	AnalyticDistribution              interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticPrecision                 *Int        `xmlrpc:"analytic_precision,omitempty"`
	CompanyId                         *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                        *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                         *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                        *Many2One   `xmlrpc:"currency_id,omitempty"`
	DateApprove                       *Time       `xmlrpc:"date_approve,omitempty"`
	DateOrder                         *Time       `xmlrpc:"date_order,omitempty"`
	DatePlanned                       *Time       `xmlrpc:"date_planned,omitempty"`
	Discount                          *Float      `xmlrpc:"discount,omitempty"`
	DisplayName                       *String     `xmlrpc:"display_name,omitempty"`
	DisplayType                       *Selection  `xmlrpc:"display_type,omitempty"`
	DistributionAnalyticAccountIds    *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	ForecastedIssue                   *Bool       `xmlrpc:"forecasted_issue,omitempty"`
	GroupId                           *Many2One   `xmlrpc:"group_id,omitempty"`
	Id                                *Int        `xmlrpc:"id,omitempty"`
	InvoiceLines                      *Relation   `xmlrpc:"invoice_lines,omitempty"`
	IsDownpayment                     *Bool       `xmlrpc:"is_downpayment,omitempty"`
	LocationFinalId                   *Many2One   `xmlrpc:"location_final_id,omitempty"`
	MoveDestIds                       *Relation   `xmlrpc:"move_dest_ids,omitempty"`
	MoveIds                           *Relation   `xmlrpc:"move_ids,omitempty"`
	Name                              *String     `xmlrpc:"name,omitempty"`
	OrderId                           *Many2One   `xmlrpc:"order_id,omitempty"`
	OrderpointId                      *Many2One   `xmlrpc:"orderpoint_id,omitempty"`
	PartnerId                         *Many2One   `xmlrpc:"partner_id,omitempty"`
	PriceSubtotal                     *Float      `xmlrpc:"price_subtotal,omitempty"`
	PriceTax                          *Float      `xmlrpc:"price_tax,omitempty"`
	PriceTotal                        *Float      `xmlrpc:"price_total,omitempty"`
	PriceUnit                         *Float      `xmlrpc:"price_unit,omitempty"`
	PriceUnitDiscounted               *Float      `xmlrpc:"price_unit_discounted,omitempty"`
	ProductDescriptionVariants        *String     `xmlrpc:"product_description_variants,omitempty"`
	ProductId                         *Many2One   `xmlrpc:"product_id,omitempty"`
	ProductNoVariantAttributeValueIds *Relation   `xmlrpc:"product_no_variant_attribute_value_ids,omitempty"`
	ProductPackagingId                *Many2One   `xmlrpc:"product_packaging_id,omitempty"`
	ProductPackagingQty               *Float      `xmlrpc:"product_packaging_qty,omitempty"`
	ProductQty                        *Float      `xmlrpc:"product_qty,omitempty"`
	ProductTemplateAttributeValueIds  *Relation   `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductType                       *Selection  `xmlrpc:"product_type,omitempty"`
	ProductUom                        *Many2One   `xmlrpc:"product_uom,omitempty"`
	ProductUomCategoryId              *Many2One   `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomQty                     *Float      `xmlrpc:"product_uom_qty,omitempty"`
	PropagateCancel                   *Bool       `xmlrpc:"propagate_cancel,omitempty"`
	QtyInvoiced                       *Float      `xmlrpc:"qty_invoiced,omitempty"`
	QtyReceived                       *Float      `xmlrpc:"qty_received,omitempty"`
	QtyReceivedManual                 *Float      `xmlrpc:"qty_received_manual,omitempty"`
	QtyReceivedMethod                 *Selection  `xmlrpc:"qty_received_method,omitempty"`
	QtyToInvoice                      *Float      `xmlrpc:"qty_to_invoice,omitempty"`
	SaleLineId                        *Many2One   `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId                       *Many2One   `xmlrpc:"sale_order_id,omitempty"`
	Sequence                          *Int        `xmlrpc:"sequence,omitempty"`
	State                             *Selection  `xmlrpc:"state,omitempty"`
	TaxCalculationRoundingMethod      *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxesId                           *Relation   `xmlrpc:"taxes_id,omitempty"`
	WriteDate                         *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                          *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PurchaseOrderLines represents array of purchase.order.line model.
type PurchaseOrderLines []PurchaseOrderLine

// PurchaseOrderLineModel is the odoo model name.
const PurchaseOrderLineModel = "purchase.order.line"

// Many2One convert PurchaseOrderLine to *Many2One.
func (pol *PurchaseOrderLine) Many2One() *Many2One {
	return NewMany2One(pol.Id.Get(), "")
}

// CreatePurchaseOrderLine creates a new purchase.order.line model and returns its id.
func (c *Client) CreatePurchaseOrderLine(pol *PurchaseOrderLine) (int64, error) {
	ids, err := c.CreatePurchaseOrderLines([]*PurchaseOrderLine{pol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseOrderLine creates a new purchase.order.line model and returns its id.
func (c *Client) CreatePurchaseOrderLines(pols []*PurchaseOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range pols {
		vv = append(vv, v)
	}
	return c.Create(PurchaseOrderLineModel, vv, nil)
}

// UpdatePurchaseOrderLine updates an existing purchase.order.line record.
func (c *Client) UpdatePurchaseOrderLine(pol *PurchaseOrderLine) error {
	return c.UpdatePurchaseOrderLines([]int64{pol.Id.Get()}, pol)
}

// UpdatePurchaseOrderLines updates existing purchase.order.line records.
// All records (represented by ids) will be updated by pol values.
func (c *Client) UpdatePurchaseOrderLines(ids []int64, pol *PurchaseOrderLine) error {
	return c.Update(PurchaseOrderLineModel, ids, pol, nil)
}

// DeletePurchaseOrderLine deletes an existing purchase.order.line record.
func (c *Client) DeletePurchaseOrderLine(id int64) error {
	return c.DeletePurchaseOrderLines([]int64{id})
}

// DeletePurchaseOrderLines deletes existing purchase.order.line records.
func (c *Client) DeletePurchaseOrderLines(ids []int64) error {
	return c.Delete(PurchaseOrderLineModel, ids)
}

// GetPurchaseOrderLine gets purchase.order.line existing record.
func (c *Client) GetPurchaseOrderLine(id int64) (*PurchaseOrderLine, error) {
	pols, err := c.GetPurchaseOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pols)[0]), nil
}

// GetPurchaseOrderLines gets purchase.order.line existing records.
func (c *Client) GetPurchaseOrderLines(ids []int64) (*PurchaseOrderLines, error) {
	pols := &PurchaseOrderLines{}
	if err := c.Read(PurchaseOrderLineModel, ids, nil, pols); err != nil {
		return nil, err
	}
	return pols, nil
}

// FindPurchaseOrderLine finds purchase.order.line record by querying it with criteria.
func (c *Client) FindPurchaseOrderLine(criteria *Criteria) (*PurchaseOrderLine, error) {
	pols := &PurchaseOrderLines{}
	if err := c.SearchRead(PurchaseOrderLineModel, criteria, NewOptions().Limit(1), pols); err != nil {
		return nil, err
	}
	return &((*pols)[0]), nil
}

// FindPurchaseOrderLines finds purchase.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseOrderLines(criteria *Criteria, options *Options) (*PurchaseOrderLines, error) {
	pols := &PurchaseOrderLines{}
	if err := c.SearchRead(PurchaseOrderLineModel, criteria, options, pols); err != nil {
		return nil, err
	}
	return pols, nil
}

// FindPurchaseOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PurchaseOrderLineModel, criteria, options)
}

// FindPurchaseOrderLineId finds record id by querying it with criteria.
func (c *Client) FindPurchaseOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
