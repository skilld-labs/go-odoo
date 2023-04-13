package odoo

import (
	"fmt"
)

// SaleOrderLine represents sale.order.line model.
type SaleOrderLine struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	AmtInvoiced            *Float     `xmlrpc:"amt_invoiced,omptempty"`
	AmtToInvoice           *Float     `xmlrpc:"amt_to_invoice,omptempty"`
	AnalyticTagIds         *Relation  `xmlrpc:"analytic_tag_ids,omptempty"`
	AutosalesBaseOrderLine *Many2One  `xmlrpc:"autosales_base_order_line,omptempty"`
	AutosalesLine          *Bool      `xmlrpc:"autosales_line,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omptempty"`
	CustomerLead           *Float     `xmlrpc:"customer_lead,omptempty"`
	Discount               *Float     `xmlrpc:"discount,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	InvoiceLines           *Relation  `xmlrpc:"invoice_lines,omptempty"`
	InvoiceStatus          *Selection `xmlrpc:"invoice_status,omptempty"`
	IsDownpayment          *Bool      `xmlrpc:"is_downpayment,omptempty"`
	IsService              *Bool      `xmlrpc:"is_service,omptempty"`
	LayoutCategoryId       *Many2One  `xmlrpc:"layout_category_id,omptempty"`
	LayoutCategorySequence *Int       `xmlrpc:"layout_category_sequence,omptempty"`
	MoveIds                *Relation  `xmlrpc:"move_ids,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	OrderId                *Many2One  `xmlrpc:"order_id,omptempty"`
	OrderPartnerId         *Many2One  `xmlrpc:"order_partner_id,omptempty"`
	PriceReduce            *Float     `xmlrpc:"price_reduce,omptempty"`
	PriceReduceTaxexcl     *Float     `xmlrpc:"price_reduce_taxexcl,omptempty"`
	PriceReduceTaxinc      *Float     `xmlrpc:"price_reduce_taxinc,omptempty"`
	PriceSubtotal          *Float     `xmlrpc:"price_subtotal,omptempty"`
	PriceTax               *Float     `xmlrpc:"price_tax,omptempty"`
	PriceTotal             *Float     `xmlrpc:"price_total,omptempty"`
	PriceUnit              *Float     `xmlrpc:"price_unit,omptempty"`
	ProductId              *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductImage           *String    `xmlrpc:"product_image,omptempty"`
	ProductPackaging       *Many2One  `xmlrpc:"product_packaging,omptempty"`
	ProductUom             *Many2One  `xmlrpc:"product_uom,omptempty"`
	ProductUomQty          *Float     `xmlrpc:"product_uom_qty,omptempty"`
	ProductUpdatable       *Bool      `xmlrpc:"product_updatable,omptempty"`
	QtyDelivered           *Float     `xmlrpc:"qty_delivered,omptempty"`
	QtyDeliveredUpdateable *Bool      `xmlrpc:"qty_delivered_updateable,omptempty"`
	QtyInvoiced            *Float     `xmlrpc:"qty_invoiced,omptempty"`
	QtyToInvoice           *Float     `xmlrpc:"qty_to_invoice,omptempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omptempty"`
	SalesmanId             *Many2One  `xmlrpc:"salesman_id,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	State                  *Selection `xmlrpc:"state,omptempty"`
	TaskId                 *Many2One  `xmlrpc:"task_id,omptempty"`
	TaxId                  *Relation  `xmlrpc:"tax_id,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(SaleOrderLineModel, vv)
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
	return nil, fmt.Errorf("sale.order.line was not found with criteria %v", criteria)
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
	return -1, fmt.Errorf("sale.order.line was not found with criteria %v and options %v", criteria, options)
}
