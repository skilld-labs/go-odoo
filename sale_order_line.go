package odoo

// SaleOrderLine represents sale.order.line model.
type SaleOrderLine struct {
	AllowedUomIds                     *Relation   `xmlrpc:"allowed_uom_ids,omitempty"`
	AmountInvoiced                    *Float      `xmlrpc:"amount_invoiced,omitempty"`
	AmountToInvoice                   *Float      `xmlrpc:"amount_to_invoice,omitempty"`
	AmountToInvoiceAtDate             *Float      `xmlrpc:"amount_to_invoice_at_date,omitempty"`
	AnalyticDistribution              interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticLineIds                   *Relation   `xmlrpc:"analytic_line_ids,omitempty"`
	AnalyticPrecision                 *Int        `xmlrpc:"analytic_precision,omitempty"`
	AutosalesBaseOrderLine            *Many2One   `xmlrpc:"autosales_base_order_line,omitempty"`
	AutosalesLine                     *Bool       `xmlrpc:"autosales_line,omitempty"`
	AvailableProductDocumentIds       *Relation   `xmlrpc:"available_product_document_ids,omitempty"`
	CategId                           *Many2One   `xmlrpc:"categ_id,omitempty"`
	CollapseComposition               *Bool       `xmlrpc:"collapse_composition,omitempty"`
	CollapsePrices                    *Bool       `xmlrpc:"collapse_prices,omitempty"`
	ComboItemId                       *Many2One   `xmlrpc:"combo_item_id,omitempty"`
	CompanyId                         *Many2One   `xmlrpc:"company_id,omitempty"`
	CompanyPriceInclude               *Selection  `xmlrpc:"company_price_include,omitempty"`
	CreateDate                        *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                         *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                        *Many2One   `xmlrpc:"currency_id,omitempty"`
	CustomerLead                      *Float      `xmlrpc:"customer_lead,omitempty"`
	Discount                          *Float      `xmlrpc:"discount,omitempty"`
	DisplayName                       *String     `xmlrpc:"display_name,omitempty"`
	DisplayQtyWidget                  *Bool       `xmlrpc:"display_qty_widget,omitempty"`
	DisplayType                       *Selection  `xmlrpc:"display_type,omitempty"`
	DistributionAnalyticAccountIds    *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	ExtraTaxData                      interface{} `xmlrpc:"extra_tax_data,omitempty"`
	ForecastExpectedDate              *Time       `xmlrpc:"forecast_expected_date,omitempty"`
	FreeQtyToday                      *Float      `xmlrpc:"free_qty_today,omitempty"`
	HasDisplayedWarningUpsell         *Bool       `xmlrpc:"has_displayed_warning_upsell,omitempty"`
	Id                                *Int        `xmlrpc:"id,omitempty"`
	InvoiceLines                      *Relation   `xmlrpc:"invoice_lines,omitempty"`
	InvoiceStatus                     *Selection  `xmlrpc:"invoice_status,omitempty"`
	IsConfigurableProduct             *Bool       `xmlrpc:"is_configurable_product,omitempty"`
	IsDownpayment                     *Bool       `xmlrpc:"is_downpayment,omitempty"`
	IsExpense                         *Bool       `xmlrpc:"is_expense,omitempty"`
	IsMto                             *Bool       `xmlrpc:"is_mto,omitempty"`
	IsOptional                        *Bool       `xmlrpc:"is_optional,omitempty"`
	IsProductArchived                 *Bool       `xmlrpc:"is_product_archived,omitempty"`
	IsService                         *Bool       `xmlrpc:"is_service,omitempty"`
	IsStorable                        *Bool       `xmlrpc:"is_storable,omitempty"`
	LinkedLineId                      *Many2One   `xmlrpc:"linked_line_id,omitempty"`
	LinkedLineIds                     *Relation   `xmlrpc:"linked_line_ids,omitempty"`
	LinkedVirtualId                   *String     `xmlrpc:"linked_virtual_id,omitempty"`
	MoveIds                           *Relation   `xmlrpc:"move_ids,omitempty"`
	Name                              *String     `xmlrpc:"name,omitempty"`
	OrderId                           *Many2One   `xmlrpc:"order_id,omitempty"`
	OrderPartnerId                    *Many2One   `xmlrpc:"order_partner_id,omitempty"`
	ParentId                          *Many2One   `xmlrpc:"parent_id,omitempty"`
	PriceReduceTaxexcl                *Float      `xmlrpc:"price_reduce_taxexcl,omitempty"`
	PriceReduceTaxinc                 *Float      `xmlrpc:"price_reduce_taxinc,omitempty"`
	PriceSubtotal                     *Float      `xmlrpc:"price_subtotal,omitempty"`
	PriceTax                          *Float      `xmlrpc:"price_tax,omitempty"`
	PriceTotal                        *Float      `xmlrpc:"price_total,omitempty"`
	PriceUnit                         *Float      `xmlrpc:"price_unit,omitempty"`
	PricelistItemId                   *Many2One   `xmlrpc:"pricelist_item_id,omitempty"`
	ProductCustomAttributeValueIds    *Relation   `xmlrpc:"product_custom_attribute_value_ids,omitempty"`
	ProductDocumentIds                *Relation   `xmlrpc:"product_document_ids,omitempty"`
	ProductId                         *Many2One   `xmlrpc:"product_id,omitempty"`
	ProductNoVariantAttributeValueIds *Relation   `xmlrpc:"product_no_variant_attribute_value_ids,omitempty"`
	ProductTemplateAttributeValueIds  *Relation   `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTemplateId                 *Many2One   `xmlrpc:"product_template_id,omitempty"`
	ProductType                       *Selection  `xmlrpc:"product_type,omitempty"`
	ProductUomId                      *Many2One   `xmlrpc:"product_uom_id,omitempty"`
	ProductUomQty                     *Float      `xmlrpc:"product_uom_qty,omitempty"`
	ProductUomReadonly                *Bool       `xmlrpc:"product_uom_readonly,omitempty"`
	ProductUpdatable                  *Bool       `xmlrpc:"product_updatable,omitempty"`
	ProjectId                         *Many2One   `xmlrpc:"project_id,omitempty"`
	PurchaseLineCount                 *Int        `xmlrpc:"purchase_line_count,omitempty"`
	PurchaseLineIds                   *Relation   `xmlrpc:"purchase_line_ids,omitempty"`
	QtyAvailableToday                 *Float      `xmlrpc:"qty_available_today,omitempty"`
	QtyDelivered                      *Float      `xmlrpc:"qty_delivered,omitempty"`
	QtyDeliveredAtDate                *Float      `xmlrpc:"qty_delivered_at_date,omitempty"`
	QtyDeliveredMethod                *Selection  `xmlrpc:"qty_delivered_method,omitempty"`
	QtyInvoiced                       *Float      `xmlrpc:"qty_invoiced,omitempty"`
	QtyInvoicedAtDate                 *Float      `xmlrpc:"qty_invoiced_at_date,omitempty"`
	QtyInvoicedPosted                 *Float      `xmlrpc:"qty_invoiced_posted,omitempty"`
	QtyToDeliver                      *Float      `xmlrpc:"qty_to_deliver,omitempty"`
	QtyToInvoice                      *Float      `xmlrpc:"qty_to_invoice,omitempty"`
	ReachedMilestonesIds              *Relation   `xmlrpc:"reached_milestones_ids,omitempty"`
	RemainingHours                    *Float      `xmlrpc:"remaining_hours,omitempty"`
	RemainingHoursAvailable           *Bool       `xmlrpc:"remaining_hours_available,omitempty"`
	RouteIds                          *Relation   `xmlrpc:"route_ids,omitempty"`
	SaleLineWarnMsg                   *String     `xmlrpc:"sale_line_warn_msg,omitempty"`
	SalesmanId                        *Many2One   `xmlrpc:"salesman_id,omitempty"`
	ScheduledDate                     *Time       `xmlrpc:"scheduled_date,omitempty"`
	SelectedComboItems                *String     `xmlrpc:"selected_combo_items,omitempty"`
	Sequence                          *Int        `xmlrpc:"sequence,omitempty"`
	ServiceTracking                   *Selection  `xmlrpc:"service_tracking,omitempty"`
	State                             *Selection  `xmlrpc:"state,omitempty"`
	TaskId                            *Many2One   `xmlrpc:"task_id,omitempty"`
	TaxCalculationRoundingMethod      *Selection  `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCountryId                      *Many2One   `xmlrpc:"tax_country_id,omitempty"`
	TaxIds                            *Relation   `xmlrpc:"tax_ids,omitempty"`
	TechnicalPriceUnit                *Float      `xmlrpc:"technical_price_unit,omitempty"`
	TimesheetIds                      *Relation   `xmlrpc:"timesheet_ids,omitempty"`
	TranslatedProductName             *String     `xmlrpc:"translated_product_name,omitempty"`
	UntaxedAmountInvoiced             *Float      `xmlrpc:"untaxed_amount_invoiced,omitempty"`
	UntaxedAmountToInvoice            *Float      `xmlrpc:"untaxed_amount_to_invoice,omitempty"`
	VirtualAvailableAtDate            *Float      `xmlrpc:"virtual_available_at_date,omitempty"`
	VirtualId                         *String     `xmlrpc:"virtual_id,omitempty"`
	WarehouseId                       *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WriteDate                         *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                          *Many2One   `xmlrpc:"write_uid,omitempty"`
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
