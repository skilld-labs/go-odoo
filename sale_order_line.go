package odoo

import (
	"fmt"
)

// SaleOrderLine represents sale.order.line model.
type SaleOrderLine struct {
	LastUpdate                        *Time      `xmlrpc:"__last_update,omitempty"`
	AnalyticLineIds                   *Relation  `xmlrpc:"analytic_line_ids,omitempty"`
	AnalyticTagIds                    *Relation  `xmlrpc:"analytic_tag_ids,omitempty"`
	CompanyId                         *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                        *Many2One  `xmlrpc:"currency_id,omitempty"`
	CustomerLead                      *Float     `xmlrpc:"customer_lead,omitempty"`
	Discount                          *Float     `xmlrpc:"discount,omitempty"`
	DisplayName                       *String    `xmlrpc:"display_name,omitempty"`
	DisplayType                       *Selection `xmlrpc:"display_type,omitempty"`
	Id                                *Int       `xmlrpc:"id,omitempty"`
	InvoiceLines                      *Relation  `xmlrpc:"invoice_lines,omitempty"`
	InvoiceStatus                     *Selection `xmlrpc:"invoice_status,omitempty"`
	IsDownpayment                     *Bool      `xmlrpc:"is_downpayment,omitempty"`
	IsExpense                         *Bool      `xmlrpc:"is_expense,omitempty"`
	Name                              *String    `xmlrpc:"name,omitempty"`
	OrderId                           *Many2One  `xmlrpc:"order_id,omitempty"`
	OrderPartnerId                    *Many2One  `xmlrpc:"order_partner_id,omitempty"`
	PriceReduce                       *Float     `xmlrpc:"price_reduce,omitempty"`
	PriceReduceTaxexcl                *Float     `xmlrpc:"price_reduce_taxexcl,omitempty"`
	PriceReduceTaxinc                 *Float     `xmlrpc:"price_reduce_taxinc,omitempty"`
	PriceSubtotal                     *Float     `xmlrpc:"price_subtotal,omitempty"`
	PriceTax                          *Float     `xmlrpc:"price_tax,omitempty"`
	PriceTotal                        *Float     `xmlrpc:"price_total,omitempty"`
	PriceUnit                         *Float     `xmlrpc:"price_unit,omitempty"`
	ProductCustomAttributeValueIds    *Relation  `xmlrpc:"product_custom_attribute_value_ids,omitempty"`
	ProductId                         *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductNoVariantAttributeValueIds *Relation  `xmlrpc:"product_no_variant_attribute_value_ids,omitempty"`
	ProductTemplateId                 *Many2One  `xmlrpc:"product_template_id,omitempty"`
	ProductUom                        *Many2One  `xmlrpc:"product_uom,omitempty"`
	ProductUomCategoryId              *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomQty                     *Float     `xmlrpc:"product_uom_qty,omitempty"`
	ProductUomReadonly                *Bool      `xmlrpc:"product_uom_readonly,omitempty"`
	ProductUpdatable                  *Bool      `xmlrpc:"product_updatable,omitempty"`
	QtyDelivered                      *Float     `xmlrpc:"qty_delivered,omitempty"`
	QtyDeliveredManual                *Float     `xmlrpc:"qty_delivered_manual,omitempty"`
	QtyDeliveredMethod                *Selection `xmlrpc:"qty_delivered_method,omitempty"`
	QtyInvoiced                       *Float     `xmlrpc:"qty_invoiced,omitempty"`
	QtyToInvoice                      *Float     `xmlrpc:"qty_to_invoice,omitempty"`
	SaleOrderOptionIds                *Relation  `xmlrpc:"sale_order_option_ids,omitempty"`
	SalesmanId                        *Many2One  `xmlrpc:"salesman_id,omitempty"`
	Sequence                          *Int       `xmlrpc:"sequence,omitempty"`
	State                             *Selection `xmlrpc:"state,omitempty"`
	TaxId                             *Relation  `xmlrpc:"tax_id,omitempty"`
	UntaxedAmountInvoiced             *Float     `xmlrpc:"untaxed_amount_invoiced,omitempty"`
	UntaxedAmountToInvoice            *Float     `xmlrpc:"untaxed_amount_to_invoice,omitempty"`
	WriteDate                         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                          *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(SaleOrderLineModel, sol)
}

// UpdateSaleOrderLine updates an existing sale.order.line record.
func (c *Client) UpdateSaleOrderLine(sol *SaleOrderLine) error {
	return c.UpdateSaleOrderLines([]int64{sol.Id.Get()}, sol)
}

// UpdateSaleOrderLines updates existing sale.order.line records.
// All records (represented by ids) will be updated by sol values.
func (c *Client) UpdateSaleOrderLines(ids []int64, sol *SaleOrderLine) error {
	return c.Update(SaleOrderLineModel, ids, sol)
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
	if sols != nil && len(*sols) > 0 {
		return &((*sols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.line not found", id)
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
	if sols != nil && len(*sols) > 0 {
		return &((*sols)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.line was not found")
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
	ids, err := c.Search(SaleOrderLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderLineId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.line was not found")
}
