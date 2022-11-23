package odoo

import (
	"fmt"
)

// StockChangeProductQty represents stock.change.product.qty model.
type StockChangeProductQty struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	LocationId          *Many2One `xmlrpc:"location_id,omptempty"`
	LotId               *Many2One `xmlrpc:"lot_id,omptempty"`
	NewQuantity         *Float    `xmlrpc:"new_quantity,omptempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omptempty"`
	ProductTmplId       *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	ProductVariantCount *Int      `xmlrpc:"product_variant_count,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockChangeProductQtys represents array of stock.change.product.qty model.
type StockChangeProductQtys []StockChangeProductQty

// StockChangeProductQtyModel is the odoo model name.
const StockChangeProductQtyModel = "stock.change.product.qty"

// Many2One convert StockChangeProductQty to *Many2One.
func (scpq *StockChangeProductQty) Many2One() *Many2One {
	return NewMany2One(scpq.Id.Get(), "")
}

// CreateStockChangeProductQty creates a new stock.change.product.qty model and returns its id.
func (c *Client) CreateStockChangeProductQty(scpq *StockChangeProductQty) (int64, error) {
	return c.Create(StockChangeProductQtyModel, scpq)
}

// UpdateStockChangeProductQty updates an existing stock.change.product.qty record.
func (c *Client) UpdateStockChangeProductQty(scpq *StockChangeProductQty) error {
	return c.UpdateStockChangeProductQtys([]int64{scpq.Id.Get()}, scpq)
}

// UpdateStockChangeProductQtys updates existing stock.change.product.qty records.
// All records (represented by ids) will be updated by scpq values.
func (c *Client) UpdateStockChangeProductQtys(ids []int64, scpq *StockChangeProductQty) error {
	return c.Update(StockChangeProductQtyModel, ids, scpq)
}

// DeleteStockChangeProductQty deletes an existing stock.change.product.qty record.
func (c *Client) DeleteStockChangeProductQty(id int64) error {
	return c.DeleteStockChangeProductQtys([]int64{id})
}

// DeleteStockChangeProductQtys deletes existing stock.change.product.qty records.
func (c *Client) DeleteStockChangeProductQtys(ids []int64) error {
	return c.Delete(StockChangeProductQtyModel, ids)
}

// GetStockChangeProductQty gets stock.change.product.qty existing record.
func (c *Client) GetStockChangeProductQty(id int64) (*StockChangeProductQty, error) {
	scpqs, err := c.GetStockChangeProductQtys([]int64{id})
	if err != nil {
		return nil, err
	}
	if scpqs != nil && len(*scpqs) > 0 {
		return &((*scpqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.change.product.qty not found", id)
}

// GetStockChangeProductQtys gets stock.change.product.qty existing records.
func (c *Client) GetStockChangeProductQtys(ids []int64) (*StockChangeProductQtys, error) {
	scpqs := &StockChangeProductQtys{}
	if err := c.Read(StockChangeProductQtyModel, ids, nil, scpqs); err != nil {
		return nil, err
	}
	return scpqs, nil
}

// FindStockChangeProductQty finds stock.change.product.qty record by querying it with criteria.
func (c *Client) FindStockChangeProductQty(criteria *Criteria) (*StockChangeProductQty, error) {
	scpqs := &StockChangeProductQtys{}
	if err := c.SearchRead(StockChangeProductQtyModel, criteria, NewOptions().Limit(1), scpqs); err != nil {
		return nil, err
	}
	if scpqs != nil && len(*scpqs) > 0 {
		return &((*scpqs)[0]), nil
	}
	return nil, fmt.Errorf("stock.change.product.qty was not found with criteria %v", criteria)
}

// FindStockChangeProductQtys finds stock.change.product.qty records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockChangeProductQtys(criteria *Criteria, options *Options) (*StockChangeProductQtys, error) {
	scpqs := &StockChangeProductQtys{}
	if err := c.SearchRead(StockChangeProductQtyModel, criteria, options, scpqs); err != nil {
		return nil, err
	}
	return scpqs, nil
}

// FindStockChangeProductQtyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockChangeProductQtyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockChangeProductQtyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockChangeProductQtyId finds record id by querying it with criteria.
func (c *Client) FindStockChangeProductQtyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockChangeProductQtyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.change.product.qty was not found with criteria %v and options %v", criteria, options)
}
