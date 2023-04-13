package odoo

import (
	"fmt"
)

// PurchaseOrderLine represents purchase.order.line model.
type PurchaseOrderLine struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticId *Many2One  `xmlrpc:"account_analytic_id,omptempty"`
	AnalyticTagIds    *Relation  `xmlrpc:"analytic_tag_ids,omptempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId        *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateOrder         *Time      `xmlrpc:"date_order,omptempty"`
	DatePlanned       *Time      `xmlrpc:"date_planned,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	InvoiceLines      *Relation  `xmlrpc:"invoice_lines,omptempty"`
	MoveDestIds       *Relation  `xmlrpc:"move_dest_ids,omptempty"`
	MoveIds           *Relation  `xmlrpc:"move_ids,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	OrderId           *Many2One  `xmlrpc:"order_id,omptempty"`
	OrderpointId      *Many2One  `xmlrpc:"orderpoint_id,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	PriceSubtotal     *Float     `xmlrpc:"price_subtotal,omptempty"`
	PriceTax          *Float     `xmlrpc:"price_tax,omptempty"`
	PriceTotal        *Float     `xmlrpc:"price_total,omptempty"`
	PriceUnit         *Float     `xmlrpc:"price_unit,omptempty"`
	ProductId         *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductImage      *String    `xmlrpc:"product_image,omptempty"`
	ProductQty        *Float     `xmlrpc:"product_qty,omptempty"`
	ProductUom        *Many2One  `xmlrpc:"product_uom,omptempty"`
	QtyInvoiced       *Float     `xmlrpc:"qty_invoiced,omptempty"`
	QtyReceived       *Float     `xmlrpc:"qty_received,omptempty"`
	Sequence          *Int       `xmlrpc:"sequence,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	TaxesId           *Relation  `xmlrpc:"taxes_id,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(PurchaseOrderLineModel, []interface{}{pol})
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
	return c.Create(PurchaseOrderLineModel, vv)
}

// UpdatePurchaseOrderLine updates an existing purchase.order.line record.
func (c *Client) UpdatePurchaseOrderLine(pol *PurchaseOrderLine) error {
	return c.UpdatePurchaseOrderLines([]int64{pol.Id.Get()}, pol)
}

// UpdatePurchaseOrderLines updates existing purchase.order.line records.
// All records (represented by ids) will be updated by pol values.
func (c *Client) UpdatePurchaseOrderLines(ids []int64, pol *PurchaseOrderLine) error {
	return c.Update(PurchaseOrderLineModel, ids, pol)
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
	if pols != nil && len(*pols) > 0 {
		return &((*pols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of purchase.order.line not found", id)
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
	if pols != nil && len(*pols) > 0 {
		return &((*pols)[0]), nil
	}
	return nil, fmt.Errorf("purchase.order.line was not found with criteria %v", criteria)
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
	ids, err := c.Search(PurchaseOrderLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPurchaseOrderLineId finds record id by querying it with criteria.
func (c *Client) FindPurchaseOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("purchase.order.line was not found with criteria %v and options %v", criteria, options)
}
